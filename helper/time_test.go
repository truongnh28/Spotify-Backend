package helper

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestIsGteEverestReleaseTime(t *testing.T) {
	tests := []struct {
		name    string
		scmDate string
		want    bool
	}{
		{
			name:    "invalid scmDate",
			scmDate: "9999/99/99",
			want:    false,
		},
		{
			name:    "success",
			scmDate: "2022-12-02T12:34:56+0000",
			want:    true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			if got := IsGteEverestReleaseTime(tt.scmDate); got != tt.want {
				t.Errorf("IsGteEverestReleaseTime() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_GetDateTimeOnly(t *testing.T) {
	tests := []struct {
		name  string
		input time.Time
		want  time.Time
	}{
		{
			name:  "success",
			input: time.Date(2022, 10, 4, 23, 19, 12, 123, LocLocal),
			want:  time.Date(2022, 10, 4, 23, 19, 12, 0, LocLocal),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			if got := GetDateTimeOnly(tt.input); got != tt.want {
				t.Errorf("GetDateTimeOnly() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_GetDateOnly(t *testing.T) {
	tests := []struct {
		name  string
		input time.Time
		want  time.Time
	}{
		{
			name:  "success",
			input: time.Date(2022, 10, 4, 23, 19, 12, 123, LocLocal),
			want:  time.Date(2022, 10, 4, 0, 0, 0, 0, LocLocal),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			if got := GetDateOnly(tt.input); got != tt.want {
				t.Errorf("GetDateTimeOnly() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_GetDatesLastYearToToday(t *testing.T) {
	var (
		end   = time.Date(2022, 10, 4, 0, 0, 0, 0, LocLocal)
		start = end.AddDate(0, 0, -364)
	)
	tests := []struct {
		name string
		want []time.Time
	}{
		{
			name: "success",
			want: []time.Time{time.Date(2021, time.October, 5, 0, 0, 0, 0, LocLocal), time.Date(2021, time.October, 6, 0, 0, 0, 0, LocLocal), time.Date(2021, time.October, 7, 0, 0, 0, 0, LocLocal), time.Date(2021, time.October, 8, 0, 0, 0, 0, LocLocal), time.Date(2021, time.October, 9, 0, 0, 0, 0, LocLocal), time.Date(2021, time.October, 10, 0, 0, 0, 0, LocLocal), time.Date(2021, time.October, 11, 0, 0, 0, 0, LocLocal), time.Date(2021, time.October, 12, 0, 0, 0, 0, LocLocal), time.Date(2021, time.October, 13, 0, 0, 0, 0, LocLocal), time.Date(2021, time.October, 14, 0, 0, 0, 0, LocLocal), time.Date(2021, time.October, 15, 0, 0, 0, 0, LocLocal), time.Date(2021, time.October, 16, 0, 0, 0, 0, LocLocal), time.Date(2021, time.October, 17, 0, 0, 0, 0, LocLocal), time.Date(2021, time.October, 18, 0, 0, 0, 0, LocLocal), time.Date(2021, time.October, 19, 0, 0, 0, 0, LocLocal), time.Date(2021, time.October, 20, 0, 0, 0, 0, LocLocal), time.Date(2021, time.October, 21, 0, 0, 0, 0, LocLocal), time.Date(2021, time.October, 22, 0, 0, 0, 0, LocLocal), time.Date(2021, time.October, 23, 0, 0, 0, 0, LocLocal), time.Date(2021, time.October, 24, 0, 0, 0, 0, LocLocal), time.Date(2021, time.October, 25, 0, 0, 0, 0, LocLocal), time.Date(2021, time.October, 26, 0, 0, 0, 0, LocLocal), time.Date(2021, time.October, 27, 0, 0, 0, 0, LocLocal), time.Date(2021, time.October, 28, 0, 0, 0, 0, LocLocal), time.Date(2021, time.October, 29, 0, 0, 0, 0, LocLocal), time.Date(2021, time.October, 30, 0, 0, 0, 0, LocLocal), time.Date(2021, time.October, 31, 0, 0, 0, 0, LocLocal), time.Date(2021, time.November, 1, 0, 0, 0, 0, LocLocal), time.Date(2021, time.November, 2, 0, 0, 0, 0, LocLocal), time.Date(2021, time.November, 3, 0, 0, 0, 0, LocLocal), time.Date(2021, time.November, 4, 0, 0, 0, 0, LocLocal), time.Date(2021, time.November, 5, 0, 0, 0, 0, LocLocal), time.Date(2021, time.November, 6, 0, 0, 0, 0, LocLocal), time.Date(2021, time.November, 7, 0, 0, 0, 0, LocLocal), time.Date(2021, time.November, 8, 0, 0, 0, 0, LocLocal), time.Date(2021, time.November, 9, 0, 0, 0, 0, LocLocal), time.Date(2021, time.November, 10, 0, 0, 0, 0, LocLocal), time.Date(2021, time.November, 11, 0, 0, 0, 0, LocLocal), time.Date(2021, time.November, 12, 0, 0, 0, 0, LocLocal), time.Date(2021, time.November, 13, 0, 0, 0, 0, LocLocal), time.Date(2021, time.November, 14, 0, 0, 0, 0, LocLocal), time.Date(2021, time.November, 15, 0, 0, 0, 0, LocLocal), time.Date(2021, time.November, 16, 0, 0, 0, 0, LocLocal), time.Date(2021, time.November, 17, 0, 0, 0, 0, LocLocal), time.Date(2021, time.November, 18, 0, 0, 0, 0, LocLocal), time.Date(2021, time.November, 19, 0, 0, 0, 0, LocLocal), time.Date(2021, time.November, 20, 0, 0, 0, 0, LocLocal), time.Date(2021, time.November, 21, 0, 0, 0, 0, LocLocal), time.Date(2021, time.November, 22, 0, 0, 0, 0, LocLocal), time.Date(2021, time.November, 23, 0, 0, 0, 0, LocLocal), time.Date(2021, time.November, 24, 0, 0, 0, 0, LocLocal), time.Date(2021, time.November, 25, 0, 0, 0, 0, LocLocal), time.Date(2021, time.November, 26, 0, 0, 0, 0, LocLocal), time.Date(2021, time.November, 27, 0, 0, 0, 0, LocLocal), time.Date(2021, time.November, 28, 0, 0, 0, 0, LocLocal), time.Date(2021, time.November, 29, 0, 0, 0, 0, LocLocal), time.Date(2021, time.November, 30, 0, 0, 0, 0, LocLocal), time.Date(2021, time.December, 1, 0, 0, 0, 0, LocLocal), time.Date(2021, time.December, 2, 0, 0, 0, 0, LocLocal), time.Date(2021, time.December, 3, 0, 0, 0, 0, LocLocal), time.Date(2021, time.December, 4, 0, 0, 0, 0, LocLocal), time.Date(2021, time.December, 5, 0, 0, 0, 0, LocLocal), time.Date(2021, time.December, 6, 0, 0, 0, 0, LocLocal), time.Date(2021, time.December, 7, 0, 0, 0, 0, LocLocal), time.Date(2021, time.December, 8, 0, 0, 0, 0, LocLocal), time.Date(2021, time.December, 9, 0, 0, 0, 0, LocLocal), time.Date(2021, time.December, 10, 0, 0, 0, 0, LocLocal), time.Date(2021, time.December, 11, 0, 0, 0, 0, LocLocal), time.Date(2021, time.December, 12, 0, 0, 0, 0, LocLocal), time.Date(2021, time.December, 13, 0, 0, 0, 0, LocLocal), time.Date(2021, time.December, 14, 0, 0, 0, 0, LocLocal), time.Date(2021, time.December, 15, 0, 0, 0, 0, LocLocal), time.Date(2021, time.December, 16, 0, 0, 0, 0, LocLocal), time.Date(2021, time.December, 17, 0, 0, 0, 0, LocLocal), time.Date(2021, time.December, 18, 0, 0, 0, 0, LocLocal), time.Date(2021, time.December, 19, 0, 0, 0, 0, LocLocal), time.Date(2021, time.December, 20, 0, 0, 0, 0, LocLocal), time.Date(2021, time.December, 21, 0, 0, 0, 0, LocLocal), time.Date(2021, time.December, 22, 0, 0, 0, 0, LocLocal), time.Date(2021, time.December, 23, 0, 0, 0, 0, LocLocal), time.Date(2021, time.December, 24, 0, 0, 0, 0, LocLocal), time.Date(2021, time.December, 25, 0, 0, 0, 0, LocLocal), time.Date(2021, time.December, 26, 0, 0, 0, 0, LocLocal), time.Date(2021, time.December, 27, 0, 0, 0, 0, LocLocal), time.Date(2021, time.December, 28, 0, 0, 0, 0, LocLocal), time.Date(2021, time.December, 29, 0, 0, 0, 0, LocLocal), time.Date(2021, time.December, 30, 0, 0, 0, 0, LocLocal), time.Date(2021, time.December, 31, 0, 0, 0, 0, LocLocal), time.Date(2022, time.January, 1, 0, 0, 0, 0, LocLocal), time.Date(2022, time.January, 2, 0, 0, 0, 0, LocLocal), time.Date(2022, time.January, 3, 0, 0, 0, 0, LocLocal), time.Date(2022, time.January, 4, 0, 0, 0, 0, LocLocal), time.Date(2022, time.January, 5, 0, 0, 0, 0, LocLocal), time.Date(2022, time.January, 6, 0, 0, 0, 0, LocLocal), time.Date(2022, time.January, 7, 0, 0, 0, 0, LocLocal), time.Date(2022, time.January, 8, 0, 0, 0, 0, LocLocal), time.Date(2022, time.January, 9, 0, 0, 0, 0, LocLocal), time.Date(2022, time.January, 10, 0, 0, 0, 0, LocLocal), time.Date(2022, time.January, 11, 0, 0, 0, 0, LocLocal), time.Date(2022, time.January, 12, 0, 0, 0, 0, LocLocal), time.Date(2022, time.January, 13, 0, 0, 0, 0, LocLocal), time.Date(2022, time.January, 14, 0, 0, 0, 0, LocLocal), time.Date(2022, time.January, 15, 0, 0, 0, 0, LocLocal), time.Date(2022, time.January, 16, 0, 0, 0, 0, LocLocal), time.Date(2022, time.January, 17, 0, 0, 0, 0, LocLocal), time.Date(2022, time.January, 18, 0, 0, 0, 0, LocLocal), time.Date(2022, time.January, 19, 0, 0, 0, 0, LocLocal), time.Date(2022, time.January, 20, 0, 0, 0, 0, LocLocal), time.Date(2022, time.January, 21, 0, 0, 0, 0, LocLocal), time.Date(2022, time.January, 22, 0, 0, 0, 0, LocLocal), time.Date(2022, time.January, 23, 0, 0, 0, 0, LocLocal), time.Date(2022, time.January, 24, 0, 0, 0, 0, LocLocal), time.Date(2022, time.January, 25, 0, 0, 0, 0, LocLocal), time.Date(2022, time.January, 26, 0, 0, 0, 0, LocLocal), time.Date(2022, time.January, 27, 0, 0, 0, 0, LocLocal), time.Date(2022, time.January, 28, 0, 0, 0, 0, LocLocal), time.Date(2022, time.January, 29, 0, 0, 0, 0, LocLocal), time.Date(2022, time.January, 30, 0, 0, 0, 0, LocLocal), time.Date(2022, time.January, 31, 0, 0, 0, 0, LocLocal), time.Date(2022, time.February, 1, 0, 0, 0, 0, LocLocal), time.Date(2022, time.February, 2, 0, 0, 0, 0, LocLocal), time.Date(2022, time.February, 3, 0, 0, 0, 0, LocLocal), time.Date(2022, time.February, 4, 0, 0, 0, 0, LocLocal), time.Date(2022, time.February, 5, 0, 0, 0, 0, LocLocal), time.Date(2022, time.February, 6, 0, 0, 0, 0, LocLocal), time.Date(2022, time.February, 7, 0, 0, 0, 0, LocLocal), time.Date(2022, time.February, 8, 0, 0, 0, 0, LocLocal), time.Date(2022, time.February, 9, 0, 0, 0, 0, LocLocal), time.Date(2022, time.February, 10, 0, 0, 0, 0, LocLocal), time.Date(2022, time.February, 11, 0, 0, 0, 0, LocLocal), time.Date(2022, time.February, 12, 0, 0, 0, 0, LocLocal), time.Date(2022, time.February, 13, 0, 0, 0, 0, LocLocal), time.Date(2022, time.February, 14, 0, 0, 0, 0, LocLocal), time.Date(2022, time.February, 15, 0, 0, 0, 0, LocLocal), time.Date(2022, time.February, 16, 0, 0, 0, 0, LocLocal), time.Date(2022, time.February, 17, 0, 0, 0, 0, LocLocal), time.Date(2022, time.February, 18, 0, 0, 0, 0, LocLocal), time.Date(2022, time.February, 19, 0, 0, 0, 0, LocLocal), time.Date(2022, time.February, 20, 0, 0, 0, 0, LocLocal), time.Date(2022, time.February, 21, 0, 0, 0, 0, LocLocal), time.Date(2022, time.February, 22, 0, 0, 0, 0, LocLocal), time.Date(2022, time.February, 23, 0, 0, 0, 0, LocLocal), time.Date(2022, time.February, 24, 0, 0, 0, 0, LocLocal), time.Date(2022, time.February, 25, 0, 0, 0, 0, LocLocal), time.Date(2022, time.February, 26, 0, 0, 0, 0, LocLocal), time.Date(2022, time.February, 27, 0, 0, 0, 0, LocLocal), time.Date(2022, time.February, 28, 0, 0, 0, 0, LocLocal), time.Date(2022, time.March, 1, 0, 0, 0, 0, LocLocal), time.Date(2022, time.March, 2, 0, 0, 0, 0, LocLocal), time.Date(2022, time.March, 3, 0, 0, 0, 0, LocLocal), time.Date(2022, time.March, 4, 0, 0, 0, 0, LocLocal), time.Date(2022, time.March, 5, 0, 0, 0, 0, LocLocal), time.Date(2022, time.March, 6, 0, 0, 0, 0, LocLocal), time.Date(2022, time.March, 7, 0, 0, 0, 0, LocLocal), time.Date(2022, time.March, 8, 0, 0, 0, 0, LocLocal), time.Date(2022, time.March, 9, 0, 0, 0, 0, LocLocal), time.Date(2022, time.March, 10, 0, 0, 0, 0, LocLocal), time.Date(2022, time.March, 11, 0, 0, 0, 0, LocLocal), time.Date(2022, time.March, 12, 0, 0, 0, 0, LocLocal), time.Date(2022, time.March, 13, 0, 0, 0, 0, LocLocal), time.Date(2022, time.March, 14, 0, 0, 0, 0, LocLocal), time.Date(2022, time.March, 15, 0, 0, 0, 0, LocLocal), time.Date(2022, time.March, 16, 0, 0, 0, 0, LocLocal), time.Date(2022, time.March, 17, 0, 0, 0, 0, LocLocal), time.Date(2022, time.March, 18, 0, 0, 0, 0, LocLocal), time.Date(2022, time.March, 19, 0, 0, 0, 0, LocLocal), time.Date(2022, time.March, 20, 0, 0, 0, 0, LocLocal), time.Date(2022, time.March, 21, 0, 0, 0, 0, LocLocal), time.Date(2022, time.March, 22, 0, 0, 0, 0, LocLocal), time.Date(2022, time.March, 23, 0, 0, 0, 0, LocLocal), time.Date(2022, time.March, 24, 0, 0, 0, 0, LocLocal), time.Date(2022, time.March, 25, 0, 0, 0, 0, LocLocal), time.Date(2022, time.March, 26, 0, 0, 0, 0, LocLocal), time.Date(2022, time.March, 27, 0, 0, 0, 0, LocLocal), time.Date(2022, time.March, 28, 0, 0, 0, 0, LocLocal), time.Date(2022, time.March, 29, 0, 0, 0, 0, LocLocal), time.Date(2022, time.March, 30, 0, 0, 0, 0, LocLocal), time.Date(2022, time.March, 31, 0, 0, 0, 0, LocLocal), time.Date(2022, time.April, 1, 0, 0, 0, 0, LocLocal), time.Date(2022, time.April, 2, 0, 0, 0, 0, LocLocal), time.Date(2022, time.April, 3, 0, 0, 0, 0, LocLocal), time.Date(2022, time.April, 4, 0, 0, 0, 0, LocLocal), time.Date(2022, time.April, 5, 0, 0, 0, 0, LocLocal), time.Date(2022, time.April, 6, 0, 0, 0, 0, LocLocal), time.Date(2022, time.April, 7, 0, 0, 0, 0, LocLocal), time.Date(2022, time.April, 8, 0, 0, 0, 0, LocLocal), time.Date(2022, time.April, 9, 0, 0, 0, 0, LocLocal), time.Date(2022, time.April, 10, 0, 0, 0, 0, LocLocal), time.Date(2022, time.April, 11, 0, 0, 0, 0, LocLocal), time.Date(2022, time.April, 12, 0, 0, 0, 0, LocLocal), time.Date(2022, time.April, 13, 0, 0, 0, 0, LocLocal), time.Date(2022, time.April, 14, 0, 0, 0, 0, LocLocal), time.Date(2022, time.April, 15, 0, 0, 0, 0, LocLocal), time.Date(2022, time.April, 16, 0, 0, 0, 0, LocLocal), time.Date(2022, time.April, 17, 0, 0, 0, 0, LocLocal), time.Date(2022, time.April, 18, 0, 0, 0, 0, LocLocal), time.Date(2022, time.April, 19, 0, 0, 0, 0, LocLocal), time.Date(2022, time.April, 20, 0, 0, 0, 0, LocLocal), time.Date(2022, time.April, 21, 0, 0, 0, 0, LocLocal), time.Date(2022, time.April, 22, 0, 0, 0, 0, LocLocal), time.Date(2022, time.April, 23, 0, 0, 0, 0, LocLocal), time.Date(2022, time.April, 24, 0, 0, 0, 0, LocLocal), time.Date(2022, time.April, 25, 0, 0, 0, 0, LocLocal), time.Date(2022, time.April, 26, 0, 0, 0, 0, LocLocal), time.Date(2022, time.April, 27, 0, 0, 0, 0, LocLocal), time.Date(2022, time.April, 28, 0, 0, 0, 0, LocLocal), time.Date(2022, time.April, 29, 0, 0, 0, 0, LocLocal), time.Date(2022, time.April, 30, 0, 0, 0, 0, LocLocal), time.Date(2022, time.May, 1, 0, 0, 0, 0, LocLocal), time.Date(2022, time.May, 2, 0, 0, 0, 0, LocLocal), time.Date(2022, time.May, 3, 0, 0, 0, 0, LocLocal), time.Date(2022, time.May, 4, 0, 0, 0, 0, LocLocal), time.Date(2022, time.May, 5, 0, 0, 0, 0, LocLocal), time.Date(2022, time.May, 6, 0, 0, 0, 0, LocLocal), time.Date(2022, time.May, 7, 0, 0, 0, 0, LocLocal), time.Date(2022, time.May, 8, 0, 0, 0, 0, LocLocal), time.Date(2022, time.May, 9, 0, 0, 0, 0, LocLocal), time.Date(2022, time.May, 10, 0, 0, 0, 0, LocLocal), time.Date(2022, time.May, 11, 0, 0, 0, 0, LocLocal), time.Date(2022, time.May, 12, 0, 0, 0, 0, LocLocal), time.Date(2022, time.May, 13, 0, 0, 0, 0, LocLocal), time.Date(2022, time.May, 14, 0, 0, 0, 0, LocLocal), time.Date(2022, time.May, 15, 0, 0, 0, 0, LocLocal), time.Date(2022, time.May, 16, 0, 0, 0, 0, LocLocal), time.Date(2022, time.May, 17, 0, 0, 0, 0, LocLocal), time.Date(2022, time.May, 18, 0, 0, 0, 0, LocLocal), time.Date(2022, time.May, 19, 0, 0, 0, 0, LocLocal), time.Date(2022, time.May, 20, 0, 0, 0, 0, LocLocal), time.Date(2022, time.May, 21, 0, 0, 0, 0, LocLocal), time.Date(2022, time.May, 22, 0, 0, 0, 0, LocLocal), time.Date(2022, time.May, 23, 0, 0, 0, 0, LocLocal), time.Date(2022, time.May, 24, 0, 0, 0, 0, LocLocal), time.Date(2022, time.May, 25, 0, 0, 0, 0, LocLocal), time.Date(2022, time.May, 26, 0, 0, 0, 0, LocLocal), time.Date(2022, time.May, 27, 0, 0, 0, 0, LocLocal), time.Date(2022, time.May, 28, 0, 0, 0, 0, LocLocal), time.Date(2022, time.May, 29, 0, 0, 0, 0, LocLocal), time.Date(2022, time.May, 30, 0, 0, 0, 0, LocLocal), time.Date(2022, time.May, 31, 0, 0, 0, 0, LocLocal), time.Date(2022, time.June, 1, 0, 0, 0, 0, LocLocal), time.Date(2022, time.June, 2, 0, 0, 0, 0, LocLocal), time.Date(2022, time.June, 3, 0, 0, 0, 0, LocLocal), time.Date(2022, time.June, 4, 0, 0, 0, 0, LocLocal), time.Date(2022, time.June, 5, 0, 0, 0, 0, LocLocal), time.Date(2022, time.June, 6, 0, 0, 0, 0, LocLocal), time.Date(2022, time.June, 7, 0, 0, 0, 0, LocLocal), time.Date(2022, time.June, 8, 0, 0, 0, 0, LocLocal), time.Date(2022, time.June, 9, 0, 0, 0, 0, LocLocal), time.Date(2022, time.June, 10, 0, 0, 0, 0, LocLocal), time.Date(2022, time.June, 11, 0, 0, 0, 0, LocLocal), time.Date(2022, time.June, 12, 0, 0, 0, 0, LocLocal), time.Date(2022, time.June, 13, 0, 0, 0, 0, LocLocal), time.Date(2022, time.June, 14, 0, 0, 0, 0, LocLocal), time.Date(2022, time.June, 15, 0, 0, 0, 0, LocLocal), time.Date(2022, time.June, 16, 0, 0, 0, 0, LocLocal), time.Date(2022, time.June, 17, 0, 0, 0, 0, LocLocal), time.Date(2022, time.June, 18, 0, 0, 0, 0, LocLocal), time.Date(2022, time.June, 19, 0, 0, 0, 0, LocLocal), time.Date(2022, time.June, 20, 0, 0, 0, 0, LocLocal), time.Date(2022, time.June, 21, 0, 0, 0, 0, LocLocal), time.Date(2022, time.June, 22, 0, 0, 0, 0, LocLocal), time.Date(2022, time.June, 23, 0, 0, 0, 0, LocLocal), time.Date(2022, time.June, 24, 0, 0, 0, 0, LocLocal), time.Date(2022, time.June, 25, 0, 0, 0, 0, LocLocal), time.Date(2022, time.June, 26, 0, 0, 0, 0, LocLocal), time.Date(2022, time.June, 27, 0, 0, 0, 0, LocLocal), time.Date(2022, time.June, 28, 0, 0, 0, 0, LocLocal), time.Date(2022, time.June, 29, 0, 0, 0, 0, LocLocal), time.Date(2022, time.June, 30, 0, 0, 0, 0, LocLocal), time.Date(2022, time.July, 1, 0, 0, 0, 0, LocLocal), time.Date(2022, time.July, 2, 0, 0, 0, 0, LocLocal), time.Date(2022, time.July, 3, 0, 0, 0, 0, LocLocal), time.Date(2022, time.July, 4, 0, 0, 0, 0, LocLocal), time.Date(2022, time.July, 5, 0, 0, 0, 0, LocLocal), time.Date(2022, time.July, 6, 0, 0, 0, 0, LocLocal), time.Date(2022, time.July, 7, 0, 0, 0, 0, LocLocal), time.Date(2022, time.July, 8, 0, 0, 0, 0, LocLocal), time.Date(2022, time.July, 9, 0, 0, 0, 0, LocLocal), time.Date(2022, time.July, 10, 0, 0, 0, 0, LocLocal), time.Date(2022, time.July, 11, 0, 0, 0, 0, LocLocal), time.Date(2022, time.July, 12, 0, 0, 0, 0, LocLocal), time.Date(2022, time.July, 13, 0, 0, 0, 0, LocLocal), time.Date(2022, time.July, 14, 0, 0, 0, 0, LocLocal), time.Date(2022, time.July, 15, 0, 0, 0, 0, LocLocal), time.Date(2022, time.July, 16, 0, 0, 0, 0, LocLocal), time.Date(2022, time.July, 17, 0, 0, 0, 0, LocLocal), time.Date(2022, time.July, 18, 0, 0, 0, 0, LocLocal), time.Date(2022, time.July, 19, 0, 0, 0, 0, LocLocal), time.Date(2022, time.July, 20, 0, 0, 0, 0, LocLocal), time.Date(2022, time.July, 21, 0, 0, 0, 0, LocLocal), time.Date(2022, time.July, 22, 0, 0, 0, 0, LocLocal), time.Date(2022, time.July, 23, 0, 0, 0, 0, LocLocal), time.Date(2022, time.July, 24, 0, 0, 0, 0, LocLocal), time.Date(2022, time.July, 25, 0, 0, 0, 0, LocLocal), time.Date(2022, time.July, 26, 0, 0, 0, 0, LocLocal), time.Date(2022, time.July, 27, 0, 0, 0, 0, LocLocal), time.Date(2022, time.July, 28, 0, 0, 0, 0, LocLocal), time.Date(2022, time.July, 29, 0, 0, 0, 0, LocLocal), time.Date(2022, time.July, 30, 0, 0, 0, 0, LocLocal), time.Date(2022, time.July, 31, 0, 0, 0, 0, LocLocal), time.Date(2022, time.August, 1, 0, 0, 0, 0, LocLocal), time.Date(2022, time.August, 2, 0, 0, 0, 0, LocLocal), time.Date(2022, time.August, 3, 0, 0, 0, 0, LocLocal), time.Date(2022, time.August, 4, 0, 0, 0, 0, LocLocal), time.Date(2022, time.August, 5, 0, 0, 0, 0, LocLocal), time.Date(2022, time.August, 6, 0, 0, 0, 0, LocLocal), time.Date(2022, time.August, 7, 0, 0, 0, 0, LocLocal), time.Date(2022, time.August, 8, 0, 0, 0, 0, LocLocal), time.Date(2022, time.August, 9, 0, 0, 0, 0, LocLocal), time.Date(2022, time.August, 10, 0, 0, 0, 0, LocLocal), time.Date(2022, time.August, 11, 0, 0, 0, 0, LocLocal), time.Date(2022, time.August, 12, 0, 0, 0, 0, LocLocal), time.Date(2022, time.August, 13, 0, 0, 0, 0, LocLocal), time.Date(2022, time.August, 14, 0, 0, 0, 0, LocLocal), time.Date(2022, time.August, 15, 0, 0, 0, 0, LocLocal), time.Date(2022, time.August, 16, 0, 0, 0, 0, LocLocal), time.Date(2022, time.August, 17, 0, 0, 0, 0, LocLocal), time.Date(2022, time.August, 18, 0, 0, 0, 0, LocLocal), time.Date(2022, time.August, 19, 0, 0, 0, 0, LocLocal), time.Date(2022, time.August, 20, 0, 0, 0, 0, LocLocal), time.Date(2022, time.August, 21, 0, 0, 0, 0, LocLocal), time.Date(2022, time.August, 22, 0, 0, 0, 0, LocLocal), time.Date(2022, time.August, 23, 0, 0, 0, 0, LocLocal), time.Date(2022, time.August, 24, 0, 0, 0, 0, LocLocal), time.Date(2022, time.August, 25, 0, 0, 0, 0, LocLocal), time.Date(2022, time.August, 26, 0, 0, 0, 0, LocLocal), time.Date(2022, time.August, 27, 0, 0, 0, 0, LocLocal), time.Date(2022, time.August, 28, 0, 0, 0, 0, LocLocal), time.Date(2022, time.August, 29, 0, 0, 0, 0, LocLocal), time.Date(2022, time.August, 30, 0, 0, 0, 0, LocLocal), time.Date(2022, time.August, 31, 0, 0, 0, 0, LocLocal), time.Date(2022, time.September, 1, 0, 0, 0, 0, LocLocal), time.Date(2022, time.September, 2, 0, 0, 0, 0, LocLocal), time.Date(2022, time.September, 3, 0, 0, 0, 0, LocLocal), time.Date(2022, time.September, 4, 0, 0, 0, 0, LocLocal), time.Date(2022, time.September, 5, 0, 0, 0, 0, LocLocal), time.Date(2022, time.September, 6, 0, 0, 0, 0, LocLocal), time.Date(2022, time.September, 7, 0, 0, 0, 0, LocLocal), time.Date(2022, time.September, 8, 0, 0, 0, 0, LocLocal), time.Date(2022, time.September, 9, 0, 0, 0, 0, LocLocal), time.Date(2022, time.September, 10, 0, 0, 0, 0, LocLocal), time.Date(2022, time.September, 11, 0, 0, 0, 0, LocLocal), time.Date(2022, time.September, 12, 0, 0, 0, 0, LocLocal), time.Date(2022, time.September, 13, 0, 0, 0, 0, LocLocal), time.Date(2022, time.September, 14, 0, 0, 0, 0, LocLocal), time.Date(2022, time.September, 15, 0, 0, 0, 0, LocLocal), time.Date(2022, time.September, 16, 0, 0, 0, 0, LocLocal), time.Date(2022, time.September, 17, 0, 0, 0, 0, LocLocal), time.Date(2022, time.September, 18, 0, 0, 0, 0, LocLocal), time.Date(2022, time.September, 19, 0, 0, 0, 0, LocLocal), time.Date(2022, time.September, 20, 0, 0, 0, 0, LocLocal), time.Date(2022, time.September, 21, 0, 0, 0, 0, LocLocal), time.Date(2022, time.September, 22, 0, 0, 0, 0, LocLocal), time.Date(2022, time.September, 23, 0, 0, 0, 0, LocLocal), time.Date(2022, time.September, 24, 0, 0, 0, 0, LocLocal), time.Date(2022, time.September, 25, 0, 0, 0, 0, LocLocal), time.Date(2022, time.September, 26, 0, 0, 0, 0, LocLocal), time.Date(2022, time.September, 27, 0, 0, 0, 0, LocLocal), time.Date(2022, time.September, 28, 0, 0, 0, 0, LocLocal), time.Date(2022, time.September, 29, 0, 0, 0, 0, LocLocal), time.Date(2022, time.September, 30, 0, 0, 0, 0, LocLocal), time.Date(2022, time.October, 1, 0, 0, 0, 0, LocLocal), time.Date(2022, time.October, 2, 0, 0, 0, 0, LocLocal), time.Date(2022, time.October, 3, 0, 0, 0, 0, LocLocal), time.Date(2022, time.October, 4, 0, 0, 0, 0, LocLocal)},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GetDatesLastYearToToday(start, end)
			if !assert.Equal(t, tt.want, got) {
				t.Errorf("GetDatesLastYearToToday() got = %v, want %v", got, tt.want)
			}
		})
	}
}
