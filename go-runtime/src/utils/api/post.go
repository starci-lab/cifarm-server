package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

func SendPostRequest[TResponseData any](url string, body string) (*TResponseData, error) {
	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBufferString(body))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if !IsStatusCode2xx(resp.StatusCode) {
		return nil, fmt.Errorf("non-OK HTTP status: %s", resp.Status)
	}

	var response *TResponseData
	err = json.Unmarshal(responseBody, &response)
	return response, err
}
