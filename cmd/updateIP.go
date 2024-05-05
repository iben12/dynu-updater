package main

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"net/http"
)

func updateIp(config Config, ip string) error {
	client := http.Client{}

	data := fmt.Sprintf(`{"ipv4Address":"%s","name":"%s","id":"%s"}`, ip, config.Domain, config.DnsId)
	request, err := http.NewRequest("POST", config.ServerURL, bytes.NewBuffer([]byte(data)))

	if err != nil {
		return err
	}

	request.Header = http.Header{
		"Accept":  {"application/json"},
		"API-Key": {config.ApiKey},
	}

	response, err := client.Do(request)

	if err != nil {
		logger.Println(err)

		return err
	}

	if response.Status != "200 OK" {
		body, _ := io.ReadAll(response.Body)

		logger.Println(string(body))

		return errors.New(response.Status)
	}

	logger.Println("Update result is:", response.Status)

	return nil
}

func doUpdate(config Config) {
	ip := getIP(config.IpServer)

	logger.Println("Current IP is", ip)
	logger.Println("Updating Dynu...")

	err := updateIp(config, ip)

	if err != nil {
		logger.Fatal(err)
	}
}
