package piglatin

import (
	"strings"
	"unicode"
)

func isVowel(in rune) bool {
	vowels := "aeiouAEIOU"
	return strings.ContainsRune(vowels, in)
}

func encodeWord(in string) string {
	//fmt.Printf("encoding %q\n",in)
	if isVowel(rune(in[0])) {
		return in + "yay"
	} else if !isVowel(rune(in[0])) && !isVowel(rune(in[1])) {
		return in[2:] + in[:2] + "ay"
	} else {
		return in[1:] + in[:1] + "ay"
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
