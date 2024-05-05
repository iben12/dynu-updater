package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

var logger = log.New(os.Stdout, "", 5)

type Config struct {
	ApiKey    string
	DnsId     string
	Domain    string
	Period    int
	ServerURL string
	IpServer  string
}

func startInterval(config Config) {
	t := time.NewTicker(time.Duration(config.Period) * time.Hour)

	defer t.Stop()
	for range t.C {
		doUpdate(config)
	}

}

func main() {
	updateConfig := getConfig()

	updateConfig.ServerURL = "https://api.dynu.com/nic/update"
	updateConfig.ServerURL = fmt.Sprintf("https://api.dynu.com/v2/dns/%s", updateConfig.DnsId)
	updateConfig.IpServer = "https://api.myip.com"

	doUpdate(updateConfig)

	logger.Printf("Starting %v-hour interval updates...\n", updateConfig.Period)

	startInterval(updateConfig)
}
