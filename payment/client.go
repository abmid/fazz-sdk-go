// File Created: Thursday, 19th October 2023 6:17:14 pm
// Author: Abdul Hamid (abdul.surel@gmail.com)
//
// Copyright (c) 2023 Author

package payment

import (
	"context"
	"net/http"
	"strings"

	"github.com/abmid/fazz-sdk-go"
	"github.com/abmid/fazz-sdk-go/request"
)

type Client struct {
	Api     request.Api
	FazzURL string
}

const (
	pathPayment = "/v4/payments"
	pathShow    = pathPayment + "/:paymentId"
	pathUpdate  = pathPayment + "/:paymentId/tasks"
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

// CreateQRIS return response from Create a Payment for QRIS type.
//
// Docs: https://docs.fazz.com/v4-ID/reference/create-payment
func (c *Client) CreateQRIS(ctx context.Context, payload fazz.PaymentCreateQRISPayload) (*PaymentCreate, *fazz.Error) {
	res := struct {
		Data PaymentCreate `json:""`
	}{}

	payload.PaymentMethodType = "qris"

	if err := c.Api.Req(ctx, http.MethodPost, c.FazzURL+pathPayment, nil, payload, nil, &res); err != nil {
		return nil, err
	}

	return &res.Data, nil
}

// CreateEwallet return response from Create a Payment for E-Wallet type.
//
// Docs: https://docs.fazz.com/v4-ID/reference/create-payment
func (c *Client) CreateEwallet(ctx context.Context, payload fazz.PaymentCreateEwalletPayload) (*PaymentCreate, *fazz.Error) {
	res := struct {
		Data PaymentCreate `json:""`
	}{}

	payload.PaymentMethodType = "e-wallet"

	if err := c.Api.Req(ctx, http.MethodPost, c.FazzURL+pathPayment, nil, payload, nil, &res); err != nil {
		return nil, err
	}

	return &res.Data, nil
}

// Payment return response from Get a Payment API.
//
// Docs: https://docs.fazz.com/v4-ID/reference/get-a-payment
func (c *Client) Payment(ctx context.Context, paymentId string) (*Payment, *fazz.Error) {
	url := strings.ReplaceAll(c.FazzURL+pathShow, ":paymentId", paymentId)
	res := struct {
		Data Payment `json:"data"`
	}{}

	if err := c.Api.Req(ctx, http.MethodGet, url, nil, nil, nil, &res); err != nil {
		return nil, err
	}

	return &res.Data, nil
}

// Payment return response from Get a List of Payments API.
//
// Docs: https://docs.fazz.com/v4-ID/reference/get-payments
func (c *Client) Payments(ctx context.Context, params fazz.FazzParams) ([]Payment, *fazz.Error) {
	res := struct {
		Data []Payment `json:"data"`
	}{}

	if err := c.Api.Req(ctx, http.MethodGet, c.FazzURL+pathPayment, params, nil, nil, &res); err != nil {
		return nil, err
	}

	return res.Data, nil
}

// Update return response from Update a Payment API.
//
// Docs: https://docs.fazz.com/v4-ID/reference/update-payment
func (c *Client) Update(ctx context.Context, paymentId string, payload fazz.PaymentUpdatePayload) (*PaymentUpdate, *fazz.Error) {
	url := strings.ReplaceAll(c.FazzURL+pathUpdate, ":paymentId", paymentId)
	res := struct {
		Data PaymentUpdate `json:"data"`
	}{}

	if err := c.Api.Req(ctx, http.MethodPost, url, nil, payload, nil, &res); err != nil {
		return nil, err
	}

	return &res.Data, nil
}
