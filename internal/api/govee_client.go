package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

const (
	baseURL = "https://openapi.api.govee.com"
)

type GoveeClient struct {
	HTTPClient *http.Client
	APIKey     string
}

func NewGoveeClient(apiKey string) *GoveeClient {
	return &GoveeClient{
		APIKey: apiKey,
		HTTPClient: &http.Client{
			Timeout: time.Second * 10,
		},
	}
}

func (g *GoveeClient) Get(endpoint string, target any) error {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/%s", baseURL, endpoint), nil)
	if err != nil {
		return err
	}
	req.Header.Add("Govee-API-Key", g.APIKey)

	resp, err := g.HTTPClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("API return status %d", resp.StatusCode)
	}

	return json.NewDecoder(resp.Body).Decode(target)
}

func (g *GoveeClient) Post(endpoint string, payload, target any) error {
	jsonData, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/%s", baseURL, endpoint), bytes.NewBuffer(jsonData))
	if err != nil {
		return err
	}

	resp, err := g.HTTPClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("API return status %d", resp.StatusCode)
	}

	return json.NewDecoder(resp.Body).Decode(target)
}
