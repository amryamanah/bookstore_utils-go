package date_utils

import "time"

const (
	//apiDateLayout = time.RFC3339
	apiDateLayout = "2006-01-02T15:04:05Z"
	apiLoc        = "UTC"
)

func GetNowString() string {
	now := time.Now()
	loc, _ := time.LoadLocation(apiLoc)
	return now.In(loc).Format(apiDateLayout)
}
