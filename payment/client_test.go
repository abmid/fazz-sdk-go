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
						if err := json.Unmarshal(testWrap.ResJSONByte(pathTest+"res_create_retail_store_201.json"), response); err != nil {
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

func TestClient_CreateVA(t *testing.T) {
	testWrap := helper.NewTestWrapper(t)
	defer testWrap.Ctrl.Finish()

	type args struct {
		ctx     context.Context
		payload fazz.PaymentCreateVAPayload
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
				payload: fazz.PaymentCreateVAPayload{
					Payment: fazz.Payment{
						Amount:      15000,
						ReferenceId: "ORDER_0002",
						ExpiredAt:   "2023-10-19T01:07:04+07:00",
						Description: "Order Number 0001",
					},
					PaymentMethodOptions: fazz.PaymentVAOptions{
						BankShortCode: "BCA",
						DisplayName:   "test name",
						SuffixNo:      "12345678",
					},
				},
			},
			prepare: func(m helper.Mocks, args args) {
				args.payload.PaymentMethodType = "virtual_bank_account"

				m.Api.EXPECT().
					Req(args.ctx, http.MethodPost, fazz.SandboxURL+pathPayment, nil, args.payload, nil, gomock.Any()).
					DoAndReturn(func(ctx context.Context, method string, url string, param any, body any, header map[string]string, response any) *fazz.Error {
						if err := json.Unmarshal(testWrap.ResJSONByte(pathTest+"res_create_va_201.json"), response); err != nil {
							panic(err)
						}

						return nil
					})
			},
			wantRes: &PaymentCreate{
				ID:   "contract_26f319c1a162448d830c08c6de1f8619",
				Type: "payment",
				Attributes: PaymentAttributes{
					Status:      "pending",
					Amount:      "15000.0",
					CreatedAt:   "2023-10-19T00:00:10+07:00",
					Description: "Order Number 0001",
					ExpiredAt:   "2023-10-19T01:07:04+07:00",
					ReferenceId: "ORDER_0002",
					Fees:        "3663.0",
					PaymentMethod: PaymentMethod{
						ID:          "va_2b257c73adf6b0598fe24301ac85061a",
						Type:        "virtual_bank_account",
						ReferenceId: "ORDER_0002",
						Instructions: Instructions{
							BankShortCode: "BCA",
							AccountNo:     "1304431712345678",
							DisplayName:   "IKN-test name",
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
				args.payload.PaymentMethodType = "virtual_bank_account"

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
				payload := fazz.PaymentCreateVAPayload{}
				tt.args.payload.PaymentMethodType = "virtual_bank_account"

				if !testWrap.DeepEqualPayload(pathTest+"payload_create_va.json", &payload, &tt.args.payload) {
					t.Errorf("Client.CreateVA() gotPayload = %v, wantPayload %v", payload, tt.args.payload)
				}
			}

			gotRes, gotErr := c.CreateVA(tt.args.ctx, tt.args.payload)
			if !reflect.DeepEqual(gotRes, tt.wantRes) {
				t.Errorf("Client.CreateVA() gotRes = %v, wantRes %v", gotRes, tt.wantRes)
			}

			if !reflect.DeepEqual(gotErr, tt.wantErr) {
				t.Errorf("Client.CreateVA() gotErr = %v, wantRes %v", gotErr, tt.wantErr)
			}
		})
	}
}

func TestClient_CreateQRIS(t *testing.T) {
	testWrap := helper.NewTestWrapper(t)
	defer testWrap.Ctrl.Finish()

	type args struct {
		ctx     context.Context
		payload fazz.PaymentCreateQRISPayload
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
				payload: fazz.PaymentCreateQRISPayload{
					Payment: fazz.Payment{
						Amount:      15000,
						ReferenceId: "ORDER_0003",
						ExpiredAt:   "2023-10-19T01:07:04+07:00",
						Description: "Order Number 0001",
					},
					PaymentMethodOptions: fazz.PaymentQRISOptions{
						DisplayName: "Your preferred name",
					},
				},
			},
			prepare: func(m helper.Mocks, args args) {
				args.payload.PaymentMethodType = "qris"

				m.Api.EXPECT().
					Req(args.ctx, http.MethodPost, fazz.SandboxURL+pathPayment, nil, args.payload, nil, gomock.Any()).
					DoAndReturn(func(ctx context.Context, method string, url string, param any, body any, header map[string]string, response any) *fazz.Error {
						if err := json.Unmarshal(testWrap.ResJSONByte(pathTest+"res_create_qris_201.json"), response); err != nil {
							panic(err)
						}

						return nil
					})
			},
			wantRes: &PaymentCreate{
				ID:   "contract_0ed566e83f864e7ba6086033dfe3791c",
				Type: "payment",
				Attributes: PaymentAttributes{
					Status:      "pending",
					Amount:      "15000.0",
					CreatedAt:   "2023-10-19T00:01:01+07:00",
					Description: "Order Number 0001",
					ExpiredAt:   "2023-10-19T01:07:04+07:00",
					ReferenceId: "ORDER_0003",
					Fees:        "116.55",
					PaymentMethod: PaymentMethod{
						ID:          "qr_acbe2c3cd932fd41f1caab",
						Type:        "qris",
						ReferenceId: "ORDER_0003",
						Instructions: Instructions{
							ImageUrl:    "https://devel.bebasbayar.com/qrgen/payment/image/full/static?sc_id=170041&bill_number=8222",
							DisplayName: "Your preferred name",
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
				args.payload.PaymentMethodType = "qris"

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
				payload := fazz.PaymentCreateQRISPayload{}
				tt.args.payload.PaymentMethodType = "qris"

				if !testWrap.DeepEqualPayload(pathTest+"payload_create_qris.json", &payload, &tt.args.payload) {
					t.Errorf("Client.CreateQRIS() gotPayload = %v, wantPayload %v", payload, tt.args.payload)
				}
			}

			gotRes, gotErr := c.CreateQRIS(tt.args.ctx, tt.args.payload)
			if !reflect.DeepEqual(gotRes, tt.wantRes) {
				t.Errorf("Client.CreateQRIS() gotRes = %v, wantRes %v", gotRes, tt.wantRes)
			}

			if !reflect.DeepEqual(gotErr, tt.wantErr) {
				t.Errorf("Client.CreateQRIS() gotErr = %v, wantRes %v", gotErr, tt.wantErr)
			}
		})
	}
}

func TestClient_CreateEwallet(t *testing.T) {
	testWrap := helper.NewTestWrapper(t)
	defer testWrap.Ctrl.Finish()

	type args struct {
		ctx     context.Context
		payload fazz.PaymentCreateEwalletPayload
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
				payload: fazz.PaymentCreateEwalletPayload{
					Payment: fazz.Payment{
						Amount:      15000,
						ReferenceId: "ORDER_0004",
						ExpiredAt:   "2023-10-19T00:20:04+07:00",
						Description: "Order Number 0001",
					},
					PaymentMethodOptions: fazz.PaymentEwalletOptions{
						ProviderCode:              "SHOPEEPAY",
						AfterSettlementReturnlUrl: "https://pay.examplessee.co.id/return-pay-here?0340450",
					},
				},
			},
			prepare: func(m helper.Mocks, args args) {
				args.payload.PaymentMethodType = "e-wallet"

				m.Api.EXPECT().
					Req(args.ctx, http.MethodPost, fazz.SandboxURL+pathPayment, nil, args.payload, nil, gomock.Any()).
					DoAndReturn(func(ctx context.Context, method string, url string, param any, body any, header map[string]string, response any) *fazz.Error {
						if err := json.Unmarshal(testWrap.ResJSONByte(pathTest+"res_create_ewallet_201.json"), response); err != nil {
							panic(err)
						}

						return nil
					})
			},
			wantRes: &PaymentCreate{
				ID:   "contract_b174839e7ddf4d8eb826df7bb3d08a55",
				Type: "payment",
				Attributes: PaymentAttributes{
					Status:      "pending",
					Amount:      "15000.0",
					CreatedAt:   "2023-10-19T00:02:32+07:00",
					Description: "Order Number 0001",
					ExpiredAt:   "2023-10-19T00:20:04+07:00",
					ReferenceId: "ORDER_0004",
					Fees:        "249.75",
					PaymentMethod: PaymentMethod{
						ID:          "shopeepay_6fcf2d462f3837dcbe1fe1f82e92e4a78f6cbb8e",
						Type:        "e-wallet",
						ReferenceId: "ORDER_0004",
						Settlement: Settement{
							HttpUrl:            "https://pay.uat.airpay.co.id/h5pay/pay?type=start&medium_index=dFhkbmR1bTBIamhW0VqULhun6fTSt7wAUTLjlGQIEcUvLx3kpjdBTcaEg3NqoJZxo1dRWshDlOl_-VNZA6TNWcFz9UC89QsV9VNdZisX&order_key=QvcorKPvdeluA8iQWsWba93giTvjQeGASbsTpbd3ofib15cwunSk1_K2zp_kApjEoD2uvhTOU_IIdQ&source=web&token=dFhkbmR1bTBIamhW0VqULhun6fTSt7wAUTLjlGQIEcUvLx3kpjdBTcaEg3NqoJZxo1dRWshDlOl_-VNZA6TNWcFz9UC89QsV9VNdZisX",
							AfterSettlementUrl: "https://pay.examplessee.co.id/return-pay-here?0340450",
							MobileUrl:          "shopeeid://main?apprl=%2Frn%2FTRANSFER_PAGE%3Fnavigate_url%3Dhttps%253A%252F%252Fwsa.uat.wallet.airpay.co.id%252Fwallet%252Fpay%253Fmedium_index%253DdFhkbmR1bTBIamhW0VqULhun6fTSt7wAUTLjlGQIEcUvLx3kpjdBTcaEg3NqoJZxo1dRWshDlOl_-VNZA6TNWcFz9UC89QsV9VNdZisX%2526order_key%253DQvcorKPvdeluA8iQWsWba93giTvjQeGASbsTpbd3ofib15cwunSk1_K2zp_kApjEoD2uvhTOU_IIdQ%2526source%253Dqr%2526token%253DdFhkbmR1bTBIamhW0VqULhun6fTSt7wAUTLjlGQIEcUvLx3kpjdBTcaEg3NqoJZxo1dRWshDlOl_-VNZA6TNWcFz9UC89QsV9VNdZisX",
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
				args.payload.PaymentMethodType = "e-wallet"

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
				payload := fazz.PaymentCreateEwalletPayload{}
				tt.args.payload.PaymentMethodType = "e-wallet"

				if !testWrap.DeepEqualPayload(pathTest+"payload_create_ewallet.json", &payload, &tt.args.payload) {
					t.Errorf("Client.CreateEwallet() gotPayload = %v, wantPayload %v", payload, tt.args.payload)
				}
			}

			gotRes, gotErr := c.CreateEwallet(tt.args.ctx, tt.args.payload)
			if !reflect.DeepEqual(gotRes, tt.wantRes) {
				t.Errorf("Client.CreateEwallet() gotRes = %v, wantRes %v", gotRes, tt.wantRes)
			}

			if !reflect.DeepEqual(gotErr, tt.wantErr) {
				t.Errorf("Client.CreateEwallet() gotErr = %v, wantRes %v", gotErr, tt.wantErr)
			}
		})
	}
}
