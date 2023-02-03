// Code generated by goctl. DO NOT EDIT!
// Source: user.proto

package server

import (
	"context"

	"middle/app/user/rpc/internal/logic"
	"middle/app/user/rpc/internal/svc"
	"middle/app/user/rpc/types/user"
)

type UserServer struct {
	svcCtx *svc.ServiceContext
	user.UnimplementedUserServer
}

func NewUserServer(svcCtx *svc.ServiceContext) *UserServer {
	return &UserServer{
		svcCtx: svcCtx,
	}
}

// -----------------------users-----------------------
func (s *UserServer) AddUsers(ctx context.Context, in *user.AddUsersReq) (*user.AddUsersResp, error) {
	l := logic.NewAddUsersLogic(ctx, s.svcCtx)
	return l.AddUsers(in)
}

func (s *UserServer) UpdateUsers(ctx context.Context, in *user.UpdateUsersReq) (*user.UpdateUsersResp, error) {
	l := logic.NewUpdateUsersLogic(ctx, s.svcCtx)
	return l.UpdateUsers(in)
}

func (s *UserServer) DelUsers(ctx context.Context, in *user.DelUsersReq) (*user.DelUsersResp, error) {
	l := logic.NewDelUsersLogic(ctx, s.svcCtx)
	return l.DelUsers(in)
}

func (s *UserServer) GetUsersById(ctx context.Context, in *user.GetUsersByIdReq) (*user.GetUsersByIdResp, error) {
	l := logic.NewGetUsersByIdLogic(ctx, s.svcCtx)
	return l.GetUsersById(in)
}

func (s *UserServer) SearchUsers(ctx context.Context, in *user.SearchUsersReq) (*user.SearchUsersResp, error) {
	l := logic.NewSearchUsersLogic(ctx, s.svcCtx)
	return l.SearchUsers(in)
}
