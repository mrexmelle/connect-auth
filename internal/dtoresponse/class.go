package dtoresponse

import (
	"encoding/json"
	"net/http"

	"github.com/mrexmelle/connect-authx/internal/errmap"
)

type ClassWithData[T any] struct {
	Data         *T
	Error        error
	ErrorMap     *map[error]errmap.CodePair
	PreWriteHook func(*T)
}

func NewWithData[T any](data *T, err error) *ClassWithData[T] {
	return &ClassWithData[T]{
		Data:         data,
		Error:        err,
		ErrorMap:     nil,
		PreWriteHook: nil,
	}
}

func (c *ClassWithData[T]) WithErrorMap(em *map[error]errmap.CodePair) *ClassWithData[T] {
	c.ErrorMap = em
	return c
}

func (c *ClassWithData[T]) WithPrewriteHook(hook func(*T)) *ClassWithData[T] {
	c.PreWriteHook = hook
	return c
}

func (c *ClassWithData[T]) RenderTo(w http.ResponseWriter) {
	errorInfo := errmap.New(c.ErrorMap).Map(c.Error)

	responseBody, _ := json.Marshal(
		&HttpResponseWithData[T]{
			Data: c.Data,
			Error: ServiceError{
				Code:    errorInfo.ServiceErrorCode,
				Message: errorInfo.ServiceErrorMessage,
			},
		},
	)
	if errorInfo.HttpStatusCode == http.StatusOK {
		if c.PreWriteHook != nil {
			c.PreWriteHook(c.Data)
		}
		w.Write(responseBody)
	} else {
		http.Error(w, string(responseBody), errorInfo.HttpStatusCode)
	}
}

type ClassWithoutData struct {
	Error    error
	ErrorMap *map[error]errmap.CodePair
}

func NewWithoutData(err error) *ClassWithoutData {
	return &ClassWithoutData{
		Error:    err,
		ErrorMap: nil,
	}
}

func (c *ClassWithoutData) WithErrorMap(em *map[error]errmap.CodePair) *ClassWithoutData {
	c.ErrorMap = em
	return c
}

func (c *ClassWithoutData) RenderTo(w http.ResponseWriter) {
	errorInfo := errmap.New(c.ErrorMap).Map(c.Error)

	responseBody, _ := json.Marshal(
		&HttpResponseWithoutData{
			Error: ServiceError{
				Code:    errorInfo.ServiceErrorCode,
				Message: errorInfo.ServiceErrorMessage,
			},
		},
	)
	if errorInfo.HttpStatusCode == http.StatusOK {
		w.Write(responseBody)
	} else {
		http.Error(w, string(responseBody), errorInfo.HttpStatusCode)
	}
}
