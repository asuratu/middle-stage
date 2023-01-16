package xerr

// OK 成功返回
const OK uint32 = 200

// 全局错误码 (前3位代表业务, 后三位代表具体功能)
const (
	ServerCommonError         uint32 = 100001
	RequestParamError         uint32 = 100002
	TokenExpireError          uint32 = 100003
	TokenGenerateError        uint32 = 100004
	DbError                   uint32 = 100005
	DbUpdateAffectedZeroError uint32 = 100006
)

//用户模块
