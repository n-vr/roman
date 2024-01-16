// Package roman implements a converter from roman numerals to decimal.
package roman

import (
	"errors"
)

var ErrInvalidRomanNumeral = errors.New("invalid roman numeral")

// numerals is a map of roman numerals to their decimal value.
var numerals = map[byte]int{
	'I': 1,
	'V': 5,
	'X': 10,
	'L': 50,
	'C': 100,
	'D': 500,
	'M': 1000,
}

// RomanToDecimal converts a roman numeral to decimal.
// It returns an error if the roman numeral is invalid.
//
// This function expects the roman numeral to be in upper case, and
// roman numerals with a value less then 1 or greater than 3999 are not supported.
//
// This implementation follows two simple rules to allow for flexible input:
//  1. Values of numerals are added together, except when a smaller numeral precedes a larger one,
//     in which case the smaller numeral is subtracted from the larger one.
//  2. The half symbols (V, L, D) can only appear once in a numeral.
//
// source (Dutch): https://nl.wikipedia.org/wiki/Romeinse_cijfers#Oudheid
//
// Edge case: an empty string is considered to be 0 and is not an error.
func RomanToDecimal(roman string) (int, error) {
	if roman == "" {
		return 0, nil
	}

	halfSymbols := newHalfSymbolCounter()

	lastSymbol := roman[len(roman)-1]
	lastValue, ok := numerals[lastSymbol]
	if !ok {
		return 0, ErrInvalidRomanNumeral
	}

	err := halfSymbols.check(lastSymbol)
	if err != nil {
		return 0, err
	}

	// The last numeral is always added.
	// This is done outside of the loop to simplify the code
	// without getting an index out of range error.
	sum := lastValue

	// Loop over the remaining roman numeral from right to left to make it easier to check for rule 1.
	for i := len(roman) - 2; i >= 0; i-- {
		current, ok := numerals[roman[i]]
		if !ok {
			return 0, ErrInvalidRomanNumeral
		}

		err := halfSymbols.check(roman[i])
		if err != nil {
			return 0, err
		}

		next, ok := numerals[roman[i+1]]
		if !ok {
			return 0, ErrInvalidRomanNumeral
		}

		if current < next {
			sum -= current
		} else {
			sum += current
		}
	}

	return sum, nil
}

// halfSymbolCounter is a counter for the half symbols (V, L, D).
// It is used to check if a half symbol has already been used.
// This is needed to check for rule 2.
//
// empty structs are used instead of bools because they don't take up any memory.
type halfSymbolCounter map[byte]struct{}

// newHalfSymbolCounter returns a new halfSymbolCounter.
func newHalfSymbolCounter() halfSymbolCounter {
	return make(halfSymbolCounter, 3)
}

// check checks if the symbol is a half symbol (V, L, D).
// if the half symbol has already been used, an error is returned.
func (h halfSymbolCounter) check(symbol byte) error {
	if symbol != 'V' && symbol != 'L' && symbol != 'D' {
		return nil
	}

	if _, ok := h[symbol]; ok {
		return ErrInvalidRomanNumeral
	}

	// Add the symbol to the counter.
	h[symbol] = struct{}{}

	return nil
}
