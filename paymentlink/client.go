// File Created: Wednesday, 18th October 2023 1:15:20 am
// Author: Abdul Hamid (abdul.surel@gmail.com)
//
// Copyright (c) 2023 Author

package paymentlink

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
	pathPaymentLink = "/v4/payment_links"
	pathShow        = pathPaymentLink + "/:paymentLinkId"
	pathUpdate      = pathShow + "/tasks"
)

// Create return response from Create a Payment Link API.
//
// Docs: https://docs.fazz.com/v4-ID/reference/create-payment-link
func (c *Client) Create(ctx context.Context, payload fazz.PaymentLinkCreatePayload) (*PaymentLink, *fazz.Error) {
	res := struct {
		Data PaymentLink `json:"data"`
	}{}

	if err := c.Api.Req(ctx, http.MethodPost, c.FazzURL+pathPaymentLink, nil, payload, nil, &res); err != nil {
		return nil, err
	}

	return &res.Data, nil
}
