// File Created: Wednesday, 18th October 2023 12:53:48 am
// Author: Abdul Hamid (abdul.surel@gmail.com)
//
// Copyright (c) 2023 Author

package fazz

// ========== Payloads ==========

// PaymentLinkCreatePayload represent request payload for Create a Payment Link API.
type PaymentLinkCreatePayload struct {
	Amount               uint32             `json:"amount"`              // Required
	ReferenceId          string             `json:"referenceId"`         // Required
	CustomerName         string             `json:"customerName"`        // Required
	CustomerEmail        string             `json:"customerEmail"`       // Required
	CustomerPhoneNumber  string             `json:"customerPhoneNumber"` // Required
	Description          string             `json:"description"`
	ExpiredAt            string             `json:"expiredAt"` // ISO601
	PaymentMethodOptions PaymentLinkOptions `json:"paymentMethodOptions"`
}

// PaymentLinkOptions is part of PaymentLinkCreatePayload
type PaymentLinkOptions struct {
	DisplayName string `json:"displayName"`
}

// PaymetLinkUpdatePayload represent request payload for Update a Payment Link API.
type PaymentLinkUpdatePayload struct {
	// cancel, receive_payment, or settle. receive_payment & settle for sandbox purpose
	Action string `json:"action"`
}
