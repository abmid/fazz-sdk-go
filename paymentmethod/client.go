// File Created: Tuesday, 17th October 2023 12:32:40 am
// Author: Abdul Hamid (abdul.surel@gmail.com)
//
// Copyright (c) 2023 Author

package paymentmethod

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
	pathPaymentMethod         = "/v4/payment_methods"
	pathPaymentMethodVA       = pathPaymentMethod + "/virtual_bank_accounts"
	pathPaymentMethodQRIS     = pathPaymentMethod + "/qris"
	pathShowPaymentMethodVA   = pathPaymentMethodVA + "/:paymentMethodId"
	pathShowPaymentMethodQRIS = pathPaymentMethodQRIS + "/:paymentMethodId"
	pathListPaymentsVA        = pathShowPaymentMethodVA + "/payments"
	pathListPaymentsQRIS      = pathShowPaymentMethodQRIS + "/payments"
	pathSimulatePaymentVA     = pathShowPaymentMethodVA + "/tasks"
	pathSimulatePaymentQRIS   = pathShowPaymentMethodQRIS + "/tasks"
)

// CreateVA return response from Create a Payment Method API for virtual_bank_accounts type.
//
// Docs: https://docs.fazz.com/v4-ID/reference/create-a-payment-method
func (c *Client) CreateVA(ctx context.Context, payload fazz.PaymentMethodCreateVAPayload) (*PaymentMethodVA, *fazz.Error) {
	res := struct {
		Data PaymentMethodVA `json:"data"`
	}{}

	if err := c.Api.Req(ctx, http.MethodPost, c.FazzURL+pathPaymentMethodVA, nil, payload, nil, &res); err != nil {
		return nil, err
	}

	return &res.Data, nil
}

// CreateQRIS return response from Create a Payment Method API for QRIS type.
//
// Docs: https://docs.fazz.com/v4-ID/reference/create-a-payment-method
func (c *Client) CreateQRIS(ctx context.Context, payload fazz.PaymentMethodCreateQRISPayload) (*PaymentMethodQRIS, *fazz.Error) {
	res := struct {
		Data PaymentMethodQRIS `json:"data"`
	}{}

	if err := c.Api.Req(ctx, http.MethodPost, c.FazzURL+pathPaymentMethodQRIS, nil, payload, nil, &res); err != nil {
		return nil, err
	}

	return &res.Data, nil
}

// PaymentMethodVA return response from Get a Payment Method API for virtual_bank_accounts type.
//
// Docs: https://docs.fazz.com/v4-ID/reference/get-payment-method
func (c *Client) PaymentMethodVA(ctx context.Context, paymentMethodId string) (*PaymentMethodVA, *fazz.Error) {
	url := strings.ReplaceAll(c.FazzURL+pathShowPaymentMethodVA, ":paymentMethodId", paymentMethodId)
	res := struct {
		Data PaymentMethodVA `json:"data"`
	}{}

	if err := c.Api.Req(ctx, http.MethodGet, url, nil, nil, nil, &res); err != nil {
		return nil, err
	}

	return &res.Data, nil
}

// PaymentMethodQRIS return response from Get a Payment Method API for QRIS type.
//
// Docs: https://docs.fazz.com/v4-ID/reference/get-payment-method
func (c *Client) PaymentMethodQRIS(ctx context.Context, paymentMethodId string) (*PaymentMethodQRIS, *fazz.Error) {
	url := strings.ReplaceAll(c.FazzURL+pathShowPaymentMethodQRIS, ":paymentMethodId", paymentMethodId)
	res := struct {
		Data PaymentMethodQRIS `json:"data"`
	}{}

	if err := c.Api.Req(ctx, http.MethodGet, url, nil, nil, nil, &res); err != nil {
		return nil, err
	}

	return &res.Data, nil
}

// ListPaymentsVA return response from Get a List of Payments for a Payment Method API (virtual_bank_accounts)
//
// Docs: https://docs.fazz.com/v4-ID/reference/get-payment-method-payments
func (c *Client) ListPaymentsVA(ctx context.Context, paymentMethodId string, params fazz.FazzParams) ([]ListPaymentVA, *fazz.Error) {
	url := strings.ReplaceAll(c.FazzURL+pathListPaymentsVA, ":paymentMethodId", paymentMethodId)
	res := struct {
		Data []ListPaymentVA `json:"data"`
	}{}

	if err := c.Api.Req(ctx, http.MethodGet, url, params, nil, nil, &res); err != nil {
		return nil, err
	}

	return res.Data, nil
}

// ListPaymentsQRIS return response from Get a List of Payments for a Payment Method API (QRIS)
//
// Docs: https://docs.fazz.com/v4-ID/reference/get-payment-method-payments
func (c *Client) ListPaymentsQRIS(ctx context.Context, paymentMethodId string, params fazz.FazzParams) ([]ListPaymentQRIS, *fazz.Error) {
	url := strings.ReplaceAll(c.FazzURL+pathListPaymentsQRIS, ":paymentMethodId", paymentMethodId)
	res := struct {
		Data []ListPaymentQRIS `json:"data"`
	}{}

	if err := c.Api.Req(ctx, http.MethodGet, url, params, nil, nil, &res); err != nil {
		return nil, err
	}

	return res.Data, nil
}
