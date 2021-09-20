package exceptions

import "net/http"

type BaseException struct {
	Error      string `json:"error"`
	Message    string `json:"message"`
	StatusCode uint16 `json:"statusCode"`
}

func NewBaseException(status int, msg string) *BaseException {
	b := new(BaseException)
	b.StatusCode = uint16(status)
	b.Error = http.StatusText(status)
	b.Message = msg
	return b
}
