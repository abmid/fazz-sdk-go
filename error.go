// File Created: Friday, 13th October 2023 12:05:56 am
// Author: Abdul Hamid (abdul.surel@gmail.com)
//
// Copyright (c) 2023 Author

package fazz

import "encoding/json"

const (
	ErrorCodeSDK = "SDK_ERROR"
	// Generic Errors
	// ErrorCode000 is used for error Unknown error
	ErrorCode000 = "000"
	// ErrorCode001 is used for error Internal server error
	ErrorCode001 = "001"
	// ErrorCode002 is used for error Access denied
	ErrorCode002 = "002"
	// ErrorCode003 is used for error Feature not supported
	ErrorCode003 = "003"
	// ErrorCode004 is used for error Record not found
	ErrorCode004 = "004"
	// ErrorCode005 is used for error Invalid parameters
	ErrorCode005 = "005"
	// ErrorCode010 is used for error Service Temporarily Unavailable
	ErrorCode010 = "010"

	// Transaction Specific Errors
	// ErrorCodeTXN0001 is used for error Account has insufficient funds
	ErrorCodeTXN0001 = "TXN0001"
	// ErrorCodeTXN0002 is used for error Invalid transaction amount
	ErrorCodeTXN0002 = "TXN0002"
	// ErrorCodeTXN0003 is used for error Transaction exceeds account transaction limit
	ErrorCodeTXN0003 = "TXN0003"
	// ErrorCodeTXN0004 is used for error Transaction exceeds account balance limit
	ErrorCodeTXN0004 = "TXN0004"
	// ErrorCodeTXN0005 is used for error Failed to create virtual account number
	ErrorCodeTXN0005 = "TXN0005"
	// ErrorCodeTXN0006 is used for error suffixNo params is not supported for this bank
	ErrorCodeTXN0006 = "TXN0006"
	// ErrorCodeTXN0007 is used for error suffixNo params exceed maximum number limit
	ErrorCodeTXN0007 = "TXN0007"
)

// Error is combined from standard error Fazz with response HTTP Status Code
type Error struct {
	HttpStatusCode int      // Response from http status code
	Errors         []Errors `json:"errors"`
}

// Errors is standard list response error from Fazz
type Errors struct {
	Code   string `json:"code"`
	Title  string `json:"title"`
	Detail string `json:"detail"`
}

// ErrFromSDK is used when get an error from the SDK
func ErrFromSDK(err error) *Error {
	return &Error{
		Errors: []Errors{
			{Code: ErrorCodeSDK},
		},
	}
}

// ErrFromAPI is used when get an error from response API
func ErrFromAPI(httpStatusCode int, responseBody []byte) *Error {
	resErr := Error{HttpStatusCode: httpStatusCode}

	err := json.Unmarshal(responseBody, &resErr)
	if err != nil {
		return ErrFromSDK(err)
	}

	return &resErr
}
