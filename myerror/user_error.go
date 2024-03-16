package myerror

const TYPE_USER_EXIST = "user_exist"
const TYPE_EMAIL_CODE_WRONG = "email_code_wrong"

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
	case TYPE_EMAIL_CODE_WRONG:
		return "Email verification code is wrong"
	default:
		return "Unknown error"
	}
}

func (ue *UserError) Unwrap() error {
	return &ue.AppError
}
