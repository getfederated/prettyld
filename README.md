# Pretty LD

_The best library for processing JSON-LD documents. Make parsing and interpreting as intuitive as parsing just plain JSON_

Working with JSON-LD is a pain. You have to expand the document, extract the first node, then interpret business domain fields as "predicates" to "objects". And these predicates aren't just familiar field names; they're URLs (IRIs, actually). And even if you absolutely know that a particular business domain field will legally always have a single node associated with the field, with JSON-LD, you will always get a slice of `any`. So verbose. So much work, for so little reward.

This is where Pretty LD comes along: you can just interpret JSON-LD documents as if they were plain old JSON documents!

It's as easy as this:

```go
type MyModel struct {
	ID string `json:"@id"`
	Type []string `json:"@type"`
	Name prettyld.String `json:"https://example.com/ns#name"`
}

var j = `
	{
		"@context": {
			"ex": "https://example.com/ns#",
			"name": "ex:name"
		}
	}
`

var myModel MyModel
err := prettyld.Unmarshal(j, &myModel, nil)
if err {
	// Do stuff with err, ending things early.
}

// Data should be in `myModel`.
```
