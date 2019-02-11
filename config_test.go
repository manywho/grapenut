package main

import (
	"os"
	"strings"
	"testing"
)

func TestConfigIsLoadedFromEnvironment(t *testing.T) {
	gURL := "https://grafanatest.com"
	gAPIKey := "abc123"
	sWebhookURL := "https://slackwebhook.com"

	os.Setenv("GRAFANA_URL", gURL)
	os.Setenv("GRAFANA_API_KEY", gAPIKey)
	os.Setenv("SLACK_WEBHOOK_URL", sWebhookURL)

	gConfig, sConfig := loadConfig()
	if gConfig.URL != gURL {
		t.Errorf("expected '%s', got '%s'", gURL, gConfig.URL)
	}
	if gConfig.APIKey != gAPIKey {
		t.Errorf("expected '%s', got '%s'", gAPIKey, gConfig.APIKey)
	}
	if sConfig.WebhookURL != sWebhookURL {
		t.Errorf("expected '%s', got '%s'", sWebhookURL, sConfig.WebhookURL)
	}
}

func TestValidateConfig(t *testing.T) {
	err := validateConfig(grafanaConfig{}, slackConfig{})

	if !strings.Contains(err.Error(), "GRAFANA_URL, GRAFANA_API_KEY, SLACK_WEBHOOK_URL") {
		t.Fatalf("expected 3 required keys, got %s", err)
	}
}
