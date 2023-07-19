package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"time"

	payu "github.com/runnart/go-payu"
)

func main() {
	ht := &http.Client{
		Timeout: time.Second * 60,
	}
	client, err := payu.NewClient(os.Getenv("SERVER_URL"), payu.WithClient(ht))
	if err != nil {
		panic(err)
	}

	resp, err := client.TokensPost(context.TODO(), payu.OptTokensPostReq{
		Value: payu.TokensPostReq{
			TokenType:      payu.TokenTypeCreditCard,
			CreditCardCvv:  "123",
			CardNumber:     "4111111111111111",
			ExpirationDate: "10/29",
			HolderName:     "John Mark",
			BillingAddress: payu.BillingAddress{
				Country: "USA",
				State:   "NY",
				City:    "NYC",
				Line1:   "fifth avenue 10th",
			},
		},
		Set: true,
	}, payu.TokensPostParams{
		AppID: payu.OptString{
			Value: os.Getenv("APP_ID"),
		},
		PublicKey: payu.OptString{
			Value: os.Getenv("PUBLIC_KEY"),
		},
		XPaymentsOsEnv: payu.OptString{
			Value: os.Getenv("PAYMENT_ENV"),
		},
		APIVersion: payu.OptString{
			Value: "1.3.0",
		},
	})
	if err != nil {
		panic(err)
	}

	fmt.Println(resp)
}
