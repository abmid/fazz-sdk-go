// File Created: Monday, 16th October 2023 11:46:02 pm
// Author: Abdul Hamid (abdul.surel@gmail.com)
//
// Copyright (c) 2023 Author

package paymentmethod

import "time"

// PaymentMethodVA represents response from Create a payment method & Get a payment method API for type Virtual Account
type PaymentMethodVA struct {
	ID         string                    `json:"id"`
	Type       string                    `json:"type"`
	Attributes PaymentMethodVAAttributes `json:"attributes"`
}

// PaymentMethodVAAttributes is part of PaymentMethodVA
type PaymentMethodVAAttributes struct {
	ReferenceId  string         `json:"referenceId"`
	Instructions VAInstructions `json:"instructions"`
}

// VAInstructions is part of PaymentMethodVA & ListPaymentMethodVA
type VAInstructions struct {
	BankShortCode string `json:"bankShortCode"`
	AccountNo     string `json:"accountNo"`
	DisplayName   string `json:"displayName"`
}

// PaymentMethodQRIS represents response from Create a payment method & Get a payment method API for type QRIS
type PaymentMethodQRIS struct {
	ID         string                      `json:"id"`
	Type       string                      `json:"type"`
	Attributes PaymentMethodQRISAttributes `json:"attributes"`
}

// PaymentMethodQRISAttributes is part of PaymentMethodQRIS
type PaymentMethodQRISAttributes struct {
	ReferenceId  string           `json:"referenceId"`
	Instructions QRISInstructions `json:"instructions"`
}

// QRISInstructions is part of PaymentMethodQRIS & ListPaymentMethodQRIS
type QRISInstructions struct {
	ImageURL    string `json:"imageUrl"`
	DisplayName string `json:"displayName"`
}

// ListPaymentVA represents response from Get a list of payments for a payment method API on type Virtual Account
type ListPaymentVA struct {
	ID         string                  `json:"id"`
	Type       string                  `json:"type"`
	Attributes ListPaymentVAAttributes `json:"attributes"`
}

// ListPaymentVAAttributes is part of ListPaymentVA
type ListPaymentVAAttributes struct {
	Status        string              `json:"status"`
	Amount        string              `json:"amount"`
	CreatedAt     time.Time           `json:"createdAt"`
	Description   string              `json:"description"`
	ExpiredAt     time.Time           `json:"expiredAt"`
	ReferenceId   string              `json:"referenceId"`
	Fees          string              `json:"fees"`
	PaymentMethod ListPaymentMethodVA `json:"paymentMethod"`
}

// ListPaymentMethodVA is part of ListPaymentVAAttributes
type ListPaymentMethodVA struct {
	ID           string         `json:"id"`
	Type         string         `json:"type"`
	ReferenceId  string         `json:"referenceId"`
	Instructions VAInstructions `json:"instructions"`
}

// ListPaymentVA represents response from Get a list of payments for a payment method API on type QRIS
type ListPaymentQRIS struct {
	ID         string                    `json:"id"`
	Type       string                    `json:"type"`
	Attributes ListPaymentQRISAttributes `json:"attributes"`
}

// ListPaymentQRISAttributes is part of ListPaymentQRIS
type ListPaymentQRISAttributes struct {
	Status        string                `json:"status"`
	Amount        string                `json:"amount"`
	CreatedAt     time.Time             `json:"createdAt"`
	Description   string                `json:"description"`
	ExpiredAt     time.Time             `json:"expiredAt"`
	ReferenceId   string                `json:"referenceId"`
	Fees          string                `json:"fees"`
	PaymentMethod ListPaymentMethodQRIS `json:"paymentMethod"`
}

// ListPaymentMethodQRIS is part of ListPaymentQRISAttributes
type ListPaymentMethodQRIS struct {
	ID           string           `json:"id"`
	Type         string           `json:"type"`
	ReferenceId  string           `json:"referenceId"`
	Instructions QRISInstructions `json:"instructions"`
}

// SimulatePaymentVA represents response from Create a mock payment for a payment method Virtual Account
type SimulatePaymentVA struct {
	Type       string                      `json:"type"`
	Attributes SimulatePaymentVAAttributes `json:"attributes"`
}

// SimulatePaymentVAAttributes is part of SimulatePaymentVA
type SimulatePaymentVAAttributes struct {
	TargetId    string                 `json:"targetId"`
	TargetType  string                 `json:"targetType"`
	ReferenceId string                 `json:"referenceId"`
	Action      string                 `json:"action"`
	Options     SimulatePaymentOptions `json:"options"`
}

// SimulatePaymentOptions is part of SimulatePaymentVAAttributes & SimulatePaymentQRISAttributes
type SimulatePaymentOptions struct {
	Amount string `json:"amount"`
}

// SimulatePaymentVA represents response from Create a mock payment for a payment method QRIS
type SimulatePaymentQRIS struct {
	Type       string                        `json:"type"`
	Attributes SimulatePaymentQRISAttributes `json:"attributes"`
}

// SimulatePaymentQRISAttributes is part of SimulatePaymentQRIS
type SimulatePaymentQRISAttributes struct {
	TargetId   string                 `json:"targetId"`
	TargetType string                 `json:"targetType"`
	Action     string                 `json:"action"`
	Options    SimulatePaymentOptions `json:"options"`
}
