// File Created: Friday, 20th October 2023 4:19:00 pm
// Author: Abdul Hamid (abdul.surel@gmail.com)
//
// Copyright (c) 2023 Author

package account

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
	pathAccountBalance = "/v4/overviews/balance_overview"
)

// Balance return response from Get a Account Balance API.
//
// Docs: https://docs.fazz.com/v4-ID/reference/get-account-balance
func (c *Client) Balance(ctx context.Context) (*Balance, *fazz.Error) {
	res := struct {
		Data struct {
			Attributes Balance `json:"attributes"`
		} `json:"data"`
	}{}

	if err := c.Api.Req(ctx, http.MethodGet, c.FazzURL+pathAccountBalance, nil, nil, nil, &res); err != nil {
		return nil, err
	}

	return &res.Data.Attributes, nil
}
