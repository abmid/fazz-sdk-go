# DurianPay SDK for Go #

[![GoDoc](https://godoc.org/github.com/abmid/fazz-sdk-go?status.svg)](https://godoc.org/github.com/abmid/fazz-sdk-go)
![Test Status](https://github.com/abmid/fazz-sdk-go/actions/workflows/test.yml/badge.svg)
[![Go Report Card](https://goreportcard.com/badge/github.com/abmid/fazz-sdk-go)](https://goreportcard.com/report/github.com/abmid/fazz-sdk-go)
[![License: MIT](https://img.shields.io/badge/license-MIT-green.svg)](https://opensource.org/licenses/MIT)

## Table of Contents

- [Overview](#overview)
- [Installation](#installation)
- [Documentation](#documentation)
- [API Supports](#api-supports)
- [Contributing](#contributing)
- [License](#license)

## Overview

🚧 *The SDK is currently undergoing heavy development with frequent changes, because of this please get the latest update of the SDK* 🚧

> Compatibility: The SDK has been developing use documentation Indonesia Payment API with API version 4.

[Fazz](https://fazz.com/) previously Xfers, is a company that provides financial services in Southeast Asia such as sending and receiving payments, grow their capital, and get funding.

Fazz provides SDKs in several programming languages but not in Go. Because of this the SDK was created.

For more information, visit the [Fazz Payments API Official documentation](https://docs.fazz.com/v4-ID/docs).

## Installation

Make sure you are using go version `1.18` or later

```bash
go get -u github.com/abmid/fazz-sdk-go
```

## Documentation

```go
package main

import (
	"context"
	"fmt"

	"github.com/abmid/fazz-sdk-go"
	"github.com/abmid/fazz-sdk-go/client"
)

func main() {
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

```

**Example Create Payment for Virtual Account**

```go
	// ========== Example Create VA ==========
	payload := fazz.PaymentCreateVAPayload{
		Payment: fazz.Payment{
			Amount:      15000,
			ReferenceId: "SDK_TEST_01",
			ExpiredAt:   time.Now().Add(70 * time.Minute).Format(time.RFC3339),
			Description: "SDK desc",
		},
		PaymentMethodOptions: fazz.PaymentVAOptions{
			BankShortCode: "BCA",
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
```



For more examples, please check directory [example](https://github.com/abmid/fazz-sdk-go/tree/master/example) and [Godoc](https://godoc.org/github.com/abmid/fazz-sdk-go)

## API Supports

- PAYMENTS
  - [x] Create a payment
  - [x] Get a payment
  - [x] Get a list of payments
  - [x] Update a payment
- PAYMENT METHODS
  - [x] Create a payment method
  - [x] Get a payment method
  - [x] Get a list of payments for a payment method
  - [x] Create a mock payment for a payment method
- PAYMENT LINKS
  - [x] Create a payment link
  - [x] Get a payment link
  - [x] Update a payment link
- DISBURSEMENTS
  - [x] Create a disbursement
  - [x] Get a disbursement
  - [x] Get a list of disbursements
  - [x] Update a disbursement
  - [ ] Get a list of banks
- ACCOUNTS
  - [x] Get account balance
- VALIDATIONS
  - [x] validate a bank account

## Contributing

We are open  and grateful for any contribution. If you want to contribute please do PR and follow the code guide.

## License

Copyright (c) 2023-present [Abdul Hamid](https://github.com/abmid) and [Contributors](https://github.com/abmid/fazz-sdk-go/graphs/contributors). This SDK is free and open-source software licensed under the [MIT License](https://github.com/abmid/fazz-sdk-go/tree/master/LICENSE).
