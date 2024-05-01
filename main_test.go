package main

import (
	"SimpleCalculator/pkg/calculator"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"strings"
	"testing"
)

// SimpleSuite defines the test suite for the Simple calculator
type MainSuite struct {
	suite.Suite
	calc    calculator.Simple
	logger  calculator.Logger
	appName string
}

// SetupTest initializes the Simple calculator and logger before each test
func (suite *MainSuite) SetupTest() {
	suite.logger = calculator.NewSimpleLogger()
	suite.calc = calculator.NewSimple(suite.logger)
}

// TestSimpleSuite runs the Simple test suite, and all the tests within the suite
func TestMainSuite(t *testing.T) {
	suite.Run(t, new(MainSuite))
}

// HelperAddAndTestResult is a helper function to test the Add function
// It takes the test suite, input string, and expected result as arguments
// TODO: Refactor this function to use a variadic parameter for input and want
func HelperAddAndTestResult(suite *MainSuite, input string, want int) {
	result, err := suite.calc.Add(input)
	if err != nil {
		suite.T().Errorf("Error: %v", err)
	}
	assert.Equal(suite.T(), want, result, "Want: %v, Got: %v", want, result)
}

// HelperAddAndTestError is a helper function to test the Add function
// It takes the test suite, input string, and expected error as arguments
// TODO: Refactor this function to use a variadic parameter for input and want
func HelperAddAndTestError(suite *MainSuite, input string, want string) {
	_, err := suite.calc.Add(input)
	if err == nil {
		suite.T().Errorf("Want: %v, Got: %v", want, err.Error())
	}
	assert.Equal(suite.T(), want, err.Error(), "Want: %v, Got: %v", want, err.Error())
}

func (suite *MainSuite) Test_Main_No_Input() {
	input := "\n\n"
	reader := strings.NewReader(input)
	NewCalc(reader, suite.calc)

	// Output:
	// Welcome to the Simple Calculator! 🧮
	//
	// 	Usage: Any numbers you write will be added together
	// 		Ex: 1,2,3,4	 will return 10
	// 		Ex: 1,2,3,4,5,6,7,8,9,10	 will return 55
	// 	You can also use a custom separator
	// 		Ex: //;<newline>1;2;3;4;5;6;7;8;9;10	 will return 55
	// 	Standard separators are: , and newline
	//
	// scalc: The result is 0

}
