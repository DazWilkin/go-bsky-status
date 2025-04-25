package client

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"strconv"
	"time"
)

const (
	baseURL = "https://status.bsky.app"
)

type Method int

func (m Method) String() string {
	switch m {
	case MethodGetEventFeed:
		return "getEventFeed"
	case MethodGetMonitorList:
		return "getMonitorList"
	case MethodInvalid:
		fallthrough
	default:
		return ""
	}
}

const (
	MethodInvalid Method = iota
	MethodGetEventFeed
	MethodGetMonitorList
)

// Client is a type that represents a client of the Status service
type Client struct {
	// Public
	// Necessary for testing
	BaseURL string

	// Private
	httpClient *http.Client
	property   string
	logger     *slog.Logger
}

func NewClient(property string, logger *slog.Logger) *Client {
	client := &http.Client{
		Timeout: time.Second * 10,
	}

	return &Client{
		BaseURL:    baseURL,
		httpClient: client,
		property:   property,
		logger:     logger,
	}
}

// getPath is a method that returns the URL Path for a given method
// This method is used by the tests to define a mock server
func (c Client) getPath(method Method) string {
	path := fmt.Sprintf("/api/%s/%s", method.String(), c.property)
	return path
}

// getURL is a method that gets the URL for a given method
func (c Client) getURL(method Method) string {
	path := c.getPath(method)

	// Path includes prefixing "/"
	url := fmt.Sprintf("%s%s", c.BaseURL, path)
	return url
}

// GetEventFeed is a method that invokes /api/getEventFeed
// https://status.bsky.app/api/getEventFeed/zwOvMT8x16?from_date=1743465600
func (c Client) GetEventFeed(from time.Time) (*EventFeed, error) {
	method := MethodGetEventFeed
	logger := c.logger.With("method", method.String())
	logger.Info("Entered")
	defer logger.Info("Done")

	result := &EventFeed{}

	url := c.getURL(method)

	rqst, err := http.NewRequest("GET", url, nil)
	if err != nil {
		msg := "unable to create request"
		logger.Error(msg, "err", err)
		return result, errors.New(msg)
	}

	q := rqst.URL.Query()
	q.Add("from_date", strconv.FormatInt(from.Unix(), 10))

	rqst.URL.RawQuery = q.Encode()

	resp, err := c.httpClient.Do(rqst)
	if err != nil {
		msg := "unable to do request"
		logger.Error(msg, "err", err)
		return result, errors.New(msg)
	}
	defer func() {
		if err := resp.Body.Close(); err != nil {
			msg := "unable to close body"
			logger.Error(msg, "err", err)
		}
	}()

	if resp.StatusCode != http.StatusOK {
		msg := "response code not okay"
		logger.Error(msg,
			"code", resp.StatusCode,
		)
		return result, errors.New(msg)
	}

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		msg := "unable to read response body"
		logger.Error(msg, "err", err)
		return result, errors.New(msg)
	}

	if err := json.Unmarshal(b, result); err != nil {
		msg := "unable to unmarshal body"
		logger.Error(msg, "err", err)
		return result, errors.New(msg)
	}

	return result, nil
}

// GetMonitorList is a method that invokes /api/getMonitorList
// https://status.bsky.app/api/getMonitorList/zwOvMT8x16
func (c Client) GetMonitorList() (*MonitorList, error) {
	method := MethodGetMonitorList
	logger := c.logger.With("method", method.String())
	logger.Info("Entered")
	defer logger.Info("Done")

	result := &MonitorList{}

	url := c.getURL(method)

	rqst, err := http.NewRequest("GET", url, nil)
	if err != nil {
		msg := "unable to create request"
		logger.Error(msg, "err", err)
		return result, errors.New(msg)
	}

	resp, err := c.httpClient.Do(rqst)
	if err != nil {
		msg := "unable to do request"
		logger.Error(msg, "err", err)
		return result, errors.New(msg)
	}
	defer func() {
		if err := resp.Body.Close(); err != nil {
			msg := "unable to close body"
			logger.Error(msg, "err", err)
		}
	}()

	if resp.StatusCode != http.StatusOK {
		msg := "response code not OK"
		logger.Error(msg,
			"code", resp.StatusCode,
		)
		return result, errors.New(msg)
	}

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		msg := "unable to read response body"
		logger.Error(msg, "err", err)
		return result, errors.New(msg)
	}

	if err := json.Unmarshal(b, result); err != nil {
		msg := "unable to unmarshal body"
		logger.Error(msg, "err", err)
		return result, errors.New(msg)
	}

	return result, nil
}
