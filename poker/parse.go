package poker

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

func Parsefile(c *gin.Context) []Table {
	dataByte, err := ioutil.ReadAll(c.Request.Body)

	// err 沒有錯誤的話會回傳 nil
	if err != nil {
		// 當有錯誤時顯示錯誤訊息並離開程式
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	dataString := string(dataByte)            // 從 byte slice 轉成 string
	data := strings.Split(dataString, "\r\n") // 從 string 轉成 string slice

	tables := []Table{}
	for line := 0; line < len(data); line++ {
		if strings.Contains(data[line], "Poker Hand") {
			tables = append(tables, ParseTable(data, &line))
		}
	}
	println(tables)
	return tables
}

func ParseTable(data []string, line *int) Table {
	table := Table{}
	table.Player = append(table.Player, Player{Name: "Hero"})

	ParseBasic(data, line, &table)

	ParsePreFlop(data, line, &table)
	if CheckEnd(data, line, &table) {
		ParseShowdown(data, line, &table)
		return table
	}

	ParseFlop(data, line, &table)
	if CheckEnd(data, line, &table) {
		ParseShowdown(data, line, &table)
		return table
	}

	ParseTurn(data, line, &table)
	if CheckEnd(data, line, &table) {
		ParseShowdown(data, line, &table)
		return table
	}

	ParseRiver(data, line, &table)
	ParseShowdown(data, line, &table)

	return table
}

func CheckEnd(data []string, line *int, table *Table) bool {
	if strings.Contains(data[*line], "SHOWDOWN") ||
		(len(table.Player[0].Action.Preflop) != 0 && table.Player[0].Action.Preflop[len(table.Player[0].Action.Preflop)-1] == "folds") ||
		(len(table.Player[0].Action.Flop) != 0 && table.Player[0].Action.Flop[len(table.Player[0].Action.Flop)-1] == "folds") ||
		(len(table.Player[0].Action.Turn) != 0 && table.Player[0].Action.Turn[len(table.Player[0].Action.Turn)-1] == "folds") ||
		(len(table.Player[0].Action.River) != 0 && table.Player[0].Action.River[len(table.Player[0].Action.River)-1] == "folds") {
		return true
	}
	return false
}

func ParseShowdown(data []string, line *int, table *Table) {
	*line++
	pay := 0.0
	if strings.Contains(data[*line], "Hero") {
		str := strings.Split(data[*line], " ")
		pay, _ = strconv.ParseFloat(str[2][1:], 64)
	}

	table.Player[0].Gain += pay
}

func ParseRiver(data []string, line *int, table *Table) {
	if strings.Contains(data[*line], "RIVER") {
		card := GetCard(strings.Split(data[*line], " ")[3][1:3])
		(*table).Card = append((*table).Card, card)
	}
	pay, action := GetPay(data, line, "SHOWDOWN")
	(*table).Player[0].Gain += pay
	(*table).Player[0].Action.River = action
}

func ParseTurn(data []string, line *int, table *Table) {
	if strings.Contains(data[*line], "TURN") {
		card := GetCard(strings.Split(data[*line], " ")[6][1:3])

		(*table).Card = append((*table).Card, card)
	}
	pay, action := GetPay(data, line, "RIVER")
	(*table).Player[0].Gain += pay
	(*table).Player[0].Action.Turn = action
}

func ParseFlop(data []string, line *int, table *Table) {
	if strings.Contains(data[*line], "FLOP") {
		card1 := GetCard(strings.Split(data[*line], " ")[3][1:3])
		card2 := GetCard(strings.Split(data[*line], " ")[4][:2])
		card3 := GetCard(strings.Split(data[*line], " ")[5][:2])

		(*table).Card = append((*table).Card, card1)
		(*table).Card = append((*table).Card, card2)
		(*table).Card = append((*table).Card, card3)
	}

	pay, action := GetPay(data, line, "TURN")
	(*table).Player[0].Gain += pay
	(*table).Player[0].Action.Flop = action

}

func ParsePreFlop(data []string, line *int, table *Table) {
	for ; !strings.Contains(data[*line], "HOLE CARDS"); *line++ {
		if strings.Contains(data[*line], "Hero") {
			if strings.Contains(data[*line], "small blind") || strings.Contains(data[*line], "big blind") {
				str := strings.Split(data[*line], " ")
				(*table).Player[0].Gain, _ = strconv.ParseFloat(str[len(str)-1][1:], 64)
				(*table).Player[0].Gain *= -1
			}
		}
	}
	for ; strings.Split(data[*line], " ")[2] != "Hero"; *line++ {
	}
	card1 := GetCard(strings.Split(data[*line], " ")[3][1:3])
	card2 := GetCard(strings.Split(data[*line], " ")[4][:2])
	(*table).Player[0].Card = append((*table).Player[0].Card, card1)
	(*table).Player[0].Card = append((*table).Player[0].Card, card2)

	*line++

	pay, action := GetPay(data, line, "FLOP")
	(*table).Player[0].Gain += pay
	(*table).Player[0].Action.Preflop = action

}

func ParseBasic(data []string, line *int, table *Table) {
	timestr := strings.Split(data[*line], " ")[9] +" " + strings.Split(data[*line], " ")[10]
	timestr = strings.ReplaceAll(timestr, "/", "-")
	fmt.Println(timestr)
	(*table).Time, _ = time.Parse("2006-01-02 15:04:05", timestr)
	fmt.Println((*table).Time)
	id := strings.Split(data[*line], " ")[2]
	(*table).ID, _ = strconv.Atoi(id[3 : len(id)-1])

	for ; !strings.Contains(data[*line], "Hero"); *line++ {
	}

	switch strings.Split(data[*line], " ")[1][0] {
	case '1':
		(*table).Player[0].Seat = "BTN"
	case '2':
		(*table).Player[0].Seat = "SB"
	case '3':
		(*table).Player[0].Seat = "BB"
	case '4':
		(*table).Player[0].Seat = "LJ"
	case '5':
		(*table).Player[0].Seat = "HJ"
	case '6':
		(*table).Player[0].Seat = "CO"
	}
	*line += 1
}

func GetPay(data []string, line *int, nextState string) (float64, []string) {
	pay := 0.0
	action := []string{}
	for ; !strings.Contains(data[*line], nextState) && !strings.Contains(data[*line], "SHOWDOWN"); *line++ {
		if strings.Contains(data[*line], "Hero") {
			act := strings.Split(data[*line], " ")[1]
			if strings.Contains(data[*line], "returned to Hero") {
				str := strings.Split(data[*line], " ")
				tmp, _ := strconv.ParseFloat(str[2][2:6], 64)
				pay += tmp
			} else {
				if act != "shows" {
					action = append(action, act)
				}
				if strings.Contains(data[*line], "calls") || strings.Contains(data[*line], "raises") || strings.Contains(data[*line], "bets") {
					str := strings.Split(data[*line], " ")
					tmp, _ := strconv.ParseFloat(str[len(str)-1][1:], 64)
					pay += (tmp * -1)
				}
			}
		}
	}

	return pay, action
}

func GetCard(str string) Card {
	card := Card{}

	switch str[0] {
	case 'A':
		card.Num = 1
	case 'K':
		card.Num = 13
	case 'Q':
		card.Num = 12
	case 'J':
		card.Num = 11
	case 'T':
		card.Num = 10
	default:
		card.Num, _ = strconv.Atoi(string(str[0]))
	}
	card.Suit = string(str[1])

	return card
}
