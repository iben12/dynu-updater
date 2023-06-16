package main

import (
	"crypto/sha256"
	"fmt"
	"os"
	"strconv"
)

func getConfig() Config {
	user := os.Getenv("USERNAME")
	password := os.Getenv("PASSWORD")
	passwordHash := os.Getenv("PASSWORD_HASH")
	domain := os.Getenv("DOMAIN")
	period := os.Getenv("PERIOD_HOURS")

	var secret string

	if user == "" {
		logger.Fatal("Please set USERNAME env variable.")
	}

	if password == "" && passwordHash == "" {
		logger.Fatal("Please set PASSWORD or PASSWORD_HASH env variable.")
	} else if passwordHash != "" {
		logger.Print("Using password hash")
		secret = passwordHash
	} else if password != "" {
		logger.Print("Using password")
		secret = fmt.Sprintf("%x", sha256.Sum256([]byte(password)))
	}

	if domain == "" {
		logger.Fatal("Please set DOMAIN env variable.")
	}

	if period == "" {
		logger.Fatal("Please set PERIOD_HOURS env variable.")
	}

	parsedPeriod, err := strconv.Atoi(period)
	if err != nil {
		logger.Fatal("Please set PERIOD_HOURS to valid integer")
	}

	updateConfig := Config{
		User:   user,
		Secret: secret,
		Domain: domain,
		Period: parsedPeriod,
	}

	return updateConfig
}
