package dateutils

import (
	"testing"
	"time"
)

func TestToday(t *testing.T) {
	today := *Today()
	if !today.Equal(*Today()) {
		t.Error("")
	}
}

func TestEoD(t *testing.T) {
	eod := *EoD()
	if eod != TodayStart().Add(time.Hour*23).Add(time.Minute*59).Add(time.Second*59) {
		t.Error("")
	}
}

func TestYesterday(t *testing.T) {
	yesterday := *Yesterday()
	if yesterday != TodayStart().AddDate(0, 0, -1) {
		t.Error("")
	}
}

func TestTomorrow(t *testing.T) {
	tomorrow := *Tomorrow()
	if tomorrow != TodayStart().AddDate(0, 0, 1) {
		t.Error("")
	}
}

func TestLastWeek(t *testing.T) {
	lastWeek := *LastWeek()
	if lastWeek != TodayStart().AddDate(0, 0, -7) {
		t.Error("")
	}
}

func TestLastMonth(t *testing.T) {
	lastMonth := *LastMonth()
	if lastMonth != TodayStart().AddDate(0, -1, 0) {
		t.Error("")
	}
}

func TestLastYear(t *testing.T) {
	lastYear := *LastYear()
	if lastYear != TodayStart().AddDate(-1, 0, 0) {
		t.Error("")
	}
}

func TestNextWeek(t *testing.T) {
	nextWeek := *NextWeek()
	if nextWeek != TodayStart().AddDate(0, 0, 7) {
		t.Error("")
	}
}

func TestNextMonth(t *testing.T) {
	nextMonth := *NextMonth()
	if nextMonth != TodayStart().AddDate(0, 1, 0) {
		t.Error("")
	}
}

func TestNextYear(t *testing.T) {
	nextYear := *NextYear()
	if nextYear != TodayStart().AddDate(1, 0, 0) {
		t.Error("")
	}
}
