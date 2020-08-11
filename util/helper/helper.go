package helper

import (
	"fmt"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/sabnak227/jwt-demo/util/constant"
)


func BuildErrorResponse(err error) (map[string]int32, error){
	e, ok := err.(validation.Errors)
	if !ok {
		return nil, fmt.Errorf("cannot cast to validation errors")
	}

	errorRes := map[string]int32{}
	for index, element := range e {
		errorObj, ok := element.(validation.ErrorObject)
		if !ok {
			return nil, fmt.Errorf("cannot cast to validation error objects")
		}
		errorRes[index] = getCode(errorObj.Code())
	}
	return errorRes, nil
}

func getCode(vcode string) int32 {
	switch vcode {
	case "validation_required":
		return constant.Required
	default:
		return constant.ValidationUndefined
	}
}