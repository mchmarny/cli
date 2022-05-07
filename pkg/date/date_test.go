package date

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestPrettyDuration(t *testing.T) {
	t.Run("pretty yesterday", func(t *testing.T) {
		v := time.Now().UTC().AddDate(0, 0, -1)
		s := DaysSince(v)
		assert.NotEmpty(t, s)
		assert.Equal(t, "1 day", s)
	})
	t.Run("pretty last year", func(t *testing.T) {
		v := time.Now().UTC().AddDate(-1, 0, 0)
		s := DaysSince(v)
		assert.NotEmpty(t, s)
		assert.Equal(t, "1 years, 0 months, and 0 days", s)
	})
	t.Run("yesterday to today", func(t *testing.T) {
		y := Yesterday()
		d := Today()
		diff := d.Sub(y)
		assert.Equal(t, float64(24), diff.Hours())
	})
	t.Run("today formats", func(t *testing.T) {
		now := time.Now().UTC()
		d1 := DateStr(now)
		d2 := TodayStr()
		assert.Equal(t, d1, d2)
	})
	t.Run("yesterday formats", func(t *testing.T) {
		ye := time.Now().UTC().AddDate(0, 0, -1)
		d1 := DateStr(ye)
		d2 := YesterdayStr()
		assert.Equal(t, d1, d2)
	})
}

func TestWeekFormating(t *testing.T) {
	t.Run("week repro", func(t *testing.T) {
		now := time.Now().UTC()
		w1 := Week(now)
		w2 := ThisWeek()
		assert.Equal(t, w1, w2)
	})
	t.Run("week formats", func(t *testing.T) {
		y := 2002
		w := 1
		w1 := fmt.Sprintf(weekFormat, y, w)
		assert.Len(t, w1, 7)

		w = 11
		w2 := fmt.Sprintf(weekFormat, y, w)
		assert.Len(t, w2, 7)
	})
}

func TestWeekParsing(t *testing.T) {
	t.Run("test now", func(t *testing.T) {
		now := time.Now().UTC()

		week := Week(now)
		assert.Len(t, week, 7)

		start, end, err := WeekToDates(week)
		assert.NoError(t, err)
		assert.False(t, start.IsZero())
		assert.False(t, end.IsZero())

		assert.Equal(t, time.Monday, start.Weekday())
		assert.Equal(t, time.Sunday, end.Weekday())

		week1 := Week(start)
		assert.Equal(t, week, week1)

		week2 := Week(end)
		assert.Equal(t, week, week2)
	})
	t.Run("test last week", func(t *testing.T) {
		start, end := LastWeek()
		assert.False(t, start.IsZero())
		assert.False(t, end.IsZero())

		assert.Equal(t, time.Monday, start.Weekday())
		assert.Equal(t, time.Sunday, end.Weekday())
	})
}
