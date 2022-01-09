package api

import (
	"context"
	"io/ioutil"
	"log"
	"net/http"

	pb "poker/api/grpc"
	"poker/api/model"
	"poker/api/token"
	"poker/poker"

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
	conn := initGrpc()
	defer conn.Close()
	client := pb.NewGetOauthCodeClient(conn)

	request := pb.Empty{}

	response, err := client.GetOauthCode(context.Background(), &request)

	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}

	c.JSON(http.StatusOK, response.GetUrl())
}

func oauthGetToken(c *gin.Context) {
	conn := initGrpc()
	defer conn.Close()
	client := pb.NewGetOauthTokenClient(conn)

	request := pb.GetOauthTokenRequest{
		Code: c.Query("code"),
	}

	_, err := client.GetOauthToken(context.Background(), &request)

	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	c.JSON(http.StatusOK, nil)
}

func oauthCheckToken(c *gin.Context) {
	conn := initGrpc()
	defer conn.Close()
	client := pb.NewCheckOauthTokenClient(conn)

	request := pb.Empty{}

	response, err := client.CheckOauthToken(context.Background(), &request)

	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	c.JSON(http.StatusOK, response.GetResult())
}

func getProfit(c *gin.Context) {

	conn := initGrpc()
	defer conn.Close()
	client := pb.NewGetProfitClient(conn)

	request := pb.GetAnalysisRequest{
		Username: c.GetString("username"),
		Player: c.Query("player"),
	}

	response, err := client.GetProfit(context.Background(), &request)

	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}

	results := []struct {
		Hand int
		Gain float64
	}{}

	for _, result := range response.GetResult() {
		results = append(results, struct {
			Hand int
			Gain float64
		}{Hand: int(result.GetHand()), Gain: result.GetGain()})
	}
	c.JSON(http.StatusOK, results)
}

func getPreflop(c *gin.Context) {
	conn := initGrpc()
	defer conn.Close()
	client := pb.NewGetPreflopClient(conn)

	request := pb.GetAnalysisRequest{
		Username: c.GetString("username"),
		Player: c.Query("player"),
	}

	response, err := client.GetPreflop(context.Background(), &request)

	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}

	result := struct {
		Raise string
		Call  string
		Fold  string
		Check string
		Bet   string
	}{
		Raise: response.GetRaise(),
		Call:  response.GetCall(),
		Fold:  response.GetFold(),
		Check: response.GetCheck(),
		Bet:   response.GetBet(),
	}

	c.JSON(http.StatusOK, result)
}

func getFlop(c *gin.Context) {
	conn := initGrpc()
	defer conn.Close()
	client := pb.NewGetFlopClient(conn)

	request := pb.GetAnalysisRequest{
		Username: c.GetString("username"),
		Player: c.Query("player"),
	}

	response, err := client.GetFlop(context.Background(), &request)

	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}

	result := struct {
		Raise string
		Call  string
		Fold  string
		Check string
		Bet   string
	}{
		Raise: response.GetRaise(),
		Call:  response.GetCall(),
		Fold:  response.GetFold(),
		Check: response.GetCheck(),
		Bet:   response.GetBet(),
	}

	c.JSON(http.StatusOK, result)
}

func getTurn(c *gin.Context) {
	conn := initGrpc()
	defer conn.Close()
	client := pb.NewGetTurnClient(conn)

	request := pb.GetAnalysisRequest{
		Username: c.GetString("username"),
		Player: c.Query("player"),
	}

	response, err := client.GetTurn(context.Background(), &request)

	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}

	result := struct {
		Raise string
		Call  string
		Fold  string
		Check string
		Bet   string
	}{
		Raise: response.GetRaise(),
		Call:  response.GetCall(),
		Fold:  response.GetFold(),
		Check: response.GetCheck(),
		Bet:   response.GetBet(),
	}

	c.JSON(http.StatusOK, result)
}

func getRiver(c *gin.Context) {
	conn := initGrpc()
	defer conn.Close()
	client := pb.NewGetRiverClient(conn)

	request := pb.GetAnalysisRequest{
		Username: c.GetString("username"),
		Player: c.Query("player"),
	}

	response, err := client.GetRiver(context.Background(), &request)

	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}

	result := struct {
		Raise string
		Call  string
		Fold  string
		Check string
		Bet   string
	}{
		Raise: response.GetRaise(),
		Call:  response.GetCall(),
		Fold:  response.GetFold(),
		Check: response.GetCheck(),
		Bet:   response.GetBet(),
	}

	c.JSON(http.StatusOK, result)
}

func getPlayer(c *gin.Context) {
	conn := initGrpc()
	defer conn.Close()
	client := pb.NewGetPlayerClient(conn)

	request := pb.GetPlayerRequest{
		Username: c.GetString("username"),
	}

	response, err := client.GetPlayer(context.Background(), &request)

	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}

	result := response.GetResult()

	c.JSON(http.StatusOK, result)
}
