package controller

import (
	"github.com/gin-gonic/gin"
	"myasd/internal/models"
	// "myasd/internal/service"
	// "myasd/internal/service"
	"net/http"
)

// signUp godoc
// @Summary      Регистрация пользователя
// @Description  Создает нового пользователя
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        user  body      models.User  true  "Данные пользователя"
// @Success      201   {object}  models.User
// @Failure      400   {object}  map[string]string
// @Failure      500   {object}  map[string]string
// @Router       /signup [post]
func (contr *ControllerStruct) signUp(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "error in request body",
		})
		return
	}

	err := contr.serv.CreateUser(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": user,
	})
}

// SignInRequest описывает тело запроса для входа
type SignInRequest struct {
	Login    string `json:"login"`
	Password string `json:"-"`
}

// signIn godoc
// @Summary      Вход пользователя
// @Description  Получение access и refresh токенов
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        credentials  body      SignInRequest  true  "Данные для входа"
// @Success      200          {object}  map[string]string
// @Failure      400          {object}  map[string]string
// @Failure      404          {object}  map[string]string
// @Router       /signin [post]
func (contr *ControllerStruct) signIn(c *gin.Context) {
	var user SignInRequest
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "error bad request body",
		})
		return
	}

	tokens, err := contr.serv.GetUser(user.Login, user.Password)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": tokens,
	})
}

type RefreshTokenRequest struct {
	RefreshToken string `json:"refresh_token"`
}

// RefreshToken godoc
// @Summary      Обновление access токена
// @Description  Генерирует новый access token по refresh token
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        token  body      RefreshTokenRequest  true  "Refresh токен"
// @Success      200    {object}  map[string]string
// @Failure      400    {object}  map[string]string
// @Failure      401    {object}  map[string]string
// @Router       /refresh [post]
func (contr *ControllerStruct) RefreshToken(c *gin.Context) {
	var input RefreshTokenRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	access, err := contr.serv.RefreshToken(input.RefreshToken)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"access_token": access,
	})
}
