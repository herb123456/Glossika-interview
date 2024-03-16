package myerror

type AppError struct {
	Code         int    // HTTP狀態碼
	ErrorCode    string // 自定義的錯誤代碼
	Message      string // 錯誤訊息給用戶看
	DebugMessage string // 錯誤訊息給開發者看
}

func (e *AppError) Error() string {
	return e.Message
}

func getErrorMsg(errtype string) string {
	switch errtype {
	case TYPE_USER_EXIST:
		return "User already exists"
	case TYPE_EMAIL_CODE_WRONG:
		return "Email verification code is wrong"
	case TYPE_EMAIL_PASSWORD_WRONG:
		return "Email or password is wrong"
	case TYPE_EMAIL_NOT_VERIFIED:
		return "Email is not verified"
	case TYPE_HEADER_MISSING:
		return "Authorization header is missing"
	case TYPE_TOKEN_INVALID:
		return "Token is invalid"
	default:
		return "Unknown error"
	}
}
