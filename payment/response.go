// File Created: Thursday, 19th October 2023 10:46:53 am
// Author: Abdul Hamid (abdul.surel@gmail.com)
//
// Copyright (c) 2023 Author

package payment

// Payment represent response from List of Payments & Get a Payment API.
type Payment struct {
	ID        string            `json:"id"`
	Type      string            `json:"type"`
	Attribute PaymentAttributes `json:"attributes"`
}

// PAymentAttributes is part of Payment
type PaymentAttributes struct {
	Status        string        `json:"status"`
	Amount        string        `json:"amount"`
	CreatedAt     string        `json:"createdAt"`
	Description   string        `json:"description"`
	ExpiredAt     string        `json:"expiredAt"`
	ReferenceId   string        `json:"referenceId"`
	Fees          string        `json:"fees"`
	PaymentMethod PaymentMethod `json:"paymentMethod"`
}

// PaymentMethod is part of PaymentAttributes
type PaymentMethod struct {
	ID           string       `json:"id"`
	Type         string       `json:"type"`
	ReferenceId  string       `json:"referenceId"`
	Instructions Instructions `json:"instructions"` // Wrap QRIS, VA and Retail Outlet in one instructions
	Settlement   Settement    `json:"settlement"`   // Wrap EWallet in one Settlement
}

// Instructions is part of PaymentMethod
type Instructions struct {
	ImageUrl         string `json:"imageUrl"`         // QRIS
	DisplayName      string `json:"displayName"`      // QRIS, VA, Retail Outlet
	BankShortCode    string `json:"bankShortCode"`    // VA
	AccountNo        string `json:"accountNo"`        // VA
	RetailOutletName string `json:"retailOutletName"` // Retail Outlet
	PaymentCode      string `json:"paymentCode"`      // Retail Outlet
}

// Settlement is part of Payment Method
type Settement struct {
	HttpUrl            string `json:"httpUrl"`            // E-Wallet
	AfterSettlementUrl string `json:"afterSettlementUrl"` // E-Wallet
	MobileUrl          string `json:"mobileUrl"`          // E-Wallet
}

// PaymentCreate represent response from Create a Payment API.
// Attribute instructions for type QRIS, VA and Retail Outlet is wrapped in one struct and also Settlement from E-Wallet.
type PaymentCreate struct {
	ID         string            `json:"id"`
	Type       string            `json:"type"`
	Attributes PaymentAttributes `json:"attributes"`
}
