package rpcErr

var (
	DataBaseError     = NewRPCErr(10001, "数据库错误")
	CacheBaseError    = NewRPCErr(10002, "缓存错误")
	MessageQueueError = NewRPCErr(10003, "消息队列错误")
)

var (
	UserNotExit = NewRPCErr(30001, "用户不存在")
)
