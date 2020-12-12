package piglatin

import (
	"strings"
	"unicode"
)

func getVowelIndex(word string) int {
	vowels := "aeiouAEIOU"
	for i, c := range word {
		if strings.ContainsRune(vowels, c) {
			return i
		}
	}
	return len(word)
}

func encodeWord(in string) string {
	index := getVowelIndex(in)
	if index == 0 {
		return in + "yay"
	} else {
		return in[index:] + in[:index] + "ay"
	}
}

func Encode(toEncode string) string {
	result := make([]rune, 0, len(toEncode)*2)
	var word string
	var wasSpecial bool

	for _, char := range toEncode {
		if !unicode.IsLetter(char) {
			if !wasSpecial {
				result = append(result, []rune(encodeWord(word))...)
			}
			result = append(result, char)
			wasSpecial = true
		} else {
			if wasSpecial {
				word = ""
			}
			wasSpecial = false
			word = word + string(char)
		}

	}

	if !wasSpecial {
		result = append(result, []rune(encodeWord(word))...)
	}

	return strings.TrimSpace(string(result))
}
