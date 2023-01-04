package date_utils

import (
	"time"
)

const apiDateLayout = "2006-01-02T15:04:05Z"
const localTimZone = "Africa/Nairobi"

func GetNowString() string {

	loc, _ := time.LoadLocation(localTimZone)

	//set timezone,
	now := time.Now().In(loc)
	return now.Format(apiDateLayout)

}

func GetNowDate() (time.Time, error) {

	loc, _ := time.LoadLocation(localTimZone)

	//set timezone,
	now := time.Now().In(loc)
	newNow := now.Format(apiDateLayout)
	output, err := time.Parse(time.RFC3339, newNow)
	if err != nil {
		return time.Now(), err
	}
	return output, nil

}
