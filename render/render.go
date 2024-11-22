package render

import "go-wallet/api/models"

type Err struct {
	Code int
	Msg  string
}

var (
	InvalidParameter = newErr(3001, "Invalid parameter")
)

func newErr(code int, msg string) Err {
	return Err{
		Code: code,
		Msg:  msg,
	}
}

func Error(err error) models.Response {
	return models.Response{
		Status:    "error",
		ErrorCode: 1,
		ErrorMsg:  err.Error(),
	}
}

func BindError(err error) models.Response {
	resp := errResponse(InvalidParameter)
	resp.ErrorMsg = err.Error()

	return resp
}

func errResponse(err Err) models.Response {
	return models.Response{
		Status:    "error",
		ErrorCode: err.Code,
		ErrorMsg:  err.Msg,
	}
}

func Success(data interface{}) models.Response {
	return models.Response{
		Status: "success",
		Data:   data,
	}
}
