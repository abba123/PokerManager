package proto

import (
	"context"
	"fmt"
	"log"
	"net"
	"poker/api/kafka"
	"poker/api/model"
	"poker/api/oauth"
	"poker/api/token"
	"poker/poker"
	"strconv"

	"github.com/shopspring/decimal"
	"google.golang.org/grpc"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

type Server struct{}

func RunGrpcSetver() {
	fmt.Println("starting gRPC server...")

	lis, err := net.Listen("tcp", "127.0.0.1:81")
	if err != nil {
		log.Fatalf("failed to listen: %v \n", err)
	}

	grpcServer := grpc.NewServer()

	RegisterGetWinRateServiceServer(grpcServer, &Server{})
	RegisterLoginServiceServer(grpcServer, &Server{})
	RegisterRegisterServiceServer(grpcServer, &Server{})
	RegisterInsertHandServiceServer(grpcServer, &Server{})
	RegisterGetHandServiceServer(grpcServer, &Server{})
	RegisterGetOauthCodeServer(grpcServer, &Server{})
	RegisterGetOauthTokenServer(grpcServer, &Server{})
	RegisterCheckOauthTokenServer(grpcServer, &Server{})
	RegisterGetProfitServer(grpcServer, &Server{})
	RegisterGetPreflopServer(grpcServer, &Server{})
	RegisterGetFlopServer(grpcServer, &Server{})
	RegisterGetTurnServer(grpcServer, &Server{})
	RegisterGetRiverServer(grpcServer, &Server{})
	RegisterGetPlayerServer(grpcServer, &Server{})

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v \n", err)
	}
}

func (*Server) GetWinRate(ctx context.Context, req *GetWinRateRequest) (*GetWinRateResponse, error) {

	players := []poker.Player{}

	for key, value := range req.GetPlayer() {
		player := poker.Player{Name: key}
		player.Card = append(player.Card, poker.Card{Num: int(value.GetCard1().GetNum()), Suit: value.GetCard1().GetSuit()})
		player.Card = append(player.Card, poker.Card{Num: int(value.GetCard2().GetNum()), Suit: value.GetCard2().GetSuit()})
		players = append(players, player)
	}

	result := poker.GetWinRate(players, 10000)

	response := &GetWinRateResponse{
		Result: result,
	}

	return response, nil
}
func (*Server) Login(ctx context.Context, req *LoginRequest) (*LoginResponse, error) {
	user := model.GetUserDB(req.GetUsername())
	tk := ""
	if user.Password == req.GetPassword() {
		tk = token.GenerateToken(user.Username)
	}

	response := &LoginResponse{
		Token: tk,
	}
	fmt.Println("GRPC LOGIN")
	return response, nil
}

func (*Server) Register(ctx context.Context, req *LoginRequest) (*Error, error) {
	err := model.InsertUserDB(req.GetUsername(), req.GetPassword())

	response := &Error{}

	if err != nil {
		response.Error = err.Error()
	}

	return response, nil
}

func (*Server) InsertHand(ctx context.Context, req *InsertHandRequest) (*Empty, error) {

	kafka.KafkaWrite([]byte(req.GetData()), []byte(req.GetUsername()))

	response := &Empty{}

	return response, nil
}

func (*Server) GetHand(ctx context.Context, req *GetHandRequest) (*GetHandResponse, error) {
	fmt.Println((req.GetUsername()))
	results := model.GetHandRedis(req.GetNum(), req.GetGain(), req.GetSeat(), req.GetUsername())
	response := &GetHandResponse{}
	for _, result := range results {
		table := &GetHandResponse_Table{}
		table.Id = int32(result.ID)
		table.Time = timestamppb.New(result.Time)
		table.Player = make(map[string]*Player)
		table.Player[result.User.Username] = &Player{
			Name:  result.User.Username,
			Seat:  result.Seat.Seat,
			Gain:  result.Gain,
			Card1: &Card{Num: int32(result.HeroCard1.Num), Suit: result.HeroCard1.Suit},
			Card2: &Card{Num: int32(result.HeroCard2.Num), Suit: result.HeroCard2.Suit},
			Action: &Player_Action{
				Preflop: result.Preflop.Action,
				Flop:    result.Flop.Action,
				Turn:    result.Turn.Action,
				River:   result.River.Action,
			},
		}
		if result.TableCard1.ID != 0 {
			table.Card = append(table.Card, &Card{})
			table.Card[0].Num = int32(result.TableCard1.Num)
			table.Card[0].Suit = result.TableCard1.Suit
		}
		if result.TableCard2.ID != 0 {
			table.Card = append(table.Card, &Card{})
			table.Card[1].Num = int32(result.TableCard2.Num)
			table.Card[1].Suit = result.TableCard2.Suit
		}
		if result.TableCard3.ID != 0 {
			table.Card = append(table.Card, &Card{})
			table.Card[2].Num = int32(result.TableCard3.Num)
			table.Card[2].Suit = result.TableCard3.Suit
		}
		if result.TableCard4.ID != 0 {
			table.Card = append(table.Card, &Card{})
			table.Card[3].Num = int32(result.TableCard4.Num)
			table.Card[3].Suit = result.TableCard4.Suit
		}
		if result.TableCard5.ID != 0 {
			table.Card = append(table.Card, &Card{})
			table.Card[4].Num = int32(result.TableCard5.Num)
			table.Card[4].Suit = result.TableCard5.Suit
		}

		response.Table = append(response.Table, table)
	}

	return response, nil
}

func (*Server) GetOauthCode(ctx context.Context, req *Empty) (*GetOauthCodeResponse, error) {

	response := &GetOauthCodeResponse{
		Url: oauth.GenerateCodeURL(),
	}

	return response, nil
}

func (*Server) GetOauthToken(ctx context.Context, req *GetOauthTokenRequest) (*Empty, error) {

	code := req.GetCode()
	oauthToken := oauth.GenerateTokenURL(code)
	username := oauth.GetUser(oauthToken)
	tk := ""
	if username != "" {
		tk = token.GenerateToken(username)
		oauth.OAuthChan <- tk
	}
	response := &Empty{}

	return response, nil
}

func (*Server) CheckOauthToken(ctx context.Context, req *Empty) (*CheckOauthTokenResponse, error) {

	response := &CheckOauthTokenResponse{}

	if len(oauth.OAuthChan) > 0 {
		response.Result = <-oauth.OAuthChan
	}

	return response, nil
}

func (*Server) GetProfit(ctx context.Context, req *GetAnalysisRequest) (*GetProfitResponse, error) {
	response := &GetProfitResponse{}

	profits := model.GetProfitRedis(req.GetUsername(), req.GetPlayer())
	total := 0.0
	for count, profit := range profits {
		num, _ := strconv.ParseFloat(profit, 64)
		total, _ = decimal.NewFromFloat(total).Add(decimal.NewFromFloat(num)).Float64()
		response.Result = append(response.Result, &GetProfitResponse_Result{
			Hand: int32(count),
			Gain: total,
		})
	}

	return response, nil
}

func (*Server) GetPreflop(ctx context.Context, req *GetAnalysisRequest) (*GetLineResponse, error) {

	username := req.GetUsername()
	player := req.GetPlayer()

	response := &GetLineResponse{
		Raise: model.GetActionRedis("pre_flop", "R", username, player),
		Call:  model.GetActionRedis("pre_flop", "C", username, player),
		Fold:  model.GetActionRedis("pre_flop", "F", username, player),
		Check: model.GetActionRedis("pre_flop", "X", username, player),
		Bet:   model.GetActionRedis("pre_flop", "B", username, player),
	}

	return response, nil
}

func (*Server) GetFlop(ctx context.Context, req *GetAnalysisRequest) (*GetLineResponse, error) {
	username := req.GetUsername()
	player := req.GetPlayer()

	response := &GetLineResponse{
		Raise: model.GetActionRedis("flop", "R", username, player),
		Call:  model.GetActionRedis("flop", "C", username, player),
		Fold:  model.GetActionRedis("flop", "F", username, player),
		Check: model.GetActionRedis("flop", "X", username, player),
		Bet:   model.GetActionRedis("flop", "B", username, player),
	}

	return response, nil
}

func (*Server) GetTurn(ctx context.Context, req *GetAnalysisRequest) (*GetLineResponse, error) {
	username := req.GetUsername()
	player := req.GetPlayer()

	response := &GetLineResponse{
		Raise: model.GetActionRedis("turn", "R", username, player),
		Call:  model.GetActionRedis("turn", "C", username, player),
		Fold:  model.GetActionRedis("turn", "F", username, player),
		Check: model.GetActionRedis("turn", "X", username, player),
		Bet:   model.GetActionRedis("turn", "B", username, player),
	}

	return response, nil
}

func (*Server) GetRiver(ctx context.Context, req *GetAnalysisRequest) (*GetLineResponse, error) {
	username := req.GetUsername()
	player := req.GetPlayer()

	response := &GetLineResponse{
		Raise: model.GetActionRedis("river", "R", username, player),
		Call:  model.GetActionRedis("river", "C", username, player),
		Fold:  model.GetActionRedis("river", "F", username, player),
		Check: model.GetActionRedis("river", "X", username, player),
		Bet:   model.GetActionRedis("river", "B", username, player),
	}

	return response, nil
}

func (*Server) GetPlayer(ctx context.Context, req *GetPlayerRequest) (*GetPlayerResponse, error) {
	username := req.GetUsername()

	response := &GetPlayerResponse{
		Result: model.GetPlayerRedis(username),
	}

	return response, nil
}
