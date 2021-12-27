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

	lis, err := net.Listen("tcp", "127.0.0.1:80")
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
	t := poker.Table{}

	p1 := poker.Player{Name: req.GetName1()}
	p2 := poker.Player{Name: req.GetName2()}

	p1.Card = []poker.Card{{}, {}}
	p1.Card[0].Num = int(req.GetP1Card1Num())
	p1.Card[0].Suit = req.GetP1Card1Suit()
	p1.Card[1].Num = int(req.GetP1Card2Num())
	p1.Card[1].Suit = req.GetP1Card2Suit()

	p2.Card = []poker.Card{{}, {}}
	p2.Card[0].Num = int(req.GetP2Card1Num())
	p2.Card[0].Suit = req.GetP2Card1Suit()
	p2.Card[1].Num = int(req.GetP2Card2Num())
	p2.Card[1].Suit = req.GetP2Card2Suit()

	result := poker.GetWinRate([]poker.Player{p1, p2}, 10000)

	res := &GetWinRateResponse{
		Result: result,
	}

	return res, nil
}
