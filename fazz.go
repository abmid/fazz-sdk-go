// File Created: Thursday, 12th October 2023 11:47:46 pm
// Author: Abdul Hamid (abdul.surel@gmail.com)
//
// Copyright (c) 2023 Author

package fazz

const (
	SandboxURL    = "https://sandbox-id.xfers.com/api"
	ProductionURL = "https://id.xfers.com/api"
)

// FazzPayload represent default format body params from Fazz
type FazzPayload struct {
	Data struct {
		Attributes any `json:"attributes"`
	} `json:"data"`
}
