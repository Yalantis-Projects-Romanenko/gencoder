package numbers

var encodeDictionary = map[string]rune{
	"a": '1',
	"e": '2',
	"i": '3',
	"o": '4',
	"u": '5',

	"A": '1',
	"E": '2',
	"I": '3',
	"O": '4',
	"U": '5',
}

var decodeDictionary = map[string]rune{
	"1": 'a',
	"2": 'e',
	"3": 'i',
	"4": 'o',
	"5": 'u',
}

func Encode(toEncode string) string {
	result := []rune(toEncode)
	for i, char := range toEncode {
		val, ok := encodeDictionary[string(char)]
		if ok {
			result[i] = val
		} else {
			result[i] = char
		}
	}
	return string(result)
}

func Decode(toEncode string) string {
	result := []rune(toEncode)
	for i, char := range toEncode {
		val, ok := decodeDictionary[string(char)]
		if ok {
			result[i] = val
		} else {
			result[i] = char
		}
	}
	return string(result)
}
