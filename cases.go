package jct

import (
	"strings"
	"unicode"
)

type Case interface {
	Split(key string) (words []string)
	Join(words []string) (key string)
}

type DelCase struct {
	Delimiter string
}

func (d DelCase) Split(key string) []string {
	return strings.Split(strings.ToLower(key), d.Delimiter)
}
func (d DelCase) Join(words []string) string {
	return strings.Join(words, d.Delimiter)
}

func SnakeCase() Case {
	return DelCase{"_"}
}
func KebabCase() Case {
	return DelCase{"-"}
}
func DotCase() Case {
	return DelCase{"."}
}

func PascalCase() Case {
	return pascalCase{}
}

type pascalCase struct{}

func (pascalCase) Split(key string) (words []string) {

	add := func(r []rune) {
		words = append(words, strings.ToLower(string(r)))
	}

	var buff []rune
	rkey := []rune(key)
	rlen := len(rkey)
	lastWasUpper := true
	for i := range rkey {
		var split bool
		var isLast = i+1 == rlen

		var cur = rkey[i]
		var next rune
		if !isLast {
			next = rkey[i+1]
		}

		curIsUpper := unicode.IsUpper(cur)
		nextIsUpper := unicode.IsUpper(next)

		if curIsUpper && len(buff) > 0 {
			if !lastWasUpper {
				split = true
				goto breakif
			}
			if isLast {
				goto breakif
			}
			if lastWasUpper && !nextIsUpper {
				split = true
				goto breakif
			}
		}
	breakif:
		if split {
			add(buff)
			buff = []rune{}
		}
		buff = append(buff, cur)
		lastWasUpper = curIsUpper
	}
	add(buff)

	return words
}
func (pascalCase) Join(words []string) string {
	var buff []rune
	for _, w := range words {
		word := []rune(strings.ToLower(w))
		if len(word) > 0 {
			word[0] = unicode.ToUpper(word[0])
		}
		buff = append(buff, word...)

	}
	return string(buff)
}

func CamelCase() Case {
	return camelCase{}
}

type camelCase struct {
	pascalCase
}

func (c camelCase) Join(words []string) string {
	key := []rune(c.pascalCase.Join(words))

	if unicode.IsUpper(key[0]) {
		key[0] = unicode.ToLower(key[0])
	}

	return string(key)
}
