package network

import (
	"fmt"
	"io"
	"net/http"
)

type RequestDep struct {
	Client http.Client
	URL string
}

func BaseRequest(rd RequestDep) ([]byte, error) {
	req, err := http.NewRequest("GET", rd.URL, nil) 
	if err != nil {
		return nil, err
	}

	resp, err := rd.Client.Do(req)
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

	return body, nil
}