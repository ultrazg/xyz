package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

var client = &http.Client{
	Timeout: time.Second * 15,
}

func Request(url, method string, body map[string]any, headers map[string]string) (*http.Response, int, error) {
	payload, err := json.Marshal(body)
	if err != nil {
		return nil, 0, fmt.Errorf("failed to marshal request body: %v", err)
	}
	req, err := http.NewRequest(method, url, bytes.NewReader(payload))
	if err != nil {
		return nil, 0, fmt.Errorf("failed to create request: %v", err)
	}

	if headers != nil {
		// 请求头
		for key, value := range headers {
			req.Header.Set(key, value)
		}
	}

	var resp *http.Response
	var retryErr error
	//for i := 0; i < 3; i++ {
	resp, retryErr = client.Do(req)
	//if retryErr == nil && resp.StatusCode >= 200 && resp.StatusCode < 300 {
	//	break
	//}

	//time.Sleep(time.Second * 2)
	//}

	//defer resp.Body.Close()

	if retryErr != nil {
		return nil, 0, fmt.Errorf("failed to send request: %v", retryErr)
	}

	// 未登录
	if resp.StatusCode == http.StatusUnauthorized {
		return nil, resp.StatusCode, fmt.Errorf("received 401 status code: %d", resp.StatusCode)
	}
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, resp.StatusCode, fmt.Errorf("received non-200 status code: %d", resp.StatusCode)
	}

	return resp, resp.StatusCode, nil
}
