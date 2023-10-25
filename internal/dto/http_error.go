package dto

type HTTPError struct {
	ErrCode    int
	ErrMessage string
}

func NewHTTPError(errCode int, errMessage string) HTTPError {
	var httpError HTTPError

	httpError.ErrMessage = errMessage
	httpError.ErrCode = errCode

	return httpError
}

func (e HTTPError) Error() string {
	return e.ErrMessage
}
