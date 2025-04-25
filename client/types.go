package client

type Counts struct {
	Up     int `json:"up"`
	Down   int `json:"down"`
	Paused int `json:"paused"`
}
type DailyRatio struct {
	Ratio string `json:"ratio"`
	Label string `json:"label"`
}
type EventFeed struct {
	Status  bool     `json:"status"`
	Results []Result `json:"results"`
	Meta    Meta     `json:"meta"`
}
type Meta struct {
	Count int `json:"count"`
	Total int `json:"total"`
}
type Monitor struct {
	MonitorID           int          `json:"monitorId"`
	CreatedAt           int          `json:"createdAt"`
	StatusClass         string       `json:"statusClass"`
	Name                string       `json:"name"`
	URL                 any          `json:"url"`
	MonitorType         string       `json:"type"`
	DailyRatios         []DailyRatio `json:"dailyRatios"`
	NinetyDayRatio      DailyRatio   `json:"90dRatio"`
	ThirtyDayRatio      DailyRatio   `json:"30dRatio"`
	HasIncidentComments bool         `json:"hasIncidentComments"`
}
type MonitorList struct {
	Status     string     `json:"status"`
	PSP        PSP        `json:"psp"`
	Days       []string   `json:"days"`
	Statistics Statistics `json:"statistics"`
}
type PSP struct {
	PerPage       int       `json:"perPage"`
	TotalMonitors int       `json:"totalMonitors"`
	Monitors      []Monitor `json:"monitors"`
	Timezone      string    `json:"timezone"`
}
type Result struct {
	EventType  string `json:"eventType"`
	Icon       string `json:"icon"`
	Title      string `json:"title"`
	Content    string `json:"content"`
	Timestamp  int    `json:"ts"`
	Date       string `json:"date"`
	Time       string `json:"time"`
	TimeGMT    string `json:"timeGMT"`
	EndDate    string `json:"endDate"`
	EndDateGMT string `json:"endDateGMT"`
}
type Statistics struct {
	Uptime         map[string]DailyRatio `json:"uptime"`
	LatestDowntime any                   `json:"latest_downtime"`
	Counts         Counts                `json:"counts"`
	CountResult    string                `json:"count_result"`
}
