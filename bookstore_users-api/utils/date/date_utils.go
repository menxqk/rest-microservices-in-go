package date

import "time"

const (
	API_DATE_LAYOUT = "2006-01-02T15:04:05Z"
	API_DB_LAYOUT   = "2006-01-02 15:04:05"
)

func GetNow() time.Time {
	return time.Now().UTC()
}

func GetNowAsString() string {
	return GetNow().Format(API_DATE_LAYOUT)
}

func GetNowDBFormat() string {
	return GetNow().Format(API_DB_LAYOUT)
}

func ParseDateString(s string) (time.Time, error) {
	return time.Parse(API_DATE_LAYOUT, s)
}
