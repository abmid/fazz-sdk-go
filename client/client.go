// File Created: Thursday, 12th October 2023 11:59:50 pm
// Author: Abdul Hamid (abdul.surel@gmail.com)
//
// Copyright (c) 2023 Author

package client

import (
	"strings"

	"github.com/abmid/fazz-sdk-go"
	"github.com/abmid/fazz-sdk-go/validationservice"
)

type Client struct {
	Opts              Options
	FazzURL           string
	ValidationService *validationservice.Client
}

type Options struct {
	ApiKey    string
	SecretKey string
}

func (c *Client) Init() {

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

	return &c
}
