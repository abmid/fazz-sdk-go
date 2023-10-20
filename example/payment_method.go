// File Created: Friday, 20th October 2023 5:29:37 pm
// Author: Abdul Hamid (abdul.surel@gmail.com)
//
// Copyright (c) 2023 Author

package example

import (
	"context"
	"fmt"

	"github.com/abmid/fazz-sdk-go"
	"github.com/abmid/fazz-sdk-go/client"
)

func PaymentMethodCreate() {
	c := client.New(client.Options{ApiKey: "test_key", SecretKey: "secretKey"})

	// ========== EXAMPLE QRIS ==========
	createQRIS := fazz.PaymentMethodCreateQRISPayload{
		ReferenceID: "SDK_TEST_01",
		DisplayName: "SDK TEST",
	}
	resQRIS, err := c.PaymentMethod.CreateQRIS(context.Background(), createQRIS)
	if err != nil {
		// Handle case error
		fmt.Println(err)
	}
	fmt.Println(resQRIS)

	// ========== EXAMPLE Virtual Account ==========
	createVA := fazz.PaymentMethodCreateVAPayload{
		ReferenceID:   "SDK_TEST_02",
		BankShortCode: "BCA",
		DisplayName:   "SDK TEST",
		SuffixNo:      "123456",
	}
	resVA, err := c.PaymentMethod.CreateVA(context.Background(), createVA)
	if err != nil {
		// Handle case error
		fmt.Println(err)
	}

	fmt.Println(resVA)
}

func PaymentMethodShow() {
	c := client.New(client.Options{ApiKey: "test_key", SecretKey: "secretKey"})

	// ========== Example Show QRIS ==========
	res, err := c.PaymentMethod.PaymentMethodQRIS(context.Background(), "123")
	if err != nil {
		// Handle case error
	}

	fmt.Println(res)

	// ========== Example Show VA ==========
	resVA, err := c.PaymentMethod.PaymentMethodVA(context.Background(), "321")
	if err != nil {
		// Handle case error
	}

	fmt.Println(resVA)
}

func PaymentMethodGetListPayments() {
	c := client.New(client.Options{ApiKey: "test_key", SecretKey: "secretKey"})

	// ========== Example get list payments for virtual account type ==========
	params := fazz.FazzParams{
		PageSize: 20,
	}
	res, err := c.PaymentMethod.ListPaymentsVA(context.Background(), "123", params)
	if err != nil {
		// Handle case error
	}

	fmt.Println(res)

	// for QRIS use c.PaymentMethod.ListPaymentsQRIS()
}

func PaymentMethodSimulatePayment() {
	c := client.New(client.Options{ApiKey: "test_key", SecretKey: "secretKey"})

	// ========== Example simulate payment for QRIS ==========
	payload := fazz.PaymentMethodSimulatePayload{
		Action: "receive_payment",
	}
	res, err := c.PaymentMethod.SimulatePaymentQRIS(context.Background(), "123", payload)
	if err != nil {
		// Handle case error
	}

	fmt.Println(res)

	// for VA use c.PaymentMethod.SimulatePaymentVA()
}
