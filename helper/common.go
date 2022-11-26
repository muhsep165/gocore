package helper

import (
	"strconv"
	"time"
)

const BE_DEFAULT_DATE_FORMAT = "2006-01-02"

func ToInt(str string) int {
	i, e := strconv.Atoi(str)
	if e != nil {
		return 0
	} else {
		return i
	}
}

func ToInt64(str string) uint64 {
	i, e := strconv.ParseUint(str, 10, 64)
	if e != nil {
		return 0
	} else {
		return i
	}
}

func ToFloat64(str string) float64 {
	i, e := strconv.ParseFloat(str, 64)
	if e != nil {
		return 0
	} else {
		return i
	}
}

func IntToString(intgr int) string {
	return strconv.Itoa(intgr)
}

func Float64ToString(flt64 float64) string {
	return strconv.FormatFloat(flt64, 'f', 3, 64)
}

func Int64ToString(intgr uint64) string {
	return strconv.FormatUint(intgr, 10)
}

func InArray(val string, array []string) (exists bool) {
	exists = false
	for _, v := range array {
		if val == v {
			exists = true
			return
		}
	}
	return
}

func GetFullDate() string {
	currentTime := time.Now()
	return currentTime.Format(BE_DEFAULT_DATE_FORMAT)
}

func GetFullTIme() string {
	currentTime := time.Now()
	return currentTime.Format("15:04:05")
}

func GetFullDateTime() string {
	return GetFullDate() + " " + GetFullTIme()
}

func GetTodayUTC(dtUtc string) string {
	dateValue, err := time.Parse(time.RFC3339, dtUtc)
	if err != nil {
		return dtUtc
	} else {
		return dateValue.Format(BE_DEFAULT_DATE_FORMAT)
	}
}
