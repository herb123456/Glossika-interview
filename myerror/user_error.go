package myerror

const TYPE_USER_EXIST = "user_exist"
const TYPE_EMAIL_CODE_WRONG = "email_code_wrong"
const TYPE_EMAIL_PASSWORD_WRONG = "email_password_wrong"
const TYPE_EMAIL_NOT_VERIFIED = "email_not_verified"

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

func (ue *UserError) Unwrap() error {
	return &ue.AppError
}
