# Simple Client

Simple client is a simple HTTP Client with some added features.

## Install 
```
go get github.com/maxidelo/simpleclient
```

## How to use it

As its name indicates, using it is quite simple. It follows the option pattern for the creation of the Request object and then uses it to execute the request.
It supports automatically JSON marshalling/unmarshalling too.

```GO
type Item struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type ResponseObject struct {
	Items []Item `json:"items"`
}

func main() {
	var response ResponseObject

	headers := map[string]string{
		"Content-Type": "application/json",
		"Accept":       "application/json;charset=UTF-8",
	}

	queryParams := map[string]string{
		"page": "1",
		"size": "25",
	}

	request := simpleclient.NewRequest(
		"https://simpleserver.com/items",
		simpleclient.GET,
		&response,
		simpleclient.WithHeaders(headers),
		simpleclient.WithQueryParams(queryParams),
	)

	if err := simpleclient.Execute(*request); err != nil {
		// do something with the response 
	}
}

```