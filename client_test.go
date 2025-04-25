package main

import (
	"log/slog"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"
)

// getTestServer is a function that creates an httptest.Server
// It takes Client because it needs the Client's getPath method
func getTestServer(client *Client, logger *slog.Logger) *httptest.Server {
	logger = logger.With("test", "getTestServer")
	
	urlGetEventFeed := client.getPath(MethodGetEventFeed)
	urlGetMonitorList := client.getPath(MethodGetMonitorList)

	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var responseBody []byte

		if r.URL.Path == urlGetEventFeed {
			logger.Info("Handling GetEventFeed")
			responseBody = jsonEventFeed
		}
		if r.URL.Path == urlGetMonitorList {
			logger.Info("Handling GetMonitorList")
			responseBody = jsonMonitorList
		}

		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		if _, err := w.Write(responseBody); err != nil {
			logger.Error("failed to write response", "err", err)
		}
	}))
}

// TestGetEventFeed tests the Client's GetEventFeed method
// It uses getTestServer to create an HTTP mock server
func TestGetEventFeed(t *testing.T) {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	// Create the client first as it's required by the server
	client := NewClient(property, logger)

	server := getTestServer(client, logger)
	defer server.Close()

	// Use the server's URL as the client's URL
	client.BaseURL = server.URL

	from := time.Date(2025, time.April, 1, 0, 0, 0, 0, time.UTC)
	_, err := client.GetEventFeed(from)
	if err != nil {
		t.Errorf("expected success")
	}
}

// TestGetMonitorList tests the Client's GetMonitorList method
// It uses getTestServer to create an HTTP mock server
func TestGetMonitorList(t *testing.T) {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	// Create the client first as it's required by the server
	client := NewClient(property, logger)

	server := getTestServer(client, logger)
	defer server.Close()

	// Use the server's URL as the client's URL
	client.BaseURL = server.URL

	_, err := client.GetMonitorList()
	if err != nil {
		t.Errorf("expected success")
	}

}
