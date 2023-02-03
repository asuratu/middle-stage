// Code generated by goctl. DO NOT EDIT.
package types

type PhoneExistReq struct {
	Phone string `json:"phone" validate:"required,startswith=1,len=11"`
}

type PhoneExistReply struct {
	Exist bool `json:"exist"`
}

type EmailExistReq struct {
	Email string `json:"email"`
}

type EmailExistReply struct {
	Exist bool `json:"exist"`
}

type RegisterReq struct {
	Username string `json:"username"`
	Mobile   string `json:"mobile"`
	Password string `json:"password"`
	Smscode  string `json:"smscode"`
}

type LoginReq struct {
	Username string `json:"username"`
	Mobile   string `json:"mobile"`
	Password string `json:"password"`
	Smscode  string `json:"smscode"`
}

type UserReply struct {
	Auid     int64  `json:"auid"`
	Uid      int64  `json:"uid"`
	Beid     int64  `json:"beid"`  //应用id
	Ptyid    int64  `json:"ptyid"` //对应平台
	Username string `json:"username"`
	Mobile   string `json:"mobile"`
	Nickname string `json:"nickname"`
	Openid   string `json:"openid"`
	Avator   string `json:"avator"`
	Gender   string `json:"gender"`
	JwtToken
}

type JwtToken struct {
	AccessToken  string `json:"accessToken,omitempty"`
	AccessExpire int64  `json:"accessExpire,omitempty"`
	RefreshAfter int64  `json:"refreshAfter,omitempty"`
}
