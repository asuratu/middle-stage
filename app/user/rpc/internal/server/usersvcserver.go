// Code generated by goctl. DO NOT EDIT.
// Source: usersvc.proto

package server

import (
	"context"

	"middle/app/user/rpc/internal/logic"
	"middle/app/user/rpc/internal/svc"
	"middle/app/user/rpc/user"
)

type UsersvcServer struct {
	svcCtx *svc.ServiceContext
	user.UnimplementedUsersvcServer
}

func NewUsersvcServer(svcCtx *svc.ServiceContext) *UsersvcServer {
	return &UsersvcServer{
		svcCtx: svcCtx,
	}
}

// -----------------------category-----------------------
func (s *UsersvcServer) AddCategory(ctx context.Context, in *user.AddCategoryReq) (*user.AddCategoryResp, error) {
	l := logic.NewAddCategoryLogic(ctx, s.svcCtx)
	return l.AddCategory(in)
}

func (s *UsersvcServer) UpdateCategory(ctx context.Context, in *user.UpdateCategoryReq) (*user.UpdateCategoryResp, error) {
	l := logic.NewUpdateCategoryLogic(ctx, s.svcCtx)
	return l.UpdateCategory(in)
}

func (s *UsersvcServer) DelCategory(ctx context.Context, in *user.DelCategoryReq) (*user.DelCategoryResp, error) {
	l := logic.NewDelCategoryLogic(ctx, s.svcCtx)
	return l.DelCategory(in)
}

func (s *UsersvcServer) GetCategoryById(ctx context.Context, in *user.GetCategoryByIdReq) (*user.GetCategoryByIdResp, error) {
	l := logic.NewGetCategoryByIdLogic(ctx, s.svcCtx)
	return l.GetCategoryById(in)
}

func (s *UsersvcServer) SearchCategory(ctx context.Context, in *user.SearchCategoryReq) (*user.SearchCategoryResp, error) {
	l := logic.NewSearchCategoryLogic(ctx, s.svcCtx)
	return l.SearchCategory(in)
}

// -----------------------topic-----------------------
func (s *UsersvcServer) AddTopic(ctx context.Context, in *user.AddTopicReq) (*user.AddTopicResp, error) {
	l := logic.NewAddTopicLogic(ctx, s.svcCtx)
	return l.AddTopic(in)
}

func (s *UsersvcServer) UpdateTopic(ctx context.Context, in *user.UpdateTopicReq) (*user.UpdateTopicResp, error) {
	l := logic.NewUpdateTopicLogic(ctx, s.svcCtx)
	return l.UpdateTopic(in)
}

func (s *UsersvcServer) DelTopic(ctx context.Context, in *user.DelTopicReq) (*user.DelTopicResp, error) {
	l := logic.NewDelTopicLogic(ctx, s.svcCtx)
	return l.DelTopic(in)
}

func (s *UsersvcServer) GetTopicById(ctx context.Context, in *user.GetTopicByIdReq) (*user.GetTopicByIdResp, error) {
	l := logic.NewGetTopicByIdLogic(ctx, s.svcCtx)
	return l.GetTopicById(in)
}

func (s *UsersvcServer) SearchTopic(ctx context.Context, in *user.SearchTopicReq) (*user.SearchTopicResp, error) {
	l := logic.NewSearchTopicLogic(ctx, s.svcCtx)
	return l.SearchTopic(in)
}

// -----------------------user-----------------------
func (s *UsersvcServer) AddUser(ctx context.Context, in *user.AddUserReq) (*user.AddUserResp, error) {
	l := logic.NewAddUserLogic(ctx, s.svcCtx)
	return l.AddUser(in)
}

func (s *UsersvcServer) UpdateUser(ctx context.Context, in *user.UpdateUserReq) (*user.UpdateUserResp, error) {
	l := logic.NewUpdateUserLogic(ctx, s.svcCtx)
	return l.UpdateUser(in)
}

func (s *UsersvcServer) DelUser(ctx context.Context, in *user.DelUserReq) (*user.DelUserResp, error) {
	l := logic.NewDelUserLogic(ctx, s.svcCtx)
	return l.DelUser(in)
}

func (s *UsersvcServer) GetUserById(ctx context.Context, in *user.GetUserByIdReq) (*user.GetUserResp, error) {
	l := logic.NewGetUserByIdLogic(ctx, s.svcCtx)
	return l.GetUserById(in)
}

func (s *UsersvcServer) GetUserByMobile(ctx context.Context, in *user.GetUserByMobileReq) (*user.GetUserResp, error) {
	l := logic.NewGetUserByMobileLogic(ctx, s.svcCtx)
	return l.GetUserByMobile(in)
}

func (s *UsersvcServer) SearchUser(ctx context.Context, in *user.SearchUserReq) (*user.SearchUserResp, error) {
	l := logic.NewSearchUserLogic(ctx, s.svcCtx)
	return l.SearchUser(in)
}

// -----------------------auth-----------------------
func (s *UsersvcServer) SendImageCaptcha(ctx context.Context, in *user.SendImageCaptchaReq) (*user.SendImageCaptchaResp, error) {
	l := logic.NewSendImageCaptchaLogic(ctx, s.svcCtx)
	return l.SendImageCaptcha(in)
}

func (s *UsersvcServer) VerifyCaptcha(ctx context.Context, in *user.VerifyCaptchaReq) (*user.VerifyCaptchaResp, error) {
	l := logic.NewVerifyCaptchaLogic(ctx, s.svcCtx)
	return l.VerifyCaptcha(in)
}

func (s *UsersvcServer) SendSmsCode(ctx context.Context, in *user.SendSmsCodeReq) (*user.SendSmsCodeResp, error) {
	l := logic.NewSendSmsCodeLogic(ctx, s.svcCtx)
	return l.SendSmsCode(in)
}
