package network

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type Request struct {
	Amount float64
	From string
	To string
}

type Response struct {
	Result float64 `json:"conversion_result"`
}

var client = http.Client{}

func MakeRequest(requestBody Request) (*Response, error) {
	apikey := os.Getenv("ExchangeRate_API_KEY")
	url := fmt.Sprintf("https://v6.exchangerate-api.com/v6/%s/pair/%s/%s/%.2f", apikey, requestBody.From, requestBody.To, requestBody.Amount)

	req, err := http.NewRequest("GET", url, nil) 
	if err != nil {
		return nil, err
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API request failed with status code: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var response Response

	error := json.Unmarshal(body, &response)
	if error != nil {
		return nil, error
	}

	return &response, nil
}