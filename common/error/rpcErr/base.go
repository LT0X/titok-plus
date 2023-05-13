package rpcErr

type RPCErr struct {
	StatusCode int64  `json:"status_code"`
	StatusMsg  string `json:"status_msg"`
}

func NewRPCErr(code int64, msg string) RPCErr {
	return RPCErr{
		StatusCode: code,
		StatusMsg:  msg,
	}
}

func (e RPCErr) Error() string {
	return e.StatusMsg
}

// WithDetails 在基础错误上追加详细信息，例如：密码错误，密码长度不足6位
func (e RPCErr) WithDetails(detail string) RPCErr {
	return RPCErr{
		StatusCode: e.StatusCode,
		StatusMsg:  e.StatusMsg + ": " + detail,
	}
}
