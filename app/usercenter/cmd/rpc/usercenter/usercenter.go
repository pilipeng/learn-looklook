// Code generated by goctl. DO NOT EDIT.
// Source: usercenter.proto

package usercenter

import (
	"context"

	"cmd/rpc/pb"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	GenerateTokenReq         = pb.GenerateTokenReq
	GenerateTokenResp        = pb.GenerateTokenResp
	GetUserAuthByAuthKeyReq  = pb.GetUserAuthByAuthKeyReq
	GetUserAuthByAuthKeyResp = pb.GetUserAuthByAuthKeyResp
	GetUserAuthByUserIdReq   = pb.GetUserAuthByUserIdReq
	GetUserAuthByUserIdResp  = pb.GetUserAuthByUserIdResp
	GetUserInfoReq           = pb.GetUserInfoReq
	GetUserInfoResp          = pb.GetUserInfoResp
	LoginReq                 = pb.LoginReq
	LoginResp                = pb.LoginResp
	RegisterReq              = pb.RegisterReq
	RegisterResp             = pb.RegisterResp
	User                     = pb.User
	UserAuth                 = pb.UserAuth

	Usercenter interface {
		Login(ctx context.Context, in *LoginReq, opts ...grpc.CallOption) (*LoginResp, error)
		Register(ctx context.Context, in *RegisterReq, opts ...grpc.CallOption) (*RegisterResp, error)
		GetUserInfo(ctx context.Context, in *GetUserInfoReq, opts ...grpc.CallOption) (*GetUserInfoResp, error)
		GetUserAuthByAuthKey(ctx context.Context, in *GetUserAuthByAuthKeyReq, opts ...grpc.CallOption) (*GetUserAuthByAuthKeyResp, error)
		GetUserAuthByUserId(ctx context.Context, in *GetUserAuthByUserIdReq, opts ...grpc.CallOption) (*GetUserAuthByUserIdResp, error)
		GenerateToken(ctx context.Context, in *GenerateTokenReq, opts ...grpc.CallOption) (*GenerateTokenResp, error)
	}

	defaultUsercenter struct {
		cli zrpc.Client
	}
)

func NewUsercenter(cli zrpc.Client) Usercenter {
	return &defaultUsercenter{
		cli: cli,
	}
}

func (m *defaultUsercenter) Login(ctx context.Context, in *LoginReq, opts ...grpc.CallOption) (*LoginResp, error) {
	client := pb.NewUsercenterClient(m.cli.Conn())
	return client.Login(ctx, in, opts...)
}

func (m *defaultUsercenter) Register(ctx context.Context, in *RegisterReq, opts ...grpc.CallOption) (*RegisterResp, error) {
	client := pb.NewUsercenterClient(m.cli.Conn())
	return client.Register(ctx, in, opts...)
}

func (m *defaultUsercenter) GetUserInfo(ctx context.Context, in *GetUserInfoReq, opts ...grpc.CallOption) (*GetUserInfoResp, error) {
	client := pb.NewUsercenterClient(m.cli.Conn())
	return client.GetUserInfo(ctx, in, opts...)
}

func (m *defaultUsercenter) GetUserAuthByAuthKey(ctx context.Context, in *GetUserAuthByAuthKeyReq, opts ...grpc.CallOption) (*GetUserAuthByAuthKeyResp, error) {
	client := pb.NewUsercenterClient(m.cli.Conn())
	return client.GetUserAuthByAuthKey(ctx, in, opts...)
}

func (m *defaultUsercenter) GetUserAuthByUserId(ctx context.Context, in *GetUserAuthByUserIdReq, opts ...grpc.CallOption) (*GetUserAuthByUserIdResp, error) {
	client := pb.NewUsercenterClient(m.cli.Conn())
	return client.GetUserAuthByUserId(ctx, in, opts...)
}

func (m *defaultUsercenter) GenerateToken(ctx context.Context, in *GenerateTokenReq, opts ...grpc.CallOption) (*GenerateTokenResp, error) {
	client := pb.NewUsercenterClient(m.cli.Conn())
	return client.GenerateToken(ctx, in, opts...)
}
