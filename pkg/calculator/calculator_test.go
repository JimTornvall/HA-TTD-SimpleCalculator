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

func TestSimple_Add_Multiple_Numbers(t *testing.T) {
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
		calc := NewSimple()
		result, err := calc.Add(test.input)
		if err != nil {
			t.Errorf("Error: %v", err)
		}
		if result != test.want {
			t.Errorf("Want: %v, Got: %v", test.want, result)
		}
	}
}
