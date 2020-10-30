package data

import (
	"encoding/json"
	"io"
)

type MetricsLog struct {
	ID        int    `json:"id"`
	TimeStamp string `json:"timestamp"`
	IP        string `json:"ip"`
	URL       string `json:"url"`
}

// Metric Logs
type MetricsLogs []*MetricsLog

// Unique Ids
type UniqueIPs struct {
	UniqueIPAddresses int `json:"unique_ip_addresses"`
}

// FromJSON decoding
func (ml *MetricsLog) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(ml)
}

// ToJSON enconding
func (ml *MetricsLogs) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(ml)
}

// ToJSON enconding
func (ui *UniqueIPs) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(ui)
}

// Add MetricsLog
func AddMetricsLog(ml *MetricsLog) {
	ml.ID = getNextID()
	metricsLogList = append(metricsLogList, ml)
}

// Simple ID Auto Generator
func getNextID() int {
	ml := metricsLogList[len(metricsLogList)-1]
	return ml.ID + 1
}

// Simple get All metrics
func GetMetrics() MetricsLogs {
	return metricsLogList
}

// Calc unique IP Addresses
func GetUniqueIPAdrresses() UniqueIPs {

	uIPs := UniqueIPs{
		UniqueIPAddresses: GetNumberOfUniqueIPs(metricsLogList),
	}

	return uIPs
}

func GetNumberOfUniqueIPs(metricsList []*MetricsLog) int {
	var unique []*MetricsLog

	for _, i := range metricsLogList {
		skip := false
		for _, u := range unique {
			if i.IP == u.IP {
				skip = true
				break
			}
		}
		if !skip {
			unique = append(unique, i)
		}
	}

	return len(unique)
}

// Mock Data
var metricsLogList = []*MetricsLog{
	{
		ID:        1,
		TimeStamp: "2020-06-24-T15:27:00.123456Z",
		IP:        "83.150.59.250",
		URL:       "www.test.com",
	},
	{
		ID:        2,
		TimeStamp: "2020-08-24-T15:27:00.123456Z",
		IP:        "83.151.59.250",
		URL:       "www.test.com",
	},
	{
		ID:        3,
		TimeStamp: "2020-07-24-T15:27:00.123456Z",
		IP:        "83.152.59.250",
		URL:       "www.test.com",
	},
	{
		ID:        4,
		TimeStamp: "2020-05-24-T15:27:00.123456Z",
		IP:        "83.150.59.250",
		URL:       "www.test.com",
	},
}
