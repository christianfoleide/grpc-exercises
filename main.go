package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"sync"

	currencypb "github.com/christianfoleide/grpcurrency/proto/currency"
	"google.golang.org/grpc"
)

func main() {

	var wg sync.WaitGroup
	wg.Add(2) //for blocking
	go func() {
		http.HandleFunc("/currencies", handler)
		log.Fatal(http.ListenAndServe("localhost:8080", nil))
	}()

	go backGroundFetch()

	wg.Wait()
}

func createGRPCClient() *grpc.ClientConn {

	clientConn, err := grpc.Dial("localhost:8081", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not dial server: %v", err)
	}
	return clientConn
}

func backGroundFetch() {

	clientConn := createGRPCClient()
	defer clientConn.Close()

	client := currencypb.NewExchangeRateServiceClient(clientConn)

	req := &currencypb.RateRequest{
		Base:   "USD",
		Target: "NOK",
	}

	stream, err := client.StreamRates(context.Background(), req)
	if err != nil {
		log.Fatalf("Error calling StreamRates: %v", err)
	}

	for {
		msg, err := stream.Recv()
		if err == io.EOF {
			log.Printf("End of stream: %v", err)
			return
		}
		if err != nil {
			log.Printf("Error: %v", err)
			return
		}

		log.Printf("background fetch: BASE=%s TARGET=%s RATE=%v", req.Base, req.Target, msg.Rate)
	}

}

func handler(rw http.ResponseWriter, r *http.Request) {

	base := r.FormValue("base")
	target := r.FormValue("target")

	if base == "" || target == "" {
		http.Error(rw, "\nMissing query parameters", http.StatusBadRequest)
		return
	}

	clientConn := createGRPCClient()
	defer clientConn.Close()

	client := currencypb.NewExchangeRateServiceClient(clientConn)

	req := &currencypb.RateRequest{
		Base:   r.FormValue("base"),
		Target: r.FormValue("target"),
	}

	resp, err := client.GetRate(context.Background(), req)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusBadRequest)
	}

	fmt.Fprintf(rw, "\n[RESPONSE] BASE=%s TARGET=%s RATE=%v", base, target, resp.Rate)
}
