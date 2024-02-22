// File Created: Friday, 23rd February 2024 12:48:00 am
// Author: Abdul Hamid (abdul.surel@gmail.com)
//
// Copyright (c) 2024 Author

package bank

import (
	"context"
	"encoding/json"
	"net/http"
	"reflect"
	"testing"

	"github.com/abmid/fazz-sdk-go"
	"github.com/abmid/fazz-sdk-go/internal/helper"
	mock_request "github.com/abmid/fazz-sdk-go/internal/test/mock"
	"github.com/golang/mock/gomock"
)

const (
	pathTest        = "../internal/test/bank/"
	pathTestInvalid = "../internal/test/invalid/"
)

func TestClient_Banks(t *testing.T) {
	testWrap := helper.NewTestWrapper(t)
	defer testWrap.Ctrl.Finish()

	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		args    args
		prepare func(m helper.Mocks, args args)
		wantRes []Bank
		wantErr *fazz.Error
	}{
		{
			name: "Success",
			args: args{
				ctx: context.Background(),
			},
			prepare: func(m helper.Mocks, args args) {
				m.Api.EXPECT().
					Req(args.ctx, http.MethodGet, fazz.SandboxURL+pathBank, nil, nil, nil, gomock.Any()).
					DoAndReturn(func(ctx context.Context, method string, url string, param any, body any, header map[string]string, response any) *fazz.Error {
						if err := json.Unmarshal(testWrap.ResJSONByte(pathTest+"res_banks_200.json"), response); err != nil {
							panic(err)
						}

						return nil
					})
			},
			wantRes: []Bank{
				{
					ID:   "bca",
					Type: "bank",
					Attributes: BankAttributes{
						Name:      "Bank Central Asia",
						ShortCode: "BCA",
						SwiftBic:  "CENAIDJA",
					},
				},
				{
					ID:   "mandiri",
					Type: "bank",
					Attributes: BankAttributes{
						Name:      "Bank Mandiri",
						ShortCode: "MANDIRI",
						SwiftBic:  "BMRIIDJA",
					},
				},
			},
		},
		{
			name: "Invalid requests",
			args: args{
				ctx: context.Background(),
			},
			prepare: func(m helper.Mocks, args args) {
				m.Api.EXPECT().
					Req(args.ctx, http.MethodGet, fazz.SandboxURL+pathBank, nil, nil, nil, gomock.Any()).
					Return(fazz.ErrFromAPI(400, testWrap.ResJSONByte(pathTestInvalid+"res_400.json")))
			},
			wantErr: fazz.ErrFromAPI(400, testWrap.ResJSONByte(pathTestInvalid+"res_400.json")),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			apiMock := mock_request.NewMockApi(testWrap.Ctrl)

			c := &Client{
				Api:     apiMock,
				FazzURL: fazz.SandboxURL,
			}

			tt.prepare(helper.Mocks{Api: apiMock}, tt.args)

			gotRes, gotErr := c.Banks(tt.args.ctx)
			if !reflect.DeepEqual(gotRes, tt.wantRes) {
				t.Errorf("Client.Banks() gotRes = %v, wantRes %v", gotRes, tt.wantRes)
			}
			if !reflect.DeepEqual(gotErr, tt.wantErr) {
				t.Errorf("Client.Banks() gotErr = %v, wantErr %v", gotErr, tt.wantErr)
			}
		})
	}
}
