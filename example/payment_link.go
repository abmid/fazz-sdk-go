// File Created: Friday, 20th October 2023 5:53:04 pm
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

func PaymentLinkCreate() {
	payload := fazz.PaymentLinkCreatePayload{
		Amount:              15000,
		ReferenceId:         "SDK_TEST_01",
		CustomerName:        "Good Man",
		CustomerEmail:       "goodman@domain.com",
		CustomerPhoneNumber: "080000000000",
		Description:         "Test desc",
		ExpiredAt:           time.Now().Add(60 * time.Minute).Format(time.RFC3339),
		PaymentMethodOptions: fazz.PaymentLinkOptions{
			DisplayName: "Good Man",
		},
	}

	res, err := c.PaymentLink.Create(context.Background(), payload)
	if err != nil {
		// handle case error
		fmt.Println(err)
	}

	fmt.Println(res)
}

func PaymentLinkShow() {
	res, err := c.PaymentLink.PaymentLink(context.Background(), "paymentlink_123")
	if err != nil {
		// handle case error
	}

	fmt.Println(res)
}

func PaymentLinkUpdate() {
	payload := fazz.PaymentLinkUpdatePayload{
		Action: "receive_payment",
	}
	res, err := c.PaymentLink.Update(context.Background(), "paymentlink_123", payload)
	if err != nil {
		// handle case error
	}

	fmt.Println(res)
}
