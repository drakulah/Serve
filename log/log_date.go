package log

import (
	"fmt"
	"serve/conv"
	"serve/str"
	"time"
)

func properNum(num int) string {
	if num >= 0 && num < 10 {
		return "0" + conv.IntToString(num)
	}
	return conv.IntToString(num)
}

func LogDate() string {
	t := time.Now().Local()
	hrs := properNum(t.Hour())
	mins := properNum(t.Minute())
	secs := properNum(t.Second())
	month := str.Substring(t.Month().String(), 0, 3)
	day := t.Day()
	year := t.Year()

	return fmt.Sprintf("%s:%s:%s %s %d, %d", hrs, mins, secs, month, day, year)
}
