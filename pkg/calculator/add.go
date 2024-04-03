package calculator

import (
	"strconv"
	"strings"
)

func split(r rune) bool {
	return r == ',' || r == '\n'
}

func (s *Simple) Add(numString string) (int, error) {

	// Return 0 if the input is empty
	if strings.TrimSpace(numString) == "" {
		return 0, nil

		// Return the sum of the numbers if the input is a custom delimiter
	} else if strings.HasPrefix(numString, "//") {
		numString := strings.TrimPrefix(numString, "//")
		delimiter := strings.Split(numString, "\n")[0]
		numString = strings.TrimPrefix(numString, delimiter+"\n")
		numString = strings.TrimPrefix(numString, "\n")
		numArr := strings.Split(numString, delimiter)
		sum := 0
		for _, numberString := range numArr {
			number, err := strconv.Atoi(numberString)
			if err != nil {
				return -1, err
			}
			sum += number
		}
		return sum, nil

		// Return the sum of the numbers if the input is a list of numbers separated by a comma or newline
	} else if strings.Contains(numString, "\n") || strings.Contains(numString, ",") {
		numArr := strings.FieldsFunc(numString, split)
		sum := 0
		for _, numberString := range numArr {
			number, err := strconv.Atoi(numberString)
			if err != nil {
				return -1, err
			}
			sum += number
		}
		return sum, nil
	}

	// Return the number if the input is a single number
	num, err := strconv.Atoi(numString)
	if err != nil {
		return -1, err
	}
	return num, nil
}
