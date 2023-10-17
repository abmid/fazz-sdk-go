// File Created: Tuesday, 17th October 2023 12:32:40 am
// Author: Abdul Hamid (abdul.surel@gmail.com)
//
// Copyright (c) 2023 Author

package paymentmethod

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
