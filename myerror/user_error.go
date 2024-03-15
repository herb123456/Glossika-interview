package myerror

const TYPE_USER_EXIST = "user_exist"

type UserError struct {
	AppError
}

func NewUserError(errtype string) *UserError {

	return &UserError{
		AppError: AppError{
			Code:      400,
			ErrorCode: "user_error",
			Message:   getErrorMsg(errtype),
		},
	}
}

func getErrorMsg(errtype string) string {
	switch errtype {
	case TYPE_USER_EXIST:
		return "User already exists"
	default:
		return "Unknown error"
	}
}

func (ue *UserError) Unwrap() error {
	return &ue.AppError
}
