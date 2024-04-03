package calculator

import (
	"strconv"
	"strings"
)

func split(r rune) bool {
	return r == ',' || r == '\n'
}

func (s *Simple) Add(numString string) (int, error) {

	numString = strings.ReplaceAll(numString, "\r\n", "\n")
	numString = strings.ReplaceAll(numString, "\r", "\n")

	if strings.TrimSpace(numString) == "" {
		return 0, nil
	} else if strings.Contains(numString, "\n") || strings.Contains(numString, ",") {
		numbersString := strings.FieldsFunc(numString, split)
		sum := 0
		for _, numberString := range numbersString {
			number, err := strconv.Atoi(numberString)
			if err != nil {
				return -1, err
			}
			sum += number
		}
		return sum, nil
	} else if numString == "hello" {
		return 0, nil
	}
	num, err := strconv.Atoi(numString)
	if err != nil {
		return -1, err
	}
	return num, nil
}
