## go-fastapi
[![Go Reference](https://pkg.go.dev/badge/github.com/sashabaranov/go-fastapi.svg)](https://pkg.go.dev/github.com/sashabaranov/go-fastapi)
[![Go Report Card](https://goreportcard.com/badge/github.com/sashabaranov/go-gpt3)](https://goreportcard.com/report/github.com/sashabaranov/go-fastapi)


go-fastapi is a library to quickly build APIs. It is inspired by Python's popular [FastAPI](https://github.com/tiangolo/fastapi) library.

Features:
* Auto-generated OpenAPI/Swagger schema without any markup
* Declare handlers using types, not just `Context`
* Based on [gin](https://github.com/gin-gonic/gin) framework

Installation: `go get github.com/sashabaranov/go-fastapi`

## Example

```go
package main

import (
	"github.com/gin-gonic/gin"
	"github.com/sashabaranov/go-fastapi"
)

type EchoInput struct {
	Phrase string `json:"phrase"`
}

type EchoOutput struct {
	OriginalInput EchoInput `json:"original_input"`
}

func EchoHandler(ctx *gin.Context, in EchoInput) (out EchoOutput, err error) {
	out.OriginalInput = in
	return
}

func main() {
	r := gin.Default()

	myRouter := fastapi.NewRouter()
	myRouter.AddCall("/echo", EchoHandler)

	r.POST("/api/*path", myRouter.GinHandler) // must have *path parameter
	r.Run()
}

// Try it:
//     $ curl -H "Content-Type: application/json" -X POST --data '{"phrase": "hello"}' localhost:8080/api/echo
//     {"response":{"original_input":{"phrase":"hello"}}}
```

To generate OpenAPI/Swagger schema:

```go
myRouter := fastapi.NewRouter()
myRouter.AddCall("/echo", EchoHandler)

swagger := myRouter.EmitOpenAPIDefinition()
swagger.Info.Title = "My awesome API"
jsonBytes, _ := json.MarshalIndent(swagger, "", "    ")
fmt.Println(string(jsonBytes))
```

<img src="https://user-images.githubusercontent.com/677093/146807480-be53b3fb-6de8-451f-8373-e8d6da54a032.png" width="400px" height="auto">
