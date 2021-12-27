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

	result := poker.GetWinRate([]poker.Player{t.Player["Player1"], t.Player["Player2"]}, 10000)
	c.JSON(http.StatusOK, result)
}

func getHand(c *gin.Context) {

	result := model.GetHandRedis(c.Query("num"), c.Query("gain"), c.Query("seat"), c.GetString("username"))
	tables := []poker.Table{}

	for _, r := range result {
		table := poker.Table{}

		table.Time = r.Time
		if r.TableCard1.ID != 0 {
			table.Card = append(table.Card, poker.Card{})
			table.Card[0].Num = r.TableCard1.Num
			table.Card[0].Suit = r.TableCard1.Suit
		}
		if r.TableCard2.ID != 0 {
			table.Card = append(table.Card, poker.Card{})
			table.Card[1].Num = r.TableCard2.Num
			table.Card[1].Suit = r.TableCard2.Suit
		}
		if r.TableCard3.ID != 0 {
			table.Card = append(table.Card, poker.Card{})
			table.Card[2].Num = r.TableCard3.Num
			table.Card[2].Suit = r.TableCard3.Suit
		}
		if r.TableCard4.ID != 0 {
			table.Card = append(table.Card, poker.Card{})
			table.Card[3].Num = r.TableCard4.Num
			table.Card[3].Suit = r.TableCard4.Suit
		}
		if r.TableCard5.ID != 0 {
			table.Card = append(table.Card, poker.Card{})
			table.Card[4].Num = r.TableCard5.Num
			table.Card[4].Suit = r.TableCard5.Suit
		}
		player := poker.Player{}

		player.Card = append(player.Card, poker.Card{})
		player.Card[0].Num = r.HeroCard1.Num
		player.Card[0].Suit = r.HeroCard1.Suit
		player.Card = append(player.Card, poker.Card{})
		player.Card[1].Num = r.HeroCard2.Num
		player.Card[1].Suit = r.HeroCard2.Suit
		player.Gain = r.Gain
		player.Seat = r.Seat.Seat
		player.Name = r.Player.Playername

		player.Action.Preflop = r.Preflop.Action
		player.Action.Flop = r.Flop.Action
		player.Action.Turn = r.Turn.Action
		player.Action.River = r.River.Action
		table.Player = make(map[string]poker.Player)
		table.Player[c.GetString("username")] = player
		tables = append(tables, table)
	}

	c.JSON(http.StatusOK, tables)
}

func insertHand(c *gin.Context) {
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
	err := model.InsertUserDB(request.Username, request.Password)
	if err != nil {
		c.JSON(http.StatusForbidden, nil)
	} else {
		c.JSON(http.StatusOK, nil)
	}
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
		Bet   string
	}{
		Raise: model.GetActionRedis("pre_flop", "R", username),
		Call:  model.GetActionRedis("pre_flop", "C", username),
		Fold:  model.GetActionRedis("pre_flop", "F", username),
		Check: model.GetActionRedis("pre_flop", "X", username),
		Bet:   model.GetActionRedis("pre_flop", "B", username),
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
		Bet   string
	}{
		Raise: model.GetActionRedis("flop", "R", username),
		Call:  model.GetActionRedis("flop", "C", username),
		Fold:  model.GetActionRedis("flop", "F", username),
		Check: model.GetActionRedis("flop", "X", username),
		Bet:   model.GetActionRedis("flop", "B", username),
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
		Bet   string
	}{
		Raise: model.GetActionRedis("turn", "R", username),
		Call:  model.GetActionRedis("turn", "C", username),
		Fold:  model.GetActionRedis("turn", "F", username),
		Check: model.GetActionRedis("turn", "X", username),
		Bet:   model.GetActionRedis("turn", "B", username),
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
		Bet   string
	}{
		Raise: model.GetActionRedis("river", "R", username),
		Call:  model.GetActionRedis("river", "C", username),
		Fold:  model.GetActionRedis("river", "F", username),
		Check: model.GetActionRedis("river", "X", username),
		Bet:   model.GetActionRedis("river", "B", username),
	}

	c.JSON(http.StatusOK, result)
}
