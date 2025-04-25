package client

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/DazWilkin/go-bsky-status/testdata"
)

var (
	TestEventFeed = &EventFeed{
		Status: true,
		Results: []Result{
			{
				EventType:  "ann",
				Icon:       "alert-triangle",
				Title:      "Update on PDS Downtime",
				Content:    "We have identified a likely root cause and are rolling out a fix to the Bluesky PDS fleet.",
				Timestamp:  1745537880,
				Date:       "April 24, 2025",
				Time:       "23:38",
				TimeGMT:    "April 24, 2025, 23:38",
				EndDate:    "",
				EndDateGMT: "",
			},
			{
				EventType:  "ann",
				Icon:       "alert-triangle",
				Title:      "Major PDS Networking Problems",
				Content:    "We are investigating a major outage with Bluesky hosted PDS instances.",
				Timestamp:  1745535300,
				Date:       "April 24, 2025",
				Time:       "22:55",
				TimeGMT:    "April 24, 2025, 22:55",
				EndDate:    "",
				EndDateGMT: "",
			},
		},
		Meta: Meta{
			Count: 2,
			Total: 2,
		},
	}
	TestMonitorList = &MonitorList{}
)

func TestUnmarshalEventFeed(t *testing.T) {
	want := TestEventFeed

	got := &EventFeed{}
	if err := json.Unmarshal(testdata.JsonEventFeed, got); err != nil {
		t.Errorf("expected to unmarshal")
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got:\n%v\nwant:\n%v\n", got, want)
	}
}
func TestUnmarshalMonitorList(t *testing.T) {
	want := TestMonitorList

	got := &MonitorList{}
	if err := json.Unmarshal(testdata.JsonMonitorList, got); err != nil {
		t.Errorf("expected to unmarshal")
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got:\n%v\nwant:\n%v\n", got, want)
	}
}
