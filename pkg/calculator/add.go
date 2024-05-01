package calculator

import (
	"errors"
	"strconv"
	"strings"
)

func split(r rune) bool {
	return r == ',' || r == '\n'
}

func negativeNumber(num int) error {
	if num < 0 {
		return errors.New("Negatives not allowed: " + strconv.Itoa(num))
	}
	return nil
}

func (s *Simple) checkIfNumberIsOver1000(num int) {
	if num > 1000 {
		s.Log(num)
	}
}

// check if a string has equal amounts of opening and closing brackets and brackets are more than 0
func hasEqualBrackets(s string) bool {
	splitString := strings.Split(s, "")
	opening := 0
	closing := 0
	for _, bracket := range splitString {
		if bracket == "[" {
			opening++
		} else if bracket == "]" {
			closing++
		}
	}
	return opening == closing && opening > 0
}

// takes a string returns an array of what is inside the brackets
func getSeparatorsFromBrackets(s string) []string {
	var brackets []string
	for i := 0; i < len(s); i++ {
		if s[i] == '[' {
			bracket := ""
			i++
			for s[i] != ']' {
				bracket += string(s[i])
				i++
			}
			brackets = append(brackets, bracket)
		}
	}
	return brackets
}

func (s *Simple) Add(numString string) (int, error) {

	// Return 0 if the input is empty
	if strings.TrimSpace(numString) == "" {
		return 0, nil

		// Return the sum of the numbers if the input is a custom delimiter
	} else if strings.HasPrefix(numString, "//") {
		numString := strings.TrimPrefix(numString, "//")

		if hasEqualBrackets(numString) {
			separators := getSeparatorsFromBrackets(numString)
			for _, separator := range separators {
				numString = strings.TrimPrefix(numString, "["+separator+"]")
				numString = strings.ReplaceAll(numString, separator, ",")
			}
			numString = strings.TrimSpace(numString)
			return sumStringOfNumbers(numString, s)
		} else {
			delimiter := strings.Split(numString, "\n")[0]
			numString = strings.TrimPrefix(numString, delimiter+"\n")
			numString = strings.TrimSpace(numString)
			numArr := strings.Split(numString, delimiter)
			sum := 0
			for _, numberString := range numArr {
				numberString = strings.TrimSpace(numberString)
				number, err := strconv.Atoi(numberString)
				if err != nil {
					return -1, err
				}

				err = negativeNumber(number)
				if err != nil {
					return -1, err
				}

				s.checkIfNumberIsOver1000(number)
				sum += number
			}
			return sum, nil
		}

		// Return the sum of the numbers if the input is a list of numbers separated by a comma or newline
	} else if strings.Contains(numString, "\n") || strings.Contains(numString, ",") {
		return sumStringOfNumbers(numString, s)
	}

	// Return the number if the input is a single number
	num, err := strconv.Atoi(numString)
	if err != nil {
		return -1, err
	}
	err = negativeNumber(num)
	if err != nil {
		return -1, err
	}

	s.checkIfNumberIsOver1000(num)
	return num, nil
}

func sumStringOfNumbers(numString string, s *Simple) (int, error) {
	numArr := strings.FieldsFunc(numString, split)
	sum := 0
	for _, numberString := range numArr {
		number, err := strconv.Atoi(numberString)
		if err != nil {
			return -1, err
		}
		err = negativeNumber(number)
		if err != nil {
			return -1, err
		}

		s.checkIfNumberIsOver1000(number)
		sum += number
	}
	return sum, nil
}
