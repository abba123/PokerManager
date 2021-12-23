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

	results := model.GetHandRedis(c.Query("num"), c.Query("gain"), c.Query("seat"), c.GetString("username"))

	tables := []poker.Table{}

	for _, result := range results {
		table := poker.Table{}

		table.Time = result.Time

		for _, card := range result.TableCard {
			table.Card = append(table.Card, poker.Card(card))
		}

		player := poker.Player{}

		player.Card = append(player.Card, poker.Card(result.HeroCard[0]))
		player.Card = append(player.Card, poker.Card(result.HeroCard[1]))

		player.Gain = result.Gain
		player.Seat = result.Seat.Location
		player.Name = result.User.Username

		player.Action.Preflop = result.Preflop.Action
		player.Action.Flop = result.Flop.Action
		player.Action.Turn = result.Turn.Action
		player.Action.River = result.River.Action

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
	}{Raise: model.GetActionRedis("Preflop", "Raise", username), Call: model.GetActionRedis("Preflop", "Call", username), Fold: model.GetActionRedis("Preflop", "Fold", username)}

	c.JSON(http.StatusOK, result)
}

func getFlop(c *gin.Context) {
	username := c.GetString("username")

	result := struct {
		Raise string
		Call  string
		Fold  string
	}{Raise: model.GetActionRedis("Flop", "Raise", username), Call: model.GetActionRedis("Flop", "Call", username), Fold: model.GetActionRedis("Flop", "Fold", username)}

	c.JSON(http.StatusOK, result)
}

func getTurn(c *gin.Context) {
	username := c.GetString("username")

	result := struct {
		Raise string
		Call  string
		Fold  string
	}{Raise: model.GetActionRedis("Turn", "Raise", username), Call: model.GetActionRedis("Turn", "Call", username), Fold: model.GetActionRedis("Turn", "Fold", username)}

	c.JSON(http.StatusOK, result)
}

func getRiver(c *gin.Context) {
	username := c.GetString("username")

	result := struct {
		Raise string
		Call  string
		Fold  string
	}{Raise: model.GetActionRedis("River", "Raise", username), Call: model.GetActionRedis("River", "Call", username), Fold: model.GetActionRedis("River", "Fold", username)}

	c.JSON(http.StatusOK, result)
}
