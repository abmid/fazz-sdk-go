// File Created: Monday, 16th October 2023 3:35:45 pm
// Author: Abdul Hamid (abdul.surel@gmail.com)
//
// Copyright (c) 2023 Author

package fazz

// ========== Payloads ==========

// PaymentMethodCreateVAPayload represent request payload for Create Payment Method type Virtual Account API.
type PaymentMethodCreateVAPayload struct {
	ReferenceID   string `json:"referenceId"`
	BankShortCode string `json:"bankShortCode"`
	DisplayName   string `json:"displayName"`
	SuffixNo      string `json:"suffixNo"`
}

// PaymentMethodCreateQRISPayload represent request payload for Create Payment Method type QRIS API.
type PaymentMethodCreateQRISPayload struct {
	ReferenceID string `json:"referenceID"`
	DisplayName string `json:"displayName"`
}

// PaymentMethodSimulatePayload represent request payload for Create a mock payment API.
type PaymentMethodSimulatePayload struct {
	Action  string               `json:"action"`
	Options PaymentMethodOptions `json:"options"`
}

// PaymentMethodOptions is part of PaymentMethodSimulatePayload
type PaymentMethodOptions struct {
	Amount float32 `json:"amount"`
}
