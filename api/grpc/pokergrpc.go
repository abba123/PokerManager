package grpc

import (
	"context"
	"fmt"
	"log"
	"net"
	"poker/poker"

	"google.golang.org/grpc"
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

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v \n", err)
	}
}

func (*Server) GetWinRate(ctx context.Context, req *GetWinRateRequest) (*GetWinRateResponse, error) {

	players := []poker.Player{}

	for key, value := range req.GetPlayers() {
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
