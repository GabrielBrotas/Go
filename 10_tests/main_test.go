package main

import "testing"

var tests = []struct {
	name     string
	divident float32
	divisor  float32
	expected float32
	isError  bool
}{
	{"valid-data", 100.0, 10.0, 10.0, false},
	{"invalid-data", 100.0, 0.0, 0, true},
}

func TestDivision(t *testing.T) {
	for _, tt := range tests {
		got, err := divide(tt.divident, tt.divisor)

		if tt.isError {
			if err == nil {
				t.Error("Did not get an error when we should have")
			}
		} else {
			if err != nil {
				t.Error("Got an error when we should not have")
			}
		}

		if got != tt.expected {
			t.Errorf("expected %f but got %f", tt.expected, got)
		}

	}
}

// the test must init with Test* otherwise it will be ignored
// func TestDivide(t *testing.T) {
// 	_, err := divide(10.0, 1.0)

// 	if err != nil {
// 		t.Error("Got an error when we should not have")
// 	}
// }

// func TestBadDivide(t *testing.T) {
// 	_, err := divide(10.0, 0)

// 	if err == nil {
// 		t.Error("Did not get an error when we should have")
// 	}
// }
