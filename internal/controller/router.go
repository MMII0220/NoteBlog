package controller

import (
	"github.com/gin-gonic/gin"
	// "myasd/internal/controller"
	"github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "myasd/docs"
	// "myasd/internal/middleware"
)

func (contr *ControllerStruct) StartRoute() error {
	r := gin.Default()

	// Swagger endpoint
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	//Loggirovanie
	r.Use(contr.LoggerMiddleware())

	r.POST("/signup", contr.signUp)
	r.POST("/signin", contr.signIn)
	r.POST("/refresh", contr.RefreshToken)

	auth := r.Group("/articles")
	auth.Use(contr.authMiddleware())
	{
		auth.GET("/", contr.getAllArticles) // только для авторизованных
		auth.POST("", contr.createArticle)
		auth.GET("/:id", contr.getArticleByID)
		auth.PATCH("/:id", contr.PatchArticle) // только для авторизованных
		auth.DELETE("/:id", contr.deleteArticle)
	}

	r.GET("/ping", contr.ping)

	err := r.Run(":7999")
	if err != nil {
		return err
	}
	return nil
}
