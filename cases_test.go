package jct

import (
	"reflect"
	"testing"
)

var casesSplitTable = []struct {
	in       string
	expected []string
	_case    Case
}{
	{
		in:       "hello-world",
		expected: []string{"hello", "world"},
		_case:    KebabCase(),
	},
	{
		in:       "hello_world",
		expected: []string{"hello", "world"},
		_case:    SnakeCase(),
	},
	{
		in:       "HelloWorld",
		expected: []string{"hello", "world"},
		_case:    CamelCase(),
	},
	{
		in:       "HelloWorld",
		expected: []string{"hello", "world"},
		_case:    PascalCase(),
	},
	{
		in:       "URL",
		expected: []string{"url"},
		_case:    PascalCase(),
	},
	{
		in:       "URLName",
		expected: []string{"url", "name"},
		_case:    PascalCase(),
	},
	{
		in:       "MSRunner",
		expected: []string{"ms", "runner"},
		_case:    PascalCase(),
	},
	{
		in:       "NameURL",
		expected: []string{"name", "url"},
		_case:    PascalCase(),
	},
	{
		in:       "AbC",
		expected: []string{"ab", "c"},
		_case:    PascalCase(),
	},
	{
		in:       "AbCd",
		expected: []string{"ab", "cd"},
		_case:    PascalCase(),
	},
	{
		in:       "AB",
		expected: []string{"ab"},
		_case:    PascalCase(),
	},
}

var casesJoinTable = []struct {
	in       []string
	expected string
	_case    Case
}{
	{
		in:       []string{"hello", "world"},
		expected: "hello-world",
		_case:    KebabCase(),
	},
	{
		in:       []string{"hello", "world"},
		expected: "hello_world",
		_case:    SnakeCase(),
	},
	{
		in:       []string{"hello", "world"},
		expected: "helloWorld",
		_case:    CamelCase(),
	},
	{
		in:       []string{"hello", "world"},
		expected: "HelloWorld",
		_case:    PascalCase(),
	},
	{
		in:       []string{"a", "b"},
		expected: "AB",
		_case:    PascalCase(),
	},
}

func TestSplitCases(t *testing.T) {
	for _, s := range casesSplitTable {
		got := s._case.Split(s.in)
		if !reflect.DeepEqual(got, s.expected) {
			t.Error("got", got, "expected", s.expected)
		}
	}
}

func TestJoinCases(t *testing.T) {

	for _, s := range casesJoinTable {
		got := s._case.Join(s.in)
		if got != s.expected {
			t.Error("got", got, "expected", s.expected)
		}
	}

}
