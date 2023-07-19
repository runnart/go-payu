// Code generated by ogen, DO NOT EDIT.

package go_payu

import (
	"bytes"
	"net/http"

	"github.com/go-faster/jx"

	ht "github.com/ogen-go/ogen/http"
)

func encodeCustomersCustomerIdPaymentMethodsTokenPostRequest(
	req *CustomersCustomerIdPaymentMethodsTokenPostReq,
	r *http.Request,
) error {
	const contentType = "application/json"
	e := jx.GetEncoder()
	{
		if req != nil {
			req.Encode(e)
		}
	}
	encoded := e.Bytes()
	ht.SetBody(r, bytes.NewReader(encoded), contentType)
	return nil
}

func encodeCustomersPostRequest(
	req *CustomersPostReq,
	r *http.Request,
) error {
	const contentType = "application/json"
	e := jx.GetEncoder()
	{
		if req != nil {
			req.Encode(e)
		}
	}
	encoded := e.Bytes()
	ht.SetBody(r, bytes.NewReader(encoded), contentType)
	return nil
}

func encodePaymentsPaymentidAuthorizationsPostRequest(
	req OptPaymentsPaymentidAuthorizationsPostReq,
	r *http.Request,
) error {
	const contentType = "application/json"
	if !req.Set {
		// Keep request with empty body if value is not set.
		return nil
	}
	e := jx.GetEncoder()
	{
		if req.Set {
			req.Encode(e)
		}
	}
	encoded := e.Bytes()
	ht.SetBody(r, bytes.NewReader(encoded), contentType)
	return nil
}

func encodePaymentsPaymentidCapturesPostRequest(
	req *PaymentsPaymentidCapturesPostReq,
	r *http.Request,
) error {
	const contentType = "application/json"
	e := jx.GetEncoder()
	{
		if req != nil {
			req.Encode(e)
		}
	}
	encoded := e.Bytes()
	ht.SetBody(r, bytes.NewReader(encoded), contentType)
	return nil
}

func encodePaymentsPaymentidChargesPostRequest(
	req *PaymentsPaymentidChargesPostReq,
	r *http.Request,
) error {
	const contentType = "application/json"
	e := jx.GetEncoder()
	{
		if req != nil {
			req.Encode(e)
		}
	}
	encoded := e.Bytes()
	ht.SetBody(r, bytes.NewReader(encoded), contentType)
	return nil
}

func encodePaymentsPaymentidCreditsPostRequest(
	req *PaymentsPaymentidCreditsPostReq,
	r *http.Request,
) error {
	const contentType = "application/json"
	e := jx.GetEncoder()
	{
		if req != nil {
			req.Encode(e)
		}
	}
	encoded := e.Bytes()
	ht.SetBody(r, bytes.NewReader(encoded), contentType)
	return nil
}

func encodePaymentsPaymentidRefundsPostRequest(
	req *PaymentsPaymentidRefundsPostReq,
	r *http.Request,
) error {
	const contentType = "application/json"
	e := jx.GetEncoder()
	{
		if req != nil {
			req.Encode(e)
		}
	}
	encoded := e.Bytes()
	ht.SetBody(r, bytes.NewReader(encoded), contentType)
	return nil
}

func encodePaymentsPaymentidVoidsPostRequest(
	req *PaymentsPaymentidVoidsPostReq,
	r *http.Request,
) error {
	const contentType = "application/json"
	e := jx.GetEncoder()
	{
		if req != nil {
			req.Encode(e)
		}
	}
	encoded := e.Bytes()
	ht.SetBody(r, bytes.NewReader(encoded), contentType)
	return nil
}

func encodePaymentsPostRequest(
	req *PaymentsPostReq,
	r *http.Request,
) error {
	const contentType = "application/json"
	e := jx.GetEncoder()
	{
		if req != nil {
			req.Encode(e)
		}
	}
	encoded := e.Bytes()
	ht.SetBody(r, bytes.NewReader(encoded), contentType)
	return nil
}

func encodeTokensPostRequest(
	req OptTokensPostReq,
	r *http.Request,
) error {
	const contentType = "application/json"
	if !req.Set {
		// Keep request with empty body if value is not set.
		return nil
	}
	e := jx.GetEncoder()
	{
		if req.Set {
			req.Encode(e)
		}
	}
	encoded := e.Bytes()
	ht.SetBody(r, bytes.NewReader(encoded), contentType)
	return nil
}
