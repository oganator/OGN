package controllers

import (
	"math"
	"strconv"
)

// Datetype -
type Datetype struct {
	Dateint   int    `json:"Dateint,omitempty"`
	Month     int    `json:"Month,omitempty"`
	Quarter   int    `json:"Quarter,omitempty"`
	Year      int    `json:"Year,omitempty"`
	MonthName string `json:"MonthName,omitempty"`
	Bool      bool   `json:"Bool,omitempty"`
}

// Dateadd -
func Dateadd(x Datetype, amt int) (y Datetype) {
	y.Month = Mod((amt+x.Month-1), 12) + 1
	switch {
	case amt < 0 && int(math.Abs(float64(amt))) >= x.Month:
		y.Year = x.Year + int((amt+x.Month)/12-1)
	default:
		y.Year = x.Year + int((x.Month+amt-1)/12)
	}
	mstring := strconv.Itoa(y.Month)
	ystring := strconv.Itoa(y.Year)
	switch {
	case y.Month < 10:
		y.Dateint, _ = strconv.Atoi(ystring + "0" + mstring)
	case y.Month >= 10:
		y.Dateint, _ = strconv.Atoi(ystring + mstring)
	}
	if y.Month == 1 {
		y.MonthName = "Jan"
	}
	if y.Month == 2 {
		y.MonthName = "Feb"
	}
	if y.Month == 3 {
		y.MonthName = "Mar"
	}
	if y.Month == 4 {
		y.MonthName = "Apr"
	}
	if y.Month == 5 {
		y.MonthName = "May"
	}
	if y.Month == 6 {
		y.MonthName = "Jun"
	}
	if y.Month == 7 {
		y.MonthName = "Jul"
	}
	if y.Month == 8 {
		y.MonthName = "Aug"
	}
	if y.Month == 9 {
		y.MonthName = "Sep"
	}
	if y.Month == 10 {
		y.MonthName = "Oct"
	}
	if y.Month == 11 {
		y.MonthName = "Nov"
	}
	if y.Month == 12 {
		y.MonthName = "Dec"
	}
	return y
}

// Add - Performs a Dateadd() on the datetype
func (date *Datetype) Add(n int) {
	*date = Dateadd(*date, n)
}

//a - b
func dateintdiff(a, b int) (result int) {
	aa := strconv.Itoa(a)
	bb := strconv.Itoa(b)
	if len(aa) < 6 || len(bb) < 6 {
	} else {
		amonth, _ := strconv.Atoi(aa[4:])
		bmonth, _ := strconv.Atoi(bb[4:])
		month := amonth - bmonth
		ayear, _ := strconv.Atoi(aa[:4])
		byear, _ := strconv.Atoi(bb[:4])
		year := ayear - byear
		result := month + year*12 + 1
		return result
	}
	return result
}

// Mod -
func Mod(a, b int) int {
	m := a % b
	if a < 0 && b < 0 {
		m -= b
	}
	if a < 0 && b > 0 {
		m += b
	}
	return m
}
