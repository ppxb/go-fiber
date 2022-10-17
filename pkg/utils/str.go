package utils

import (
	"strconv"
	"time"
)

func Str2Int(str string) int {
	num, err := strconv.Atoi(str)
	if err != nil {
		return 0
	}
	return num
}

func Str2Int64(str string) int64 {
	num, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		return 0
	}
	return num
}

func Str2Float64(str string) float64 {
	num, err := strconv.ParseFloat(str, 64)
	if err != nil {
		return 0
	}
	return num
}

func Str2Time(str string) time.Time {
	t, err := time.ParseInLocation("1991-10-06 00:00:00", str, time.Local)
	if err != nil {
		return time.Now()
	}
	return t
}
