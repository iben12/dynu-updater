package main

import (
	"crypto/sha256"
	"fmt"
	"os"
	"strconv"
)

func getConfig() UpdateConfig {
	user := os.Getenv("USERNAME")
	password := os.Getenv("PASSWORD")
	domain := os.Getenv("DOMAIN")
	period := os.Getenv("PERIOD_HOURS")

	if user == "" {
		logger.Fatal("Please set USERNAME env variable.")
	} else if password == "" {
		logger.Fatal("Please set PASSWORD env variable.")
	} else if domain == "" {
		logger.Fatal("Please set DOMAIN env variable.")
	} else if period == "" {
		logger.Fatal("Please set PERIOD_HOURS env variable.")
	}

	parsedPeriod, err := strconv.Atoi(period)
	if err != nil {
		logger.Fatal("Please set PERIOD_HOURS to valid integer")
	}

	updateConfig := UpdateConfig{
		User:      user,
		Password:  fmt.Sprintf("%v", sha256.Sum256([]byte(password))),
		Domain:    domain,
		Period:    parsedPeriod,
		ServerURL: "",
		IpServer:  "",
	}

	return updateConfig
}
