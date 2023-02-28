package xerr

var message map[uint32]string

func init() {
	message = make(map[uint32]string)
	// 全局错误码
	message[OK] = "SUCCESS"
	message[ServerCommonError] = "服务器开小差啦,稍后再来试一试"
	message[RequestParamError] = "参数错误"
	message[TokenExpireError] = "token失效，请重新登陆"
	message[TokenGenerateError] = "生成token失败"
	message[DbError] = "数据库繁忙,请稍后再试"
	message[DbUpdateAffectedZeroError] = "更新数据影响行数为0"
	message[ParamValidateError] = "参数校验失败"
	message[PermissionDenied] = "权限不足"
	message[JsonMarshalError] = "json marshal fail"
	message[AsynqEnqueueError] = "异步任务入队失败"
	message[ElasticSearchError] = "es操作失败"

	// 用户模块
	message[UserNotFound] = "用户不存在"
	message[UserIsBlack] = "用户已被拉黑"
	message[PhoneNotFound] = "手机号不存在"
	message[EmailNotFound] = "邮箱不存在"
	message[SendSmsError] = "发送短信失败"
	message[GraphCaptchaError] = "图形验证码错误"
	message[PhoneRegistered] = "手机号已注册"
	message[PhoneNotRegistered] = "手机号未注册"
	message[PasswordError] = "密码错误"
	message[ErrGenerateTokenError] = "生成token失败"
	message[RefreshTokenError] = "刷新token失败"
	message[VerifyCodeError] = "验证码错误"
}

func MapErrMsg(errcode uint32) string {
	if msg, ok := message[errcode]; ok {
		return msg
	} else {
		return "服务器开小差啦,稍后再来试一试1"
	}
}

func IsCodeErr(errcode uint32) bool {
	if _, ok := message[errcode]; ok {
		return true
	} else {
		return false
	}
}
