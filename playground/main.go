package main

import (
	"context"
	"fmt"
	"github.com/imshakthi/lets-go/playground/account"
)

func main() {
	fmt.Println("test")
	getAccount(nil, request{})
}

// grpc message
type request struct {
	AccountID string
}

// grpc message
type response struct {
	Status  string
	Holding Holdings
}

// grpc message
type Holdings struct {
	Text1 string
	Text2 string
}

func getAccount(ctx context.Context, req request) response {
	//req.Validate()
	sReq := account.Request{AccountID: req.AccountID}


	sResp := getResponse(sReq)

	holdings := transform(sResp.Holdings)




	return response{
		Status:  "d",
		Holding: holdings,
	}
}

func transform(holdings account.Holdings) Holdings {

	return Holdings{
		Text1: holdings.Text1,
		Text2: holdings.Text2,
	}
}

func getResponse(req account.Request) account.Response {

	fmt.Println(req)

	return account.Response{}
}
