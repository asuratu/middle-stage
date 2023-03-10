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
	ParamValidateError        uint32 = 100007
	PermissionDenied          uint32 = 100008
	JsonMarshalError          uint32 = 100009
	AsynqEnqueueError         uint32 = 100010
	ElasticSearchError        uint32 = 100011
	RequestLimitError         uint32 = 100012
)

// 用户模块
const (
	UserNotFound          uint32 = 200001
	UserIsBlack           uint32 = 200002
	PhoneNotFound         uint32 = 200003
	EmailNotFound         uint32 = 200004
	SendSmsError          uint32 = 200005
	GraphCaptchaError     uint32 = 200006
	PhoneRegistered       uint32 = 200007
	PhoneNotRegistered    uint32 = 200008
	PasswordError         uint32 = 200009
	ErrGenerateTokenError uint32 = 200010
	RefreshTokenError     uint32 = 200011
	VerifyCodeError       uint32 = 200012
)
