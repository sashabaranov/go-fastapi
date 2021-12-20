package fastapi

import (
	"log"
	"net/http"
	"reflect"

	"github.com/gin-gonic/gin"
)

type Router struct {
	routesMap map[string]interface{}
}

func NewRouter() *Router {
	return &Router{
		routesMap: make(map[string]interface{}),
	}
}

func (r *Router) AddCall(path string, handler interface{}) {
	handlerType := reflect.TypeOf(handler)

	if handlerType.NumIn() != 2 {
		panic("Wrong number of arguments")
	}
	if handlerType.NumOut() != 2 {
		panic("Wrong number of return values")
	}

	ginCtxType := reflect.TypeOf(&gin.Context{})
	if !handlerType.In(0).ConvertibleTo(ginCtxType) {
		panic("First argument should be *gin.Context!")
	}
	// fmt.Println(handlerType.In(1).Kind() == reflect.Struct)
	if handlerType.In(1).Kind() != reflect.Struct {
		panic("Second argument must be a struct")
	}

	errorInterface := reflect.TypeOf((*error)(nil)).Elem()
	if !handlerType.Out(1).Implements(errorInterface) {
		panic("Second return value should be an error")
	}
	if handlerType.Out(0).Kind() != reflect.Struct {
		panic("First return value be a struct")
	}

	r.routesMap[path] = handler
}

func (r *Router) GinHandler(c *gin.Context) {
	path := c.Param("path")
	log.Print(path)
	handlerFuncPtr, present := r.routesMap[path]
	if !present {
		c.JSON(http.StatusNotFound, gin.H{"error": "handler not found"})
		return
	}

	handlerType := reflect.TypeOf(handlerFuncPtr)
	inputType := handlerType.In(1)
	inputVal := reflect.New(inputType).Interface()
	err := c.BindJSON(inputVal)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	toCall := reflect.ValueOf(handlerFuncPtr)
	outputVal := toCall.Call(
		[]reflect.Value{
			reflect.ValueOf(c),
			reflect.ValueOf(inputVal).Elem(),
		},
	)

	returnedErr := outputVal[1].Interface()
	if returnedErr != nil || !outputVal[1].IsNil() {
		c.JSON(http.StatusInternalServerError, gin.H{"error": returnedErr})
		return
	}

	c.JSON(http.StatusOK, gin.H{"response": outputVal[0].Interface()})
}
