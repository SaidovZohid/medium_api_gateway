package grpc_client

import (
	"fmt"

	"github.com/SaidovZohid/medium_api_gateway/config"
	pbp "github.com/SaidovZohid/medium_api_gateway/genproto/post_service"
	pbu "github.com/SaidovZohid/medium_api_gateway/genproto/user_service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type GrpcClientI interface {
	UserService() pbu.UserServiceClient
	AuthService() pbu.AuthServiceClient
	PostService() pbp.PostServiceClient
	LikeService() pbp.LikeServiceClient
	CommentService() pbp.CommentServiceClient
	CategoryService() pbp.CategoryServiceClient
}

type GrpcClient struct {
	cfg         config.Config
	connections map[string]interface{}
}

func New(cfg config.Config) (GrpcClientI, error) {
	connUserService, err := grpc.Dial(
		fmt.Sprintf("%s%s", cfg.UserServiceHost, cfg.UserServiceGrpcPort),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return nil, fmt.Errorf("user service dial host: %v port: %v", cfg.UserServiceHost, cfg.UserServiceGrpcPort)
	}
	connPostService, err := grpc.Dial(
		fmt.Sprintf("%s%s", cfg.PostServiceHost, cfg.PostServiceGrpcPort),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return nil, fmt.Errorf("post service dial host: %v port: %v", cfg.UserServiceHost, cfg.UserServiceGrpcPort)
	}
	connLikeService, err := grpc.Dial(
		fmt.Sprintf("%s%s", cfg.PostServiceHost, cfg.PostServiceGrpcPort),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return nil, fmt.Errorf("like service dial host: %v port: %v", cfg.UserServiceHost, cfg.UserServiceGrpcPort)
	}

	connCategoryService, err := grpc.Dial(
		fmt.Sprintf("%s%s", cfg.PostServiceHost, cfg.PostServiceGrpcPort),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return nil, fmt.Errorf("category service dial host: %v port: %v", cfg.UserServiceHost, cfg.UserServiceGrpcPort)
	}

	connCommentService, err := grpc.Dial(
		fmt.Sprintf("%s%s", cfg.PostServiceHost, cfg.PostServiceGrpcPort),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return nil, fmt.Errorf("comment service dial host: %v port: %v", cfg.UserServiceHost, cfg.UserServiceGrpcPort)
	}

	return &GrpcClient{
		cfg: cfg,
		connections: map[string]interface{}{
			"user_service":     pbu.NewUserServiceClient(connUserService),
			"auth_service":     pbu.NewAuthServiceClient(connUserService),
			"post_service":     pbp.NewPostServiceClient(connPostService),
			"like_service":     pbp.NewLikeServiceClient(connLikeService),
			"category_service": pbp.NewCategoryServiceClient(connCategoryService),
			"comment_service":  pbp.NewCommentServiceClient(connCommentService),
		},
	}, nil
}

func (g *GrpcClient) UserService() pbu.UserServiceClient {
	return g.connections["user_service"].(pbu.UserServiceClient)
}

func (g *GrpcClient) AuthService() pbu.AuthServiceClient {
	return g.connections["auth_service"].(pbu.AuthServiceClient)
}

func (g *GrpcClient) PostService() pbp.PostServiceClient {
	return g.connections["post_service"].(pbp.PostServiceClient)
}

func (g *GrpcClient) LikeService() pbp.LikeServiceClient {
	return g.connections["like_service"].(pbp.LikeServiceClient)
}

func (g *GrpcClient) CommentService() pbp.CommentServiceClient {
	return g.connections["comment_service"].(pbp.CommentServiceClient)
}

func (g *GrpcClient) CategoryService() pbp.CategoryServiceClient {
	return g.connections["category_service"].(pbp.CategoryServiceClient)
}
