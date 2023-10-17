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
	"strings"
	"testing"
	"time"

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

func TestClient_CreateQRIS(t *testing.T) {
	testWrap := helper.NewTestWrapper(t)
	defer testWrap.Ctrl.Finish()

	type args struct {
		ctx     context.Context
		payload fazz.PaymentMethodCreateQRISPayload
	}
	tests := []struct {
		name    string
		args    args
		prepare func(m helper.Mocks, args args)
		wantRes *PaymentMethodQRIS
		wantErr *fazz.Error
	}{
		{
			name: "Success",
			args: args{
				ctx: context.Background(),
				payload: fazz.PaymentMethodCreateQRISPayload{
					ReferenceID: "order_123",
					DisplayName: "Your preferred name",
				},
			},
			prepare: func(m helper.Mocks, args args) {
				m.Api.EXPECT().
					Req(args.ctx, http.MethodPost, fazz.SandboxURL+pathPaymentMethodQRIS, nil, args.payload, nil, gomock.Any()).
					DoAndReturn(func(ctx context.Context, method string, url string, param any, body any, header map[string]string, response any) *fazz.Error {
						if err := json.Unmarshal(testWrap.ResJSONByte(pathTest+"res_create_qris_201.json"), response); err != nil {
							panic(err)
						}

						return nil
					})
			},
			wantRes: &PaymentMethodQRIS{
				ID:   "qr_f0d07206381b2c69a52647",
				Type: "qris",
				Attributes: PaymentMethodQRISAttributes{
					ReferenceId: "order_123",
					Instructions: QRISInstructions{
						ImageURL:    "https://devel.bebasbayar.com/qrgen/payment/image/full/static?sc_id=170041&bill_number=8222",
						DisplayName: "Your preferred name",
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
					Req(args.ctx, http.MethodPost, fazz.SandboxURL+pathPaymentMethodQRIS, nil, args.payload, nil, gomock.Any()).
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
				payload := fazz.PaymentMethodCreateQRISPayload{}
				if !testWrap.DeepEqualPayload(pathTest+"payload_create_qris.json", &payload, &tt.args.payload) {
					t.Errorf("Client.CreateQRIS() gotPayload = %v, wantPayload %v", payload, tt.args.payload)
				}
			}

			gotRes, gotErr := c.CreateQRIS(tt.args.ctx, tt.args.payload)
			if !reflect.DeepEqual(gotRes, tt.wantRes) {
				t.Errorf("Client.CreateQRIS() gotRes = %v, wantRes %v", gotRes, tt.wantRes)
			}
			if !reflect.DeepEqual(gotErr, tt.wantErr) {
				t.Errorf("Client.CreateQRIS() gotErr = %v, wantErr %v", gotErr, tt.wantErr)
			}
		})
	}
}

func TestClient_PaymentMethodVA(t *testing.T) {
	testWrap := helper.NewTestWrapper(t)
	defer testWrap.Ctrl.Finish()

	type args struct {
		ctx             context.Context
		paymentMethodId string
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
				ctx:             context.Background(),
				paymentMethodId: "va_3148fe52bb6d17e4d90a6d0e55d08fb6",
			},
			prepare: func(m helper.Mocks, args args) {
				url := strings.ReplaceAll(fazz.SandboxURL+pathShowPaymentMethodVA, ":paymentMethodId", args.paymentMethodId)
				m.Api.EXPECT().
					Req(args.ctx, http.MethodGet, url, nil, nil, nil, gomock.Any()).
					DoAndReturn(func(ctx context.Context, method string, url string, param any, body any, header map[string]string, response any) *fazz.Error {
						if err := json.Unmarshal(testWrap.ResJSONByte(pathTest+"res_payment_method_va_200.json"), response); err != nil {
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
				url := strings.ReplaceAll(fazz.SandboxURL+pathShowPaymentMethodVA, ":paymentMethodId", args.paymentMethodId)
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

			gotRes, gotErr := c.PaymentMethodVA(tt.args.ctx, tt.args.paymentMethodId)
			if !reflect.DeepEqual(gotRes, tt.wantRes) {
				t.Errorf("Client.PaymentMethodVA() gotRes = %v, wantRes %v", gotRes, tt.wantRes)
			}
			if !reflect.DeepEqual(gotErr, tt.wantErr) {
				t.Errorf("Client.PaymentMethodVA() gotErr = %v, wantErr %v", gotErr, tt.wantErr)
			}
		})
	}
}

func TestClient_PaymentMethodQRIS(t *testing.T) {
	testWrap := helper.NewTestWrapper(t)
	defer testWrap.Ctrl.Finish()

	type args struct {
		ctx             context.Context
		paymentMethodId string
	}
	tests := []struct {
		name    string
		args    args
		prepare func(m helper.Mocks, args args)
		wantRes *PaymentMethodQRIS
		wantErr *fazz.Error
	}{
		{
			name: "Success",
			args: args{
				ctx:             context.Background(),
				paymentMethodId: "qr_f0d07206381b2c69a52647",
			},
			prepare: func(m helper.Mocks, args args) {
				url := strings.ReplaceAll(fazz.SandboxURL+pathShowPaymentMethodQRIS, ":paymentMethodId", args.paymentMethodId)
				m.Api.EXPECT().
					Req(args.ctx, http.MethodGet, url, nil, nil, nil, gomock.Any()).
					DoAndReturn(func(ctx context.Context, method string, url string, param any, body any, header map[string]string, response any) *fazz.Error {
						if err := json.Unmarshal(testWrap.ResJSONByte(pathTest+"res_payment_method_qris_200.json"), response); err != nil {
							panic(err)
						}

						return nil
					})
			},
			wantRes: &PaymentMethodQRIS{
				ID:   "qr_f0d07206381b2c69a52647",
				Type: "qris",
				Attributes: PaymentMethodQRISAttributes{
					ReferenceId: "order_123",
					Instructions: QRISInstructions{
						ImageURL:    "https://devel.bebasbayar.com/qrgen/payment/image/full/static?sc_id=170041&bill_number=8222",
						DisplayName: "Your preferred name",
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
				url := strings.ReplaceAll(fazz.SandboxURL+pathShowPaymentMethodQRIS, ":paymentMethodId", args.paymentMethodId)
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

			gotRes, gotErr := c.PaymentMethodQRIS(tt.args.ctx, tt.args.paymentMethodId)
			if !reflect.DeepEqual(gotRes, tt.wantRes) {
				t.Errorf("Client.PaymentMethodQRIS() gotRes = %v, wantRes %v", gotRes, tt.wantRes)
			}
			if !reflect.DeepEqual(gotErr, tt.wantErr) {
				t.Errorf("Client.PaymentMethodQRIS() gotErr = %v, want %v", gotErr, tt.wantErr)
			}
		})
	}
}

func TestClient_ListPaymentsVA(t *testing.T) {
	testWrap := helper.NewTestWrapper(t)
	defer testWrap.Ctrl.Finish()

	type args struct {
		ctx             context.Context
		paymentMethodId string
		params          fazz.FazzParams
	}
	tests := []struct {
		name    string
		args    args
		prepare func(m helper.Mocks, args args)
		wantRes []ListPaymentVA
		wantErr *fazz.Error
	}{
		{
			name: "Success",
			args: args{
				ctx:             context.Background(),
				paymentMethodId: "va_3148fe52bb6d17e4d90a6d0e55d08fb6",
				params: fazz.FazzParams{
					CreatedAfter: helper.ToPtr(time.Now().String()),
				},
			},
			prepare: func(m helper.Mocks, args args) {
				url := strings.ReplaceAll(fazz.SandboxURL+pathListPaymentsVA, ":paymentMethodId", args.paymentMethodId)
				m.Api.EXPECT().
					Req(args.ctx, http.MethodGet, url, args.params, nil, nil, gomock.Any()).
					DoAndReturn(func(ctx context.Context, method string, url string, param any, body any, header map[string]string, response any) *fazz.Error {
						if err := json.Unmarshal(testWrap.ResJSONByte(pathTest+"res_list_payments_va_200.json"), response); err != nil {
							panic(err)
						}

						return nil
					})
			},
			wantRes: []ListPaymentVA{
				{
					ID:   "contract_ba18fa4c8a0f4bdea9a6b582bb343cd3",
					Type: "payment",
					Attributes: ListPaymentVAAttributes{
						Status:      "completed",
						Amount:      "99000.2",
						CreatedAt:   helper.StringToTime("2023-10-16T15:52:15+07:00"),
						ReferenceId: "external_id_aae726a9b7",
						Fees:        "3663.0",
						PaymentMethod: ListPaymentMethodVA{
							ID:          "va_3148fe52bb6d17e4d90a6d0e55d08fb6",
							Type:        "virtual_bank_account",
							ReferenceId: "customer_id_123456",
							Instructions: VAInstructions{
								BankShortCode: "BRI",
								AccountNo:     "1269611512345678",
								DisplayName:   "IKN-My Company Name",
							},
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
				url := strings.ReplaceAll(fazz.SandboxURL+pathListPaymentsVA, ":paymentMethodId", args.paymentMethodId)
				m.Api.EXPECT().
					Req(args.ctx, http.MethodGet, url, args.params, nil, nil, gomock.Any()).
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

			gotRes, gotErr := c.ListPaymentsVA(tt.args.ctx, tt.args.paymentMethodId, tt.args.params)
			if !reflect.DeepEqual(gotRes, tt.wantRes) {
				t.Errorf("Client.ListPaymentsVA() gotRes = %v, wantRes %v", gotRes, tt.wantRes)
			}
			if !reflect.DeepEqual(gotErr, tt.wantErr) {
				t.Errorf("Client.ListPaymentsVA() gotErr = %v, wantErr %v", gotErr, tt.wantErr)
			}
		})
	}
}
