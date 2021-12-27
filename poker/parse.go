package poker

import (
	"strconv"
	"strings"
	"time"
)

func Parsefile(username string, data string) []Table {
	// 從 byte slice 轉成 string
	data = strings.Replace(data, "\r\n", "\n", -1)
	dataSlice := strings.Split(data, "\n") // 從 string 轉成 string slice

	tables := []Table{}

	for line := 4; line < len(dataSlice); line++ {
		if strings.Contains(dataSlice[line], "Poker Hand") {
			tables = append(tables, ParseTable(dataSlice, &line, username))
		}
	}
	return tables
}

func ParseTable(data []string, line *int, name string) Table {
	table := Table{}
	ParseBasic(data, line, &table)

	ParsePreFlop(data, line, name, &table)
	if CheckEnd(data, line, &table) {
		ParseShowdown(data, line, name, &table)
		return table
	}

	ParseFlop(data, line, name, &table)
	if CheckEnd(data, line, &table) {
		ParseShowdown(data, line, name, &table)
		return table
	}

	ParseTurn(data, line, name, &table)
	if CheckEnd(data, line, &table) {
		ParseShowdown(data, line, name, &table)
		return table
	}

	ParseRiver(data, line, name, &table)
	ParseShowdown(data, line, name, &table)

	return table
}

func CheckEnd(data []string, line *int, table *Table) bool {
	return  strings.Contains(data[*line], "SHOWDOWN")
}

func ParseShowdown(data []string, line *int, name string, table *Table) {
	*line++
	pay := 0.0
	s := strings.Split(data[*line], " ")
	playerName := s[0]
	if playerName == "Hero" {
		playerName = name
	}
	player := table.Player[playerName]
	pay, _ = strconv.ParseFloat(s[2][1:], 64)
	player.Gain += pay
	table.Player[playerName] = player
}

func ParseRiver(data []string, line *int, name string, table *Table) {
	if strings.Contains(data[*line], "RIVER") {
		s := strings.Split(data[*line], " ")
		card := GetCard(s[len(s)-1][1:3])
		(*table).Card = append((*table).Card, card)
	}
	GetPay(data, line, "SHOWDOWN", name, table)
}

func ParseTurn(data []string, line *int, name string, table *Table) {
	if strings.Contains(data[*line], "TURN") {
		s := strings.Split(data[*line], " ")
		card := GetCard(s[len(s)-1][1:3])
		(*table).Card = append((*table).Card, card)
	}
	GetPay(data, line, "RIVER", name, table)
}

func ParseFlop(data []string, line *int, name string, table *Table) {
	if strings.Contains(data[*line], "FLOP") {
		card1 := GetCard(strings.Split(data[*line], " ")[3][1:3])
		card2 := GetCard(strings.Split(data[*line], " ")[4][:2])
		card3 := GetCard(strings.Split(data[*line], " ")[5][:2])

		(*table).Card = append((*table).Card, card1)
		(*table).Card = append((*table).Card, card2)
		(*table).Card = append((*table).Card, card3)
	}

	GetPay(data, line, "TURN", name, table)

}

func ParsePreFlop(data []string, line *int, name string, table *Table) {
	for i := 0; i < 6; i++ {
		playerName := strings.Split(data[*line], " ")[2]
		if playerName == "Hero" {
			player := table.Player[name]
			card1 := GetCard(strings.Split(data[*line], " ")[3][1:3])
			card2 := GetCard(strings.Split(data[*line], " ")[4][:2])
			player.Card = append(player.Card, card1)
			player.Card = append(player.Card, card2)
			(*table).Player[playerName] = player
		} else {
			player := table.Player[playerName]
			(*table).Player[playerName] = player
		}
		*line += 1
	}
	GetPay(data, line, "FLOP", name, table)
}

func ParseBasic(data []string, line *int, table *Table) {
	timestr := strings.Split(data[*line], " ")[9] + " " + strings.Split(data[*line], " ")[10]
	timestr = strings.ReplaceAll(timestr, "/", "-")
	(*table).Time, _ = time.Parse("2006-01-02 15:04:05", timestr)
	id := strings.Split(data[*line], " ")[2]
	(*table).ID, _ = strconv.Atoi(id[3 : len(id)-1])

	for ; strings.Contains(data[*line], "Seat"); *line++ {
		s := strings.Split(data[*line], " ")
		name := s[2]
		player := Player{Name: name}
		switch s[1][0] {
		case '1':
			player.Seat = "BTN"
		case '2':
			player.Seat = "SB"
		case '3':
			player.Seat = "BB"
		case '4':
			player.Seat = "LJ"
		case '5':
			player.Seat = "HJ"
		case '6':
			player.Seat = "CO"
		}
		table.Player[name] = player
	}
	for i := 0; i < 2; i++ {
		s := strings.Split(data[*line], " ")
		name := s[2]
		gain, _ := strconv.ParseFloat(s[len(s)-1][1:], 64)
		gain *= -1
		player := table.Player[name]
		player.Gain = gain
		table.Player[name] = player
		*line += 1
	}
}

func GetPay(data []string, line *int, nextState string, name string, table *Table) {
	state := strings.Split(data[*line], " ")[1]
	for ; !strings.Contains(data[*line], nextState) && !strings.Contains(data[*line], "SHOWDOWN"); *line++ {
		s := strings.Split(data[*line], " ")
		act := s[1]
		playerName := ""
		action := ""
		pay := 0.0
		if strings.Contains(data[*line], "returned to") {
			playerName = s[len(s)-1]
			tmp, _ := strconv.ParseFloat(s[2][2:6], 64)
			pay += tmp
		} else {
			playerName = s[0][:len(s[0])-1]
			if playerName == "Hero" {
				playerName = name
			}
			if strings.Contains(data[*line], "calls") || strings.Contains(data[*line], "raises") || strings.Contains(data[*line], "bets") {
				str := strings.Split(data[*line], " ")
				tmp, _ := strconv.ParseFloat(str[len(str)-1][1:], 64)
				pay += (tmp * -1)
			}
		}

		if act != "shows" {
			switch act {
			case "raises":
				action += "R"
			case "calls":
				action += "C"
			case "folds":
				action += "F"
			case "checks":
				action += "X"
			case "bets":
				action += "B"
			}
		}
		player := table.Player[playerName]
		switch state {
		case "HOLE":
			player.Action.Preflop = action
		case "FLOP":
			player.Action.Flop = action
		case "TURN":
			player.Action.Turn = action
		case "RIVER":
			player.Action.River = action
		}
		player.Gain += pay

		table.Player[playerName] = player
	}
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
