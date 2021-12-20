package fastapi

import (
	"github.com/gin-gonic/gin"
	"testing"
)

type InnerStruct struct {
	XYZ string `json:"XYZ"`
}

type In struct {
	Input string          `json:"input"`
	X     int             `json:"x"`
	Y     float32         `json:"y"`
	Z     bool            `json:"z"`
	I     []string        `json:"i"`
	J     map[string]int8 `json:"j"`
}

type In2 struct {
	InputTwo string      `json:"input_two"`
	Inner    InnerStruct `json:"-"`
}

type Out struct {
	Output string `json:"output"`
}

func RequestHandler(ctx *gin.Context, in In) (out Out, err error) {
	return
}

func RequestHandlerTwo(ctx *gin.Context, in In2) (out Out, err error) {
	return
}

func TestOpenAPIDefinition(t *testing.T) {
	myRouter := NewRouter()
	myRouter.AddCall("/ping", RequestHandler)
	myRouter.AddCall("/pong", RequestHandlerTwo)
	sw := myRouter.EmitOpenAPIDefinition()

	if len(sw.Paths.Paths) != 2 {
		t.Fatal("Wrong number of pathes")
	}

	if len(sw.Definitions) != 4 {
		t.Fatal("Wrong number of definitions")
	}

	_, innerPresent := sw.Definitions["InnerStruct"]
	if !innerPresent {
		t.Fatal("Nested structure is not present in definitions")
	}
}
