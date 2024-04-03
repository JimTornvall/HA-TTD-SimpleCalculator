package calculator

import "testing"

func TestSimple_Add_Empty_String(t *testing.T) {
	input := ""
	want := 0
	calc := NewSimple()
	result, err := calc.Add(input)
	if err != nil {
		t.Errorf("Error: %v", err)
	}
	if result != want {
		t.Errorf("Want: %v, Got: %v", want, result)
	}
}

func TestSimple_Add_Single_Number(t *testing.T) {
	input := "1"
	want := 1
	calc := NewSimple()
	result, err := calc.Add(input)
	if err != nil {
		t.Errorf("Error: %v", err)
	}
	if result != want {
		t.Errorf("Want: %v, Got: %v", want, result)
	}
}

func TestSimple_Add_Two_Numbers(t *testing.T) {
	input := "1,2"
	want := 3
	calc := NewSimple()
	result, err := calc.Add(input)
	if err != nil {
		t.Errorf("Error: %v", err)
	}
	if result != want {
		t.Errorf("Want: %v, Got: %v", want, result)
	}
}
