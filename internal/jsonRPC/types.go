package jsonRPC

import (
	"encoding/json"
	"fmt"
)

type Request struct {
	JSONRPC string          `json:"jsonrpc"`
	Method  string          `json:"method"`
	Params  json.RawMessage `json:"params,omitempty"`
	ID      any             `json:"id,omitempty"` // can be string, number, or null. No fractional numbers allowed. e.g 1.2 is not allowed
}

type Response struct {
	JSONRPC string `json:"jsonrpc"`
	Result  any    `json:"result,omitempty"`
	Error   *Error `json:"error,omitempty"`
	ID      any    `json:"id,omitempty"`
}

type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}

func (e *Error) Error() string {
	return fmt.Sprintf("Code: %d, Message: %s, Data: %v", e.Code, e.Message, e.Data)
}

func (r *Request) IsNotification() bool {
	return r.ID == nil
}

func (r *Request) Validate() error {
	if r.JSONRPC != "2.0" {
		return &Error{
			Code:    ErrorInvalidRequest,
			Message: "Invalid JSON-RPC version, must be 2.0",
		}
	}

	if r.Method == "" {
		return &Error{
			Code:    ErrorInvalidRequest,
			Message: "Method is required",
		}
	}

	return nil
}

// NewResponse creates a new JSON-RPC response
func NewResponse(result any, id any, err *Error) *Response {
	return &Response{
		JSONRPC: "2.0",
		Result:  result,
		Error:   err,
		ID:      id,
	}
}

// NewErrorResponse creates a new JSON-RPC error response
func NewErrorResponse(id any, err *Error) *Response {
	return NewResponse(nil, id, err)
}

// NewSuccessResponse creates a new JSON-RPC success response
func NewSuccessResponse(result, id any) *Response {
	return NewResponse(result, id, nil)
}
