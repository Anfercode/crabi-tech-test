package infrastructure

import (
	"bytes"
	"encoding/json"
	"net/http"
)

type HTTPBlacklistValidator struct {
	endpoint string
}

func NewHTTPBlacklistValidator(endpoint string) *HTTPBlacklistValidator {
	return &HTTPBlacklistValidator{endpoint: endpoint}
}

func (api *HTTPBlacklistValidator) Validate(firstName string, lastName string, email string) (bool, error) {
	data := map[string]string{
		"first_name": firstName,
		"last_name":  lastName,
		"email":      email,
	}

	body, err := json.Marshal(data)
	if err != nil {
		return false, err
	}

	resp, err := http.Post(api.endpoint, "application/json", bytes.NewBuffer(body))
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()

	var respData struct {
		Blacklisted bool `json:"is_in_blacklist"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&respData); err != nil {
		return false, err
	}

	return respData.Blacklisted, nil
}
