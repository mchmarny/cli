package date

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/pkg/errors"
)

const (
	daysInMonth  = 32
	monthsInYear = 12
	weekFormat   = "%dW%02d"
)

func Yesterday() time.Time {
	return ZeroTime(time.Now().UTC().AddDate(0, 0, -1))
}

func LastWeek() (start, end time.Time) {
	now := time.Now().UTC()
	year, week := now.ISOWeek()
	if week == 1 {
		year, week = now.AddDate(0, 0, -weekDaysFromMonday).ISOWeek()
	}
	start = ISOWeekToFirstWeekDay(year, week)
	end = start.AddDate(0, 0, weekDaysFromMonday)

	return ZeroTime(start), ZeroTime(end)
}

func YesterdayStr() string {
	return DateStr(Yesterday())
}

func Today() time.Time {
	return ZeroTime(time.Now().UTC())
}

func TodayStr() string {
	return DateStr(Today())
}

func ZeroTime(v time.Time) time.Time {
	return time.Date(v.Year(), v.Month(), v.Day(), 0, 0, 0, 0, time.UTC)
}

func DateStr(v time.Time) string {
	return v.Format("2006-01-02")
}

func ThisWeek() string {
	return Week(time.Now().UTC())
}

// Week returns the ISO year and week number of the given time.
// For example, 2006W02 for second week of year 2006.
// Weeks start on Monday.
func Week(v time.Time) string {
	year, week := v.ISOWeek()
	return fmt.Sprintf(weekFormat, year, week)
}

const (
	weekDaysFromMonday = 6
	expectedWeekParts  = 2
	expectedYearLen    = 4
	expectedMaxWeekLen = 2
	expectedMinWeekLen = 1

	// DaysInWeek is the number of days in a week.
	DaysInWeek = 7
)

// WeekToDates returns the Mon date that start that week and Sunday date that end that week.
func WeekToDates(isoWeek string) (start time.Time, end time.Time, err error) {
	if isoWeek == "" {
		return time.Time{}, time.Time{}, fmt.Errorf("isoWeek is empty")
	}

	p := strings.Split(isoWeek, "W")
	if len(p) != expectedWeekParts {
		return time.Time{}, time.Time{}, fmt.Errorf("isoWeek is invalid")
	}

	y := strings.TrimSpace(p[0])
	if len(y) != expectedYearLen {
		return time.Time{}, time.Time{}, fmt.Errorf("isoWeek year is invalid")
	}

	year, err := strconv.Atoi(y)
	if err != nil {
		return time.Time{}, time.Time{}, errors.Wrap(err, "failed to parse year")
	}

	w := strings.TrimPrefix(strings.TrimSpace(p[1]), "0")
	if len(w) < expectedMinWeekLen || len(w) > expectedMaxWeekLen {
		return time.Time{}, time.Time{}, fmt.Errorf("isoWeek week is invalid")
	}

	week, err := strconv.Atoi(w)
	if err != nil {
		return time.Time{}, time.Time{}, errors.Wrap(err, "failed to parse year")
	}

	start = ISOWeekToFirstWeekDay(year, week)
	end = start.AddDate(0, 0, weekDaysFromMonday)

	return start, end, nil
}

func FirstDayThisWeek() time.Time {
	return ISOWeekToFirstWeekDay(time.Now().UTC().ISOWeek())
}

func ISOWeekToFirstWeekDay(year int, week int) time.Time {
	date := time.Date(year, 0, 0, 0, 0, 0, 0, time.UTC)
	isoYear, isoWeek := date.ISOWeek()

	// iterate back to Monday
	for date.Weekday() != time.Monday {
		date = date.AddDate(0, 0, -1)
		isoYear, isoWeek = date.ISOWeek()
	}

	// iterate forward to the first day of the first week
	for isoYear < year {
		date = date.AddDate(0, 0, DaysInWeek)
		isoYear, isoWeek = date.ISOWeek()
	}

	// iterate forward to the first day of the given week
	for isoWeek < week {
		date = date.AddDate(0, 0, DaysInWeek)
		_, isoWeek = date.ISOWeek()
	}

	return date
}

// DaysSince prints pretty duration since date (now is in UTC)
func DaysSince(a time.Time) string {
	b := time.Now().UTC()

	if a.Location() != b.Location() {
		b = b.In(a.Location())
	}
	if a.After(b) {
		a, b = b, a
	}

	y1, M1, d1 := a.Date()
	y2, M2, d2 := b.Date()

	year := y2 - y1
	month := int(M2 - M1)
	day := d2 - d1

	if day < 0 {
		// days in month:
		t := time.Date(y1, M1, daysInMonth, 0, 0, 0, 0, time.UTC)
		day += daysInMonth - t.Day()
		month--
	}

	if month < 0 {
		month += monthsInYear
		year--
	}

	if year > 0 {
		return fmt.Sprintf("%d years, %d months, and %d days", year, month, day)
	}

	if month > 1 {
		return fmt.Sprintf("%d months and %d days", month, day)
	}

	if day == 1 {
		return fmt.Sprintf("%d day", day)
	}

	return fmt.Sprintf("%d days", day)
}
