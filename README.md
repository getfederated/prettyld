# Pretty LD

_The best JSON-LD unmarshalling library for Go._

Wouldn't it be nice if you can just unmarshal JSON-LD documents into a struct? Then Pretty LD is just the library for this!

It's as easy as first defining your structure:

```go
type MyModel struct {
	ID string `json:"@id"`
	Type []string `json:"@type"`
	Name prettyld.String `json:"https://example.com/ns#name"`
}
```

Receiving your JSON-LD input

```go
var j = `
	{
		"@context": {
			"ex": "https://example.com/ns#",
			"name": "ex:name",
			"Person": "ex:Person"
		},
		"@type": "Person",
		"name": "Alice"
	}
`
```

And then parsing it using the `Unmarshal` function

```go
var myModel MyModel

err := prettyld.Unmarshal(j, &myModel, nil)
if err {
	// Do stuff with err, ending things early.
}
```

You should be able to see the output

```go
// Data should be in `myModel`.
fmt.Println(myModel.ID)
fmt.Println(myModel.Type)
fmt.Println(myModel.Name)
```
