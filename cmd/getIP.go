package main

import (
	"encoding/json"
	"io"
	"net/http"
)

type IPResponse struct {
	Ip string `json:"ip"`
}

func getIP(ipApiUrl string) string {
	resp, err := http.Get(ipApiUrl)
	if err != nil {
		logger.Fatal(err)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		logger.Fatal(err)
	}

	var ipResponse IPResponse

	json.Unmarshal([]byte(body), &ipResponse)

	return ipResponse.Ip
}
