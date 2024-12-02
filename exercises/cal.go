package exercises

import (
	"fmt"
	"time"
)

const empty = "   "
const separator = " "
const smallSeparator = "  "
const monthSeparator = ""

func separators() string {
	separator := " "
	for i := 0; i < 10; i++ {
		separator += separator
	}
	return separator
}

func month(offset string) string {
	output := ""

	// header
	now := time.Now()
	month := now.Month().String()
	year := now.Year()
	output += offset
	output += fmt.Sprintf("%v%v%v%v", empty, month, separator, year)
	output += "\n"

	// header day names
	dayNames := []string{"Mo", "Tu", "We", "Th", "Fr", "Sa", "Su"}
	for _, day := range dayNames {
		output += fmt.Sprintf(" %v", day)
	}
	output += "\n"

	// week days
	firstOfMonth := time.Date(now.Year(),
		now.Month(),
		1, 0, 0, 0, 0,
		now.Location())
	startDay := firstOfMonth.Weekday() - 1

	const rows = 5
	const week = 7

	var days = [7 * rows]int{}
	for i := 0; i < int(startDay); i++ {
		days[i] = 0
	}

	for i := int(startDay); i < week*rows; i++ {
		days[i] = int(i-int(startDay)) + 1
	}

	for irow := 0; irow < rows; irow++ {
		for day := 0; day < week; day++ {
			if days[day+irow*week] != 0 {
				if days[day+irow*week] <= 31 {
					if days[day+irow*week] < 10 {
						output += fmt.Sprintf("%v%v", smallSeparator, days[day+irow*week])
					} else {
						output += fmt.Sprintf("%v%v", separator, days[day+irow*week])
					}
				}
			} else {
				output += fmt.Sprintf("%v", empty)
			}
		}
		output += "\n"

	}

	return output
}

func Cal() string {
	separator := separators()
	return month(separator)
}
