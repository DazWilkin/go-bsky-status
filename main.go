package main

import (
	"log/slog"
	"os"
	"time"

	"github.com/DazWilkin/go-bsky-status/client"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	property := os.Getenv("PROPERTY")
	if property == "" {
		slog.Error("expected 'PROPERTY' in the environment")
		return
	}

	c := client.NewClient(property, logger)

	eventfeed, err := c.GetEventFeed(time.Date(2025, time.April, 1, 0, 0, 0, 0, time.UTC))
	if err != nil {
		slog.Error("unable to GetEventFeed", "err", err)
		return
	}

	slog.Info("Result", "eventfeed", eventfeed)

	monitorlist, err := c.GetMonitorList()
	if err != nil {
		slog.Error("unable to GetMonitorList", "err", err)
		return
	}

	slog.Info("Result", "monitorlist", monitorlist)
}
