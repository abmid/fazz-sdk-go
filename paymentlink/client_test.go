// File Created: Wednesday, 18th October 2023 1:15:28 am
// Author: Abdul Hamid (abdul.surel@gmail.com)
//
// Copyright (c) 2023 Author

package paymentlink

import (
	"context"
	"encoding/json"
	"net/http"
	"reflect"
	"strings"
	"testing"

	"github.com/abmid/fazz-sdk-go"
	"github.com/abmid/fazz-sdk-go/internal/helper"
	mock_request "github.com/abmid/fazz-sdk-go/internal/test/mock"
	"github.com/golang/mock/gomock"
)

const (
	pathTest        = "../internal/test/paymentlink/"
	pathTestInvalid = "../internal/test/invalid/"
)

func TestClient_Create(t *testing.T) {
	testWrap := helper.NewTestWrapper(t)
	defer testWrap.Ctrl.Finish()

	type args struct {
		ctx     context.Context
		payload fazz.PaymentLinkCreatePayload
	}
	tests := []struct {
		name    string
		args    args
		prepare func(m helper.Mocks, args args)
		wantRes *PaymentLink
		wantErr *fazz.Error
	}{
		{
			name: "Success",
			args: args{
				ctx: context.Background(),
				payload: fazz.PaymentLinkCreatePayload{
					Amount:              15000,
					ReferenceId:         "ref_123",
					CustomerName:        "John Doe",
					CustomerEmail:       "john.doe@gmail.com",
					CustomerPhoneNumber: "080000000000",
					Description:         "Order Number 0001",
					ExpiredAt:           "2023-12-06T16:00:00+07:00",
					PaymentMethodOptions: fazz.PaymentLinkOptions{
						DisplayName: "Nama Tampilan",
					},
				},
			},
			prepare: func(m helper.Mocks, args args) {
				m.Api.EXPECT().
					Req(args.ctx, http.MethodPost, fazz.SandboxURL+pathPaymentLink, nil, args.payload, nil, gomock.Any()).
					DoAndReturn(func(ctx context.Context, method string, url string, param any, body any, header map[string]string, response any) *fazz.Error {
						if err := json.Unmarshal(testWrap.ResJSONByte(pathTest+"res_create_201.json"), response); err != nil {
							panic(err)
						}

						return nil
					})
			},
			wantRes: &PaymentLink{
				ID:   "paymentlink_fca0bdbfa8cf663f50a53199e0a4c5c7",
				Type: "payment_links",
				Attributes: PaymentLinkAttributes{
					Status:              "pending",
					Amount:              "15000.0",
					ReferenceId:         "ref_123",
					CreatedAt:           "2023-10-18 00:57:45 +0700",
					Description:         "Order Number 0001",
					ExpiredAt:           "2023-12-06 16:00:00 +0700",
					PaymentLinkUrl:      "1",
					CustomerName:        "John Doe",
					CustomerEmail:       "john.doe@gmail.com",
					CustomerPhoneNumber: "080000000000",
					DisplayName:         "Nama Tampilan",
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
					Req(args.ctx, http.MethodPost, fazz.SandboxURL+pathPaymentLink, nil, args.payload, nil, gomock.Any()).
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
				payload := fazz.PaymentLinkCreatePayload{}
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

func TestClient_PaymentLink(t *testing.T) {
	testWrap := helper.NewTestWrapper(t)
	defer testWrap.Ctrl.Finish()

	type args struct {
		ctx           context.Context
		paymentLinkId string
	}
	tests := []struct {
		name    string
		args    args
		prepare func(m helper.Mocks, args args)
		wantRes *PaymentLink
		wantErr *fazz.Error
	}{
		{
			name: "Success",
			args: args{
				ctx:           context.Background(),
				paymentLinkId: "paymentlink_fca0bdbfa8cf663f50a53199e0a4c5c7",
			},
			prepare: func(m helper.Mocks, args args) {
				url := strings.ReplaceAll(fazz.SandboxURL+pathShow, ":paymentLinkId", args.paymentLinkId)
				m.Api.EXPECT().
					Req(args.ctx, http.MethodGet, url, nil, nil, nil, gomock.Any()).
					DoAndReturn(func(ctx context.Context, method string, url string, param any, body any, header map[string]string, response any) *fazz.Error {
						if err := json.Unmarshal(testWrap.ResJSONByte(pathTest+"res_payment_200.json"), response); err != nil {
							panic(err)
						}

						return nil
					})
			},
			wantRes: &PaymentLink{
				ID:   "paymentlink_fca0bdbfa8cf663f50a53199e0a4c5c7",
				Type: "payment_links",
				Attributes: PaymentLinkAttributes{
					Status:              "pending",
					Amount:              "15000.0",
					ReferenceId:         "ref_123",
					CreatedAt:           "2023-10-18 00:57:45 +0700",
					Description:         "Order Number 0001",
					ExpiredAt:           "2023-12-06 16:00:00 +0700",
					PaymentLinkUrl:      "1",
					CustomerName:        "John Doe",
					CustomerEmail:       "john.doe@gmail.com",
					CustomerPhoneNumber: "080000000000",
					DisplayName:         "Nama Tampilan",
				},
			},
		},
		{
			name: "Invalid requests",
			args: args{
				ctx: context.Background(),
			},
			prepare: func(m helper.Mocks, args args) {
				url := strings.ReplaceAll(fazz.SandboxURL+pathShow, ":paymentLinkId", args.paymentLinkId)
				m.Api.EXPECT().
					Req(args.ctx, http.MethodGet, url, nil, nil, nil, gomock.Any()).
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

			gotRes, gotErr := c.PaymentLink(tt.args.ctx, tt.args.paymentLinkId)
			if !reflect.DeepEqual(gotRes, tt.wantRes) {
				t.Errorf("Client.PaymentLink() gotRes = %v, wantRes %v", gotRes, tt.wantRes)
			}
			if !reflect.DeepEqual(gotErr, tt.wantErr) {
				t.Errorf("Client.PaymentLink() gotErr = %v, wantErr %v", gotErr, tt.wantErr)
			}
		})
	}
}
