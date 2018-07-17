package defs

type Err struct {
	Error string `json:"error"`
	ErrorCode string `json:"error_code"`
}

type ErrResponse struct {
	HttpSC int
	Error Err
}

var (
	ErrorREquestBodyParseFailed = ErrResponse{HttpSC: 400, Error: Err{Error: "Request Body is not correct.", ErrorCode: "001"}}
	ErrorNotAuthUser = ErrResponse{HttpSC: 401, Error: Err{Error: "User authentification failed.", ErrorCode: "002"}}

	)