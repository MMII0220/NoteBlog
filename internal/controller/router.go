package controller

import (
	"github.com/gin-gonic/gin"
	// "myasd/internal/controller"
	"github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "myasd/docs"
	"myasd/internal/middleware"
)

func StartRoute() error {
	r := gin.Default()

	// Swagger endpoint
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.POST("/signup", signUp)
	r.POST("/signin", signIn)
	r.POST("/refresh", RefreshToken)

	auth := r.Group("/articles")
	auth.Use(middleware.AuthMiddleware())
	{
		auth.GET("/", getAllArticles) // только для авторизованных
		auth.POST("/", createArticle)
	}

	r.GET("/ping", ping)

	err := r.Run(":7999")
	if err != nil {
		return err
	}
	return nil
}
