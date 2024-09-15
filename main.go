package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"log"
	"math/rand"
	"net"

	pb "hoshimachi/proto"

	"google.golang.org/grpc"
)

const (
	sui_port string = ":22318"
)

type SuiQuote struct {
	Quote  string `json:"quote"`
	Type   uint32 `json:"type"`
	Source string `json:"source"`
}

type HoshimachiQuoteGenServer struct {
	pb.UnimplementedHoshimachiQuoteGenServer
}

func HoshimachiQuoteFromJSON(path string, arr *[]SuiQuote) error {
	json_data, json_err := os.ReadFile(path)
	if json_err != nil {
		return fmt.Errorf("failed to read file: %v", json_err)
	}

	parse_err := json.Unmarshal(json_data, &arr)
	if parse_err != nil {
		return fmt.Errorf("failed to parse: %v", parse_err)
	}

	return nil
}

var suisei_quotes []SuiQuote

func (s *HoshimachiQuoteGenServer) GetRandomQuote(ctx context.Context, req *pb.SuiQuoteReq) (*pb.SuiQuoteResp, error) {
	rand_index := rand.Intn(len(suisei_quotes))

	selected_quote := suisei_quotes[rand_index]

	return &pb.SuiQuoteResp{
		Quote:  selected_quote.Quote,
		Type:   selected_quote.Type,
		Source: selected_quote.Source,
	}, nil
}

func main() {
	HoshimachiQuoteFromJSON("./quotes/Sui_EN.json", &suisei_quotes)

	lis, lis_err := net.Listen("tcp", sui_port)
	if lis_err != nil {
		log.Fatal("❌🌠 Hoshimachi failed to listen:", lis_err)
	}

	sui_grpc := grpc.NewServer()

	pb.RegisterHoshimachiQuoteGenServer(sui_grpc, &HoshimachiQuoteGenServer{})

	fmt.Println("🌠 Hoshimachi's running on port", sui_port)

	if serve_err := sui_grpc.Serve(lis); serve_err != nil {
		log.Fatal("❌🌠 Hoshimachi's failed to serve:", serve_err)
	}
}
