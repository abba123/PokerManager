package proto

import (
	"context"
	"fmt"
	"log"
	"net"
	"poker/api/kafka"
	"poker/api/model"
	"poker/api/token"
	"poker/poker"

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

	res := &GetWinRateResponse{
		Result: result,
	}

	return res, nil
}
func (*Server) Login(ctx context.Context, req *LoginRequest) (*LoginResponse, error) {
	user := model.GetUserDB(req.GetUsername())
	tk := ""
	if user.Password == req.GetPassword() {
		tk = token.GenerateToken(user.Username)
	}

	res := &LoginResponse{
		Token: tk,
	}
	fmt.Println("GRPC LOGIN")
	return res, nil
}

func (*Server) Register(ctx context.Context, req *LoginRequest) (*Error, error) {
	err := model.InsertUserDB(req.GetUsername(), req.GetPassword())

	res := &Error{}

	if err != nil {
		res.Error = err.Error()
	}

	return res, nil
}

func (*Server) InsertHand(ctx context.Context, req *InsertHandRequest) (*Empty, error) {

	kafka.KafkaWrite([]byte(req.GetData()), []byte(req.GetUsername()))

	res := &Empty{}

	return res, nil
}

func (*Server) GetHand(ctx context.Context, req *GetHandRequest) (*GetHandResponse, error) {

	results := model.GetHandRedis(req.GetNum(), req.GetGain(), req.GetSeat(), req.GetUsername())
	res := &GetHandResponse{}
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

		res.Table = append(res.Table, table)
	}

	return res, nil
}
