package calcdni

import "unicode"

type Word struct {
	List []rune
}

func LetrasListCalculation(index int) rune {
	words := Word{
		List: []rune{
			'T', 'R', 'W', 'A', 'G', 'M', 'Y', 'F', 'P', 'D', 'X', 'B', 'N', 'J', 'Z', 'S', 'Q', 'V', 'H', 'L', 'C', 'K', 'E',
		},
	}
	rest := index % len(words.List)
	word := words.List[rest]
	return word
}

func IsNumber(s string) bool {
	for _, r := range s {
		if !unicode.IsDigit(r) {
			return false
		}
	}
	return true
}
