// File Created: Friday, 23rd February 2024 1:02:16 am
// Author: Abdul Hamid (abdul.surel@gmail.com)
//
// Copyright (c) 2024 Author

package example

import (
	"context"
	"fmt"
)

func BankGetListBanks() {
	// ========== EXAMPLE Get a list of banks ==========
	res, err := c.Bank.Banks(context.Background())
	if err != nil {
		// handle case error
	}

	fmt.Println(res)
}
