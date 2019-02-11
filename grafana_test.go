package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestGrafanaAuthHeaderIsSet(t *testing.T) {

	var authHeader string

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader = r.Header.Get("Authorization")
		w.Write([]byte("[]"))
	}))

	cfg := grafanaConfig{
		URL:    srv.URL,
		APIKey: "abc123",
	}

	_, err := getActiveAlertCount(cfg)
	if err != nil {
		t.Fatalf("got unexpected error: %s", err)
	}

	if authHeader != "Bearer abc123" {
		t.Fatalf("expected 'Bearer abc123', got '%s'", authHeader)
	}
}

func TestGrafanaURLIsValid(t *testing.T) {
	cfg := grafanaConfig{
		URL: "missing.protocol.com",
	}

	_, err := getActiveAlertCount(cfg)
	if err == nil {
		t.Fatal("expected an error but didn't get one")
	}

	if !strings.Contains(err.Error(), "unsupported protocol scheme") {
		t.Errorf("expected unsupported protocol error, got '%s'", err)
	}
}

func TestGrafanaCountReturned(t *testing.T) {
	resp := `[
		{
			"name": "test1"
		},
		{
			"name": "test2"
		}
	]`

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(resp))
	}))

	cfg := grafanaConfig{
		URL:    srv.URL,
		APIKey: "abc123",
	}

	i, err := getActiveAlertCount(cfg)
	if err != nil {
		t.Fatalf("got unexpected error: %s", err)
	}

	if i != 2 {
		t.Fatalf("expected 2, got %d", i)
	}
}

func TestGrafanaCountHandleEmpty(t *testing.T) {
	resp := `[]`

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(resp))
	}))

	cfg := grafanaConfig{
		URL: srv.URL,
	}

	i, err := getActiveAlertCount(cfg)
	if err != nil {
		t.Fatalf("got unexpected error: %s", err)
	}

	if i != 0 {
		t.Fatalf("expected 0, got %d", i)
	}
}

func TestGrafanaErrorsOnEmptyResponse(t *testing.T) {
	resp := ``

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(resp))
	}))

	cfg := grafanaConfig{
		URL: srv.URL,
	}

	_, err := getActiveAlertCount(cfg)
	if err == nil {
		t.Fatal("expected an error but didn't get one")
	}

	if !strings.Contains(err.Error(), "end of JSON input") {
		t.Fatalf("expected JSON error, got '%s'", err)
	}

}

func TestGrafanaCountUnauthorized(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusUnauthorized)
	}))

	cfg := grafanaConfig{
		URL: srv.URL,
	}

	i, err := getActiveAlertCount(cfg)

	if i != 0 {
		t.Fatalf("expected 0, got %d", i)
	}

	if err == nil {
		t.Fatal("expected an error but didn't recieve one")
	}

	switch errType := err.(type) {
	case *grafanaError:
		break
	default:
		t.Fatalf("expected grafanaError type, got %s", errType)
	}

	if !strings.Contains(err.Error(), "401") {
		t.Fatalf("expected error to contain '401', got '%s'", err)
	}
}
