package jct

import (
	"encoding/json"
	"testing"
)

var casesToggleTable = []struct {
	in       string
	from     Case
	to       Case
	expected string
}{
	{
		in:       `{"hello-world":{}}`,
		from:     KebabCase(),
		to:       SnakeCase(),
		expected: `{"hello_world":{}}`,
	},
	{
		in:       `{"hello-world":1}`,
		from:     KebabCase(),
		to:       SnakeCase(),
		expected: `{"hello_world":1}`,
	},
	{
		in:       `[{"hello-world":1}]`,
		from:     KebabCase(),
		to:       SnakeCase(),
		expected: `[{"hello_world":1}]`,
	},
	{
		in:       `[{"hello-world":[ {"a-b":[]} ] } ]`,
		from:     KebabCase(),
		to:       SnakeCase(),
		expected: `[{"hello_world":[{"a_b":[]}]}]`,
	},
	{
		in:       `[{"hello-world":["a-string", {"a-b":[]} ] } ]`,
		from:     KebabCase(),
		to:       SnakeCase(),
		expected: `[{"hello_world":["a-string",{"a_b":[]}]}]`,
	},
	{
		in:       `{"hello-world": ["a-b", {"a-b":1}, {"a-c": ["1", 2, 3.0]} ], "some-thing": "any-thing" }`,
		from:     KebabCase(),
		to:       SnakeCase(),
		expected: `{"hello_world":["a-b",{"a_b":1},{"a_c":["1",2,3]}],"some_thing":"any-thing"}`,
	},
	{
		in:       `{"hello-world": ["a-b", {"a-b":1}, {"a-c": ["1", 2, 3.0]} ], "some-thing": "any-thing" }`,
		from:     KebabCase(),
		to:       CamelCase(),
		expected: `{"helloWorld":["a-b",{"aB":1},{"aC":["1",2,3]}],"someThing":"any-thing"}`,
	},
}

func TestToggleCases(t *testing.T) {
	for _, s := range casesToggleTable {
		j := json.RawMessage(s.in)

		got, err := Toggle(j, s.from, s.to)

		if err != nil {
			t.Error(err)
		}

		if string(got) != s.expected {
			t.Error("got", string(got), "expected", s.expected)
		}
	}
}
