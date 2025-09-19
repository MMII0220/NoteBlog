package controller

import (
	"github.com/gin-gonic/gin"
	"myasd/internal/models"
	// "myasd/internal/service"
	"net/http"
)

// @Summary Get all articles of the user
// @Description Returns articles of the authenticated user
// @Tags articles
// @Accept json
// @Produce json
// @Success 200 {array} models.Article
// @Failure 401 {object} map[string]string
// @Router /articles [get]
// @Security BearerAuth
func (contr *ControllerStruct) getAllArticles(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "unauthorized",
		})
		return
	}

	// Приводим userID к int
	id, ok := userID.(int)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "invalid user id type",
		})
		return
	}

	allArticles, err := contr.serv.GetAllArticles(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"success": allArticles,
	})
}

// @Summary Create a new article
// @Description Create a new article for the authenticated user
// @Tags articles
// @Accept json
// @Produce json
// @Param article body models.Article true "Article data"
// @Success 201 {object} models.Article
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Router /articles [post]
// @Security BearerAuth
func (contr *ControllerStruct) createArticle(c *gin.Context) {
	var input models.Article
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "unauthorized",
		})
		return
	}

	input.UserID = userID.(int)

	err := contr.serv.CreateArticle(input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"message": input,
	})
}

// @Summary Get all articles of the user
// @Description Returns articles of the authenticated user
// @Tags articles
// @Accept json
// @Produce json
// @Success 200 {array} models.Article
// @Failure 500 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Router /articles [get]
// @Security BearerAuth
func (contr *ControllerStruct) getArticleByID(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "unauthorized",
		})
		return
	}
	id, ok := userID.(int)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "invalid user id type",
		})
		return
	}

	articleID := c.Param("id")

	article, err := contr.serv.GetArticleByID(id, articleID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": article,
	})
}

// PatchArticle godoc
// @Summary Update an article
// @Description Update fields of an article. JSON body must contain fields to update.
// @Tags articles
// @Accept json
// @Produce json
// @Param id path string true "Article ID"
// @Param body body map[string]interface{} true "Fields to update"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /articles/{id} [patch]
func (ctr *ControllerStruct) PatchArticle(c *gin.Context) {
	userIDInterface, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	userID, ok := userIDInterface.(int)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "invalid user id"})
		return
	}

	articleID := c.Param("id")

	var updates map[string]interface{}
	if err := c.BindJSON(&updates); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid json"})
		return
	}

	if err := ctr.serv.PatchArticle(userID, articleID, updates); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "article updated"})
}

// DeleteArticle godoc
// @Summary Delete an article
// @Description Delete an article by ID for the authenticated user
// @Tags articles
// @Accept json
// @Produce json
// @Param id path string true "Article ID"
// @Success 204 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /articles/{id} [delete]
func (contr *ControllerStruct) deleteArticle(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "unauthorized",
		})
		return
	}
	id, ok := userID.(int)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid user id type",
		})
		return
	}

	articleID := c.Param("id")

	err := contr.serv.DeleteArticle(id, articleID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusNoContent, gin.H{
		"message": "successfull deleted",
	})
}

func (contr *ControllerStruct) ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "successfull",
	})
}
