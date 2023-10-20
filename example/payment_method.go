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
