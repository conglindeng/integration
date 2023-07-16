package core

type Result struct {
	Sucess bool
	Data
	Error error
	Code  int
}

func Sucess_Result(data Data) Result {
	return Result{Sucess: true, Data: data}
}

func Fail_Result(e error, code int) Result {
	return Result{Sucess: false, Code: code, Error: e}
}
