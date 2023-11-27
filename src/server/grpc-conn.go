package server

import (
	"github.com/Hack-Hack-geek-Vol10/graph-gateway/cmd/config"
	grpcclient "github.com/Hack-Hack-geek-Vol10/graph-gateway/pkg/grpc-client"
	"github.com/Hack-Hack-geek-Vol10/graph-gateway/src/gateways"
	"github.com/Hack-Hack-geek-Vol10/graph-gateway/src/graph"
	"github.com/Hack-Hack-geek-Vol10/graph-gateway/src/services"

	imageService "github.com/Hack-Hack-geek-Vol10/graph-gateway/pkg/grpc/image-service"
	memberService "github.com/Hack-Hack-geek-Vol10/graph-gateway/pkg/grpc/member-service"
	projectService "github.com/Hack-Hack-geek-Vol10/graph-gateway/pkg/grpc/project-service"
	tokenService "github.com/Hack-Hack-geek-Vol10/graph-gateway/pkg/grpc/token-service"
	userService "github.com/Hack-Hack-geek-Vol10/graph-gateway/pkg/grpc/user-service"
)

func NewResolver() (*graph.Resolver, error) {
	userConn, err := grpcclient.Connect(config.Config.Service.UserServiceAddr)
	if err != nil {
		return nil, err
	}

	tokenConn, err := grpcclient.Connect(config.Config.Service.TokenServiceAddr)
	if err != nil {
		return nil, err
	}

	projectConn, err := grpcclient.Connect(config.Config.Service.ProjectServiceAddr)
	if err != nil {
		return nil, err
	}

	imageConn, err := grpcclient.Connect(config.Config.Service.ImageServiceAddr)
	if err != nil {
		return nil, err
	}

	memberConn, err := grpcclient.Connect(config.Config.Service.MemberServiceAddr)
	if err != nil {
		return nil, err
	}

	return &graph.Resolver{
		UserService: services.NewUserService(gateways.NewUserClien(userService.NewUserServiceClient(userConn))),
		ProjectService: services.NewProjectService(
			gateways.NewProjectClient(projectService.NewProjectServiceClient(projectConn)),
			gateways.NewMemberClient(memberService.NewMemberServiceClient(memberConn)),
			gateways.NewTokenClient(tokenService.NewTokenServiceClient(tokenConn)),
			gateways.NewImageClient(imageService.NewImageServiceClient(imageConn)),
		),
		MemberService: services.NewMemberService(
			gateways.NewMemberClient(memberService.NewMemberServiceClient(memberConn)),
			gateways.NewTokenClient(tokenService.NewTokenServiceClient(tokenConn)),
		),
	}, nil
}
