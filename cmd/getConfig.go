package main

import (
	"os"
	"strconv"
)

func getConfig() Config {
	apiKey := os.Getenv("API_KEY")
	dnsId := os.Getenv("DNS_ID")
	domain := os.Getenv("DOMAIN")
	period := os.Getenv("PERIOD_HOURS")

	if apiKey == "" {
		logger.Fatal("Please set API_KEY env variable.")
	}

	if dnsId == "" {
		logger.Fatal("Please set DNS_ID env variable.")
	}

	if period == "" {
		logger.Fatal("Please set PERIOD_HOURS env variable.")
	}

	parsedPeriod, err := strconv.Atoi(period)
	if err != nil {
		logger.Fatal("Please set PERIOD_HOURS to valid integer")
	}

	config := Config{
		ApiKey: apiKey,
		DnsId:  dnsId,
		Domain: domain,
		Period: parsedPeriod,
	}

	return config
}
