package fastapi

import (
	"github.com/gin-gonic/gin"
	"testing"
)

func TestInvalidHandler1(t *testing.T) {
	defer func() { recover() }()
	NewRouter().AddCall("x", func(in struct{}) (out struct{}) { return })
	t.Errorf("Did not panic")
}

func TestInvalidHandler2(t *testing.T) {
	defer func() { recover() }()
	NewRouter().AddCall("x", func(ctx *gin.Context, in struct{}) (out struct{}) { return })
	t.Errorf("Did not panic")
}

func TestInvalidHandler3(t *testing.T) {
	defer func() { recover() }()
	NewRouter().AddCall("x", func(in struct{}) (out struct{}, err error) { return })
	t.Errorf("Did not panic")
}

func TestInvalidHandler4(t *testing.T) {
	defer func() { recover() }()
	NewRouter().AddCall("x", func(ctx struct{}, in struct{}) (out struct{}, err error) { return })
	t.Errorf("Did not panic")
}

func TestInvalidHandler5(t *testing.T) {
	defer func() { recover() }()
	NewRouter().AddCall("x", func(ctx *gin.Context, in string) (out struct{}, err error) { return })
	t.Errorf("Did not panic")
}

func TestInvalidHandler6(t *testing.T) {
	defer func() { recover() }()
	NewRouter().AddCall("x", func(ctx *gin.Context, in struct{}) (out string, err error) { return })
	t.Errorf("Did not panic")
}

func TestCorrectHandler(t *testing.T) {
	NewRouter().AddCall("x", func(ctx *gin.Context, in struct{}) (out struct{}, err error) { return })
}
