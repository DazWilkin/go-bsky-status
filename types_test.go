package main

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestUnmarshalEventFeed(t *testing.T) {
	want := &EventFeed{}

	got := &EventFeed{}
	if err := json.Unmarshal(jsonEventFeed, got); err != nil {
		t.Errorf("expected to unmarshal")
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got:\n%v\nwant:\n%v\n", got, want)
	}
}
func TestUnmarshalMonitorList(t *testing.T) {
	want := &MonitorList{}

	got := &MonitorList{}
	if err := json.Unmarshal(jsonMonitorList, got); err != nil {
		t.Errorf("expected to unmarshal")
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got:\n%v\nwant:\n%v\n", got, want)
	}
}
