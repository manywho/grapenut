package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	log "github.com/sirupsen/logrus"
)

type grafanaConfig struct {
	URL    string
	APIKey string
}

type grafanaError struct {
	code int
	msg  string
}

func (e *grafanaError) Error() string {
	return fmt.Sprintf("Received a %d status from grafana: %s", e.code, e.msg)
}

func getActiveAlertCount(cfg grafanaConfig) (count int, err error) {
	type alertBody struct {
		Name string `json:"name"`
	}

	req, err := http.NewRequest(http.MethodGet, cfg.URL+"/api/alerts?state=alerting", nil)
	if err != nil {
		return
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", cfg.APIKey))

	log.Debugf("Making request to grafana: %s", req.URL)
	resp, err := httpClient.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	log.Debugf("Got response: %s", body)

	if resp.StatusCode != 200 {
		err = &grafanaError{
			resp.StatusCode,
			string(body),
		}
		return
	}

	respJSON := []alertBody{}

	if err = json.Unmarshal(body, &respJSON); err != nil {
		return
	}

	count = len(respJSON)

	return
}
