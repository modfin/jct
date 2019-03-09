

# JSON Case Toggle

A simple tool for toggling different cases in json data.

eg. 


```go
j := []byte(`{"ToggleCase": 1}`)
j := jct.Toggle(j, jct.PascalCase(), jct.SnakeCase())
fmt.Println(string(j))
// {"toggle_case": 1}

```

or

```go
j := []byte(`{"some-thing": ["else", {"but-is": "needed"}]}`)
j := jct.Toggle(j, jct.KebabCase(), jct.CamelCase())
fmt.Println(string(j))
// {"someThing": ["else", {"butIs": "needed"}]}

```


## Restrictions
Camel and Pascal case contains more information then delimitede cases such as kebab 
or snake case. Camel case contain information about abbreviations which delimitede cases does not.
eg `URLName -> url_name` and the reverse `url_name -> UrlName`

In this sence information is destoyed when converting from Camel or Pascal case to a delimitede case.

