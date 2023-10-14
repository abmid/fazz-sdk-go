// File Created: Saturday, 14th October 2023 11:29:17 pm
// Author: Abdul Hamid (abdul.surel@gmail.com)
//
// Copyright (c) 2023 Author

package disbursement

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
	pathCreate = "/v4/disbursements"
)

// BankAccount return response from Create a Disbursement API
//
// Docs: https://docs.fazz.com/v4-ID/reference/create-disbursement
func (c *Client) Create(ctx context.Context, payload fazz.DisbursementCreatePayload) (*Disbursement, *fazz.Error) {
	res := struct {
		Data Disbursement `json:"data"`
	}{}

	if err := c.Api.Req(ctx, http.MethodPost, c.FazzURL+pathCreate, nil, payload, nil, &res); err != nil {
		return nil, err
	}

	return &res.Data, nil
}
