package api

import (
	v1 "github.com/SaidovZohid/medium_api_gateway/api/v1"
	"github.com/SaidovZohid/medium_api_gateway/config"
	grpcPkg "github.com/SaidovZohid/medium_api_gateway/pkg/grpc_client"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type RoutetOptions struct {
	Cfg        *config.Config
	GrpcClient grpcPkg.GrpcClientI
	Logger     *logrus.Logger
}

// New @title           Swagger for blog api
// @version         2.0
// @description     This is a blog service api.
// @host      		localhost:8000
// @BasePath  		/v1
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
// @Security ApiKeyAuth
func New(opt *RoutetOptions) *gin.Engine {
	router := gin.Default()

	handlerV1 := v1.New(&v1.HandlerV1Options{
		Cfg:        opt.Cfg,
		GrpcClient: opt.GrpcClient,
		Logger:     opt.Logger,
	})

	apiV1 := router.Group("/v1")
	// router.Static("/medias", "./media")

	apiV1.POST("/users",  handlerV1.AuthMiddleWare("users", "create"), handlerV1.CreateUser)
	apiV1.GET("/users/:id", handlerV1.GetUser)
	apiV1.GET("/users/me", handlerV1.AuthMiddleWare("users", "get"), handlerV1.GetUserProfile)
	apiV1.PUT("/users/:id",  handlerV1.AuthMiddleWare("users", "update"), handlerV1.UpdateUser)
	apiV1.DELETE("/users/:id", handlerV1.AuthMiddleWare("users", "delete"), handlerV1.DeleteUser)
	apiV1.GET("/users", handlerV1.GetAllUsers)

	// apiV1.POST("/categories", handlerV1.AuthMiddleWare, handlerV1.CreateCategory)
	// apiV1.GET("/categories/:id", handlerV1.GetCategory)
	// apiV1.PUT("/categories/:id", handlerV1.AuthMiddleWare, handlerV1.UpdateCategory)
	// apiV1.DELETE("/categories/:id", handlerV1.AuthMiddleWare, handlerV1.DeleteCategory)
	// apiV1.GET("/categories", handlerV1.GetAllCategories)

	apiV1.POST("/posts", handlerV1.AuthMiddleWare("posts", "create"), handlerV1.CreatePost)
	apiV1.GET("/posts/:id", handlerV1.GetPost)
	apiV1.PUT("/posts/:id", handlerV1.AuthMiddleWare("posts", "update"), handlerV1.UpdatePost)
	apiV1.DELETE("/posts/:id", handlerV1.AuthMiddleWare("posts", "delete"), handlerV1.DeletePost)
	apiV1.GET("/posts", handlerV1.GetAllPosts)

	// apiV1.POST("/comments", handlerV1.AuthMiddleWare, handlerV1.CreateComment)
	// apiV1.PUT("/comments/:id", handlerV1.AuthMiddleWare, handlerV1.UpdateComment)
	// apiV1.DELETE("/comments/:id", handlerV1.AuthMiddleWare, handlerV1.DeleteComment)
	// apiV1.GET("/comments", handlerV1.GetAllComments)

	// apiV1.POST("/likes", handlerV1.AuthMiddleWare, handlerV1.CreateOrUpdateLike)
	// apiV1.GET("/likes/user-post", handlerV1.AuthMiddleWare, handlerV1.GetLike)

	apiV1.POST("/auth/register", handlerV1.Register)
	apiV1.POST("/auth/login", handlerV1.Login)
	apiV1.POST("/auth/verify", handlerV1.Verify)
	apiV1.POST("/auth/forgot-password", handlerV1.ForgotPassword)
	apiV1.POST("/auth/update-password", handlerV1.UpdatePassword)
	apiV1.POST("/auth/verify-forgot-password", handlerV1.VerifyForgotPassword)

	// apiV1.POST("/file_upload", handlerV1.AuthMiddleWare, handlerV1.UploadFile)

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return router
}
