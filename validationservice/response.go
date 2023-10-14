// File Created: Friday, 13th October 2023 1:27:14 am
// Author: Abdul Hamid (abdul.surel@gmail.com)
//
// Copyright (c) 2023 Author

package validationservice

// BankAccount represents response from Validate a bank account API
type BankAccount struct {
	AccountName   string `json:"accountName"`
	AccountNo     string `json:"accountNo"`
	BankShortCode string `json:"bankShortCode"`
}
