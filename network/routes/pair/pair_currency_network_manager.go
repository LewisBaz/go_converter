package pair_currency_nm

import (
	"encoding/json"
	"fmt"
	"go_converter/network"
	"go_converter/network/env"
	"net/http"
)

type PairCurrencyNM struct {
	Client http.Client
}

type Request struct {
	Amount float64
	From string
	To string
}

type Response struct {
	Result float64 `json:"conversion_result"`
}

func MakeRequest(nm PairCurrencyNM, r Request) (*Response, error) {
	apikey, err := env.Getapikey()
	if err != nil {
		return nil, err
	}

	url := fmt.Sprintf("https://v6.exchangerate-api.com/v6/%s/pair/%s/%s/%.2f", apikey, r.From, r.To, r.Amount)
	body, err := network.BaseRequest(network.RequestDep{ Client: nm.Client, URL: url }) 
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