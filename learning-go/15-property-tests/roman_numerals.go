package roman_numerals

import (
	"errors"
	"strings"
)

var AboveAllowedError = errors.New("Input above allowed limit")

func ConvertToRoman(arabic uint16) (string, error) {
    if (arabic > 3999) {
        return "", AboveAllowedError
    }
	var result strings.Builder

	for _, numeral := range allRomanNumerals {
		for arabic >= numeral.Value {
			result.WriteString(numeral.Symbol)
			arabic -= numeral.Value
		}
	}

	return result.String(), nil
}

func ConvertToArabic(roman string) (total uint16) {
	numToVals := allRomanNumerals.RomanToValueMap()
	for _, symbols := range windowedRoman(roman).Symbols(numToVals) {
		total += numToVals.ValueOf(symbols...)
	}
	return total
}

type RomanNumeral struct {
	Value  uint16
	Symbol string
}

type RomanNumerals []RomanNumeral

var allRomanNumerals = RomanNumerals{
	{1000, "M"},
	{900, "CM"},
	{500, "D"},
	{400, "CD"},
	{100, "C"},
	{90, "XC"},
	{50, "L"},
	{40, "XL"},
	{10, "X"},
	{9, "IX"},
	{5, "V"},
	{4, "IV"},
	{1, "I"},
}

type RomanNumeralMap map[string]uint16

func (r RomanNumerals) RomanToValueMap() RomanNumeralMap {
	romanNumeralMap := make(RomanNumeralMap)
	for _, numeral := range r {
		romanNumeralMap[numeral.Symbol] = numeral.Value
	}
	return romanNumeralMap
}

func (r RomanNumeralMap) ValueOf(symbol ...byte) uint16 {
	val, exists := r[string(symbol)]
	if exists {
		return val
	}
	return 0
}

type windowedRoman string

func (w windowedRoman) Symbols(romanNumeralMap RomanNumeralMap) (symbols [][]byte) {
	for i := 0; i < len(w); i++ {
		symbol := w[i]
		notAtEnd := i+1 < len(w)

		if notAtEnd && isSubtractiveSymbol(symbol) && romanNumeralMap.ValueOf(symbol, w[i+1]) > 0 {
			symbols = append(symbols, []byte{symbol, w[i+1]})
			i++
		} else {
			symbols = append(symbols, []byte{symbol})
		}
	}

	return symbols
}

func isSubtractiveSymbol(symbol byte) bool {
	return symbol == 'I' || symbol == 'X' || symbol == 'C'
}
