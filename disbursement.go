// File Created: Saturday, 14th October 2023 11:16:44 pm
// Author: Abdul Hamid (abdul.surel@gmail.com)
//
// Copyright (c) 2023 Author

package fazz

// ========== Payloads ==========

// DisbursementCreatePayload is payload for create disbursement API
type DisbursementCreatePayload struct {
	Amount             uint32             `json:"amount"`      // Required
	ReferenceID        string             `json:"referenceId"` // Required
	Description        string             `json:"description"`
	DisbursementMethod DisbursementMethod `json:"disbursementMethod"` // Required
}

// DisbursementMethod is part of payload create disbursement
type DisbursementMethod struct {
	Type                  string `json:"type"`          // Required
	BankShortCode         string `json:"bankShortCode"` // Required
	BankAccountNo         string `json:"bankAccountNo"` // Required
	BankAccountHolderName string `json:"bankAccountHolderName"`
}

// DisbursementUpdatePayload is payload for update disbursement API
type DisbursementUpdatePayload struct {
	Action string `json:"action"`
}
