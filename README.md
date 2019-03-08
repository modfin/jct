

# Json Case Toggle

A simple tool for toggling different cases in json data.

eg. 


```go
j := []byte(`{"ToggleCase": 1})`
j := jsc.Toggle(j, jct.PascalCase(), jsc.SnakeCase()))
fmt.Pringln(string(j))
// {"toggle_case": 1}

```

or

```go
j := []byte(`{"some-thing": ["else", {"but-is": "needed"}]})`
j := jsc.Toggle(j, jct.KebabCase(), jsc.CamelCase()))
fmt.Pringln(string(j))
// {"someThing": ["else", {"butIs": "needed"}]}

```


## Restriction 
Camel and Pascal case contains more information then delimitede cases such as kebab 
or snake case. Camel case contain information about abbreviations which del delimitede cases does not.
eg `URLName -> url_name` and the reverse `url_name -> UrlName`

