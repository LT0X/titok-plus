package apiErr

// Success 根据官方文档 0 代表成功
var Success = NewApiErr(0, "Success")

//错误状态码

//400系列错误码

// 400 坏请求
var (
	InvalidParameter = NewApiErr(40001, "参数错误")
	InvalidAccount   = NewApiErr(40002, "用户名或者密码错误")
	FileIsNotVideo   = NewApiErr(40003, "文件类型错误")
)

// 401 身份验证错误
var (
	NotLogin     = NewApiErr(401001, "没有登录")
	InvalidToken = NewApiErr(401002, "Token无效")
)

// 404 未找到数据
var (
	UserNotExit = NewApiErr(404001, "用户不存在")
)

// 405 冲突
var (
	UserNameConflict = NewApiErr(405001, "用户名重复")
)

// 500 服务端错误
var (
	ServerInternal    = NewApiErr(50001, "服务器内部错误")
	EncryptionFailed  = NewApiErr(50002, "密码加密失败")
	CreateTokenFailed = NewApiErr(50003, "颁发token失败")
	FileUploadFailed  = NewApiErr(50004, "上传视频失败")
)
