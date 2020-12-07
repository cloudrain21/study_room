package numeral

import (
    "strings"
)

type RomanNumeral struct {
    Value int
    Symbol string
}

var allRomanNumerals = romanNumerals {
    {1000, "M"},
    {900, "CM"},
    {500, "D"},
    {400, "CD"},
    {100, "C"},
    {90, "XC"},
    {50, "L"},
    {40, "XL"},
    {10, "X"},
    {9,  "IX"},
    {5,  "V"},
    {4,  "IV"},
    {1,  "I"},
}

func ConvertToRoman(a int) string {
    var s strings.Builder

    for _, n := range allRomanNumerals {
        for a >= n.Value {
            s.WriteString(n.Symbol)
            a -= n.Value
        }
    }

    return s.String()
}

func ConvertToArabic(roman string) (total int) {
    for _, symbols := range windowedRoman(roman).Symbols() {
        total += allRomanNumerals.ValueOf(symbols...)
    }
    return
}

type romanNumeral struct {
    Value  int
    Symbol string
}

type romanNumerals []romanNumeral

func (r romanNumerals) ValueOf(symbols ...byte) int {
    symbol := string(symbols)
    for _, s := range r {
        if s.Symbol == symbol {
            return s.Value
        }
    }

    return 0
}

func (r romanNumerals) Exists(symbols ...byte) bool {
    symbol := string(symbols)
    for _, s := range r {
        if s.Symbol == symbol {
            return true
        }
    }
    return false
}

type windowedRoman string

func (w windowedRoman) Symbols() (symbols [][]byte) {
    for i := 0; i < len(w); i++ {
        symbol := w[i]
        notAtEnd := i+1 < len(w)

        if notAtEnd && isSubtractive(symbol) && allRomanNumerals.Exists(symbol, w[i+1]) {
            symbols = append(symbols, []byte{byte(symbol), byte(w[i+1])})
            i++
        } else {
            symbols = append(symbols, []byte{byte(symbol)})
        }
    }
    return
}

func isSubtractive(symbol uint8) bool {
    return symbol == 'I' || symbol == 'X' || symbol == 'C'
}
