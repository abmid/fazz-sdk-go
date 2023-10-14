// File Created: Friday, 13th October 2023 12:47:04 am
// Author: Abdul Hamid (abdul.surel@gmail.com)
//
// Copyright (c) 2023 Author

package helper

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
	"reflect"
	"testing"
	"time"

	mock_request "github.com/abmid/fazz-sdk-go/internal/test/mock"
	"github.com/golang/mock/gomock"
	"github.com/jarcoal/httpmock"
)

type Mocks struct {
	Api *mock_request.MockApi
}

type TestWrapper struct {
	Ctrl      *gomock.Controller
	ServerKey string
}

func NewTestWrapper(t *testing.T) *TestWrapper {
	ctrl := gomock.NewController(t)

	return &TestWrapper{
		Ctrl:      ctrl,
		ServerKey: "dpay_test_xxx",
	}
}

// ResJSONByte is used for convert file json to []byte
func (w *TestWrapper) ResJSONByte(jsonFile string) []byte {
	file, err := os.ReadFile(jsonFile)
	if err != nil {
		panic(err)
	}

	return file
}

// DeepEqualPayload checks whether the payload for the request matches the example json or not.
// Value payload and argPayload must be assign as pointer
func (w *TestWrapper) DeepEqualPayload(fileJson string, payload any, argPayload any) bool {
	type fazzPayload struct {
		Data struct {
			Attributes any `json:"attributes"`
		} `json:"data"`
	}

	payloadToFazz := fazzPayload{
		Data: struct {
			Attributes any `json:"attributes"`
		}{
			Attributes: payload,
		},
	}
	json.Unmarshal(w.ResJSONByte(fileJson), &payloadToFazz)

	if reflect.DeepEqual(payloadToFazz.Data.Attributes, argPayload) {
		return true
	}

	var castPayload, castArgsPayload any

	bytesCastPayload, err := json.Marshal(payloadToFazz)
	if err != nil {
		panic(fmt.Sprintf("DeepEqualPayload: %v", err))
	}

	bytesCastArgsPayload, err := json.Marshal(argPayload)
	if err != nil {
		panic(fmt.Sprintf("DeepEqualPayload: %v", err))
	}

	err = json.Unmarshal(bytesCastPayload, &castPayload)
	if err != nil {
		panic(fmt.Sprintf("DeepEqualPayload: %v", err))
	}

	err = json.Unmarshal(bytesCastArgsPayload, &castArgsPayload)
	if err != nil {
		panic(fmt.Sprintf("DeepEqualPayload: %v", err))
	}

	if reflect.DeepEqual(castPayload, castArgsPayload) {
		return true
	}

	return false
}

// DeepEqualResponse is used to check responses that have interface values
func (w *TestWrapper) DeepEqualResponse(gotRes any, wantRes any) bool {
	if reflect.DeepEqual(wantRes, gotRes) {
		return true
	}

	var castWantRes, castGotRes any

	bytesCastWantRes, err := json.Marshal(wantRes)
	if err != nil {
		panic(fmt.Sprintf("DeepEqualPayload: %v", err))
	}

	bytesCastGotRes, err := json.Marshal(gotRes)
	if err != nil {
		panic(fmt.Sprintf("DeepEqualPayload: %v", err))
	}

	err = json.Unmarshal(bytesCastWantRes, &castWantRes)
	if err != nil {
		panic(fmt.Sprintf("DeepEqualPayload: %v", err))
	}

	err = json.Unmarshal(bytesCastGotRes, &castGotRes)
	if err != nil {
		panic(fmt.Sprintf("DeepEqualPayload: %v", err))
	}

	if reflect.DeepEqual(castWantRes, castGotRes) {
		return true
	}

	return false
}

func HttpMockResJSON(statusCode int, filePath string, headers map[string]string) httpmock.Responder {
	return func(r *http.Request) (*http.Response, error) {
		if r.Header.Get("Authorization") == "" {
			return nil, errors.New("access_key not present!")
		}

		if r.Header.Get("Content-Type") != "application/vnd.api+json" {
			return nil, errors.New("Invalid content-type != application/vnd.api+json")
		}

		if headers != nil {
			for key, value := range headers {
				if r.Header.Get(key) != value {
					return nil, errors.New(fmt.Sprintf("Headers: Invalid %s != %s", key, value))
				}
			}
		}

		resp, _ := httpmock.NewJsonResponderOrPanic(statusCode, httpmock.File(filePath))(r)

		return resp, nil
	}
}

// ToPtr return value pointer for anything data types.
func ToPtr[V any](value V) *V {
	return &value
}

// StringToTime return string to time without return error.
// If when parsing encounters an error, it will return the default value
func StringToTime(timeString string) time.Time {
	parse, err := time.Parse(time.RFC3339, timeString)
	if err != nil {
		return time.Time{}
	}

	return parse
}
