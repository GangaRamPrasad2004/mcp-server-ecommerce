package jsonRPC

import (
	"encoding/json"
	"fmt"
)

type Request struct {
	JSONRPC string          `json:"jsonrpc"`
	Method  string          `json:"method"`
	Params  json.RawMessage `json:"params"`
	ID      any             `json:"id"`
}

type Response struct {
	JSONRPC string `json:"jsonrpc"`
	Result  any    `json:"result"`
	Error   *Error `json:"error"`
	ID      any    `json:"id"`
}
type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

// implemented the interface  for Error
func (e *Error) Error() string {
	return fmt.Sprintf(" code=%d, message=%s,Data:%v", e.Code, e.Message, e.Data)
}

func (r *Request) IsNotification() bool {
	return r.ID == nil
}

func (r *Request) Validate() error {
	if r.JSONRPC != "2.0" {
		return &Error{
			Code:    ErrorInvalidRequest,
			Message: "invalid JSON RPC version ,must be 2.0",
		}
	}
	if r.Method == "" {
		return &Error{
			Code:    ErrorInvalidRequest,
			Message: "missing method",
		}
	}
	return nil
}

func NewResponse(id, result any, err *Error) *Response {
	return &Response{
		JSONRPC: "2.0",
		Result:  result,
		Error:   err,
		ID:      id,
	}
}

func NewErrorResponse(id any, err *error) *Response {
	return NewResponse(id, err, nil)

}

func NewSuccessResponse(result, id any) *Response {
	return NewResponse(result, id, nil)
}
