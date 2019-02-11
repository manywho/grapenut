package main

import (
	"fmt"
	"os"
	"strings"
)

type validationRequiredError struct {
	keys []string
}

func (e validationRequiredError) Error() string {
	return fmt.Sprintf("The following environment variables are required but missing: %s", strings.Join(e.keys, ", "))
}

func validateConfig(gConfig grafanaConfig, sConfig slackConfig) error {
	var keys []string

	if gConfig.URL == "" {
		keys = append(keys, "GRAFANA_URL")
	}
	if gConfig.APIKey == "" {
		keys = append(keys, "GRAFANA_API_KEY")
	}

	if sConfig.WebhookURL == "" {
		keys = append(keys, "SLACK_WEBHOOK_URL")
	}

	if len(keys) > 0 {
		return &validationRequiredError{keys}
	}

	return nil
}

func loadConfig() (grafanaConfig, slackConfig) {
	return grafanaConfig{
			URL:    os.Getenv("GRAFANA_URL"),
			APIKey: os.Getenv("GRAFANA_API_KEY"),
		},
		slackConfig{
			WebhookURL: os.Getenv("SLACK_WEBHOOK_URL"),
		}
}
