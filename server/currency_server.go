package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"strings"
	"time"

	currencypb "github.com/christianfoleide/grpcurrency/proto/currency"
	"google.golang.org/grpc"
)

var (
	defaultURI = "https://api.exchangeratesapi.io/latest?base="
)

func main() {

	lis, err := net.Listen("tcp", "localhost:8081")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	srv := &Server{}
	grpcServer := grpc.NewServer()

	currencypb.RegisterExchangeRateServiceServer(grpcServer, srv)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}

//Server ...
type Server struct{}

//APIResponse ...
type APIResponse struct {
	Rates map[string]interface{} `json: "rates"`
	Base  string                 `json: "base"`
	Date  string                 `json: "date"`
}

//GetRate ...
func (s *Server) GetRate(ctx context.Context, r *currencypb.RateRequest) (*currencypb.RateResponse, error) {

	base := strings.ToUpper(r.Base)
	target := strings.ToUpper(r.Target)

	req, err := http.NewRequest(http.MethodGet, fmt.Sprint(defaultURI+base), nil)
	if err != nil {
		log.Printf("Error: %v", err)
		return nil, err
	}

	client := &http.Client{
		Timeout: 20 * time.Second,
	}

	res, err := client.Do(req)
	if err != nil {
		log.Printf("Error: %v", err)
		return nil, err
	}

	defer res.Body.Close()

	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Printf("Error parsing the result: %v", err)
	}

	var result APIResponse
	json.Unmarshal(b, &result)

	return &currencypb.RateResponse{
		Rate: float32(result.Rates[target].(float64)),
	}, nil
}

//StreamRates received a RateRequest and returns a stream of RateResponse
func (s *Server) StreamRates(r *currencypb.RateRequest, stream currencypb.ExchangeRateService_StreamRatesServer) error {

	base := strings.ToUpper(r.Base)
	target := strings.ToUpper(r.Target)

	for {
		resp, err := s.GetRate(stream.Context(), &currencypb.RateRequest{
			Base:   base,
			Target: target,
		})
		if err != nil {
			return err
		}

		stream.Send(resp)
		time.Sleep(20 * time.Second)
	}

}
