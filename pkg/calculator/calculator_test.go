package calculator

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
)

// New test cases using suite, assert and mockio

// SimpleSuite defines the test suite for the Simple calculator
type SimpleSuite struct {
	suite.Suite
	calc   Simple
	logger Logger
}

// SetupTest initializes the Simple calculator and logger before each test
func (suite *SimpleSuite) SetupTest() {
	suite.logger = NewSimpleLogger()
	suite.calc = NewSimple(suite.logger)
}

// TestSimpleSuite runs the Simple test suite, and all the tests within the suite
func TestSimpleSuite(t *testing.T) {
	suite.Run(t, new(SimpleSuite))
}

// HelperAddAndTestResult is a helper function to test the Add function
// It takes the test suite, input string, and expected result as arguments
// TODO: Refactor this function to use a variadic parameter for input and want
func HelperAddAndTestResult(suite *SimpleSuite, input string, want int) {
	result, err := suite.calc.Add(input)
	if err != nil {
		suite.T().Errorf("Error: %v", err)
	}
	assert.Equal(suite.T(), want, result, "Want: %v, Got: %v", want, result)
}

// HelperAddAndTestError is a helper function to test the Add function
// It takes the test suite, input string, and expected error as arguments
// TODO: Refactor this function to use a variadic parameter for input and want
func HelperAddAndTestError(suite *SimpleSuite, input string, want string) {
	_, err := suite.calc.Add(input)
	if err == nil {
		suite.T().Errorf("Want: %v, Got: %v", want, err.Error())
	}
	assert.Equal(suite.T(), want, err.Error(), "Want: %v, Got: %v", want, err.Error())
}

func (suite *SimpleSuite) Test_Add_Empty_String() {
	input := ""
	want := 0

	HelperAddAndTestResult(suite, input, want)
}

func (suite *SimpleSuite) Test_Add_Single_Number() {
	input := "1"
	want := 1

	HelperAddAndTestResult(suite, input, want)
}

func (suite *SimpleSuite) Test_Add_Two_Numbers() {
	input := "1,2"
	want := 3

	HelperAddAndTestResult(suite, input, want)
}

func (suite *SimpleSuite) Test_Add_Multiple_Numbers() {
	tests := []struct {
		input string
		want  int
	}{
		{"1,2,3", 6},
		{"1,2,3,4", 10},
		{"1,2,3,4,5", 15},
		{"1,2,3,4,5,6", 21},
		{"1,2,3,4,5,6,7", 28},
		{"1,2,3,4,5,6,7,8", 36},
		{"1\n2\n3\n4\n5\n6\n7\n8", 36},
		{"1\n2\n3\n4\n5\n6\n7\n8\n9", 45},
		{"1\n2\n3\n4\n5,6\n7\n8,9\n10", 55},
	}
	for _, test := range tests {
		HelperAddAndTestResult(suite, test.input, test.want)
	}
}

func (suite *SimpleSuite) Test_Add_Special_Separator() {
	tests := []struct {
		input string
		want  int
	}{
		{"//.\n1.2.3", 6},
		{"//B\n1B2B3B4", 10},
		{"//;\n1;2;3;4;5", 15},
		{"//;\n1;2;3;4;5;6", 21},
		{"//;\n1;2;3;4;5;6;7", 28},
		{"//;\n1;2;3;4;5;6;7;8", 36},
		{"//sep\n1sep2sep3sep4sep5sep6sep7sep8", 36},
		{"//\t\n1\t2\t3\t4\t5\t6\t7\t8\t9", 45},
		{"//\t\n1\t2\t3\t4\t5\t6\t7\t8\t9\t10", 55},
	}
	for _, test := range tests {
		HelperAddAndTestResult(suite, test.input, test.want)
	}
}

func (suite *SimpleSuite) Test_Add_Negative_Exception() {
	input := "1,-2,3"
	want := "Negatives not allowed: -2"

	HelperAddAndTestError(suite, input, want)
}

func (suite *SimpleSuite) Test_Add_Multiple_Negative_Exceptions() {
	tests := []struct {
		input string
		want  string
	}{
		{"-1", "Negatives not allowed: -1"},
		{"1,-2,3,-4", "Negatives not allowed: -2"},
		{"1\n2\n3\n4\n5\n-6\n7\n8\n9", "Negatives not allowed: -6"},
		{"1\n2\n3,4\n5,6\n7\n-8,9\n10", "Negatives not allowed: -8"},
		{"//;\n1;-2;3;-4;5", "Negatives not allowed: -2"},
		{"//;\n1;2;3;-4;5;-6", "Negatives not allowed: -4"},
	}
	for _, test := range tests {
		HelperAddAndTestError(suite, test.input, test.want)
	}
}
