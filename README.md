## go-fastapi

go-fastapi is a library to quickly build APIs. It is inspired by Python's popular [FastAPI](https://github.com/tiangolo/fastapi) library.

Features:
* Auto-generated OpenAPI/Swagger schema without any markup
* Declare handlers using types, not just `Context`
* [gin](https://github.com/gin-gonic/gin)-based

## Example

```go
package main

import (
	"github.com/gin-gonic/gin"
	"github.com/sashabaranov/fastapi"

	"encoding/json"
	"fmt"
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
