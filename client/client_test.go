package client

import (
	"log/slog"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"

	"github.com/DazWilkin/go-bsky-status/testdata"
)

// getTestServer is a function that creates an httptest.Server
// It takes Client because it needs the Client's getPath method
func getTestServer(t *testing.T, client *Client, logger *slog.Logger) *httptest.Server {
	t.Helper()

	logger = logger.With("test", "getTestServer")

	urlGetEventFeed := client.getPath(MethodGetEventFeed)
	urlGetMonitorList := client.getPath(MethodGetMonitorList)

	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var responseBody []byte

		if r.URL.Path == urlGetEventFeed {
			logger.Info("Handling GetEventFeed")
			responseBody = testdata.JsonEventFeed
		}
		if r.URL.Path == urlGetMonitorList {
			logger.Info("Handling GetMonitorList")
			responseBody = testdata.JsonMonitorList
		}

		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		if _, err := w.Write(responseBody); err != nil {
			logger.Error("failed to write response", "err", err)
		}
	}))
}

// setup is a function
func setup(t *testing.T) (client *Client, server *httptest.Server, teardown func(), logger *slog.Logger) {
	t.Helper()

	logger = slog.New(slog.NewJSONHandler(os.Stdout, nil))

	property := os.Getenv("PROPERTY")
	if property == "" {
		t.Errorf("expected 'PROPERTY' in the environment")
	}

	// Create the client first as it's required by the server
	client = NewClient(property, logger)

	// Create the HTTP test server
	server = getTestServer(t, client, logger)

	// Use the server's URL as the client's URL
	client.BaseURL = server.URL

	// Return a function that closes the server
	teardown = func() {
		server.Close()
	}

	return client, server, teardown, logger
}

// TestGetEventFeed tests the Client's GetEventFeed method
// It uses getTestServer to create an HTTP mock server
func TestGetEventFeed(t *testing.T) {
	client, _, teardown, _ := setup(t)
	defer teardown()

	from := time.Date(2025, time.April, 1, 0, 0, 0, 0, time.UTC)
	_, err := client.GetEventFeed(from)
	if err != nil {
		t.Errorf("expected success")
	}
}

// TestGetMonitorList tests the Client's GetMonitorList method
// It uses getTestServer to create an HTTP mock server
func TestGetMonitorList(t *testing.T) {
	client, _, teardown, _ := setup(t)
	defer teardown()

	_, err := client.GetMonitorList()
	if err != nil {
		t.Errorf("expected success")
	}

}
