// File Created: Friday, 13th October 2023 12:44:29 am
// Author: Abdul Hamid (abdul.surel@gmail.com)
//
// Copyright (c) 2023 Author

package request

import (
	"context"
	"reflect"
	"testing"

	"github.com/abmid/fazz-sdk-go"
	"github.com/abmid/fazz-sdk-go/internal/helper"
	"github.com/jarcoal/httpmock"
)

func TestApiImplement_Req(t *testing.T) {
	type httpParam struct {
		AnotherParam string `url:"another_param"`
	}

	type httpResponse struct {
		AccountName   string `json:"accountName"`
		AccountNo     string `json:"accountNo"`
		BankShortCode string `json:"bankShortCode"`
	}

	type fields struct {
		ApiKey    string
		SecretKey string
	}

	type args struct {
		ctx      context.Context
		method   string
		url      string
		param    any
		body     any
		headers  map[string]string
		response any
	}

	tests := []struct {
		name    string
		fields  fields
		args    args
		prepare func(args args)
		wantRes any
		wantErr *fazz.Error
	}{
		{
			name: "Success http response 200",
			fields: fields{
				ApiKey:    "test_api_key",
				SecretKey: "secret_key",
			},
			args: args{
				ctx:    context.Background(),
				method: "POST",
				url:    fazz.SandboxURL,
				param: httpParam{
					AnotherParam: "another_param",
				},
				body: fazz.ValidateBankAccountPayload{
					AccountNo:     "000501003219303",
					BankShortCode: "BRI",
				},
				response: httpResponse{},
			},
			prepare: func(args args) {
				query := map[string]string{}
				query["another_param"] = "another_param"

				httpmock.RegisterMatcherResponderWithQuery(args.method, args.url, query, httpmock.Matcher{},
					helper.HttpMockResJSON(201, "../internal/test/validatebankaccount/res_201.json", args.headers))
			},
			wantRes: httpResponse{
				AccountName:   "PROD ONLY",
				AccountNo:     "000501003219303",
				BankShortCode: "BRI",
			},
			wantErr: nil,
		},
		{
			name: "Failed http response 4xx",
			fields: fields{
				ApiKey:    "test_api_key",
				SecretKey: "secret_key",
			},
			args: args{
				ctx:    context.Background(),
				method: "POST",
				url:    fazz.SandboxURL,
			},
			prepare: func(args args) {
				httpmock.RegisterResponder(args.method, args.url,
					helper.HttpMockResJSON(400, "../internal/test/validatebankaccount/res_400.json", args.headers))
			},
			wantRes: nil,
			wantErr: &fazz.Error{
				HttpStatusCode: 400,
				Errors: []fazz.Errors{
					{
						Code:   fazz.ErrorCode005,
						Title:  "Invalid parameter",
						Detail: "account_no must all be numbers",
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			httpmock.Activate()
			defer httpmock.DeactivateAndReset()

			c := &ApiImplement{
				ApiKey:    tt.fields.ApiKey,
				SecretKey: tt.fields.SecretKey,
			}

			tt.prepare(tt.args)

			res := struct {
				Data struct {
					Attributes httpResponse `json:"attributes"`
				} `json:"data"`
			}{}
			gotErr := c.Req(tt.args.ctx, tt.args.method, tt.args.url, tt.args.param, tt.args.body, tt.args.headers, &res)

			if tt.wantRes != nil {
				if !reflect.DeepEqual(res.Data.Attributes, tt.wantRes) {
					t.Errorf("ApiImplement.Req() gotRes = %v, want %v", res.Data.Attributes, tt.wantRes)
				}
			}

			if !reflect.DeepEqual(gotErr, tt.wantErr) {
				t.Errorf("ApiImplement.Req() gotErr = %v, want %v", gotErr, tt.wantErr)
			}
		})
	}
}
