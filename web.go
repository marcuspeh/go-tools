package tools

import (
	"encoding/json"
	"fmt"

	"github.com/go-resty/resty/v2"
)

func GetRequest[reqStruct, respStruct any](url string, req *reqStruct) (*respStruct, error) {
	client := resty.New()

	marshalledReq, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	var queryParams map[string]interface{}
	err = json.Unmarshal(marshalledReq, &queryParams)
	if err != nil {
		return nil, err
	}

	params := map[string]string{}
	for k, v := range queryParams {
		params[k] = fmt.Sprintf("%v", v)
	}

	var respModel respStruct
	resp, err := client.R().
		SetQueryParams(params).
		SetResult(&respModel).
		Get(url)
	if err != nil {
		return nil, err
	}
	if !resp.IsSuccess() {
		return nil, fmt.Errorf("request failed with status code: %d", resp.StatusCode())
	}
	return &respModel, nil
}

func PostRequest[reqStruct, respStruct any](url string, req *reqStruct) (*respStruct, error) {
	client := resty.New()

	var respModel respStruct
	resp, err := client.R().
		SetBody(req).
		SetResult(&respModel).
		Post(url)
	if err != nil {
		return nil, err
	}
	if !resp.IsSuccess() {
		return nil, fmt.Errorf("request failed with status code: %d", resp.StatusCode())
	}
	return &respModel, nil
}

func DeleteRequest[reqStruct, respStruct any](url string, req *reqStruct) (*respStruct, error) {
	client := resty.New()

	var respModel respStruct
	resp, err := client.R().
		SetBody(req).
		SetResult(&respModel).
		Delete(url)
	if err != nil {
		return nil, err
	}
	if !resp.IsSuccess() {
		return nil, fmt.Errorf("request failed with status code: %d", resp.StatusCode())
	}
	return &respModel, nil
}
