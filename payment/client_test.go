// File Created: Thursday, 19th October 2023 6:17:44 pm
// Author: Abdul Hamid (abdul.surel@gmail.com)
//
// Copyright (c) 2023 Author

package payment

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
	pathTest        = "../internal/test/payment/"
	pathTestInvalid = "../internal/test/invalid/"
)

func TestClient_CreateRetailOutlet(t *testing.T) {
	testWrap := helper.NewTestWrapper(t)
	defer testWrap.Ctrl.Finish()

	type args struct {
		ctx     context.Context
		payload fazz.PaymentCreateRetailPayload
	}
	tests := []struct {
		name    string
		args    args
		prepare func(m helper.Mocks, args args)
		wantRes *PaymentCreate
		wantErr *fazz.Error
	}{
		{
			name: "Success",
			args: args{
				ctx: context.Background(),
				payload: fazz.PaymentCreateRetailPayload{
					Payment: fazz.Payment{
						Amount:      15000,
						ReferenceId: "ORDER_0001",
						ExpiredAt:   "2023-10-19T01:07:04+07:00",
						Description: "Order Number 0001",
					},
					PaymentMethodOptions: fazz.PaymentRetailOptions{
						RetailOutletName: "ALFAMART",
						DisplayName:      "Nama Tampilan",
					},
				},
			},
			prepare: func(m helper.Mocks, args args) {
				args.payload.PaymentMethodType = "retail_outlet"

				m.Api.EXPECT().
					Req(args.ctx, http.MethodPost, fazz.SandboxURL+pathPayment, nil, args.payload, nil, gomock.Any()).
					DoAndReturn(func(ctx context.Context, method string, url string, param any, body any, header map[string]string, response any) *fazz.Error {
						if err := json.Unmarshal(testWrap.ResJSONByte(pathTest+"res_create_retail_store_200.json"), response); err != nil {
							panic(err)
						}

						return nil
					})
			},
			wantRes: &PaymentCreate{
				ID:   "contract_b31655f4827845c9809214bf0fc1f1ab",
				Type: "payment",
				Attributes: PaymentAttributes{
					Status:      "pending",
					Amount:      "15000.0",
					CreatedAt:   "2023-10-18T23:57:48+07:00",
					Description: "Order Number 0001",
					ExpiredAt:   "2023-10-19T01:07:04+07:00",
					ReferenceId: "ORDER_0001",
					Fees:        "4440.0",
					PaymentMethod: PaymentMethod{
						ID:          "mst_XXruOhVf8PjHeaw2SLsd",
						Type:        "retail_outlet",
						ReferenceId: "ORDER_0001",
						Instructions: Instructions{
							RetailOutletName: "ALFAMART",
							PaymentCode:      "81879760",
							DisplayName:      "Nama Tampilan",
						},
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
				args.payload.PaymentMethodType = "retail_outlet"

				m.Api.EXPECT().
					Req(args.ctx, http.MethodPost, fazz.SandboxURL+pathPayment, nil, args.payload, nil, gomock.Any()).
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
				payload := fazz.PaymentCreateRetailPayload{}
				tt.args.payload.PaymentMethodType = "retail_outlet"

				if !testWrap.DeepEqualPayload(pathTest+"payload_create_retail_store.json", &payload, &tt.args.payload) {
					t.Errorf("Client.CreateRetailOutlet() gotPayload = %v, wantPayload %v", payload, tt.args.payload)
				}
			}

			gotRes, gotErr := c.CreateRetailOutlet(tt.args.ctx, tt.args.payload)
			if !reflect.DeepEqual(gotRes, tt.wantRes) {
				t.Errorf("Client.CreateRetailOutlet() gotRes = %v, wantRes %v", gotRes, tt.wantRes)
			}

			if !reflect.DeepEqual(gotErr, tt.wantErr) {
				t.Errorf("Client.CreateRetailOutlet() gotErr = %v, wantRes %v", gotErr, tt.wantErr)
			}
		})
	}
}
