package myerror

const TYPE_HEADER_MISSING = "header_missing"
const TYPE_TOKEN_INVALID = "token_invalid"

type TokenError struct {
	AppError
}

func NewTokenError(errtype string) *TokenError {
	return &TokenError{
		AppError: AppError{
			Code:      401,
			ErrorCode: "token_error",
			Message:   getErrorMsg(errtype),
		},
	}
}

func (te *TokenError) Unwrap() error {
	return &te.AppError
}
