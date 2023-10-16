// File Created: Saturday, 14th October 2023 11:29:33 pm
// Author: Abdul Hamid (abdul.surel@gmail.com)
//
// Copyright (c) 2023 Author

package disbursement

import "time"

// Disbursement represents response from Create a Disbursement API
type Disbursement struct {
	ID         string                 `json:"id"`
	Type       string                 `json:"type"`
	Attributes DisbursementAttributes `json:"attributes"`
}

// DisbursementAttributes is part of Disbursement
type DisbursementAttributes struct {
	Amount             string             `json:"amount"`
	ReferenceID        string             `json:"referenceId"`
	Description        string             `json:"description"`
	Status             string             `json:"status"`
	CreatedAt          time.Time          `json:"createdAt"`
	Fees               string             `json:"fees"`
	DisbursementMethod DisbursementMethod `json:"disbursementMethod"`
}

// DisbursementMethod is part of DisbursementAttributes
type DisbursementMethod struct {
	Type                        string `json:"type"`
	BankAccountNo               string `json:"bankAccountNo"`
	BankShortCode               string `json:"bankShortCode"`
	BankName                    string `json:"bankName"`
	BankAccountHolderName       string `json:"bankAccountHolderName"`
	BankReferenceNumber         string `json:"bankReferenceNumber"`
	ServerBankAccountHolderName string `json:"serverBankAccountHolderName"`
	FailureReason               string `json:"failureReason"`
}

// DisbursementUpdate represent response from Update a Disbursement API.
type DisbursementUpdate struct {
	Type       string                       `json:"type"`
	Attributes DisbursementUpdateAttributes `json:"attributes"`
}

// DisbursementUpdateAttributes is part of DisbursementUpdate.
type DisbursementUpdateAttributes struct {
	TargetID   string `json:"targetId"`
	TargetType string `json:"targetType"`
	Action     string `json:"action"`
}
