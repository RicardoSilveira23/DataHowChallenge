package data

import (
	"testing"
)

func TestCalcNumberOfUniqueIPs(t *testing.T){


	metricsLogList := []*MetricsLog{
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

	number := GetNumberOfUniqueIPs(metricsLogList)

	if number != 3{
		t.Fatal()
	}
	
}