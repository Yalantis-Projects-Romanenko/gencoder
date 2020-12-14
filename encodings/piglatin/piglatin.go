package piglatin

import (
	"strings"
	"unicode"
)

func getPrefixIndex(word string) int {
	vowels := "aeiouAEIOU"
	for i, c := range word {
		if strings.ContainsRune(vowels, c) {
			return i
		}
	}
	return len(word)
}

func encodeWord(in string) string {
	prefixIndex := getPrefixIndex(in)
	if prefixIndex == 0 {
		return in + "yay"
	}
	return in[prefixIndex:] + in[:prefixIndex] + "ay"

}

// Encode encodes string to piglatin
func Encode(toEncode string) string {
	result := make([]rune, 0, len(toEncode)*2)
	var word string
	var wasLetter bool // flag that marks if last char in text was a letter

	for _, char := range toEncode {
		if !unicode.IsLetter(char) {
			if wasLetter {
				result = append(result, []rune(encodeWord(word))...)
			}
			result = append(result, char)
			wasLetter = false
		} else {
			// reset word string after not letter sign
			if !wasLetter {
				word = ""
			}
			wasLetter = true
			word = word + string(char)
		}
	}

	// add word last word if there was no special sign in the end of the wor
	if wasLetter {
		result = append(result, []rune(encodeWord(word))...)
	}

	return strings.TrimSpace(string(result))
}
