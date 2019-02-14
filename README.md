# Grapenut

[![Build Status](https://travis-ci.org/manywho/grapenut.svg?branch=master)](https://travis-ci.org/manywho/grapenut)
[![Go Report Card](https://goreportcard.com/badge/github.com/manywho/grapenut)](https://goreportcard.com/report/github.com/manywho/grapenut)
[![Docker Repository on Quay](https://quay.io/repository/manywho/grapenut/status "Docker Repository on Quay")](https://quay.io/repository/manywho/grapenut)

Grapenut is a simple program for sending a count of the currently active alerts in grafana to slack. Currently, if there are no active alerts, Grapenut will skip sending the notification to slack to avoid spam.

## Configuration
Grapenut is configured via environment variables. The following environment variables can be set.

| Key               | Required | Description                                                                                              |
|-------------------|----------|----------------------------------------------------------------------------------------------------------|
| GRAFANA_URL       | Yes      | The protocol prefixed address to use when connecting to Grafana (example: https://grafana.mysite.net     |
| GRAFANA_API_KEY   | Yes      | The API key to use when connecting to Grafana. The key only requires the **Viewer** permission           |
| SLACK_WEBHOOK_URL | Yes      | The Slack webhook to send notifications to                                                               |
| DEBUG_LOG         | No       | Enable the Debug log by setting this to `true`. By default, there is no output unless there was an error |

## Running the tests
To run the tests, run the following command:

```
# This is the default target so either of these commands will run the tests
make
make test
```

## Build the binary
To build the binary, you can use the Makefile build target:

```
make build
```

## Running the Docker Image
To run the docker image, set the environment variables either with the `-e` or `--env-file` flags:

```
docker run -ti --rm --env-file /tmp/grapenut.env quay.io/manywho/grapenut:$VERSION
```