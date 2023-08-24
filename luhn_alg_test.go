package goluhnalg

import (
	"log"
	"strconv"
	"testing"
)

func TestValidate(t *testing.T) {
	tests := []struct {
		number   string
		mustBeOk bool
	}{
		{"7992739871212", false},
		{"79927398713", true},
		{"4539 3195 0343 6467", true},
		{"8273 1232 7352 0569", false},
	}

	for _, test := range tests {
		t.Run(test.number, func(t *testing.T) {
			res := Validate(test.number)
			if res != nil {
				if test.mustBeOk {
					log.Printf("Luhn check unsuccessful for %s.", test.number)
					t.Fail()
				}
			}
		})
	}
}

func TestCalculate(t *testing.T) {
	tests := []struct {
		number    string
		luhnDigit string
		expected  string
	}{
		{"123456781234567", "0", "1234567812345670"},
		{"111122223333444", "4", "1111222233334444"},
		{"7992739871", "3", "79927398713"},
	}
	for _, test := range tests {
		t.Run(test.expected, func(t *testing.T) {
			l, n, err := Calculate(test.number)
			if err != nil {
				log.Printf("Unexpected err %+v", err)
				t.Fail()
			}
			if test.luhnDigit != l {
				log.Printf("Expected luhn digit %s. Actual luhn digit %s", test.luhnDigit, l)
				t.Fail()
			}
			if n != test.expected {
				log.Printf("Expected %s to generate luhn number %s", test.number, test.expected)
				t.Fail()
			}
			err = Validate(n)
			if err != nil {
				log.Printf("Cannot validate derive luhn number %s. Error: %+v", n, err)
				t.Fail()
			}
		})
	}
}

func TestGenerate(t *testing.T) {
	tests := []struct {
		numberSize int
		sampleSize int
	}{
		{1, 100},
		{10, 1000},
		{100, 1000},
		{1000, 1000},
	}

	for _, test := range tests {
		t.Run(strconv.Itoa(test.numberSize), func(t *testing.T) {
			for i := 0; i < test.sampleSize; i++ {
				err := Validate(Generate(test.numberSize))
				if err != nil {
					log.Printf("Unexpected err %+v", err)
					t.Fail()
				}
			}
		})
	}
}
