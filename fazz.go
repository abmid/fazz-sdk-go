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

// FazzParams represent default pagination, sorting & filtering from Fazz
type FazzParams struct {
	PageNumber    *uint32 `url:"page[number]"`
	PageSize      *uint32 `url:"page[size]"`
	Sort          *string `url:"sort"`
	CreatedAfter  *string `url:"filter[createdAfter]"`  // Time.iso8601
	CreatedBefore *string `url:"filter[createdBefore]"` // Time.iso8601
	Status        *string `url:"filter[status]"`
	ReferenceID   *string `url:"filter[referenceId]"`
}
