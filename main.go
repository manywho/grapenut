package main

import (
	"net/http"
	"os"

	log "github.com/sirupsen/logrus"
)

var (
	httpClient = http.Client{}
)

func main() {
	if os.Getenv("DEBUG_LOG") == "true" {
		log.SetLevel(log.DebugLevel)
	}

	log.Debug("Loading and validating configuration from env")
	gConfig, sConfig := loadConfig()

	err := validateConfig(gConfig, sConfig)
	if err != nil {
		log.Fatal(err)
	}

	log.Debug("Fetching active alert count from Grafana")
	count, err := getActiveAlertCount(gConfig)
	if err != nil {
		log.Fatal(err)
	}

	if count == 0 {
		log.Debug("No active alerts, nothing to do")
		os.Exit(0)
	}

	log.Debugf("Found %d alerts. Sending message to slack", count)

	if err = postNotification(sConfig, gConfig.URL, count); err != nil {
		log.Fatal(err)
	}
}
