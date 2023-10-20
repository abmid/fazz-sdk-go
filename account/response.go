// File Created: Friday, 20th October 2023 4:19:19 pm
// Author: Abdul Hamid (abdul.surel@gmail.com)
//
// Copyright (c) 2023 Author

package account

// Balance represent response from Get Account Balance API.
type Balance struct {
	TotalBalance     string `json:"totalBalance"`
	AvailableBalance string `json:"availableBalance"`
	PendingBalance   string `json:"pendingBalance"`
}
