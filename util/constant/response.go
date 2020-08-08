package constant

const (
	SuccessCode = int32(1000)
	FailCode = int32(2000)
	WrongPasswordCode = int32(123123)
)

type Response struct {
	Code int32
	Message string
}

var Success Response = Response{
	Code:    SuccessCode,
	Message: "success",
}

var Fail Response = Response{
	Code:    FailCode,
	Message: "fail",
}

