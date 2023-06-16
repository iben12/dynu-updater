package main

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

var logger = log.New(os.Stdout, "", 5)

type IPResponse struct {
	Ip string `json:"ip"`
}

type UpdateConfig struct {
	User      string
	Password  string
	Domain    string
	Period    int
	ServerURL string
	IpServer  string
}

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
		User:     user,
		Password: fmt.Sprintf("%v", sha256.Sum256([]byte(password))),
		Domain:   domain,
		Period:   parsedPeriod,
	}

	return updateConfig
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

func updateIp(config UpdateConfig, ip string) error {
	query := fmt.Sprintf(`?hostname=%s&myip=%s&username=%s&password=%s`, config.Domain, ip, config.User, config.Password)
	resp, err := http.Get(fmt.Sprintf(`%s%s`, config.ServerURL, query))
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

func interval(config UpdateConfig) {
	t := time.NewTicker(time.Duration(config.Period) * time.Hour)

	defer t.Stop()
	for range t.C {
		doUpdate(config)
	}

}

func main() {
	updateConfig := getConfig()

	updateConfig.ServerURL = "https://api.dynu.com/nic/update"
	updateConfig.IpServer = "https://api.myip.com"

	doUpdate(updateConfig)

	logger.Printf("Starting %v-hour interval updates...\n", updateConfig.Period)

	interval(updateConfig)
}
