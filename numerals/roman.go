package numerals

import (
	"errors"
)

var (
	errorInvalidRomanNumeral = errors.New("invalid roman numeral")
	// https://medium.com/nerd-for-tech/leetcode-roman-to-integer-94db5376ce3
	romanMap = map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}
	// https://rosettacode.org/wiki/Roman_numerals/Encode#Go
	m0 = []string{"", "I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX"}
	m1 = []string{"", "X", "XX", "XXX", "XL", "L", "LX", "LXX", "LXXX", "XC"}
	m2 = []string{"", "C", "CC", "CCC", "CD", "D", "DC", "DCC", "DCCC", "CM"}
	m3 = []string{"", "M", "MM", "MMM", "I̅V̅", "V̅", "V̅I̅", "V̅I̅I̅", "V̅I̅I̅I̅", "I̅X̅"}
	m4 = []string{"", "X̅", "X̅X̅", "X̅X̅X̅", "X̅L̅", "L̅", "L̅X̅", "L̅X̅X̅", "L̅X̅X̅X̅", "X̅C̅"}
	m5 = []string{"", "C̅", "C̅C̅", "C̅C̅C̅", "C̅D̅", "D̅", "D̅C̅", "D̅C̅C̅", "D̅C̅C̅C̅", "C̅M̅"}
	m6 = []string{"", "M̅", "M̅M̅", "M̅M̅M̅"}
)

func Itor(i int) string {
	return m6[i/1e6] + m5[i%1e6/1e5] + m4[i%1e5/1e4] + m3[i%1e4/1e3] + m2[i%1e3/1e2] + m1[i%100/10] + m0[i%10]
}

func Rtoi(s string) (int, error) {
	length := len(s)

	if length == 0 {
		return 0, errorInvalidRomanNumeral
	}

	if length == 1 {
		if _, exist := romanMap[s[0]]; exist {
			return romanMap[s[0]], nil
		}
		return 0, errorInvalidRomanNumeral
	}

	sum := romanMap[s[length-1]]

	for i := length - 2; i >= 0; i-- {
		if _, exist := romanMap[s[i+1]]; !exist {
			return 0, errorInvalidRomanNumeral
		}
		if romanMap[s[i]] < romanMap[s[i+1]] {
			sum -= romanMap[s[i]]
		} else {
			sum += romanMap[s[i]]
		}
	}

	return sum, nil
}
