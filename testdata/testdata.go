package testdata

var (
	JsonEventFeed = []byte(`
{
  "status": true,
  "results": [
    {
      "eventType": "ann",
      "icon": "alert-triangle",
      "title": "Update on PDS Downtime",
      "content": "We have identified a likely root cause and are rolling out a fix to the Bluesky PDS fleet.",
      "ts": 1745537880,
      "date": "April 24, 2025",
      "time": "23:38",
      "timeGMT": "April 24, 2025, 23:38",
      "endDate": null,
      "endDateGMT": null
    },
    {
      "eventType": "ann",
      "icon": "alert-triangle",
      "title": "Major PDS Networking Problems",
      "content": "We are investigating a major outage with Bluesky hosted PDS instances.",
      "ts": 1745535300,
      "date": "April 24, 2025",
      "time": "22:55",
      "timeGMT": "April 24, 2025, 22:55",
      "endDate": null,
      "endDateGMT": null
    }
  ],
  "meta": {
    "count": 2,
    "total": 2
  }
}`)
	JsonMonitorList = []byte(`
{
  "status": "ok",
  "psp": {
    "perPage": 30,
    "totalMonitors": 52,
    "monitors": [
      {
        "monitorId": 796288296,
        "createdAt": 1706748387,
        "statusClass": "success",
        "name": "agaric.us-west.host.bsky.network",
        "url": null,
        "type": "HTTP(s)",
        "dailyRatios": [
          {
            "ratio": "99.546",
            "label": "warning"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "99.933",
            "label": "warning"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          }
        ],
        "90dRatio": {
          "ratio": "99.994",
          "label": "warning"
        },
        "30dRatio": {
          "ratio": "99.983",
          "label": "warning"
        },
        "hasIncidentComments": false
      },
      {
        "monitorId": 795277183,
        "createdAt": 1694622216,
        "statusClass": "black",
        "name": "api.events.bsky.app",
        "url": null,
        "type": "HTTP(s)",
        "dailyRatios": [
          {
            "ratio": "0.000",
            "label": "black"
          },
          {
            "ratio": "0.000",
            "label": "black"
          },
          {
            "ratio": "0.000",
            "label": "black"
          },
          {
            "ratio": "0.000",
            "label": "black"
          },
          {
            "ratio": "0.000",
            "label": "black"
          },
          {
            "ratio": "0.000",
            "label": "black"
          },
          {
            "ratio": "0.000",
            "label": "black"
          },
          {
            "ratio": "0.000",
            "label": "black"
          },
          {
            "ratio": "0.000",
            "label": "black"
          },
          {
            "ratio": "0.000",
            "label": "black"
          },
          {
            "ratio": "0.000",
            "label": "black"
          },
          {
            "ratio": "0.000",
            "label": "black"
          },
          {
            "ratio": "0.000",
            "label": "black"
          },
          {
            "ratio": "0.000",
            "label": "black"
          },
          {
            "ratio": "0.000",
            "label": "black"
          },
          {
            "ratio": "0.000",
            "label": "black"
          },
          {
            "ratio": "0.000",
            "label": "black"
          },
          {
            "ratio": "0.000",
            "label": "black"
          },
          {
            "ratio": "0.000",
            "label": "black"
          },
          {
            "ratio": "0.000",
            "label": "black"
          },
          {
            "ratio": "0.000",
            "label": "black"
          },
          {
            "ratio": "0.000",
            "label": "black"
          },
          {
            "ratio": "0.000",
            "label": "black"
          },
          {
            "ratio": "0.000",
            "label": "black"
          },
          {
            "ratio": "0.000",
            "label": "black"
          },
          {
            "ratio": "0.000",
            "label": "black"
          },
          {
            "ratio": "0.000",
            "label": "black"
          },
          {
            "ratio": "0.000",
            "label": "black"
          },
          {
            "ratio": "0.000",
            "label": "black"
          },
          {
            "ratio": "0.000",
            "label": "black"
          },
          {
            "ratio": "0.000",
            "label": "black"
          },
          {
            "ratio": "0.000",
            "label": "black"
          },
          {
            "ratio": "0.000",
            "label": "black"
          },
          {
            "ratio": "0.000",
            "label": "black"
          },
          {
            "ratio": "0.000",
            "label": "black"
          },
          {
            "ratio": "0.000",
            "label": "black"
          },
          {
            "ratio": "0.000",
            "label": "black"
          },
          {
            "ratio": "0.000",
            "label": "black"
          },
          {
            "ratio": "0.000",
            "label": "black"
          },
          {
            "ratio": "0.000",
            "label": "black"
          },
          {
            "ratio": "0.000",
            "label": "black"
          },
          {
            "ratio": "0.000",
            "label": "black"
          },
          {
            "ratio": "0.000",
            "label": "black"
          },
          {
            "ratio": "0.000",
            "label": "black"
          },
          {
            "ratio": "0.000",
            "label": "black"
          },
          {
            "ratio": "0.000",
            "label": "black"
          },
          {
            "ratio": "0.000",
            "label": "black"
          },
          {
            "ratio": "0.000",
            "label": "black"
          },
          {
            "ratio": "0.000",
            "label": "black"
          },
          {
            "ratio": "0.000",
            "label": "black"
          },
          {
            "ratio": "0.000",
            "label": "black"
          },
          {
            "ratio": "0.000",
            "label": "black"
          },
          {
            "ratio": "0.000",
            "label": "black"
          },
          {
            "ratio": "0.000",
            "label": "black"
          },
          {
            "ratio": "0.000",
            "label": "black"
          },
          {
            "ratio": "0.000",
            "label": "black"
          },
          {
            "ratio": "0.000",
            "label": "black"
          },
          {
            "ratio": "0.000",
            "label": "black"
          },
          {
            "ratio": "0.000",
            "label": "black"
          },
          {
            "ratio": "0.000",
            "label": "black"
          },
          {
            "ratio": "0.000",
            "label": "black"
          },
          {
            "ratio": "0.000",
            "label": "black"
          },
          {
            "ratio": "0.000",
            "label": "black"
          },
          {
            "ratio": "0.000",
            "label": "black"
          },
          {
            "ratio": "0.000",
            "label": "black"
          },
          {
            "ratio": "0.000",
            "label": "black"
          },
          {
            "ratio": "0.000",
            "label": "black"
          },
          {
            "ratio": "0.000",
            "label": "black"
          },
          {
            "ratio": "0.000",
            "label": "black"
          },
          {
            "ratio": "0.000",
            "label": "black"
          },
          {
            "ratio": "0.000",
            "label": "black"
          },
          {
            "ratio": "0.000",
            "label": "black"
          },
          {
            "ratio": "0.000",
            "label": "black"
          },
          {
            "ratio": "0.000",
            "label": "black"
          },
          {
            "ratio": "0.000",
            "label": "black"
          },
          {
            "ratio": "0.000",
            "label": "black"
          },
          {
            "ratio": "0.000",
            "label": "black"
          },
          {
            "ratio": "0.000",
            "label": "black"
          },
          {
            "ratio": "0.000",
            "label": "black"
          },
          {
            "ratio": "0.000",
            "label": "black"
          },
          {
            "ratio": "0.000",
            "label": "black"
          },
          {
            "ratio": "0.000",
            "label": "black"
          },
          {
            "ratio": "0.000",
            "label": "black"
          },
          {
            "ratio": "0.000",
            "label": "black"
          },
          {
            "ratio": "0.000",
            "label": "black"
          },
          {
            "ratio": "0.000",
            "label": "black"
          },
          {
            "ratio": "0.000",
            "label": "black"
          },
          {
            "ratio": "0.000",
            "label": "black"
          },
          {
            "ratio": "0.000",
            "label": "black"
          },
          {
            "ratio": "0.000",
            "label": "black"
          }
        ],
        "90dRatio": {
          "ratio": "0.000",
          "label": "black"
        },
        "30dRatio": {
          "ratio": "0.000",
          "label": "black"
        },
        "hasIncidentComments": false
      },
      {
        "monitorId": 796270522,
        "createdAt": 1706555794,
        "statusClass": "success",
        "name": "api.pop1.bsky.app",
        "url": null,
        "type": "HTTP(s)",
        "dailyRatios": [
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "99.762",
            "label": "warning"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          }
        ],
        "90dRatio": {
          "ratio": "99.997",
          "label": "warning"
        },
        "30dRatio": {
          "ratio": "100.00",
          "label": "success"
        },
        "hasIncidentComments": false
      },
      {
        "monitorId": 796270524,
        "createdAt": 1706555816,
        "statusClass": "success",
        "name": "api.pop2.bsky.app",
        "url": null,
        "type": "HTTP(s)",
        "dailyRatios": [
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          }
        ],
        "90dRatio": {
          "ratio": "100.00",
          "label": "success"
        },
        "30dRatio": {
          "ratio": "100.00",
          "label": "success"
        },
        "hasIncidentComments": false
      },
      {
        "monitorId": 794203327,
        "createdAt": 1682063868,
        "statusClass": "success",
        "name": "atproto.com",
        "url": null,
        "type": "HTTP(s)",
        "dailyRatios": [
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "99.962",
            "label": "warning"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "93.488",
            "label": "warning"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "99.476",
            "label": "warning"
          },
          {
            "ratio": "99.925",
            "label": "warning"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "99.956",
            "label": "warning"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          }
        ],
        "90dRatio": {
          "ratio": "99.920",
          "label": "warning"
        },
        "30dRatio": {
          "ratio": "99.762",
          "label": "warning"
        },
        "hasIncidentComments": false
      },
      {
        "monitorId": 796288289,
        "createdAt": 1706748318,
        "statusClass": "success",
        "name": "blewit.us-west.host.bsky.network",
        "url": null,
        "type": "HTTP(s)",
        "dailyRatios": [
          {
            "ratio": "99.425",
            "label": "warning"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "95.767",
            "label": "warning"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          }
        ],
        "90dRatio": {
          "ratio": "99.947",
          "label": "warning"
        },
        "30dRatio": {
          "ratio": "99.840",
          "label": "warning"
        },
        "hasIncidentComments": false
      },
      {
        "monitorId": 796512805,
        "createdAt": 1709657621,
        "statusClass": "success",
        "name": "blobber1..us-west.cdn.bsky.network",
        "url": null,
        "type": "HTTP(s)",
        "dailyRatios": [
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "99.884",
            "label": "warning"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "99.444",
            "label": "warning"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          }
        ],
        "90dRatio": {
          "ratio": "99.993",
          "label": "warning"
        },
        "30dRatio": {
          "ratio": "99.996",
          "label": "warning"
        },
        "hasIncidentComments": false
      },
      {
        "monitorId": 796512771,
        "createdAt": 1709657319,
        "statusClass": "success",
        "name": "blobber1.us-east.cdn.bsky.network",
        "url": null,
        "type": "HTTP(s)",
        "dailyRatios": [
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          }
        ],
        "90dRatio": {
          "ratio": "100.00",
          "label": "success"
        },
        "30dRatio": {
          "ratio": "100.00",
          "label": "success"
        },
        "hasIncidentComments": false
      },
      {
        "monitorId": 796512775,
        "createdAt": 1709657343,
        "statusClass": "success",
        "name": "blobber2.us-east.cdn.bsky.network",
        "url": null,
        "type": "HTTP(s)",
        "dailyRatios": [
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          }
        ],
        "90dRatio": {
          "ratio": "100.00",
          "label": "success"
        },
        "30dRatio": {
          "ratio": "100.00",
          "label": "success"
        },
        "hasIncidentComments": false
      },
      {
        "monitorId": 796512808,
        "createdAt": 1709657639,
        "statusClass": "success",
        "name": "blobber2.us-west.cdn.bsky.network",
        "url": null,
        "type": "HTTP(s)",
        "dailyRatios": [
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "99.950",
            "label": "warning"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "99.763",
            "label": "warning"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          }
        ],
        "90dRatio": {
          "ratio": "99.997",
          "label": "warning"
        },
        "30dRatio": {
          "ratio": "99.998",
          "label": "warning"
        },
        "hasIncidentComments": false
      },
      {
        "monitorId": 796512778,
        "createdAt": 1709657360,
        "statusClass": "success",
        "name": "blobber3.us-east.cdn.bsky.network",
        "url": null,
        "type": "HTTP(s)",
        "dailyRatios": [
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          }
        ],
        "90dRatio": {
          "ratio": "100.00",
          "label": "success"
        },
        "30dRatio": {
          "ratio": "100.00",
          "label": "success"
        },
        "hasIncidentComments": false
      },
      {
        "monitorId": 796512809,
        "createdAt": 1709657653,
        "statusClass": "success",
        "name": "blobber3.us-west.cdn.bsky.network",
        "url": null,
        "type": "HTTP(s)",
        "dailyRatios": [
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "99.884",
            "label": "warning"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "99.931",
            "label": "warning"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          }
        ],
        "90dRatio": {
          "ratio": "99.998",
          "label": "warning"
        },
        "30dRatio": {
          "ratio": "99.996",
          "label": "warning"
        },
        "hasIncidentComments": false
      },
      {
        "monitorId": 794729006,
        "createdAt": 1688510816,
        "statusClass": "success",
        "name": "Bluesky Account",
        "url": null,
        "type": "HTTP(s)",
        "dailyRatios": [
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "99.947",
            "label": "warning"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          }
        ],
        "90dRatio": {
          "ratio": "99.999",
          "label": "warning"
        },
        "30dRatio": {
          "ratio": "99.998",
          "label": "warning"
        },
        "hasIncidentComments": false
      },
      {
        "monitorId": 794203409,
        "createdAt": 1682064443,
        "statusClass": "success",
        "name": "blueskyweb.org",
        "url": null,
        "type": "HTTP(s)",
        "dailyRatios": [
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "99.785",
            "label": "warning"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "99.442",
            "label": "warning"
          },
          {
            "ratio": "99.259",
            "label": "warning"
          },
          {
            "ratio": "99.476",
            "label": "warning"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          }
        ],
        "90dRatio": {
          "ratio": "99.977",
          "label": "warning"
        },
        "30dRatio": {
          "ratio": "99.993",
          "label": "warning"
        },
        "hasIncidentComments": false
      },
      {
        "monitorId": 794203315,
        "createdAt": 1682063773,
        "statusClass": "success",
        "name": "blueskyweb.xyz",
        "url": null,
        "type": "HTTP(s)",
        "dailyRatios": [
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "99.914",
            "label": "warning"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "99.525",
            "label": "warning"
          },
          {
            "ratio": "99.531",
            "label": "warning"
          },
          {
            "ratio": "99.032",
            "label": "warning"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "99.838",
            "label": "warning"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          }
        ],
        "90dRatio": {
          "ratio": "99.976",
          "label": "warning"
        },
        "30dRatio": {
          "ratio": "99.997",
          "label": "warning"
        },
        "hasIncidentComments": false
      },
      {
        "monitorId": 796288314,
        "createdAt": 1706748519,
        "statusClass": "success",
        "name": "boletus.us-west.host.bsky.network",
        "url": null,
        "type": "HTTP(s)",
        "dailyRatios": [
          {
            "ratio": "99.871",
            "label": "warning"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "99.887",
            "label": "warning"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "99.931",
            "label": "warning"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          }
        ],
        "90dRatio": {
          "ratio": "99.997",
          "label": "warning"
        },
        "30dRatio": {
          "ratio": "99.992",
          "label": "warning"
        },
        "hasIncidentComments": false
      },
      {
        "monitorId": 797864563,
        "createdAt": 1729236409,
        "statusClass": "success",
        "name": "bracket.us-west.host.bsky.network/xrpc/_health",
        "url": null,
        "type": "HTTP(s)",
        "dailyRatios": [
          {
            "ratio": "99.612",
            "label": "warning"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "99.885",
            "label": "warning"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          }
        ],
        "90dRatio": {
          "ratio": "99.994",
          "label": "warning"
        },
        "30dRatio": {
          "ratio": "99.983",
          "label": "warning"
        },
        "hasIncidentComments": false
      },
      {
        "monitorId": 794203332,
        "createdAt": 1682063899,
        "statusClass": "success",
        "name": "bsky.app",
        "url": null,
        "type": "HTTP(s)",
        "dailyRatios": [
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "99.884",
            "label": "warning"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "99.687",
            "label": "warning"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "99.696",
            "label": "warning"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "99.735",
            "label": "warning"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "99.904",
            "label": "warning"
          },
          {
            "ratio": "100.00",
            "label": "success"
          }
        ],
        "90dRatio": {
          "ratio": "99.988",
          "label": "warning"
        },
        "30dRatio": {
          "ratio": "100.00",
          "label": "success"
        },
        "hasIncidentComments": false
      },
      {
        "monitorId": 797864566,
        "createdAt": 1729236431,
        "statusClass": "success",
        "name": "button.us-west.host.bsky.network/xrpc/_health",
        "url": null,
        "type": "HTTP(s)",
        "dailyRatios": [
          {
            "ratio": "99.744",
            "label": "warning"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "99.845",
            "label": "warning"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "99.958",
            "label": "warning"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "99.925",
            "label": "warning"
          },
          {
            "ratio": "99.925",
            "label": "warning"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "99.921",
            "label": "warning"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "99.931",
            "label": "warning"
          },
          {
            "ratio": "99.793",
            "label": "warning"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "99.928",
            "label": "warning"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "99.934",
            "label": "warning"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "99.932",
            "label": "warning"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          }
        ],
        "90dRatio": {
          "ratio": "99.987",
          "label": "warning"
        },
        "30dRatio": {
          "ratio": "99.968",
          "label": "warning"
        },
        "hasIncidentComments": false
      },
      {
        "monitorId": 796288317,
        "createdAt": 1706748542,
        "statusClass": "danger",
        "name": "chaga.us-west.host.bsky.network",
        "url": null,
        "type": "HTTP(s)",
        "dailyRatios": [
          {
            "ratio": "99.565",
            "label": "warning"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "99.885",
            "label": "warning"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          }
        ],
        "90dRatio": {
          "ratio": "99.994",
          "label": "warning"
        },
        "30dRatio": {
          "ratio": "99.982",
          "label": "warning"
        },
        "hasIncidentComments": false
      },
      {
        "monitorId": 797864567,
        "createdAt": 1729236444,
        "statusClass": "success",
        "name": "chanterelle.us-west.host.bsky.network/xrpc/_health",
        "url": null,
        "type": "HTTP(s)",
        "dailyRatios": [
          {
            "ratio": "99.619",
            "label": "warning"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "99.932",
            "label": "warning"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "99.920",
            "label": "warning"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "99.878",
            "label": "warning"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "99.740",
            "label": "warning"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "99.792",
            "label": "warning"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "99.931",
            "label": "warning"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "99.931",
            "label": "warning"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "99.931",
            "label": "warning"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          }
        ],
        "90dRatio": {
          "ratio": "99.985",
          "label": "warning"
        },
        "30dRatio": {
          "ratio": "99.963",
          "label": "warning"
        },
        "hasIncidentComments": false
      },
      {
        "monitorId": 796288319,
        "createdAt": 1706748563,
        "statusClass": "success",
        "name": "conocybe.us-west.host.bsky.network",
        "url": null,
        "type": "HTTP(s)",
        "dailyRatios": [
          {
            "ratio": "99.416",
            "label": "warning"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "99.884",
            "label": "warning"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "99.926",
            "label": "warning"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          }
        ],
        "90dRatio": {
          "ratio": "99.991",
          "label": "warning"
        },
        "30dRatio": {
          "ratio": "99.974",
          "label": "warning"
        },
        "hasIncidentComments": false
      },
      {
        "monitorId": 797863383,
        "createdAt": 1729211449,
        "statusClass": "danger",
        "name": "coral.us-east.host.bsky.network/xrpc/_health",
        "url": null,
        "type": "HTTP(s)",
        "dailyRatios": [
          {
            "ratio": "97.018",
            "label": "warning"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          }
        ],
        "90dRatio": {
          "ratio": "99.967",
          "label": "warning"
        },
        "30dRatio": {
          "ratio": "99.902",
          "label": "warning"
        },
        "hasIncidentComments": false
      },
      {
        "monitorId": 797863396,
        "createdAt": 1729211622,
        "statusClass": "success",
        "name": "cordyceps.us-west.host.bsky.network/xrpc/_health",
        "url": null,
        "type": "HTTP(s)",
        "dailyRatios": [
          {
            "ratio": "99.549",
            "label": "warning"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "99.956",
            "label": "warning"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "99.920",
            "label": "warning"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          }
        ],
        "90dRatio": {
          "ratio": "99.994",
          "label": "warning"
        },
        "30dRatio": {
          "ratio": "99.984",
          "label": "warning"
        },
        "hasIncidentComments": false
      },
      {
        "monitorId": 797864569,
        "createdAt": 1729236462,
        "statusClass": "success",
        "name": "cremini.us-west.host.bsky.network/xrpc/_health",
        "url": null,
        "type": "HTTP(s)",
        "dailyRatios": [
          {
            "ratio": "99.462",
            "label": "warning"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "99.934",
            "label": "warning"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "99.961",
            "label": "warning"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "99.851",
            "label": "warning"
          },
          {
            "ratio": "99.851",
            "label": "warning"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "99.861",
            "label": "warning"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "99.931",
            "label": "warning"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "99.931",
            "label": "warning"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "99.932",
            "label": "warning"
          },
          {
            "ratio": "99.932",
            "label": "warning"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "99.932",
            "label": "warning"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "99.927",
            "label": "warning"
          },
          {
            "ratio": "99.722",
            "label": "warning"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          }
        ],
        "90dRatio": {
          "ratio": "99.980",
          "label": "warning"
        },
        "30dRatio": {
          "ratio": "99.964",
          "label": "warning"
        },
        "hasIncidentComments": false
      },
      {
        "monitorId": 796267619,
        "createdAt": 1706539912,
        "statusClass": "success",
        "name": "discover.bsky.app",
        "url": null,
        "type": "HTTP(s)",
        "dailyRatios": [
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "99.947",
            "label": "warning"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "99.921",
            "label": "warning"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          }
        ],
        "90dRatio": {
          "ratio": "99.999",
          "label": "warning"
        },
        "30dRatio": {
          "ratio": "99.998",
          "label": "warning"
        },
        "hasIncidentComments": false
      },
      {
        "monitorId": 796746678,
        "createdAt": 1713193574,
        "statusClass": "success",
        "name": "embed.bsky.app",
        "url": null,
        "type": "HTTP(s)",
        "dailyRatios": [
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          }
        ],
        "90dRatio": {
          "ratio": "100.00",
          "label": "success"
        },
        "30dRatio": {
          "ratio": "100.00",
          "label": "success"
        },
        "hasIncidentComments": false
      },
      {
        "monitorId": 795638023,
        "createdAt": 1699580198,
        "statusClass": "danger",
        "name": "enoki.us-east.host.bsky.network",
        "url": null,
        "type": "HTTP(s)",
        "dailyRatios": [
          {
            "ratio": "97.853",
            "label": "warning"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "99.633",
            "label": "warning"
          },
          {
            "ratio": "99.317",
            "label": "warning"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          }
        ],
        "90dRatio": {
          "ratio": "99.965",
          "label": "warning"
        },
        "30dRatio": {
          "ratio": "99.894",
          "label": "warning"
        },
        "hasIncidentComments": false
      },
      {
        "monitorId": 797863398,
        "createdAt": 1729211639,
        "statusClass": "danger",
        "name": "ganoderma.us-west.host.bsky.network/xrpc/_health",
        "url": null,
        "type": "HTTP(s)",
        "dailyRatios": [
          {
            "ratio": "99.594",
            "label": "warning"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "99.951",
            "label": "warning"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          }
        ],
        "90dRatio": {
          "ratio": "99.995",
          "label": "warning"
        },
        "30dRatio": {
          "ratio": "99.985",
          "label": "warning"
        },
        "hasIncidentComments": false
      },
      {
        "monitorId": 796789653,
        "createdAt": 1713817157,
        "statusClass": "success",
        "name": "gifs.bsky.app",
        "url": null,
        "type": "HTTP(s)",
        "dailyRatios": [
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          },
          {
            "ratio": "100.00",
            "label": "success"
          }
        ],
        "90dRatio": {
          "ratio": "100.00",
          "label": "success"
        },
        "30dRatio": {
          "ratio": "100.00",
          "label": "success"
        },
        "hasIncidentComments": false
      }
    ],
    "timezone": "+00:00"
  },
  "days": [
    "2025-04-24",
    "2025-04-23",
    "2025-04-22",
    "2025-04-21",
    "2025-04-20",
    "2025-04-19",
    "2025-04-18",
    "2025-04-17",
    "2025-04-16",
    "2025-04-15",
    "2025-04-14",
    "2025-04-13",
    "2025-04-12",
    "2025-04-11",
    "2025-04-10",
    "2025-04-09",
    "2025-04-08",
    "2025-04-07",
    "2025-04-06",
    "2025-04-05",
    "2025-04-04",
    "2025-04-03",
    "2025-04-02",
    "2025-04-01",
    "2025-03-31",
    "2025-03-30",
    "2025-03-29",
    "2025-03-28",
    "2025-03-27",
    "2025-03-26",
    "2025-03-25",
    "2025-03-24",
    "2025-03-23",
    "2025-03-22",
    "2025-03-21",
    "2025-03-20",
    "2025-03-19",
    "2025-03-18",
    "2025-03-17",
    "2025-03-16",
    "2025-03-15",
    "2025-03-14",
    "2025-03-13",
    "2025-03-12",
    "2025-03-11",
    "2025-03-10",
    "2025-03-09",
    "2025-03-08",
    "2025-03-07",
    "2025-03-06",
    "2025-03-05",
    "2025-03-04",
    "2025-03-03",
    "2025-03-02",
    "2025-03-01",
    "2025-02-28",
    "2025-02-27",
    "2025-02-26",
    "2025-02-25",
    "2025-02-24",
    "2025-02-23",
    "2025-02-22",
    "2025-02-21",
    "2025-02-20",
    "2025-02-19",
    "2025-02-18",
    "2025-02-17",
    "2025-02-16",
    "2025-02-15",
    "2025-02-14",
    "2025-02-13",
    "2025-02-12",
    "2025-02-11",
    "2025-02-10",
    "2025-02-09",
    "2025-02-08",
    "2025-02-07",
    "2025-02-06",
    "2025-02-05",
    "2025-02-04",
    "2025-02-03",
    "2025-02-02",
    "2025-02-01",
    "2025-01-31",
    "2025-01-30",
    "2025-01-29",
    "2025-01-28",
    "2025-01-27",
    "2025-01-26",
    "2025-01-25"
  ],
  "statistics": {
    "uptime": {
      "l1": {
        "label": "warning",
        "ratio": "99.443"
      },
      "l7": {
        "label": "warning",
        "ratio": "99.889"
      },
      "l30": {
        "label": "warning",
        "ratio": "99.965"
      },
      "l90": {
        "label": "warning",
        "ratio": "99.985"
      }
    },
    "latest_downtime": null,
    "counts": {
      "up": 42,
      "down": 9,
      "paused": 1
    },
    "count_result": "9 monitors down"
  }
}`)
)
