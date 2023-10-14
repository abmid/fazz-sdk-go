// File Created: Friday, 13th October 2023 1:26:41 am
// Author: Abdul Hamid (abdul.surel@gmail.com)
//
// Copyright (c) 2023 Author

package validationservice

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
	pathValidateBankAccount = "/v4/validation_services/bank_account_validation"
)

// BankAccount return response from Validate a bank account API
//
// Docs: https://docs.fazz.com/v4-ID/reference/validate-bank-account
func (c *Client) BankAccount(ctx context.Context, payload fazz.ValidateBankAccountPayload) (*BankAccount, *fazz.Error) {
	res := struct {
		Data struct {
			Attributes BankAccount `json:"attributes"`
		} `json:"data"`
	}{}

	err := c.Api.Req(ctx, http.MethodPost, c.FazzURL+pathValidateBankAccount, nil, payload, nil, &res)
	if err != nil {
		return nil, err
	}

	return &res.Data.Attributes, nil
}
