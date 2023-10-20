// File Created: Friday, 20th October 2023 6:04:07 pm
// Author: Abdul Hamid (abdul.surel@gmail.com)
//
// Copyright (c) 2023 Author

package example

import (
	"context"
	"fmt"
	"time"

	"github.com/abmid/fazz-sdk-go"
)

func PaymentCreate() {
	// ========== Example Create VA ==========
	payload := fazz.PaymentCreateVAPayload{
		Payment: fazz.Payment{
			Amount:      15000,
			ReferenceId: "SDK_TEST_01",
			ExpiredAt:   time.Now().Add(70 * time.Minute).Format(time.RFC3339),
			Description: "SDK desc",
		},
		PaymentMethodOptions: fazz.PaymentVAOptions{
			BankShortCode: "BRI",
			DisplayName:   "GOOD Man",
			SuffixNo:      "123456",
		},
	}
	res, err := c.Payment.CreateVA(context.Background(), payload)
	if err != nil {
		// handle case error
		fmt.Println(err)
	}

	fmt.Println(res)

	// For e-wallet use method c.Payment.CreateEwallet()
	// For QRIS use method c.Payment.CreateQRIS()
	// For Retail outlet use method c.Payment.CreateRetailOutlet()
}

func PaymentShow() {
	res, err := c.Payment.Payment(context.Background(), "contract_123")
	if err != nil {
		// handle case error
		fmt.Println(err)
	}

	fmt.Println(res)
}

func PaymentGetListPayments() {
	params := fazz.FazzParams{}
	res, err := c.Payment.Payments(context.Background(), params)
	if err != nil {
		// handle case error
	}

	fmt.Println(res)
}

func PaymentUpdate() {
	payload := fazz.PaymentUpdatePayload{
		Action: "receive_payment",
	}
	res, err := c.Payment.Update(context.Background(), "contract_123", payload)
	if err != nil {
		// handle case error
	}

	fmt.Println(res)
}
