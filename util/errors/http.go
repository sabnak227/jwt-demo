package errors

import (
	"encoding/json"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/sabnak227/jwt-demo/util/constant"
	"net/http"
)

type ResponseError struct{
	Err error
	HttpCode int
	ErrorCode int
	Message string
	ValidationErrors map[string]int32
}

func NewResponseError(err error, msg string) *ResponseError {
	r := ResponseError{
		Err: err,
		ErrorCode: constant.FailCode,
		Message: msg,
		ValidationErrors: buildValidationError(err),
	}
	if err == nil {
		r.ErrorCode = constant.SuccessCode
	}
	if r.ValidationErrors != nil {
		r.HttpCode = http.StatusBadRequest
	} else {
		r.HttpCode = http.StatusInternalServerError
	}
	return &r
}

type errorResponse struct {
	Code     int    `json:"code"`
	Message  string `json:"message"`
	Error	 string  `json:"error"`
	ValidationErrors map[string]int32 `json:"validation_errors,omitempty"`
}

func (e *ResponseError) Error() string {
	return e.Err.Error()
}

func (e *ResponseError) SetStatusCode(c int) *ResponseError {
	e.HttpCode = c
	return e
}

func (e *ResponseError) SetErrorCode(c int) *ResponseError {
	e.ErrorCode = c
	return e
}

func (e *ResponseError) StatusCode() int {
	return e.HttpCode
}

func (e *ResponseError) MarshalJSON() ([]byte, error) {
	return json.Marshal(errorResponse{
		Code: e.ErrorCode,
		Message: e.Message,
		Error: e.Err.Error(),
		ValidationErrors: e.ValidationErrors,
	})
}



func buildValidationError(err error) map[string]int32 {
	e, ok := err.(validation.Errors)
	if !ok {
		return nil
	}

	errorRes := map[string]int32{}
	for index, element := range e {
		errorObj, ok := element.(validation.ErrorObject)
		if !ok {
			return nil
		}
		errorRes[index] = getCode(errorObj.Code())
	}
	return errorRes
}

func getCode(vcode string) int32 {
	switch vcode {
	case "validation_required":
		return constant.Required
	default:
		return constant.ValidationUndefined
	}
}