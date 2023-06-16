package main

import (
	"fmt"
	"io"
	"net/http"
)

func updateIp(config Config, ip string) error {
	query := fmt.Sprintf(`hostname=%s&myip=%s&username=%s&password=%s`, config.Domain, ip, config.User, config.Secret)
	resp, err := http.Get(fmt.Sprintf(`%s?%s`, config.ServerURL, query))

	if err != nil {
		return err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	code := string(body)

	if code != "good" && code != "nochg" {
		logger.Fatal("Update failed. Error code: ", code)
	}

	logger.Println("Update result is:", code)

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
