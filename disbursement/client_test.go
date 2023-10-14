// File Created: Saturday, 14th October 2023 11:29:25 pm
// Author: Abdul Hamid (abdul.surel@gmail.com)
//
// Copyright (c) 2023 Author

package disbursement

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
	pathTest        = "../internal/test/disbursement/"
	pathTestInvalid = "../internal/test/invalid/"
)

func TestClient_Create(t *testing.T) {
	testWrap := helper.NewTestWrapper(t)
	defer testWrap.Ctrl.Finish()

	type args struct {
		ctx     context.Context
		payload fazz.DisbursementCreatePayload
	}
	tests := []struct {
		name    string
		args    args
		prepare func(m helper.Mocks, args args)
		wantRes *Disbursement
		wantErr *fazz.Error
	}{
		{
			name: "Success",
			args: args{
				ctx: context.Background(),
				payload: fazz.DisbursementCreatePayload{
					Amount:      10000,
					ReferenceID: "order_id_123456",
					Description: "Your delivery payout.",
					DisbursementMethod: fazz.DisbursementMethod{
						Type:                  "bank_transfer",
						BankShortCode:         "BRI",
						BankAccountNo:         "0102030405",
						BankAccountHolderName: "John Doe",
					},
				},
			},
			prepare: func(m helper.Mocks, args args) {
				m.Api.EXPECT().
					Req(args.ctx, http.MethodPost, fazz.SandboxURL+pathCreate, nil, args.payload, nil, gomock.Any()).
					DoAndReturn(func(ctx context.Context, method string, url string, param any, body any, header map[string]string, response any) *fazz.Error {
						if err := json.Unmarshal(testWrap.ResJSONByte(pathTest+"res_create_201.json"), response); err != nil {
							panic(err)
						}

						return nil
					})
			},
			wantRes: &Disbursement{
				ID:   "contract_1a2b3c4d5e6f7890",
				Type: "disbursement",
				Attributes: DisbursementAttributes{
					Amount:      "10000.0",
					ReferenceID: "order_id_123456",
					Description: "Your delivery payout",
					Status:      "processing",
					CreatedAt:   helper.StringToTime("2020-03-27T23:59:59+07:00"),
					Fees:        "200.0",
					DisbursementMethod: DisbursementMethod{
						Type:                        "bank_transfer",
						BankAccountNo:               "0102030405",
						BankShortCode:               "BRI",
						BankName:                    "Bank Rakyat Indonesia",
						BankAccountHolderName:       "John Doe",
						ServerBankAccountHolderName: "J Doe",
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
					Req(args.ctx, http.MethodPost, fazz.SandboxURL+pathCreate, nil, args.payload, nil, gomock.Any()).
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
				payload := fazz.DisbursementCreatePayload{}
				if !testWrap.DeepEqualPayload(pathTest+"payload_create.json", &payload, &tt.args.payload) {
					t.Errorf("Client.Create() gotPayload = %v, wantPayload %v", payload, tt.args.payload)
				}
			}

			gotRes, gotErr := c.Create(tt.args.ctx, tt.args.payload)
			if !reflect.DeepEqual(gotRes, tt.wantRes) {
				t.Errorf("Client.Create() gotRes = %v, wantRes %v", gotRes, tt.wantRes)
			}
			if !reflect.DeepEqual(gotErr, tt.wantErr) {
				t.Errorf("Client.Create() gotErr = %v, wantErr %v", gotErr, tt.wantErr)
			}
		})
	}
}
