// File Created: Friday, 13th October 2023 8:33:28 pm
// Author: Abdul Hamid (abdul.surel@gmail.com)
//
// Copyright (c) 2023 Author

package client

import (
	"testing"

	"github.com/abmid/fazz-sdk-go"
	"github.com/abmid/fazz-sdk-go/validationservice"
)

func TestNew(t *testing.T) {
	type args struct {
		opts Options
	}
	tests := []struct {
		name string
		args args
		want *Client
	}{
		{
			name: "Sandbox",
			args: args{
				opts: Options{
					ApiKey:    "test_key123",
					SecretKey: "123abc",
				},
			},
			want: &Client{
				Opts: Options{
					ApiKey:    "test_key123",
					SecretKey: "123abc",
				},
				FazzURL:           fazz.SandboxURL,
				ValidationService: &validationservice.Client{},
			},
		},
		{
			name: "Production",
			args: args{
				opts: Options{
					ApiKey:    "prod_key123",
					SecretKey: "123abc",
				},
			},
			want: &Client{
				Opts: Options{
					ApiKey:    "prod_key123",
					SecretKey: "123abc",
				},
				FazzURL: fazz.ProductionURL,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(tt.args.opts); got.FazzURL != tt.want.FazzURL {
				t.Errorf("New() = %v, want %v", got.FazzURL, tt.want.FazzURL)
			}
		})
	}
}
