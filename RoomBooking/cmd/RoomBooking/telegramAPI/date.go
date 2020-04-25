package telegramAPI

import (
	"strconv"
	"time"
)

var (
	weekDays = map[string]int{
		"Monday":    1,
		"Tuesday":   2,
		"Wednesday": 3,
		"Thursday":  4,
		"Friday":    5,
	}
	weekEndDays = map[string]int{
		"Saturday": 6,
		"Sunday":   7,
	}
)

//Generate slice with days for menu buttons
func GenerateDays() (days [][]string) {
	today := time.Now()
	if weekEnd(today) {
		days = generateDaysForButtonsWeekEnd(today)
	} else {
		days = generateDaysForButtonsWeekDay(today)
	}
	return days
}

//Return true if today is a weeekend day,
//return false if today is a weekday
func weekEnd(today time.Time) bool {
	todayWeekDay := today.Weekday().String()
	if _, ok := weekEndDays[todayWeekDay]; ok {
		return true
	} else {
		return false
	}
}

//Generates days for buttons for the next week, if today is a weekday
func generateNextWeek(today time.Time) (daysNextWeek [][]string) {
	todayWeekDay := today.Weekday().String()
	l := weekEndDays["Saturday"] - weekDays[todayWeekDay]
	buf := today.AddDate(0, 0, l)
	daysNextWeek = generateDaysForButtonsWeekEnd(buf)
	return
}

//Generates days for buttons till the end of this week, if today is a weekday
func generateDaysForButtonsWeekDay(today time.Time) (days [][]string) {
	var daysThisWeek [][]string
	todayWeekDay := today.Weekday().String()
	dayKey := weekDays[todayWeekDay]
	for i := 0; i < (weekDays["Friday"] - dayKey + 1); i++ {
		weekDayWithDate := nextDay(today, i)
		daysThisWeek = append(daysThisWeek, weekDayWithDate)
	}
	daysNextWeek := generateNextWeek(today)
	days = append(days, daysThisWeek...)
	days = append(days, daysNextWeek...)
	return
}

//Generates days for buttons for the whole next week
//if the day is the weekend day
func generateDaysWeekEnd(today time.Time, k int) (days [][]string) {
	for i := k; i < (k + 5); i++ {
		weekDayWithDate := nextDay(today, i)
		days = append(days, weekDayWithDate)
	}
	return
}

//Generate k variable for generateDaysWeekEnd function
//depending on the weekend day: Saturday k=2, Sunday k=1
func generateDaysForButtonsWeekEnd(today time.Time) (days [][]string) {
	todayWeekDay := today.Weekday().String()
	dayKey := weekEndDays[todayWeekDay]
	if dayKey == weekEndDays["Saturday"] {
		days = generateDaysWeekEnd(today, 2)
	} else if dayKey == weekEndDays["Sunday"] {
		days = generateDaysWeekEnd(today, 1)
	}
	return
}

//Returns (today + n) day's weekday and date in such format: ["Friday","24 April"]
func nextDay(today time.Time, n int) (weekDayWithDate []string) {
	weekDayWithDate = make([]string, 0, 2)
	nextDay := today.AddDate(0, 0, n)

	nextDayWeekDay := nextDay.Weekday().String()

	_, month, date := nextDay.Date()
	nextDayDateMonth := strconv.Itoa(date) + " " + month.String()

	weekDayWithDate = append(weekDayWithDate, nextDayWeekDay, nextDayDateMonth)
	return
}
