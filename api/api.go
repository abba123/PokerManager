package api

import (
	"context"
	"io/ioutil"
	"log"
	"net/http"

	pb "poker/api/grpc"
	"poker/api/model"
	"poker/api/oauth"
	"poker/api/token"
	"poker/poker"
	"strconv"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func initGrpc() *grpc.ClientConn {
	conn, err := grpc.Dial("127.0.0.1:81", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	return conn
}

func getWinRate(c *gin.Context) {

	// Set up a connection to the server.
	conn := initGrpc()
	defer conn.Close()
	client := pb.NewGetWinRateServiceClient(conn)

	players := map[string]poker.Player{}
	c.Bind(&players)

	request := pb.GetWinRateRequest{}
	request.Player = make(map[string]*pb.Player)

	for key, value := range players {
		request.Player[key] = &pb.Player{Name: key}
		request.Player[key].Card1 = &pb.Card{Num: int32(value.Card[0].Num), Suit: value.Card[0].Suit}
		request.Player[key].Card2 = &pb.Card{Num: int32(value.Card[1].Num), Suit: value.Card[1].Suit}
	}

	response, err := client.GetWinRate(context.Background(), &request)

	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}

	c.JSON(http.StatusOK, response.GetResult())
}

func getHand(c *gin.Context) {

	conn := initGrpc()
	defer conn.Close()
	client := pb.NewGetHandServiceClient(conn)

	request := pb.GetHandRequest{
		Num:      c.Query("num"),
		Gain:     c.Query("gain"),
		Seat:     c.Query("seat"),
		Username: c.GetString("username"),
	}

	response, err := client.GetHand(context.Background(), &request)
	tables := []poker.Table{}
	for _, res := range response.GetTable() {
		table := poker.Table{
			ID:   int(res.GetId()),
			Time: res.GetTime().AsTime(),
		}

		table.Player = make(map[string]poker.Player)
		for key, value := range res.GetPlayer() {
			player := poker.Player{
				Name: value.GetName(),
				Seat: value.GetSeat(),
				Gain: value.GetGain(),
				Action: struct {
					Preflop string
					Flop    string
					Turn    string
					River   string
				}{
					Preflop: value.Action.GetPreflop(),
					Flop:    value.GetAction().Flop,
					Turn:    value.Action.GetTurn(),
					River:   value.Action.River,
				},
				Card: []poker.Card{
					{Num: int(value.GetCard1().GetNum()), Suit: value.GetCard2().GetSuit()},
					{Num: int(value.GetCard2().GetNum()), Suit: value.GetCard2().GetSuit()},
				},
			}
			table.Player[key] = player
		}
		tables = append(tables, table)
	}

	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}

	c.JSON(http.StatusOK, tables)
}

func insertHand(c *gin.Context) {

	conn := initGrpc()
	defer conn.Close()
	client := pb.NewInsertHandServiceClient(conn)

	dataByte, _ := ioutil.ReadAll(c.Request.Body)
	username := c.GetString("username")

	request := pb.InsertHandRequest{
		Data:     string(dataByte),
		Username: username,
	}

	_, err := client.InsertHand(context.Background(), &request)
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}

	c.JSON(http.StatusOK, nil)
}

func login(c *gin.Context) {

	conn := initGrpc()
	defer conn.Close()
	client := pb.NewLoginServiceClient(conn)

	var user model.User
	c.BindJSON(&user)

	request := pb.LoginRequest{
		Username: user.Username,
		Password: user.Password,
	}

	response, err := client.Login(context.Background(), &request)
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}

	token := response.GetToken()
	if token != "" {
		c.JSON(http.StatusOK, token)
	} else {
		c.AbortWithStatus(http.StatusForbidden)
	}
}

func register(c *gin.Context) {
	conn := initGrpc()
	defer conn.Close()
	client := pb.NewRegisterServiceClient(conn)

	var user model.User
	c.BindJSON(&user)

	request := pb.LoginRequest{
		Username: user.Username,
		Password: user.Password,
	}

	response, err := client.Register(context.Background(), &request)
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}

	if response.GetError() != "" {
		c.JSON(http.StatusForbidden, nil)
	} else {
		c.JSON(http.StatusOK, nil)
	}
}

func logout(c *gin.Context) {
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

	profits := model.GetProfitRedis(c.GetString("username"), c.Query("player"))

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
	player := c.Query("player")

	result := struct {
		Raise string
		Call  string
		Fold  string
		Check string
		Bet   string
	}{
		Raise: model.GetActionRedis("pre_flop", "R", username, player),
		Call:  model.GetActionRedis("pre_flop", "C", username, player),
		Fold:  model.GetActionRedis("pre_flop", "F", username, player),
		Check: model.GetActionRedis("pre_flop", "X", username, player),
		Bet:   model.GetActionRedis("pre_flop", "B", username, player),
	}

	c.JSON(http.StatusOK, result)
}

func getFlop(c *gin.Context) {
	username := c.GetString("username")
	player := c.Query("player")

	result := struct {
		Raise string
		Call  string
		Fold  string
		Check string
		Bet   string
	}{
		Raise: model.GetActionRedis("flop", "R", username, player),
		Call:  model.GetActionRedis("flop", "C", username, player),
		Fold:  model.GetActionRedis("flop", "F", username, player),
		Check: model.GetActionRedis("flop", "X", username, player),
		Bet:   model.GetActionRedis("flop", "B", username, player),
	}

	c.JSON(http.StatusOK, result)
}

func getTurn(c *gin.Context) {
	username := c.GetString("username")
	player := c.Query("player")

	result := struct {
		Raise string
		Call  string
		Fold  string
		Check string
		Bet   string
	}{
		Raise: model.GetActionRedis("turn", "R", username, player),
		Call:  model.GetActionRedis("turn", "C", username, player),
		Fold:  model.GetActionRedis("turn", "F", username, player),
		Check: model.GetActionRedis("turn", "X", username, player),
		Bet:   model.GetActionRedis("turn", "B", username, player),
	}

	c.JSON(http.StatusOK, result)
}

func getRiver(c *gin.Context) {
	username := c.GetString("username")
	player := c.Query("player")

	result := struct {
		Raise string
		Call  string
		Fold  string
		Check string
		Bet   string
	}{
		Raise: model.GetActionRedis("river", "R", username, player),
		Call:  model.GetActionRedis("river", "C", username, player),
		Fold:  model.GetActionRedis("river", "F", username, player),
		Check: model.GetActionRedis("river", "X", username, player),
		Bet:   model.GetActionRedis("river", "B", username, player),
	}

	c.JSON(http.StatusOK, result)
}

func getPlayer(c *gin.Context) {
	username := c.GetString("username")

	result := model.GetPlayerRedis(username)

	c.JSON(http.StatusOK, result)
}
