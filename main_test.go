package main

import (
	"SimpleCalculator/pkg/calculator"
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"io"
	"os"
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

func HelperCaptureOutput(f func()) (string, error) {
	orig := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w
	f()
	os.Stdout = orig
	err := w.Close()
	if err != nil {
		return "", err
	}
	out, _ := io.ReadAll(r)
	return string(out), err
}

func HelperExpectedOutput(result int) string {
	expected := fmt.Sprintf(`Welcome to the Simple Calculator! 🧮

	Usage: Any numbers you write will be added together
		Ex: 1,2,3,4	 will return 10
		Ex: 1,2,3,4,5,6,7,8,9,10	 will return 55
	You can also use a custom separator
		Ex: //;<newline>1;2;3;4;5;6;7;8;9;10	 will return 55
	Standard separators are: , and newline

scalc: The result is %d
`, result)
	return expected
}

func (suite *MainSuite) Test_Main_No_Input() {
	input := "\n"
	expected := HelperExpectedOutput(0)
	reader := strings.NewReader(input)

	output, err := HelperCaptureOutput(func() {
		NewCalc(reader, suite.calc)
	})

	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), expected, output)
}

func (suite *MainSuite) Test_Main_Single_Number() {
	input := "1\n\n"
	expected := HelperExpectedOutput(1)
	reader := strings.NewReader(input)

	output, err := HelperCaptureOutput(func() {
		NewCalc(reader, suite.calc)
	})

	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), expected, output)
}

func (suite *MainSuite) Test_Main_Multiple_Numbers() {
	input := "1,2,3,4,5,\n6,7,8,9,10\n\n"
	expected := HelperExpectedOutput(55)
	reader := strings.NewReader(input)

	output, err := HelperCaptureOutput(func() {
		NewCalc(reader, suite.calc)
	})

	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), expected, output)
}

func (suite *MainSuite) Test_Main_Custom_Separator() {
	input := "//sep\n1sep2sep3sep\n4sep5sep6sep7sep8sep9sep10\n\n"
	expected := HelperExpectedOutput(55)
	reader := strings.NewReader(input)

	output, err := HelperCaptureOutput(func() {
		NewCalc(reader, suite.calc)
	})

	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), expected, output)
}

// test for complex input with custom separators
// input: ‘//[***][%%%]\n1***2%%%4’"
// expected output: 7
func (suite *MainSuite) Test_Main_Complex_Custom_Separator() {
	input := "//[***][%%%]\n1***2%%%4\n\n"
	expected := HelperExpectedOutput(7)
	reader := strings.NewReader(input)

	output, err := HelperCaptureOutput(func() {
		NewCalc(reader, suite.calc)
	})

	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), expected, output)
}
