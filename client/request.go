package client

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// APIError is an error type that exposes additional information about why an API request failed.
type APIError struct {
	Code       string `json:"code"`
	Message    string `json:"message"`
	StatusCode int
	RawMessage []byte
}

// Error provides a user friendly error message.
func (e APIError) Error() string {
	return fmt.Sprintf("%s - %s", e.Code, e.Message)
}

// doRequest is a helper function for consistently requesting data from vercel.
// This manages:
// - Authorization via the Bearer token
// - Converting error responses into an inspectable type
// - Unmarshaling responses
func (c *Client) doRequest(req *http.Request, v interface{}) error {
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.token))
	resp, err := c.http().Do(req)
	if err != nil {
		return fmt.Errorf("error doing http request: %w", err)
	}

	defer resp.Body.Close()
	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("error reading response body: %w", err)
	}

	if resp.StatusCode >= 300 {
		var errorResponse APIError
		err = json.Unmarshal(responseBody, &struct {
			Error *APIError `json:"error"`
		}{
			Error: &errorResponse,
		})
		if err != nil {
			return fmt.Errorf("error unmarshaling response: %w", err)
		}
		errorResponse.StatusCode = resp.StatusCode
		errorResponse.RawMessage = responseBody
		return errorResponse
	}

	if v == nil {
		return nil
	}

	err = json.Unmarshal(responseBody, v)
	if err != nil {
		return fmt.Errorf("error unmarshaling response: %w", err)
	}

	return nil
}
