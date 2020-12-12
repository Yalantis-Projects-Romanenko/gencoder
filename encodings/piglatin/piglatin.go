package piglatin

func isVowel(in rune) bool {
	switch in {
	case 'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U':
		return true
	default:
		return false
	}
}

func encodeWord(in string) string {
	if isVowel(rune(in[0])) {
		return in + "yay"
	} else if !isVowel(rune(in[0])) && !isVowel(rune(in[1])) {
		return in[2:] + in[:2] + "ay"
	} else {
		return in[1:] + in[:1] + "ay"
	}
}

func Encode(toEncode string) string {
	result := make([]rune, len(toEncode)*2)
	var word string
	var wasSpecial bool
	for _, char := range toEncode {
		if char < 'A' || char > 'z' {
			if !wasSpecial {
				word = encodeWord(word)
				result = append(result, []rune(word)...)
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
		word = encodeWord(word)
		result = append(result, []rune(word)...)
	}

	return string(result)
}
