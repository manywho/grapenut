package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestSlackIsSent(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("ok"))
	}))

	cfg := slackConfig{
		WebhookURL: srv.URL,
	}

	err := postNotification(cfg, "https://notrealgrafana.fake", 1)
	if err != nil {
		t.Fatalf("got unexpected error: %s", err)
	}
}

func TestSlackHandlesNon200(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("testing error"))
	}))

	cfg := slackConfig{
		WebhookURL: srv.URL,
	}

	err := postNotification(cfg, "https://notrealgrafana.fake", 1)
	if err == nil {
		t.Fatal("expected an error but didn't get one")
	}

	if !strings.Contains(err.Error(), "500") {
		t.Errorf("expected error to contain status code but it didn't: %s", err)
	}

	if !strings.Contains(err.Error(), "testing error") {
		t.Errorf("expected error to contain response body but it didn't: %s", err)
	}
}
