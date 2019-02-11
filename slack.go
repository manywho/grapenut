package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	log "github.com/sirupsen/logrus"
)

type slackConfig struct {
	WebhookURL string
}

type slackError struct {
	code int
	msg  string
}

func (e *slackError) Error() string {
	return fmt.Sprintf("Recieved a %d status from slack: %s", e.code, e.msg)
}

func postNotification(cfg slackConfig, grafanaURL string, alertCount int) (err error) {

	postString := fmt.Sprintf(`{
		"attachments": [
			{
				"color": "#ff0000",
				"fallback": "Grafana has %[1]d active alerts",
				"title": "Grafana has active alerts!",
				"title_link": "%s/alerting/list?state=alerting",
				"text": "Grafana has %[1]d active alerts"
			}
		]
	}`, alertCount, grafanaURL)

	log.Debugf("Slack JSON payload: %s", strings.TrimSpace(postString))

	req, err := http.NewRequest(http.MethodPost, cfg.WebhookURL, bytes.NewBufferString(postString))
	if err != nil {
		return
	}

	log.Debugf("Sending request to %s", cfg.WebhookURL)
	resp, err := httpClient.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	if resp.StatusCode != 200 {
		err = &slackError{
			resp.StatusCode,
			string(body),
		}
		return
	}

	return
}
