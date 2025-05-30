package iteration

import "strings"

func Repeat(text string, times int) string {
	var repeated string
	for i := 0; i < times; i++ {
		repeated += text
	}

	return repeated
}

func RepeatWithStringsBuilder(text string, times int) string {
	var repeated strings.Builder

	for i := 0; i < times; i++ {
		repeated.WriteString(text)
	}

	return repeated.String()
}
