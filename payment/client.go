// File Created: Thursday, 19th October 2023 6:17:14 pm
// Author: Abdul Hamid (abdul.surel@gmail.com)
//
// Copyright (c) 2023 Author

package payment

import (
	"context"
	"net/http"

	"github.com/abmid/fazz-sdk-go"
	"github.com/abmid/fazz-sdk-go/request"
)

type Client struct {
	Api     request.Api
	FazzURL string
}

const (
	pathPayment = "/v4/payments"
	pathShow    = "/v1/payments/:paymentId"
	pathUpdate  = "/v1/payments/:paymentId/tasks"
)

// CreateRetailOutlet return response from Create a Payment for Retail Outlet type.
//
// Docs: https://docs.fazz.com/v4-ID/reference/create-payment
func (c *Client) CreateRetailOutlet(ctx context.Context, payload fazz.PaymentCreateRetailPayload) (*PaymentCreate, *fazz.Error) {
	res := struct {
		Data PaymentCreate `json:""`
	}{}

	payload.PaymentMethodType = "retail_outlet"

	if err := c.Api.Req(ctx, http.MethodPost, c.FazzURL+pathPayment, nil, payload, nil, &res); err != nil {
		return nil, err
	}

	return &res.Data, nil
}

// CreateVA return response from Create a Payment for Virtual Account type.
//
// Docs: https://docs.fazz.com/v4-ID/reference/create-payment
func (c *Client) CreateVA(ctx context.Context, payload fazz.PaymentCreateVAPayload) (*PaymentCreate, *fazz.Error) {
	res := struct {
		Data PaymentCreate `json:""`
	}{}

	payload.PaymentMethodType = "virtual_bank_account"

	if err := c.Api.Req(ctx, http.MethodPost, c.FazzURL+pathPayment, nil, payload, nil, &res); err != nil {
		return nil, err
	}

	return &res.Data, nil
}
