package defs

type Err struct {
	Error   string `json:"erroe"`
	ErrCode string `json:"error_code"`
}

type ErrorResponse struct {
	HttpSc int
	Error  Err
}

var (
	ErrorRequestBodyParseFailed = ErrorResponse{HttpSc: 400, Err: {Error: "Request bode is not correct", ErrCode: ""}}
	ErrorNotAuthUser            = ErrorResponse{HttpSc: 401, Err: {Error: "User auth failed", ErrCode: "002"}}
)
