package request

import (
	"errors"
	"fmt"
	"io"
	"net/http"
)

type Request struct {
	Method string
	Body   io.ReadCloser
}

func Init(r *http.Request) *Request {
	method := r.Method
	body := r.Body
	return &Request{Method: method, Body: body}
}

func (req *Request) CheckMethod(method string) error {
	errorMessage := fmt.Sprintf("Method %s Not Allowed", req.Method)
	if req.Method != method {
		return errors.New(errorMessage)
	}
	return nil
}
