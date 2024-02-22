// File Created: Friday, 23rd February 2024 12:44:51 am
// Author: Abdul Hamid (abdul.surel@gmail.com)
//
// Copyright (c) 2024 Author

package bank

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
	pathBank = "/v4/banks"
)

// Banks return response from Get a list of banks API.
//
// Docs: https://docs.fazz.com/v4-ID/reference/get-banks
func (c *Client) Banks(ctx context.Context) ([]Bank, *fazz.Error) {
	res := struct {
		Data []Bank `json:"data"`
	}{}

	if err := c.Api.Req(ctx, http.MethodGet, c.FazzURL+pathBank, nil, nil, nil, &res); err != nil {
		return nil, err
	}

	return res.Data, nil
}
