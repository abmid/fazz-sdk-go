// File Created: Wednesday, 18th October 2023 11:33:16 am
// Author: Abdul Hamid (abdul.surel@gmail.com)
//
// Copyright (c) 2023 Author

package fazz

// ========== Payloads ==========

type Payment struct {
	PaymentMethodType string `json:"paymentMethodType"` // Required
	Amount            uint32 `json:"amount"`            // Required
	ReferenceId       string `json:"referenceId"`       // Required
	ExpiredAt         string `json:"expiredAt,omitempty"`
	Description       string `json:"description,omitempty"`
}

// PaymentCreateRetailPayload represent for request payload Create a Payment (Retail Outlet)
type PaymentCreateRetailPayload struct {
	Payment
	PaymentMethodOptions PaymentRetailOptions `json:"paymentMethodOptions"`
}

// PaymentRetailOptions is part of PaymentCreateRetailPayload
type PaymentRetailOptions struct {
	RetailOutletName string `json:"retailOutletName"`
	DisplayName      string `json:"displayName"`
}

// PaymentCreateVAPayload represent for request payload Create a Payment (Virtual Account)
type PaymentCreateVAPayload struct {
	Payment
	PaymentMethodOptions PaymentVAOptions `json:"paymentMethodOptions"`
}

// PaymentVAOptions is part of PaymentCreateVAPayload
type PaymentVAOptions struct {
	BankShortCode string `json:"bankShortCode"`
	DisplayName   string `json:"displayName"`
	SuffixNo      string `json:"suffixNo,omitempty"`
}

// PaymentCreateQRISPayload represent for request payload Create a Payment (QRIS)
type PaymentCreateQRISPayload struct {
	Payment
	PaymentMethodOptions PaymentQRISOptions `json:"paymentMethodOptions"`
}

// PaymentQRISOptoiins is part of PaymentCreateQRISPayload
type PaymentQRISOptions struct {
	DisplayName string `json:"displayName"`
}

// PaymentCreateQRISPayload represent for request payload Create a Payment (E-Wallet)
type PaymentCreateEwalletPayload struct {
	Payment
	PaymentMethodOptions PaymentEwalletOptions `json:"paymentMethodOptions"`
}

// PaymentEwalletOptions is part of PaymentCreateEwalletPayload
type PaymentEwalletOptions struct {
	ProviderCode              string `json:"providerCode"`
	AfterSettlementReturnlUrl string `json:"afterSettlementReturnUrl"`
}

// PaymentUpdatePayload represent for request payload Update a Payment API.
type PaymentUpdatePayload struct {
	Action string `json:"action"`
}
