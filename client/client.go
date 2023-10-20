// File Created: Thursday, 12th October 2023 11:59:50 pm
// Author: Abdul Hamid (abdul.surel@gmail.com)
//
// Copyright (c) 2023 Author

package client

import (
	"strings"

	"github.com/abmid/fazz-sdk-go"
	"github.com/abmid/fazz-sdk-go/account"
	"github.com/abmid/fazz-sdk-go/disbursement"
	"github.com/abmid/fazz-sdk-go/payment"
	"github.com/abmid/fazz-sdk-go/paymentlink"
	"github.com/abmid/fazz-sdk-go/paymentmethod"
	"github.com/abmid/fazz-sdk-go/request"
	"github.com/abmid/fazz-sdk-go/validationservice"
)

type Client struct {
	Opts              Options
	FazzURL           string
	ValidationService *validationservice.Client
	Account           *account.Client
	Disbursement      *disbursement.Client
	Payment           *payment.Client
	PaymentLink       *paymentlink.Client
	PaymentMethod     *paymentmethod.Client
}

type Options struct {
	ApiKey    string
	SecretKey string
}

func (c *Client) Init() {
	requestApi := request.NewAPI(c.Opts.ApiKey, c.Opts.SecretKey)
	c.ValidationService = &validationservice.Client{Api: requestApi, FazzURL: c.FazzURL}
	c.Account = &account.Client{Api: requestApi, FazzURL: c.FazzURL}
	c.Disbursement = &disbursement.Client{Api: requestApi, FazzURL: c.FazzURL}
	c.Payment = &payment.Client{Api: requestApi, FazzURL: c.FazzURL}
	c.PaymentLink = &paymentlink.Client{Api: requestApi, FazzURL: c.FazzURL}
	c.PaymentMethod = &paymentmethod.Client{Api: requestApi, FazzURL: c.FazzURL}
}

func New(opts Options) *Client {
	var fazzURL string

	is_sandbox := strings.Contains(opts.ApiKey, "test_")
	if is_sandbox {
		fazzURL = fazz.SandboxURL
	} else {
		fazzURL = fazz.ProductionURL
	}

	c := Client{
		Opts:    opts,
		FazzURL: fazzURL,
	}

	c.Init()

	return &c
}
