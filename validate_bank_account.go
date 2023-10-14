// File Created: Friday, 13th October 2023 1:03:29 am
// Author: Abdul Hamid (abdul.surel@gmail.com)
//
// Copyright (c) 2023 Author

package fazz

// ValidateBankAccount represent for payload request Validate a bank account API
//
// Docs: https://docs.fazz.com/v4-ID/reference/validate-bank-account
type ValidateBankAccountPayload struct {
	AccountNo     string `json:"accountNo"`
	BankShortCode string `json:"bankShortCode"`
}
