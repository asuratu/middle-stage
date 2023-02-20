// Code generated by goctl. DO NOT EDIT.
// Source: usersvc.proto

package usersvc

import (
	"context"

	"middle/app/user/rpc/user"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	AddCategoryReq       = user.AddCategoryReq
	AddCategoryResp      = user.AddCategoryResp
	AddTopicReq          = user.AddTopicReq
	AddTopicResp         = user.AddTopicResp
	Category             = user.Category
	DelCategoryReq       = user.DelCategoryReq
	DelCategoryResp      = user.DelCategoryResp
	DelTopicReq          = user.DelTopicReq
	DelTopicResp         = user.DelTopicResp
	DelUserReq           = user.DelUserReq
	DelUserResp          = user.DelUserResp
	GenerateTokenReq     = user.GenerateTokenReq
	GetCategoryByIdReq   = user.GetCategoryByIdReq
	GetCategoryByIdResp  = user.GetCategoryByIdResp
	GetTopicByIdReq      = user.GetTopicByIdReq
	GetTopicByIdResp     = user.GetTopicByIdResp
	GetUserByIdReq       = user.GetUserByIdReq
	GetUserByMobileReq   = user.GetUserByMobileReq
	GetUserListReq       = user.GetUserListReq
	GetUserListResp      = user.GetUserListResp
	GetUserResp          = user.GetUserResp
	LoginByPhoneReq      = user.LoginByPhoneReq
	RefreshTokenReq      = user.RefreshTokenReq
	RegisterReq          = user.RegisterReq
	SearchCategoryReq    = user.SearchCategoryReq
	SearchCategoryResp   = user.SearchCategoryResp
	SearchTopicReq       = user.SearchTopicReq
	SearchTopicResp      = user.SearchTopicResp
	SearchUserReq        = user.SearchUserReq
	SearchUserResp       = user.SearchUserResp
	SendImageCaptchaReq  = user.SendImageCaptchaReq
	SendImageCaptchaResp = user.SendImageCaptchaResp
	SendSmsCodeReq       = user.SendSmsCodeReq
	SendSmsCodeResp      = user.SendSmsCodeResp
	SimpleUser           = user.SimpleUser
	TokenResp            = user.TokenResp
	Topic                = user.Topic
	UpdateCategoryReq    = user.UpdateCategoryReq
	UpdateCategoryResp   = user.UpdateCategoryResp
	UpdateTopicReq       = user.UpdateTopicReq
	UpdateTopicResp      = user.UpdateTopicResp
	UpdateUserReq        = user.UpdateUserReq
	UpdateUserResp       = user.UpdateUserResp
	User                 = user.User
	VerifyCaptchaReq     = user.VerifyCaptchaReq
	VerifyCaptchaResp    = user.VerifyCaptchaResp
	VerifySmsCodeReq     = user.VerifySmsCodeReq
	VerifySmsCodeResp    = user.VerifySmsCodeResp

	Usersvc interface {
		// -----------------------category-----------------------
		AddCategory(ctx context.Context, in *AddCategoryReq, opts ...grpc.CallOption) (*AddCategoryResp, error)
		UpdateCategory(ctx context.Context, in *UpdateCategoryReq, opts ...grpc.CallOption) (*UpdateCategoryResp, error)
		DelCategory(ctx context.Context, in *DelCategoryReq, opts ...grpc.CallOption) (*DelCategoryResp, error)
		GetCategoryById(ctx context.Context, in *GetCategoryByIdReq, opts ...grpc.CallOption) (*GetCategoryByIdResp, error)
		SearchCategory(ctx context.Context, in *SearchCategoryReq, opts ...grpc.CallOption) (*SearchCategoryResp, error)
		// -----------------------topic-----------------------
		AddTopic(ctx context.Context, in *AddTopicReq, opts ...grpc.CallOption) (*AddTopicResp, error)
		UpdateTopic(ctx context.Context, in *UpdateTopicReq, opts ...grpc.CallOption) (*UpdateTopicResp, error)
		DelTopic(ctx context.Context, in *DelTopicReq, opts ...grpc.CallOption) (*DelTopicResp, error)
		GetTopicById(ctx context.Context, in *GetTopicByIdReq, opts ...grpc.CallOption) (*GetTopicByIdResp, error)
		SearchTopic(ctx context.Context, in *SearchTopicReq, opts ...grpc.CallOption) (*SearchTopicResp, error)
		// -----------------------user-----------------------
		LoginByPhone(ctx context.Context, in *LoginByPhoneReq, opts ...grpc.CallOption) (*TokenResp, error)
		Register(ctx context.Context, in *RegisterReq, opts ...grpc.CallOption) (*TokenResp, error)
		GenerateToken(ctx context.Context, in *GenerateTokenReq, opts ...grpc.CallOption) (*TokenResp, error)
		RefreshToken(ctx context.Context, in *RefreshTokenReq, opts ...grpc.CallOption) (*TokenResp, error)
		UpdateUser(ctx context.Context, in *UpdateUserReq, opts ...grpc.CallOption) (*UpdateUserResp, error)
		DelUser(ctx context.Context, in *DelUserReq, opts ...grpc.CallOption) (*DelUserResp, error)
		GetUserById(ctx context.Context, in *GetUserByIdReq, opts ...grpc.CallOption) (*GetUserResp, error)
		GetUserByMobile(ctx context.Context, in *GetUserByMobileReq, opts ...grpc.CallOption) (*GetUserResp, error)
		SearchUser(ctx context.Context, in *SearchUserReq, opts ...grpc.CallOption) (*SearchUserResp, error)
		GetUserList(ctx context.Context, in *GetUserListReq, opts ...grpc.CallOption) (*GetUserListResp, error)
		// -----------------------auth-----------------------
		SendImageCaptcha(ctx context.Context, in *SendImageCaptchaReq, opts ...grpc.CallOption) (*SendImageCaptchaResp, error)
		VerifyCaptcha(ctx context.Context, in *VerifyCaptchaReq, opts ...grpc.CallOption) (*VerifyCaptchaResp, error)
		SendSmsCode(ctx context.Context, in *SendSmsCodeReq, opts ...grpc.CallOption) (*SendSmsCodeResp, error)
		VerifySmsCode(ctx context.Context, in *VerifySmsCodeReq, opts ...grpc.CallOption) (*VerifySmsCodeResp, error)
	}

	defaultUsersvc struct {
		cli zrpc.Client
	}
)

func NewUsersvc(cli zrpc.Client) Usersvc {
	return &defaultUsersvc{
		cli: cli,
	}
}

// -----------------------category-----------------------
func (m *defaultUsersvc) AddCategory(ctx context.Context, in *AddCategoryReq, opts ...grpc.CallOption) (*AddCategoryResp, error) {
	client := user.NewUsersvcClient(m.cli.Conn())
	return client.AddCategory(ctx, in, opts...)
}

func (m *defaultUsersvc) UpdateCategory(ctx context.Context, in *UpdateCategoryReq, opts ...grpc.CallOption) (*UpdateCategoryResp, error) {
	client := user.NewUsersvcClient(m.cli.Conn())
	return client.UpdateCategory(ctx, in, opts...)
}

func (m *defaultUsersvc) DelCategory(ctx context.Context, in *DelCategoryReq, opts ...grpc.CallOption) (*DelCategoryResp, error) {
	client := user.NewUsersvcClient(m.cli.Conn())
	return client.DelCategory(ctx, in, opts...)
}

func (m *defaultUsersvc) GetCategoryById(ctx context.Context, in *GetCategoryByIdReq, opts ...grpc.CallOption) (*GetCategoryByIdResp, error) {
	client := user.NewUsersvcClient(m.cli.Conn())
	return client.GetCategoryById(ctx, in, opts...)
}

func (m *defaultUsersvc) SearchCategory(ctx context.Context, in *SearchCategoryReq, opts ...grpc.CallOption) (*SearchCategoryResp, error) {
	client := user.NewUsersvcClient(m.cli.Conn())
	return client.SearchCategory(ctx, in, opts...)
}

// -----------------------topic-----------------------
func (m *defaultUsersvc) AddTopic(ctx context.Context, in *AddTopicReq, opts ...grpc.CallOption) (*AddTopicResp, error) {
	client := user.NewUsersvcClient(m.cli.Conn())
	return client.AddTopic(ctx, in, opts...)
}

func (m *defaultUsersvc) UpdateTopic(ctx context.Context, in *UpdateTopicReq, opts ...grpc.CallOption) (*UpdateTopicResp, error) {
	client := user.NewUsersvcClient(m.cli.Conn())
	return client.UpdateTopic(ctx, in, opts...)
}

func (m *defaultUsersvc) DelTopic(ctx context.Context, in *DelTopicReq, opts ...grpc.CallOption) (*DelTopicResp, error) {
	client := user.NewUsersvcClient(m.cli.Conn())
	return client.DelTopic(ctx, in, opts...)
}

func (m *defaultUsersvc) GetTopicById(ctx context.Context, in *GetTopicByIdReq, opts ...grpc.CallOption) (*GetTopicByIdResp, error) {
	client := user.NewUsersvcClient(m.cli.Conn())
	return client.GetTopicById(ctx, in, opts...)
}

func (m *defaultUsersvc) SearchTopic(ctx context.Context, in *SearchTopicReq, opts ...grpc.CallOption) (*SearchTopicResp, error) {
	client := user.NewUsersvcClient(m.cli.Conn())
	return client.SearchTopic(ctx, in, opts...)
}

// -----------------------user-----------------------
func (m *defaultUsersvc) LoginByPhone(ctx context.Context, in *LoginByPhoneReq, opts ...grpc.CallOption) (*TokenResp, error) {
	client := user.NewUsersvcClient(m.cli.Conn())
	return client.LoginByPhone(ctx, in, opts...)
}

func (m *defaultUsersvc) Register(ctx context.Context, in *RegisterReq, opts ...grpc.CallOption) (*TokenResp, error) {
	client := user.NewUsersvcClient(m.cli.Conn())
	return client.Register(ctx, in, opts...)
}

func (m *defaultUsersvc) GenerateToken(ctx context.Context, in *GenerateTokenReq, opts ...grpc.CallOption) (*TokenResp, error) {
	client := user.NewUsersvcClient(m.cli.Conn())
	return client.GenerateToken(ctx, in, opts...)
}

func (m *defaultUsersvc) RefreshToken(ctx context.Context, in *RefreshTokenReq, opts ...grpc.CallOption) (*TokenResp, error) {
	client := user.NewUsersvcClient(m.cli.Conn())
	return client.RefreshToken(ctx, in, opts...)
}

func (m *defaultUsersvc) UpdateUser(ctx context.Context, in *UpdateUserReq, opts ...grpc.CallOption) (*UpdateUserResp, error) {
	client := user.NewUsersvcClient(m.cli.Conn())
	return client.UpdateUser(ctx, in, opts...)
}

func (m *defaultUsersvc) DelUser(ctx context.Context, in *DelUserReq, opts ...grpc.CallOption) (*DelUserResp, error) {
	client := user.NewUsersvcClient(m.cli.Conn())
	return client.DelUser(ctx, in, opts...)
}

func (m *defaultUsersvc) GetUserById(ctx context.Context, in *GetUserByIdReq, opts ...grpc.CallOption) (*GetUserResp, error) {
	client := user.NewUsersvcClient(m.cli.Conn())
	return client.GetUserById(ctx, in, opts...)
}

func (m *defaultUsersvc) GetUserByMobile(ctx context.Context, in *GetUserByMobileReq, opts ...grpc.CallOption) (*GetUserResp, error) {
	client := user.NewUsersvcClient(m.cli.Conn())
	return client.GetUserByMobile(ctx, in, opts...)
}

func (m *defaultUsersvc) SearchUser(ctx context.Context, in *SearchUserReq, opts ...grpc.CallOption) (*SearchUserResp, error) {
	client := user.NewUsersvcClient(m.cli.Conn())
	return client.SearchUser(ctx, in, opts...)
}

func (m *defaultUsersvc) GetUserList(ctx context.Context, in *GetUserListReq, opts ...grpc.CallOption) (*GetUserListResp, error) {
	client := user.NewUsersvcClient(m.cli.Conn())
	return client.GetUserList(ctx, in, opts...)
}

// -----------------------auth-----------------------
func (m *defaultUsersvc) SendImageCaptcha(ctx context.Context, in *SendImageCaptchaReq, opts ...grpc.CallOption) (*SendImageCaptchaResp, error) {
	client := user.NewUsersvcClient(m.cli.Conn())
	return client.SendImageCaptcha(ctx, in, opts...)
}

func (m *defaultUsersvc) VerifyCaptcha(ctx context.Context, in *VerifyCaptchaReq, opts ...grpc.CallOption) (*VerifyCaptchaResp, error) {
	client := user.NewUsersvcClient(m.cli.Conn())
	return client.VerifyCaptcha(ctx, in, opts...)
}

func (m *defaultUsersvc) SendSmsCode(ctx context.Context, in *SendSmsCodeReq, opts ...grpc.CallOption) (*SendSmsCodeResp, error) {
	client := user.NewUsersvcClient(m.cli.Conn())
	return client.SendSmsCode(ctx, in, opts...)
}

func (m *defaultUsersvc) VerifySmsCode(ctx context.Context, in *VerifySmsCodeReq, opts ...grpc.CallOption) (*VerifySmsCodeResp, error) {
	client := user.NewUsersvcClient(m.cli.Conn())
	return client.VerifySmsCode(ctx, in, opts...)
}
