package api

import (
	"io/ioutil"
	"net/http"

	"poker/api/kafka"
	"poker/api/model"
	"poker/api/oauth"
	"poker/api/token"
	"poker/poker"
	"strconv"

	"github.com/gin-gonic/gin"
)

var Tokens map[string]bool

func getWinRate(c *gin.Context) {

	t := poker.Table{}
	c.Bind(&t.Player)

	result := poker.GetWinRate(t.Player, 10000)
	c.JSON(http.StatusOK, result)
}

func getHand(c *gin.Context) {

	result := model.GetHandRedis(c.Query("num"), c.Query("gain"), c.Query("seat"), c.GetString("username"))

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
		player.Name = r.Player.Username

		player.Action.Preflop = r.Preflop
		player.Action.Flop = r.Flop
		player.Action.Turn = r.Turn
		player.Action.River = r.River

		table.Player = append(table.Player, player)
		tables = append(tables, table)
	}

	c.JSON(http.StatusOK, tables)
}

func putHand(c *gin.Context) {
	dataByte, _ := ioutil.ReadAll(c.Request.Body)
	username := c.GetString("username")
	kafka.KafkaWrite(dataByte, []byte(username))
	c.JSON(http.StatusOK, nil)
}

func login(c *gin.Context) {
	var request model.User
	c.BindJSON(&request)
	user := model.GetUserDB(request.Username)
	if user.Password == request.Password {
		tk := token.GenerateToken(user.Username)
		c.JSON(http.StatusOK, tk)
	} else {
		c.AbortWithStatus(http.StatusForbidden)
	}
}

func register(c *gin.Context) {
	var request model.User
	c.BindJSON(&request)
	model.InsertUserDB(request.Username, request.Password)
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
	url := oauth.GenerateCodeURL()
	c.JSON(http.StatusOK, url)
}

func oauthGetToken(c *gin.Context) {
	code := c.Query("code")
	oauthToken := oauth.GenerateTokenURL(code)
	username := oauth.GetUser(oauthToken)
	token := token.GenerateToken(username)
	oauth.OAuthChan <- token
}

func oauthCheckToken(c *gin.Context) {
	if len(oauth.OAuthChan) > 0 {
		result := <-oauth.OAuthChan
		c.JSON(http.StatusOK, result)
	}
}

func getPorfit(c *gin.Context) {

	profits := model.GetProfitRedis(c.GetString("username"))

	result := []struct {
		Hand int
		Gain float64
	}{}
	total := 0.0
	for count, profit := range profits {
		num, _ := strconv.ParseFloat(profit, 64)
		total += num
		result = append(result, struct {
			Hand int
			Gain float64
		}{Hand: count, Gain: total})
	}
	c.JSON(http.StatusOK, result)
}

func getPreflop(c *gin.Context) {
	username := c.GetString("username")

	result := struct {
		Raise string
		Call  string
		Fold  string
		Check string
		Bet string
	}{
		Raise: model.GetActionRedis("Preflop", "R", username),
		Call:  model.GetActionRedis("Preflop", "C", username),
		Fold:  model.GetActionRedis("Preflop", "F", username),
		Check: model.GetActionRedis("Preflop", "X", username),
		Bet: model.GetActionRedis("Preflop", "B", username),
	}

	c.JSON(http.StatusOK, result)
}

func getFlop(c *gin.Context) {
	username := c.GetString("username")

	result := struct {
		Raise string
		Call  string
		Fold  string
		Check string
		Bet string
	}{
		Raise: model.GetActionRedis("Flop", "R", username),
		Call:  model.GetActionRedis("Flop", "C", username),
		Fold:  model.GetActionRedis("Flop", "F", username),
		Check: model.GetActionRedis("Flop", "X", username),
		Bet: model.GetActionRedis("Flop", "B", username),
	}

	c.JSON(http.StatusOK, result)
}

func getTurn(c *gin.Context) {
	username := c.GetString("username")

	result := struct {
		Raise string
		Call  string
		Fold  string
		Check string
		Bet string
	}{
		Raise: model.GetActionRedis("Turn", "R", username),
		Call:  model.GetActionRedis("Turn", "C", username),
		Fold:  model.GetActionRedis("Turn", "F", username),
		Check: model.GetActionRedis("Turn", "X", username),
		Bet: model.GetActionRedis("Turn", "B", username),
	}

	c.JSON(http.StatusOK, result)
}

func getRiver(c *gin.Context) {
	username := c.GetString("username")

	result := struct {
		Raise string
		Call  string
		Fold  string
		Check string
		Bet string
	}{
		Raise: model.GetActionRedis("River", "R", username),
		Call:  model.GetActionRedis("River", "C", username),
		Fold:  model.GetActionRedis("River", "F", username),
		Check: model.GetActionRedis("River", "X", username),
		Bet: model.GetActionRedis("River", "B", username),
	}

	c.JSON(http.StatusOK, result)
}
