package jct

import (
	"encoding/json"
	"testing"
	"fmt"
)

var casesToggleTable = []struct {
	in       string
	from     Case
	to       Case
	expected string
}{
	{
		in:       `{"HelloWorld":{}}`,
		from:     PascalCase(),
		to:       CamelCase(),
		expected: `{"helloWorld":{}}`,
	},
	{
		in:       `{"ABC":{}}`,
		from:     PascalCase(),
		to:       CamelCase(),
		expected: `{"abc":{}}`,
	},
	{
		in:       `{"AbC":{}}`,
		from:     PascalCase(),
		to:       CamelCase(),
		expected: `{"abC":{}}`,
	},
	{
		in:       `{"AbCd":{}}`,
		from:     PascalCase(),
		to:       CamelCase(),
		expected: `{"abCd":{}}`,
	},
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
		expected: `{"helloWorld": ["a-b",{"aB":1},{"aC":["1",2,3]}],"someThing":"any-thing"}`,
	},
	{
		in:       `{"hello-world": ["a-b", {"a-b":1}, {"a-c": ["1", 2, [[{"k-l": 1}]]]} ], "some-thing": "any-thing" }`,
		from:     KebabCase(),
		to:       CamelCase(),
		expected: `{"helloWorld": ["a-b", {"aB":1}, {"aC": ["1", 2, [[{"kL": 1}]]]} ], "someThing": "any-thing" }`,
	},
	{
		in:       `{"hello-world":"{\"do-not-toggel\": \"some other text\"}"}`,
		from:     KebabCase(),
		to:       SnakeCase(),
		expected: `{"hello_world":"{\"do-not-toggel\": \"some other text\"}"}`,
	},
	{
		in:       `[{"hello-world":"{\"do-not-toggel\": [1,2, {\"a-b\": 1} ]}"}]`,
		from:     KebabCase(),
		to:       SnakeCase(),
		expected: `[{"hello_world":"{\"do-not-toggel\": [1,2, {\"a-b\": 1} ]}"}]`,
	},
	{
		in:       `[{"hello-world":"{\"do-not-toggel\": [1,2, {\"a-b\": 1} ]}"}]`,
		from:     KebabCase(),
		to:       CamelCase(),
		expected: `[{"helloWorld":"{\"do-not-toggel\": [1,2, {\"a-b\": 1} ]}"}]`,
	},
}

func TestToggleCases(t *testing.T) {
	for _, s := range casesToggleTable {
		j := []byte(s.in)

		got := Toggle(j, s.from, s.to)

		var tmp0 interface{}
		err := json.Unmarshal([]byte(got), &tmp0)
		if err != nil {
			t.Error(err)
		}
		gdata, err := json.Marshal(tmp0)
		if err != nil {
			t.Error(err)
		}

		var tmp1 interface{}
		err = json.Unmarshal([]byte(s.expected), &tmp1)
		if err != nil {
			t.Error(err)
		}

		edata, err := json.Marshal(tmp1)
		if err != nil {
			t.Error(err)
		}

		if string(gdata) != string(edata) {
			t.Error("got", string(gdata), "expected", string(edata))
		}
	}
}

func BenchmarkToggleToDel(b *testing.B) {
	j := []byte(`{"hello-world": ["a-b", "{\"a-b\": 123}", {"a-b":1}, {"a-c": ["1", 2, [[{"k-l": 1}]]]} ], "some-thing": "any-thing" }`)

	from := KebabCase()
	to := SnakeCase()

	for n := 0; n < b.N; n++ {
		Toggle(j, from, to)
	}
}

func BenchmarkToggleToCamel(b *testing.B) {
	j := []byte(`{"hello-world": ["a-b", "{\"a-b\": 123}", {"a-b":1}, {"a-c": ["1", 2, [[{"k-l": 1}]]]} ], "some-thing": "any-thing" }`)

	from := KebabCase()
	to := CamelCase()

	for n := 0; n < b.N; n++ {
		Toggle(j, from, to)
	}
}

func BenchmarkTogglePascalToCamel(b *testing.B) {
	j := []byte(`{"HelloWorld": ["a-b", "{\"a-b\": 123}", {"AaB":1}, {"AaC": ["1", 2, [[{"KkL": 1}]]]} ], "SomeThing": "AnyThing" }`)

	from := PascalCase()
	to := CamelCase()

	for n := 0; n < b.N; n++ {
		Toggle(j, from, to)
	}
}


func aa(){
	j := []byte(`{"ToggleCase": 1}`)
	j := Toggle(j, PascalCase(), SnakeCase())
	fmt.Println(string(j))
	// {"toggle_case": 1}

		j := []byte(`{"some-thing": ["else", {"but-is": "needed"}]}`)
	j := Toggle(j, KebabCase(), CamelCase()))
	fmt.Println(string(j))
	// {"someThing": ["else", {"butIs": "needed"}]}
}
