// File Created: Friday, 20th October 2023 5:26:14 pm
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

func ValidationBankAccount() {
	c := client.New(client.Options{ApiKey: "test_apikey", SecretKey: "secretkey"})

	payload := fazz.ValidateBankAccountPayload{
		AccountNo:     "000501003219303",
		BankShortCode: "BRI",
	}

	res, err := c.ValidationService.BankAccount(context.Background(), payload)
	if err != nil {
		// Handle case error
	}

	fmt.Println(res)
}
