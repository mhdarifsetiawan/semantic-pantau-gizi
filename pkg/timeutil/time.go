package timeutil

import "time"

var WIBLocation *time.Location

func init() {
	var err error
	WIBLocation, err = time.LoadLocation("Asia/Jakarta")
	if err != nil {
		WIBLocation = time.FixedZone("WIB", 7*60*60)
	}
}

func CorrectDatabaseTimeToUTC(fieldTime *time.Time) time.Time {
	wibTime := time.Date(
		fieldTime.Year(), fieldTime.Month(), fieldTime.Day(),
		fieldTime.Hour(), fieldTime.Minute(), fieldTime.Second(),
		fieldTime.Nanosecond(), WIBLocation,
	)

	return wibTime.UTC()
}

func IsSameDay(t1, t2 time.Time) bool {
	y1, m1, d1 := t1.Date()
	y2, m2, d2 := t2.Date()
	return y1 == y2 && m1 == m2 && d1 == d2
}

func GetCurrentTimestamp() string {
	return time.Now().In(WIBLocation).Format("20060102150405")
}
