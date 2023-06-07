package main

import "testing"

var tests = []struct {
	name     string
	dividend float32
	divisor  float32
	expected float32
	isErr    bool
}{
	{"valid-data", 100.0, 10.0, 10.0, false},
	{"invalid-data", 100.0, 0.0, 0.0, true},
	{"expect-5", 50.0, 10.0, 5.0, false},
	{"expect-fraction", -1.0, -777.0, 0.001287001287001, false},
}

//there are two ways to test this code:

//1. function test

/*
func TestDivide(t *testing.T) {
	_, err := divide(10.0, 1.0)

	if err != nil {
		t.Error("Got an error when we should not have")
	}
}

func TestBadDivide(t *testing.T) {
	_, err := divide(10.0, 0)

	if err == nil {
		t.Error("Did not get an error when we should have")
	}
}
*/

//2. table test

func TestDivision(t *testing.T) {
	for _, tt := range tests { //tt is very common when you are ranging over tests
		got, err := divide(tt.dividend, tt.divisor)
		if tt.isErr {
			if err == nil {
				t.Error("expected an error but did not get one.")
			}
		} else {
			if err != nil {
				t.Error("did not expect an error but got one", err.Error())
			}
		}

		if got != tt.expected {
			t.Errorf("expected %f but got %f", tt.expected, got)
		}
	}
}

// $go test or $go test -v // to run all tests
// $go test -run TestDivision // to run a specific test
// $go test -cover // to see the coverage report in terminal
// $go test -coverprofile=coverage.out && go tool cover -html=coverage.out // to see the coverage report in browser
