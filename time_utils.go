/**
* @Author: Cooper
* @Date: 2019/10/16 15:20
 */

package utils

import (
	"fmt"
	"time"
)

// get the current time full format
func GetCurrentTimeFullFormat() string {
	return TimeFullFormat(time.Now())
}

// get the current time format to second
func GetCurrentTimeFormatToSecond() string {
	return TimeFormatToSecond(time.Now())
}

// get the current time format to day
func GetCurrentTimeFormatToDay() string {
	return TimeFormatToDay(time.Now())
}

// get full format for [t]
func TimeFullFormat(t time.Time) string {
	return t.Format("2006-01-02 15:04:05.999")
}

// get format to second for [t]
func TimeFormatToSecond(t time.Time) string {
	return t.Format("2006-01-02 15:04:05")
}

// get format to day for [t]
func TimeFormatToDay(t time.Time) string {
	return t.Format("2006-01-02")
}


// get a Time next second of [t]
func TimeNextSecond(t time.Time) time.Time {
	return TimeNextUnit(t , time.Second)
}

// get a Time next minute of [t]
func TimeNextMinute(t time.Time) time.Time {
	return TimeNextUnit(t , time.Minute)
}

// get a Time next hour of [t]
func TimeNextHour(t time.Time) time.Time {
	return TimeNextUnit(t , time.Hour)
}

// get a Time next day of [t]
func TimeNextDay(t time.Time) time.Time {
	return TimeNextUnit(t , time.Hour*24)
}

// get a Time next month of [t]
func TimeNextMonth(t time.Time) time.Time {
	month := t.Month()
	month++
	format := fmt.Sprintf("%d-%d-%d %d:%d:%d.%d",t.Year(),month,t.Day(),t.Hour(),t.Minute(),t.Second(),t.Nanosecond())
	loc , _ := time.LoadLocation("Local")
	afterTime , _ := time.ParseInLocation("2006-01-02 15:04:05.999" , format , loc)
	return afterTime
}

// get a Time next year of [t]
func TimeNextYear(t time.Time) time.Time {
	year := t.Year()
	year++
	format := fmt.Sprintf("%d-%d-%d %d:%d:%d.%d",year,t.Month(),t.Day(),t.Hour(),t.Minute(),t.Second(),t.Nanosecond())
	loc , _ := time.LoadLocation("Local")
	afterTime , _ := time.ParseInLocation("2006-01-02 15:04:05.999" , format , loc)
	return afterTime
}

// get a Time next unit of [t]
func TimeNextUnit(t time.Time , unit time.Duration) time.Time {
	return TimeAfterDate(t , unit , 1)
}

// get a Time after [num] count unit of [t]
func TimeAfterDate(t time.Time, unit time.Duration, num int64) time.Time {
	return t.Add(unit*time.Duration(num))
}