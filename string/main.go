package string

import (
	"fmt"
	"unicode"
)

func ContainsSpace(s string) bool {
	for _, r := range s {
        if unicode.IsSpace(r) {
            return true
        }
    }
	return false
}

func ReplaceByPred(s string, new string, predicate func(r rune) bool) string {
	result := ""
	for _, r := range s {
        if predicate(r) {
			result += new
        } else {
			result += string(r)
		}
    }
	return result
}
func IterConvertRunes(s string, convert func(r rune) string) string {
	result := ""
	for _, r := range s {
        result += convert(r)
    }
	return result
}

func Stringify(s string) string {
	ss := IterConvertRunes(s, func(r rune) string {
		if r == '"' || r == '\\' {
			return fmt.Sprintf("\\%c", r)
		}
		return string(r)
	})
	return fmt.Sprintf("\"%v\"", ss)
}
