package myerror

import (
	"github.com/go-playground/validator/v10"
	"strings"
)

type ValidationError struct {
	AppError
}

func NewValidationError(e error) *ValidationError {
	errs, ok := e.(validator.ValidationErrors)
	if !ok {
		return &ValidationError{
			AppError: AppError{
				Code:         400,
				ErrorCode:    "validation_error",
				Message:      "Validation error",
				DebugMessage: e.Error(),
			},
		}
	}

	out := make([]string, len(errs))
	for _, e := range errs {
		out = append(out, e.Field()+":"+customValidatorErrorMsg(e))
	}

	return &ValidationError{
		AppError: AppError{
			Code:      400,
			ErrorCode: "validation_error",
			Message:   strings.Join(out, ","),
		},
	}
}

func customValidatorErrorMsg(err validator.FieldError) string {
	switch err.Tag() {
	case "passwd":
		return "Password must be at least 6 characters long, contain at least one uppercase letter, one lowercase letter, and one special character"
	default:
		return err.Error()
	}
}

func (ve *ValidationError) Unwrap() error {
	return &ve.AppError
}
