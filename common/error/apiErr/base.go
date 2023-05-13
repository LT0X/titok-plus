package apiErr

type ApiErr struct {
	StatusCode int64  `json:"status_code"`
	StatusMsg  string `json:"status_msg"`
}

func NewApiErr(Code int64, msg string) ApiErr {
	return ApiErr{
		StatusCode: Code,
		StatusMsg:  msg,
	}
}

func (e ApiErr) Error() string {
	return e.StatusMsg
}

// WithDetails 在基础错误上追加详细信息，例如：密码错误，密码长度不足6位
func (e ApiErr) WithDetails(detail string) ApiErr {
	return ApiErr{
		StatusCode: e.StatusCode,
		StatusMsg:  e.StatusMsg + ": " + detail,
	}
}
