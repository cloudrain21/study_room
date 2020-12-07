package numeral

import (
    //"log"
    "testing"
    "testing/quick"
)

var cases = []struct {
    Description string
    Arabic int
    Roman string
} {
    {"test1", 1, "I"},
    {"test2", 2, "II"},
    {"test3", 3, "III"},
    {"test4", 4, "IV"},
    {"test5", 5, "V"},
    {"test6", 6, "VI"},
    {"test8", 8, "VIII"},
    {"test9", 9, "IX"},
    {"test10", 10, "X"},
    {"test14", 14, "XIV"},
    {"test18", 18, "XVIII"},
    {"test20", 20, "XX"},
    {"test39", 39, "XXXIX"},
    {"test40", 40, "XL"},
    {"test47", 47, "XLVII"},
    {"test49", 49, "XLIX"},
    {"test50", 50, "L"},
    {"test90", 90, "XC"},
    {"test99", 99, "XCIX"},
    {"test100", 100, "C"},
    {"test400", 400, "CD"},
    {"test499", 499, "CDXCIX"},
    {"test500", 500, "D"},
    {"test900", 900, "CM"},
    {"test999", 999, "CMXCIX"},
    {"test1000", 1000, "M"},
    {"test1984", 1984, "MCMLXXXIV"},
}

func TestRomanNumerals(t *testing.T) {
    t.Run("test_1", func(t *testing.T) {
        got := ConvertToRoman(1)
        want := "I"

        if got != want {
            t.Errorf("got (%q) want (%q)", got, want)
        }
    })

    t.Run("test_2", func(t *testing.T) {
        got := ConvertToRoman(2)
        want := "II"

        if got != want {
            t.Errorf("got (%q) want (%q)", got, want)
        }
    })

    for i, v := range cases {
        t.Run(v.Description, func(t *testing.T) {
            got := ConvertToRoman(v.Arabic)
            want := v.Roman

            if got != want {
                t.Errorf("[%d] got (%q) want (%q)", i, got, want)
            }
        })
    }
}

func TestConvertingToArabic(t *testing.T) {
    for i, v := range cases[:1] {
        t.Run(v.Description, func(t *testing.T) {
            got := ConvertToArabic(v.Roman)
            want := v.Arabic

            if got != want {
                t.Errorf("[%d] got (%d) want (%d)", i, got, want)
            }
        })
    }
}

func TestPropertiesOfConversion(t *testing.T) {
    assertion := func(arabic uint16) bool {
        if arabic > 1984 {
            //log.Println(arabic)
            return true
        }

        roman := ConvertToRoman(int(arabic))
        fromRoman := ConvertToArabic(roman)
        return fromRoman == int(arabic)
    }

    if err := quick.Check(assertion, &quick.Config{
        MaxCount:1000,
    }); err != nil {
        t.Error("failed check", err)
    }
}
