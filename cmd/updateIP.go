package main

import (
	"fmt"
	"io"
	"net/http"
)

func updateIp(config UpdateConfig, ip string) error {
	query := fmt.Sprintf(`hostname=%s&myip=%s&username=%s&password=%s`, config.Domain, ip, config.User, config.Password)
	resp, err := http.Get(fmt.Sprintf(`%s?%s`, config.ServerURL, query))
	if err != nil {
		return err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	logger.Println("Update result is:", string(body))

	return nil
}

func doUpdate(config UpdateConfig) {
	ip := getIP(config.IpServer)

	logger.Println("Current IP is", ip)
	logger.Println("Updating Dynu...")

	err := updateIp(config, ip)

	if err != nil {
		logger.Fatal(err)
	}
}
