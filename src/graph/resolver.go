package graph

import (
	"log"

	"github.com/schema-creator/graph-gateway/cmd/config"
	grpcclient "github.com/schema-creator/graph-gateway/pkg/grpc-client"
	"github.com/schema-creator/graph-gateway/src/gateways"
	"github.com/schema-creator/graph-gateway/src/services"

	imageService "github.com/schema-creator/graph-gateway/pkg/grpc/image-service/v1"
	memberService "github.com/schema-creator/graph-gateway/pkg/grpc/member-service/v1"
	projectService "github.com/schema-creator/graph-gateway/pkg/grpc/project-service/v1"
	saveService "github.com/schema-creator/graph-gateway/pkg/grpc/save-service/v1"
	tokenService "github.com/schema-creator/graph-gateway/pkg/grpc/token-service/v1"
	userService "github.com/schema-creator/graph-gateway/pkg/grpc/user-service/v1"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	UserService    services.UserService
	ProjectService services.ProjectService
	MemberService  services.MemberService
	SaveService    services.SaveService
}

func NewResolver() *Resolver {
	log.Printf(`
		[server] user-service: %v
		[server] project-service: %v
		[server] image-service: %v
		[server] member-service: %v
		[server] save-service: %v	
	`,
		config.Config.Service.UserServiceAddr,
		config.Config.Service.ProjectServiceAddr,
		config.Config.Service.ImageServiceAddr,
		config.Config.Service.MemberServiceAddr,
		config.Config.Service.SaveServiceAddr,
	)

	userConn, err := grpcclient.Connect(config.Config.Service.UserServiceAddr)
	if err != nil {
		log.Fatalf("failed to connect user-service: %v", err)
	}
	log.Println("userConn ok")
	tokenConn, err := grpcclient.Connect(config.Config.Service.TokenServiceAddr)
	if err != nil {
		log.Fatalf("failed to connect token-service: %v", err)
	}

	log.Println("tokenConn ok")
	projectConn, err := grpcclient.Connect(config.Config.Service.ProjectServiceAddr)
	if err != nil {
		log.Fatalf("failed to connect project-service: %v", err)
	}

	log.Println("projectConn ok")
	imageConn, err := grpcclient.Connect(config.Config.Service.ImageServiceAddr)
	if err != nil {
		log.Fatalf("failed to connect image-service: %v", err)
	}

	log.Println("imageConn ok")
	memberConn, err := grpcclient.Connect(config.Config.Service.MemberServiceAddr)
	if err != nil {
		log.Fatalf("failed to connect member-service: %v", err)
	}

	log.Println("memberConn ok")
	saveConn, err := grpcclient.Connect(config.Config.Service.SaveServiceAddr)
	if err != nil {
		log.Fatalf("failed to connect save-service: %v", err)
	}

	log.Println("saveConn ok")
	return &Resolver{
		UserService: services.NewUserService(gateways.NewUserClient(userService.NewUserClient(userConn))),
		ProjectService: services.NewProjectService(
			gateways.NewProjectClient(projectService.NewProjectClient(projectConn)),
			gateways.NewMemberClient(memberService.NewMemberClient(memberConn)),
			gateways.NewTokenClient(tokenService.NewTokenClient(tokenConn)),
			gateways.NewImageClient(imageService.NewImageClient(imageConn)),
		),
		MemberService: services.NewMemberService(
			gateways.NewMemberClient(memberService.NewMemberClient(memberConn)),
			gateways.NewTokenClient(tokenService.NewTokenClient(tokenConn)),
		),
		SaveService: services.NewSaveService(gateways.NewSaveClient(saveService.NewSaveClient(saveConn))),
	}
}
