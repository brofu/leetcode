package str

import "strings"

func reformatDate(date string) string {

	dateStrs := strings.Split(date, " ")

	day, _, _ := dateStrs[0], dateStrs[1], dateStrs[2]

	_ = map[string]string{
		"Jan": "01",
		"Feb": "02",
		"May": "03",
	}[day]

	return ""
}
