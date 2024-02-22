// File Created: Friday, 23rd February 2024 12:41:15 am
// Author: Abdul Hamid (abdul.surel@gmail.com)
//
// Copyright (c) 2024 Author

package bank

// Bank is represent response from get a list of banks.
type Bank struct {
	ID         string         `json:"id"`
	Type       string         `json:"type"`
	Attributes BankAttributes `json:"attributes"`
}

// BankAttributes is part of Bank.
type BankAttributes struct {
	Name      string `json:"name"`
	ShortCode string `json:"shortCode"`
	SwiftBic  string `json:"swiftBic"`
}
