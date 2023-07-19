# Go-Payu - GoLang Client for PayU Payment Gateway

![Go version](https://img.shields.io/github/go-mod/go-version/runnart/go-payu)
[![GoDoc](https://godoc.org/github.com/runnart/go-payu?status.svg)](https://pkg.go.dev/github.com/runnart/go-payu)
[![License](https://img.shields.io/badge/license-MIT-blue.svg)](https://github.com/runnart/go-payu/blob/main/LICENSE)

## Overview

Go-Payu is a feature-rich GoLang client library that provides seamless integration with the PayU payment gateway. It allows developers to easily process online payments, securely handle transactions, and streamline payment workflows within their Go applications. With Go-Payu, building robust and reliable payment solutions becomes a breeze.

## Features

- Simple and intuitive API for PayU integration
- Support for multiple payment methods, including credit/debit cards, online banking, and more
- Tokenization for secure and quick recurring payments
- Easy management of refunds and cancellations
- Query transaction status and history
- Customizable request and response handling
- Comprehensive documentation and examples

## Installation

To use Go-Payu in your Go project, you need to install it using `go get`:

```bash
go get github.com/runnart/go-payu
```

## Usage

1. Import the Go-Payu package into your Go code:

```go
import payu "github.com/runnart/go-payu"
```

2. Initialize the PayU client with your credentials:

```go
client, err := payu.NewClient(os.Getenv("SERVER_URL"))
if err != nil {
    panic(err)
}
```

3. Make requests and handle responses:

```go
// TokensPost creates a new token for credit card information on the PayU server.
// If there is an error, the code panics. In a production application, handle errors gracefully.
// The response (resp) contains the token information returned from the PayU server.
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

// Handle errors by panicking. In real-world scenarios, implement appropriate error handling.
if err != nil {
    panic(err)
}

```

For more detailed examples and API documentation, please refer to the [GoDoc](https://pkg.go.dev/github.com/runnart/go-payu) page.

## Contributing

We welcome contributions from the community! If you have suggestions, bug reports, or new features to add, please open an issue or submit a pull request.


## License

Go-Payu is licensed under the MIT License. See the [LICENSE](https://github.com/runnart/go-payu/blob/main/LICENSE) file for details.

