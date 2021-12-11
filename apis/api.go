package api

import (
	"fmt"
	"net/http"

	oauth "poker/apis/OAuth"
	"poker/apis/token"
	"poker/poker"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

var Tokens map[string]bool

func getWinRate(c *gin.Context) {

	t := poker.Table{}

	p1 := poker.Player{Name: c.Query("name1")}
	p2 := poker.Player{Name: c.Query("name2")}

	p1.Card = []poker.Card{{}, {}}
	p1.Card[0].Num, _ = strconv.Atoi(c.Query("p1Card1Num"))
	p1.Card[0].Suit = c.Query("p1Card1Suit")
	p1.Card[1].Num, _ = strconv.Atoi(c.Query("p1Card2Num"))
	p1.Card[1].Suit = c.Query("p1Card2Suit")

	p2.Card = []poker.Card{{}, {}}
	p2.Card[0].Num, _ = strconv.Atoi(c.Query("p2Card1Num"))
	p2.Card[0].Suit = c.Query("p2Card1Suit")
	p2.Card[1].Num, _ = strconv.Atoi(c.Query("p2Card2Num"))
	p2.Card[1].Suit = c.Query("p2Card2Suit")

	t.Player = append(t.Player, p1)
	t.Player = append(t.Player, p2)

	result := poker.GetWinRate(t.Player, 10000)

	c.JSON(http.StatusOK, result)
}

func getHand(c *gin.Context) {

	result := getHandDB(c.Query("num"), c.Query("gain"), c.Query("seat"), c.GetString("username"))

	tables := []poker.Table{}

	for _, r := range result {
		table := poker.Table{}

		table.Time = r.Time
		if r.TableCard1 != "" {
			table.Card = append(table.Card, poker.Card{})
			table.Card[0].Num, _ = strconv.Atoi(string(r.TableCard1[:len(r.TableCard1)-1]))
			table.Card[0].Suit = string(r.TableCard1[len(r.TableCard1)-1:])
		}
		if r.TableCard2 != "" {
			table.Card = append(table.Card, poker.Card{})
			table.Card[1].Num, _ = strconv.Atoi(string(r.TableCard2[:len(r.TableCard1)-1]))
			table.Card[1].Suit = string(r.TableCard2[len(r.TableCard1)-1:])
		}
		if r.TableCard3 != "" {
			table.Card = append(table.Card, poker.Card{})
			table.Card[2].Num, _ = strconv.Atoi(string(r.TableCard3[:len(r.TableCard1)-1]))
			table.Card[2].Suit = string(r.TableCard3[len(r.TableCard1)-1:])
		}
		if r.TableCard4 != "" {
			table.Card = append(table.Card, poker.Card{})
			table.Card[3].Num, _ = strconv.Atoi(string(r.TableCard4[:len(r.TableCard1)-1]))
			table.Card[3].Suit = string(r.TableCard4[len(r.TableCard1)-1:])
		}
		if r.TableCard5 != "" {
			table.Card = append(table.Card, poker.Card{})
			table.Card[4].Num, _ = strconv.Atoi(string(r.TableCard5[:len(r.TableCard1)-1]))
			table.Card[4].Suit = string(r.TableCard5[len(r.TableCard1)-1:])
		}
		player := poker.Player{}

		player.Card = append(player.Card, poker.Card{})
		player.Card[0].Num, _ = strconv.Atoi(string(r.HeroCard1[:len(r.HeroCard1)-1]))
		player.Card[0].Suit = string(r.HeroCard1[len(r.HeroCard1)-1:])
		player.Card = append(player.Card, poker.Card{})
		player.Card[1].Num, _ = strconv.Atoi(string(r.HeroCard2[:len(r.HeroCard2)-1]))
		player.Card[1].Suit = string(r.HeroCard2[len(r.HeroCard2)-1:])
		player.Gain = r.Gain
		player.Seat = r.Seat
		player.Name = r.Player

		if r.Preflop != "" {
			player.Action.Preflop = strings.Split(r.Preflop, " ")
		}
		if r.Flop != "" {
			player.Action.Flop = strings.Split(r.Flop, " ")
		}
		if r.Turn != "" {
			player.Action.Turn = strings.Split(r.Turn, " ")
		}
		if r.River != "" {
			player.Action.River = strings.Split(r.River, " ")
		}

		table.Player = append(table.Player, player)
		tables = append(tables, table)
	}

	c.JSON(http.StatusOK, tables)
}

func putHand(c *gin.Context) {
	table := poker.Parsefile(c)
	InsertHandDB(table)
	c.JSON(http.StatusOK, table)
}

func login(c *gin.Context) {
	var request user
	c.BindJSON(&request)
	user := GetUserDB(request.Username)
	if user.Password == request.Password {
		tk := token.GenerateToken(user.Username)
		c.JSON(http.StatusOK, tk)
	} else {
		c.AbortWithStatus(http.StatusForbidden)
	}
}

func register(c *gin.Context) {
	var request user
	c.BindJSON(&request)
	InsertUserDB(request.Username, request.Password)
}

func logout(c *gin.Context) {
	token := c.Request.Header.Get("Authorization")
	delete(Tokens, token)
	c.JSON(http.StatusOK, nil)
}

func middlewaree(c *gin.Context) {
	tk := c.Request.Header.Get("Authorization")
	claim, err := token.ValidToken(tk)

	if err != nil || claim.Authority != 0 {
		c.AbortWithStatus(http.StatusForbidden)
	}

	c.Set("username", claim.Username)
}

func oauthGetCode(c *gin.Context) {
	fmt.Println("get code")
	url := oauth.GenerateCodeURL()
	c.JSON(http.StatusOK, url)
}

func oauthGetToken(c *gin.Context) {
	fmt.Println("get token")
	code := c.Query("code")
	token := oauth.GenerateTokenURL(code)
	Tokens[token] = true

	oauth.OAuthChan <- token
	fmt.Println("finish token")

}

func oauthCheckToken(c *gin.Context) {
	if len(oauth.OAuthChan) > 0 {
		result := <-oauth.OAuthChan
		fmt.Println("get check")
		fmt.Println(result)
		c.JSON(http.StatusOK, result)
		fmt.Println("finish check")
	}
}

func getAnalysis(c *gin.Context) {

	profits := getProfitDB(c.GetString("username"))

	result := []struct {
		Hand int
		Gain float64
	}{}
	total := 0.0
	for count, profit := range profits {
		total += profit
		result = append(result, struct {
			Hand int
			Gain float64
		}{Hand: count, Gain: total})
	}
	c.JSON(http.StatusOK, result)
}
