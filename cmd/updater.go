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

type IPResponse struct {
	Ip string `json:"ip"`
}

type UpdateConfig struct {
	User     string
	Password string
	Domain   string
	Period   int
}

func getConfig() UpdateConfig {
	user := os.Getenv("USERNAME")
	password := os.Getenv("PASSWORD")
	domain := os.Getenv("DOMAIN")
	period := os.Getenv("PERIOD_HOURS")

	if user == "" {
		log.Fatal("Please set USERNAME env variable.")
	} else if password == "" {
		log.Fatal("Please set PASSWORD env variable.")
	} else if domain == "" {
		log.Fatal("Please set DOMAIN env variable.")
	} else if period == "" {
		log.Fatal("Please set PERIOD_HOURS env variable.")
	}

	parsedPeriod, err := strconv.Atoi(period)
	if err != nil {
		log.Fatal("Please set PERIOD_HOURS to valid integer")
	}

	updateConfig := UpdateConfig{
		User:     user,
		Password: fmt.Sprintf("%v", sha256.Sum256([]byte(password))),
		Domain:   domain,
		Period:   parsedPeriod,
	}

	return updateConfig
}

func getIP() string {
	resp, err := http.Get("https://api.myip.com")
	if err != nil {
		log.Fatal(err)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	var ipResponse IPResponse

	json.Unmarshal([]byte(body), &ipResponse)

	return ipResponse.Ip
}

func updateIp(config UpdateConfig, ip string) error {
	resp, err := http.Get(fmt.Sprintf(`https://api.dynu.com/nic/update?hostname=%s&myip=%s&username=%s&password=%s`, config.Domain, ip, config.User, config.Password))
	if err != nil {
		return err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	fmt.Println("Update result is:", string(body))

	return nil
}

func doUpdate(config UpdateConfig) {
	ip := getIP()

	fmt.Println("Current IP is", ip)
	fmt.Println("Updating Dynu...")

	err := updateIp(config, ip)

	if err != nil {
		log.Fatal(err)
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

	doUpdate(updateConfig)

	fmt.Printf("Starting %v-hour interval updates...\n", updateConfig.Period)

	interval(updateConfig)
}
