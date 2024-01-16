package roman_test

import (
	"fmt"
	"testing"

	"github.com/n-vr/roman"
)

func ExampleRomanToDecimal() {
	value, err := roman.RomanToDecimal("MMXXIV")
	if err != nil {
		panic(err)
	}

	fmt.Println(value)
	// Output: 2024
}

func TestRomanToDecimal_valid_cases(t *testing.T) {
	testCases := []struct {
		roman string
		want  int
	}{
		{"I", 1},
		{"II", 2},
		{"IV", 4},
		{"V", 5},
		{"VI", 6},
		{"IX", 9},
		{"X", 10},
		{"XIV", 14},
		{"XLV", 45},
		{"XLIX", 49},
		{"XCV", 95},
		{"CXXVI", 126},
		{"MCMXCVIII", 1998},
		{"MMXIV", 2014},
		{"MMXVIII", 2018},
		{"MMMXLV", 3045},
		{"MMMCMXCIX", 3999},
	}

	for _, tc := range testCases {
		t.Run(tc.roman, func(t *testing.T) {
			got, err := roman.RomanToDecimal(tc.roman)
			if err != nil {
				t.Fatalf("roman.RomanToDecimal(%q) err = %v, want nil", tc.roman, err)
			}

			if got != tc.want {
				t.Errorf("roman.RomanToDecimal(%q) = %d, want %d", tc.roman, got, tc.want)
			}
		})
	}
}

func TestRomanToDecimal_invalid_cases(t *testing.T) {
	testCases := []struct {
		name  string
		roman string
	}{
		{"InvalidSymbol", "A"},
		{"DoubleHalfSymbol", "XVV"},
		{"DoubleHalfSymbol", "CLL"},
		{"DoubleHalfSymbol", "MDD"},
	}

	for _, tc := range testCases {
		t.Run(tc.name+"_"+tc.roman, func(t *testing.T) {
			value, err := roman.RomanToDecimal(tc.roman)
			if err == nil {
				t.Fatalf("roman.RomanToDecimal(%q) value = %d, err = nil, want ErrInvalidRomanNumeral", tc, value)
			}
		})
	}
}
