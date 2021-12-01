package api

import (
	"fmt"
	"net/http"
	"poker/poker"
	"strconv"

	"github.com/gin-gonic/gin"
)

var Token []string

func getWinRate(c *gin.Context) {
	fmt.Println(Token)

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
	t := poker.Table{}
	t.Time = "2014-Feb-04"

	p1 := poker.Player{Name: c.Query("name1")}
	p1.Card[0].Num, _ = strconv.Atoi(c.Query("p1Card1Num"))
	p1.Card[0].Suit = c.Query("p1Card1Suit")
	p1.Card[1].Num, _ = strconv.Atoi(c.Query("p1Card2Num"))
	p1.Card[1].Suit = c.Query("p1Card2Suit")

	t.Player = append(t.Player, p1)

	c.JSON(http.StatusOK, t)
}

func putHand(c *gin.Context) {
	table := poker.Parse(c)
	go InsertDB(table)
	c.JSON(http.StatusOK, table)

}

func login(c *gin.Context) {
	if c.Query("Authorization") == "" {
		Token = append(Token, "123456")
	}
	c.JSON(http.StatusOK, Token)
}

func middlewaree(c *gin.Context){
	if len(Token) == 0{
		c.AbortWithStatus(http.StatusForbidden)
	}
}