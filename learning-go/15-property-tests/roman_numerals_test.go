package roman_numerals

import (
	"fmt"
	"testing"
	"testing/quick"
)


var cases = []struct {
    Arabic uint16
    Roman string
}{
    {1, "I"},
    {2, "II"},
    {3, "III"},
    {4, "IV"},
    {5, "V"},
    {6, "VI"},
    {7, "VII"},
    {9, "IX"},
    {10, "X"},
    {11, "XI"},
    {14, "XIV"},
    {15, "XV"},
    {20, "XX"},
    {24, "XXIV"},
    {25, "XXV"},
    {39, "XXXIX"},
    {40, "XL"},
    {47, "XLVII"},
    {49, "XLIX"},
    {50, "L"},
    {90, "XC"},
    {100, "C"},
    {90, "XC"},
    {400, "CD"},
    {500, "D"},
    {900, "CM"},
    {1000, "M"},
    {1984, "MCMLXXXIV"},
    {3999, "MMMCMXCIX"},
    {2014, "MMXIV"},
    {1006, "MVI"},
    {798, "DCCXCVIII"},
}

func TestRomanNumerals(t *testing.T) {
	for _, test := range cases {
		t.Run(fmt.Sprintf("%d gets converted to %s", test.Arabic, test.Roman), func(t *testing.T) {
			got, err := ConvertToRoman(test.Arabic)
            if err != nil {
                t.Fatal("Unexpected error")
            }
			if got != test.Roman {
				t.Errorf("got %s, want %s", got, test.Roman)
			}
		})
	}
}

func TestConvertingToArabic(t *testing.T) {
    for _, test := range cases {
		t.Run(fmt.Sprintf("%s gets converted to %d", test.Roman, test.Arabic), func(t *testing.T) {
			got := ConvertToArabic(test.Roman)
			if got != test.Arabic {
				t.Errorf("got %d, want %d", got, test.Arabic)
			}
		})
	}
}

func TestPropertiesOfConversion(t *testing.T) {
    assertion := func(arabic uint16) bool {
        if arabic > 3999 {
            return true
        }
        roman, _ := ConvertToRoman(arabic)
        fromRoman := ConvertToArabic(roman)
        return fromRoman == arabic
    }

    if err := quick.Check(assertion, &quick.Config{
        MaxCount: 1000,
    }); err != nil {
        t.Error("failed checks", err)
    }
}
