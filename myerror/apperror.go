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
