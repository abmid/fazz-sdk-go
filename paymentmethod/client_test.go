// File Created: Tuesday, 17th October 2023 12:32:47 am
// Author: Abdul Hamid (abdul.surel@gmail.com)
//
// Copyright (c) 2023 Author

package paymentmethod

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
	pathTest        = "../internal/test/paymentmethod/"
	pathTestInvalid = "../internal/test/invalid/"
)

func TestClient_CreateVA(t *testing.T) {
	testWrap := helper.NewTestWrapper(t)
	defer testWrap.Ctrl.Finish()

	type args struct {
		ctx     context.Context
		payload fazz.PaymentMethodCreateVAPayload
	}
	tests := []struct {
		name    string
		args    args
		prepare func(m helper.Mocks, args args)
		wantRes *PaymentMethodVA
		wantErr *fazz.Error
	}{
		{
			name: "Success",
			args: args{
				ctx: context.Background(),
				payload: fazz.PaymentMethodCreateVAPayload{
					ReferenceID:   "customer_id_123456",
					BankShortCode: "BRI",
					DisplayName:   "My Company Name",
					SuffixNo:      "12345678",
				},
			},
			prepare: func(m helper.Mocks, args args) {
				m.Api.EXPECT().
					Req(args.ctx, http.MethodPost, fazz.SandboxURL+pathPaymentMethodVA, nil, args.payload, nil, gomock.Any()).
					DoAndReturn(func(ctx context.Context, method string, url string, param any, body any, header map[string]string, response any) *fazz.Error {
						if err := json.Unmarshal(testWrap.ResJSONByte(pathTest+"res_create_va_201.json"), response); err != nil {
							panic(err)
						}

						return nil
					})
			},
			wantRes: &PaymentMethodVA{
				ID:   "va_3148fe52bb6d17e4d90a6d0e55d08fb6",
				Type: "virtual_bank_account",
				Attributes: PaymentMethodVAAttributes{
					ReferenceId: "customer_id_123456",
					Instructions: VAInstructions{
						BankShortCode: "BRI",
						AccountNo:     "1269611512345678",
						DisplayName:   "IKN-My Company Name",
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
					Req(args.ctx, http.MethodPost, fazz.SandboxURL+pathPaymentMethodVA, nil, args.payload, nil, gomock.Any()).
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

			if tt.wantRes != nil {
				payload := fazz.PaymentMethodCreateVAPayload{}
				if !testWrap.DeepEqualPayload(pathTest+"payload_create_va.json", &payload, &tt.args.payload) {
					t.Errorf("Client.Create() gotPayload = %v, wantPayload %v", payload, tt.args.payload)
				}
			}

			gotRes, gotErr := c.CreateVA(tt.args.ctx, tt.args.payload)
			if !reflect.DeepEqual(gotRes, tt.wantRes) {
				t.Errorf("Client.CreateVA() gotRes = %v, wantRes %v", gotRes, tt.wantRes)
			}
			if !reflect.DeepEqual(gotErr, tt.wantErr) {
				t.Errorf("Client.CreateVA() gotErr = %v, wantErr %v", gotErr, tt.wantErr)
			}
		})
	}
}
