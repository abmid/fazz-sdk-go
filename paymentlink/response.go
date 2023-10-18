// File Created: Wednesday, 18th October 2023 1:15:35 am
// Author: Abdul Hamid (abdul.surel@gmail.com)
//
// Copyright (c) 2023 Author

package paymentlink

// PaymentLink represent response from Create & Show Payment Link API.
type PaymentLink struct {
	ID         string                `json:"id"`
	Type       string                `json:"type"`
	Attributes PaymentLinkAttributes `json:"attributes"`
}

// PaymentLinkAttributes is part of PaymentLink.
type PaymentLinkAttributes struct {
	Status              string `json:"status"`
	Amount              string `json:"amount"`
	ReferenceId         string `json:"referenceId"`
	CreatedAt           string `json:"createdAt"`
	Description         string `json:"description"`
	ExpiredAt           string `json:"expiredAt"`
	PaymentLinkUrl      string `json:"paymentLinkUrl"`
	CustomerName        string `json:"customerName"`
	CustomerEmail       string `json:"customerEmail"`
	CustomerPhoneNumber string `json:"customerPhoneNumber"`
	DisplayName         string `json:"displayName"`
}

// PaymentLinkUpdate represent response from Update Payment Link API.
type PaymentLinkUpdate struct {
	Type       string                      `json:"type"`
	Attributes PaymentLinkUpdateAttributes `json:"attributes"`
}

// PaymentLinkUpdateAttributes is part of PaymentLinkUpdate.
type PaymentLinkUpdateAttributes struct {
	TargetId   string `json:"targetId"`
	TargetType string `json:"targetType"`
	Action     string `json:"action"`
}
