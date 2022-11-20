package helper

import (
	"time"

	"github.com/golang/glog"
)

var (
	LocLocal *time.Location
)

const (
	FormatDateTime = "02-01-2006 15:04:05"

	FormatSonarDateTime = "2006-01-02T15:04:05+0000"
	EverestReleaseTime  = "29-08-2022 01:16:30"
)

func init() {
	l, err := time.LoadLocation("Asia/Bangkok")
	if err != nil {
		panic(err)
	}
	LocLocal = l
}

func TimeParseLocal(layout, value string) (time.Time, error) {
	return time.ParseInLocation(layout, value, LocLocal)
}

func IsGteEverestReleaseTime(scmDate string) bool {
	t2, err := TimeParseLocal(FormatSonarDateTime, scmDate)
	if err != nil {
		glog.Errorf("scmDate: %s with err: %v", scmDate, err)
		return false
	}
	t1, _ := TimeParseLocal(FormatDateTime, EverestReleaseTime) // not error
	return t2.Sub(t1) > 0
}

func rangeDate(start, end time.Time) func() time.Time {
	y, m, d := start.Date()
	start = time.Date(y, m, d, 0, 0, 0, 0, LocLocal)
	y, m, d = end.Date()
	end = time.Date(y, m, d, 0, 0, 0, 0, LocLocal)

	return func() time.Time {
		if start.After(end) {
			return time.Time{}
		}
		date := start
		start = start.AddDate(0, 0, 1)
		return date
	}
}

func GetDatesLastYearToToday(start time.Time, end time.Time) []time.Time {
	var res = make([]time.Time, 0)
	//end := time.Now()
	//start := end.AddDate(0, 0, -364)

	for rd := rangeDate(start, end); ; {
		date := rd()
		if date.IsZero() {
			break
		}
		res = append(res, date)
	}
	return res
}

func GetDateOnly(value time.Time) time.Time {
	year, month, day := value.Date()
	return time.Date(year, month, day, 0, 0, 0, 0, LocLocal)
}

func GetDateTimeOnly(value time.Time) time.Time {
	year, month, day := value.Date()
	h := value.Hour()
	m := value.Minute()
	s := value.Second()
	return time.Date(year, month, day, h, m, s, 0, LocLocal)
}
