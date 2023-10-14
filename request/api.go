// File Created: Friday, 13th October 2023 12:03:48 am
// Author: Abdul Hamid (abdul.surel@gmail.com)
//
// Copyright (c) 2023 Author

package request

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/abmid/fazz-sdk-go"
	goquery "github.com/google/go-querystring/query"
)

// Api interface build for DI and testing.
type Api interface {
	Req(ctx context.Context, method string, url string, param any, body any, headers map[string]string, response any) *fazz.Error
}

type ApiImplement struct {
	ApiKey    string
	SecretKey string
}

func NewAPI(apiKey, secretKey string) *ApiImplement {
	return &ApiImplement{
		ApiKey:    apiKey,
		SecretKey: secretKey,
	}
}

// Req is an http request made specifically to hit the Fazz endpoint.
// If the HTTP status code return is not 2xx, response error will be returned
func (c *ApiImplement) Req(ctx context.Context, method string, url string, param any, body any, headers map[string]string, response any) *fazz.Error {
	fazzPayload := fazz.FazzPayload{
		Data: struct {
			Attributes any `json:"attributes"`
		}{
			Attributes: body,
		},
	}

	parseBody, err := json.Marshal(fazzPayload)
	if err != nil {
		return fazz.ErrFromSDK(err)
	}

	base64SecretKey := base64.StdEncoding.EncodeToString([]byte(c.ApiKey + ":" + c.SecretKey))
	httpReq, err := http.NewRequestWithContext(ctx, method, url, bytes.NewReader(parseBody))
	httpReq.Header.Add("Content-Type", "application/vnd.api+json")
	httpReq.Header.Add("Authorization", fmt.Sprintf("Basic %s", base64SecretKey))

	if headers != nil {
		for key, value := range headers {
			httpReq.Header.Add(key, value)
		}
	}

	if param != nil {
		parseParam, err := goquery.Values(param)
		if err != nil {
			return fazz.ErrFromSDK(err)
		}

		httpReq.URL.RawQuery = parseParam.Encode()
	}

	httpRes, err := http.DefaultClient.Do(httpReq)
	if err != nil {
		return fazz.ErrFromSDK(err)
	}
	defer httpRes.Body.Close()

	resBody, err := io.ReadAll(httpRes.Body)
	if err != nil {
		return fazz.ErrFromSDK(err)
	}

	isStatusCodeSuccess := (httpRes.StatusCode >= 200) && (httpRes.StatusCode < 300)

	if !isStatusCodeSuccess {
		return fazz.ErrFromAPI(httpRes.StatusCode, resBody)
	}

	if response != nil {
		jsonErr := json.Unmarshal(resBody, response)
		if jsonErr != nil {
			return fazz.ErrFromSDK(jsonErr)
		}
	}

	return nil
}
