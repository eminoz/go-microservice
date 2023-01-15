package utilities

type Result struct {
	Success bool
	Message string
}

type DataResult struct {
	Result
	Data    interface{}
	Success bool
}

func SuccessDataResult(msg string, d interface{}) *DataResult {
	return &DataResult{
		Result: Result{
			Message: msg,
		},
		Data:    d,
		Success: true,
	}
}

func ErrorDataResult(msg string, d interface{}) *DataResult {
	return &DataResult{

		Result: Result{
			Message: msg,
		},
		Data: d,

		Success: false,
	}
}

type ResultSuccess struct {
	Success bool
	Message string
}

func SuccessResult(msg string) *ResultSuccess {
	return &ResultSuccess{
		Success: true,
		Message: msg,
	}
}

type ResultError struct {
	Success bool
	Message string
}

func ErrorResult(msg string) *ResultError {
	return &ResultError{
		Success: false,
		Message: msg,
	}
}
