package myerror

type GeneralError struct {
	AppError
}

func NewGeneralError(message string) *GeneralError {
	return &GeneralError{
		AppError: AppError{
			Code:         500,
			ErrorCode:    "general_error",
			Message:      "Unknown error",
			DebugMessage: message,
		},
	}
}
