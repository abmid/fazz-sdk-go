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

🚧 *The SDK is currently undergoing heavy development with frequent changes, because of this please get the latest update SDK* 🚧

[Fazz](https://fazz.com/) previously Xfers, is a company that provides financial services in Southeast Asia such as sending and receiving payments, grow their capital, and get funding.

For more information, visit the [Fazz Payments API Official documentation](https://docs.fazz.com/v4-ID/docs).

## Installation

Make sure you are using go version `1.18` or later

```bash
go get -u github.com/abmid/fazz-sdk-go
```

## Documentation


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
