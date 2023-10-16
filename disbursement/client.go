// File Created: Saturday, 14th October 2023 11:29:17 pm
// Author: Abdul Hamid (abdul.surel@gmail.com)
//
// Copyright (c) 2023 Author

package disbursement

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
	pathDisbursement = "/v4/disbursements"
	pathShow         = pathDisbursement + "/:id"
	pathUpdate       = pathDisbursement + "/:id/tasks"
)

// BankAccount return response from Create a Disbursement API
//
// Docs: https://docs.fazz.com/v4-ID/reference/create-disbursement
func (c *Client) Create(ctx context.Context, payload fazz.DisbursementCreatePayload) (*Disbursement, *fazz.Error) {
	res := struct {
		Data Disbursement `json:"data"`
	}{}

	if err := c.Api.Req(ctx, http.MethodPost, c.FazzURL+pathDisbursement, nil, payload, nil, &res); err != nil {
		return nil, err
	}

	return &res.Data, nil
}

// Disbursement return response from Get a Disbursement API
//
// Docs: https://docs.fazz.com/v4-ID/reference/get-disbursement
func (c *Client) Disbursement(ctx context.Context, disbursementId string) (*Disbursement, *fazz.Error) {
	url := strings.ReplaceAll(c.FazzURL+pathShow, ":id", disbursementId)
	res := struct {
		Data Disbursement `json:"data"`
	}{}

	if err := c.Api.Req(ctx, http.MethodGet, url, nil, nil, nil, &res); err != nil {
		return nil, err
	}

	return &res.Data, nil
}

// Disbursements return response from Get a List of Disbursement API.
// For pagination, sorting & filtering you can use FazzParams
//
// Docs: https://docs.fazz.com/v4-ID/reference/get-disbursements
func (c *Client) Disbursements(ctx context.Context, params *fazz.FazzParams) ([]Disbursement, *fazz.Error) {
	res := struct {
		Data []Disbursement `json:"data"`
	}{}

	if err := c.Api.Req(ctx, http.MethodGet, c.FazzURL+pathDisbursement, params, nil, nil, &res); err != nil {
		return []Disbursement{}, err
	}

	return res.Data, nil
}

// Update return response from Update a Disbursement API.
//
// Docs: https://docs.fazz.com/v4-ID/reference/update-disbursement
func (c *Client) Update(ctx context.Context, disbursementId string, payload fazz.DisbursementUpdatePayload) (*DisbursementUpdate, *fazz.Error) {
	url := strings.ReplaceAll(c.FazzURL+pathUpdate, ":id", disbursementId)
	res := struct {
		Data DisbursementUpdate `json:"data"`
	}{}

	if err := c.Api.Req(ctx, http.MethodPost, url, nil, payload, nil, &res); err != nil {
		return nil, err
	}

	return &res.Data, nil
}
