package strings

import (
	"fmt"
	"regexp"
	"strings"
	"unicode"
	"unicode/utf8"
)

// Camelize camelizes a string
func Camelize(word string) string {
	return camelize(split(word), func(word string) string {
		word = strings.ToLower(word)
		fchar, _ := utf8.DecodeLastRuneInString(word[:1])
		return fmt.Sprintf("%s%s", string(unicode.ToUpper(fchar)), word[1:])
	})
}

// Pathify converts a word to a restfull resource path
func Pathify(word string) string {
	return Pluralize(strings.ToLower(strings.Join(split(word), "-")))
}

// Underscored split camel cased and hyphenated word and replace
// with underscores(_)
func Underscored(word string) string {
	return strings.Join(split(word), "_")
}

// Hyphenized split camel cased and hyphenated word and replace
// with hyphens(_)
func Hyphenized(word string) string {
	return strings.Join(split(word), "-")
}

// List converts a word name into a list name
// e.g: article becomes articles
func List(word string) string {
	return Pluralize(word)
}

// Objectify converts a word into a souce code object name
func Objectify(word string) string {
	return strings.ToLower(strings.Join(split(word), "_"))
}

// Pluralize pluralizes a word
func Pluralize(word string) string {
	ending, _ := utf8.DecodeLastRuneInString(word)
	switch ending {
	case 'y', 'Y', 'i', 'I':
		return fmt.Sprintf("%s%s", word[:len(word)-1], "ies")
	case 's', 'S':
		return word
	default:
		return fmt.Sprintf("%s%s", word, "s")
	}
}

// HasSpaces check if word has white spaces
func HasSpaces(word string) bool {
	return len(strings.Split(word, " ")) > 1
}

// SingleCasedWord determines if a word is made up of all uppercase
// or all lowercase characters
func SingleCasedWord(word string) bool {
	return AllUpperCase(word) || AllLowerCase(word)
}

// AllUpperCase determines if a word is made up of all uppcase
func AllUpperCase(word string) bool {
	result, _ := regexp.MatchString("^[A-Z]+$", word)
	return result
}

// AllLowerCase determines if a word is made up of all lowercase characters
func AllLowerCase(word string) bool {
	result, _ := regexp.MatchString("^[a-z]+$", word)
	return result
}

func camelize(words []string, upper func(word string) string) string {
	var _word string
	for _, val := range words {
		_word += upper(val)
	}
	return _word
}

func split(word string) []string {
	regx1, _ := regexp.Compile("[^a-zA-Z]+")
	_words := regx1.Split(word, -1)
	words := make([]string, 0)
	for _, val := range _words {
		if SingleCasedWord(val) {
			words = append(words, strings.ToLower(val))
		} else {
			words = append(words, strings.Fields(next(val, "", 0))[:]...)
		}
	}
	return words
}

func next(word, prev string, level int) string {
	if _next, pos := nextUpperCase(word); pos > 0 {
		if level == 0 {
			prev += _next
		} else {
			prev += fmt.Sprintf(" %s", _next)
		}
		return next(word[pos:], prev, level+1)
	}
	return fmt.Sprintf("%s %s", prev, word)

}

func nextUpperCase(word string) (string, int) {
	for i, char := range word[1:] {
		if unicode.IsUpper(char) {
			return word[:i+1], i + 1
		}
	}
	return word, -1
}
