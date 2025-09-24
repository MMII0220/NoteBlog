package controller

import (
	"myasd/internal/errs"
	"myasd/internal/models"

	"github.com/gin-gonic/gin"

	"errors"
	"net/http"
	"strconv"
	"time"
)

// @Summary Get all articles of the user
// @Description Returns articles of the authenticated user
// @Tags articles
// @Accept json
// @Produce json
// @Success 200 {array} models.Article
// @Failure 401 {object} CommonError
// @Router /articles [get]
// @Security BearerAuth
func (contr *ControllerStruct) getAllArticles(c *gin.Context) {
	start := time.Now()

	userID, exists := c.Get("user_id")
	if !exists {
		contr.logger.Warn().
			Str("method", "getAllArticles").
			Msg("user_id not found in context")
		contr.handleError(c, errs.ErrUserIDNotFoundInContext)
		return
	}

	id, ok := userID.(int)
	if !ok {
		contr.logger.Error().
			Str("method", "getAllArticles").
			Msg("invalid user id type")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "invalid user id type"})
		return
	}

	articles, err := contr.serv.GetAllArticles(id)
	duration := time.Since(start)

	if err != nil {
		contr.logger.Error().
			Str("method", "getAllArticles").
			Int("user_id", id).
			Dur("latency", duration).
			Err(err).
			Msg("failed to fetch articles")
		contr.handleError(c, err)
		return
	}

	contr.logger.Info().
		Str("method", "getAllArticles").
		Int("user_id", id).
		Int("articles_count", len(articles)).
		Dur("latency", duration).
		Msg("fetched all articles successfully")

	c.JSON(http.StatusOK, gin.H{"success": articles})
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
	start := time.Now()

	var input models.Article
	if err := c.ShouldBindJSON(&input); err != nil {
		contr.logger.Warn().
			Str("method", "createArticle").
			Err(err).
			Msg("invalid request body")
		contr.handleError(c, errors.Join(errs.ErrInvalidRequestBody, err))
		return
	}

	userID, exists := c.Get("user_id")
	if !exists {
		contr.logger.Warn().
			Str("method", "createArticle").
			Msg("user_id not found in context")
		contr.handleError(c, errs.ErrUserIDNotFoundInContext)
		return
	}

	input.UserID = userID.(int)

	if err := contr.serv.CreateArticle(input); err != nil {
		contr.logger.Error().
			Str("method", "createArticle").
			Int("user_id", input.UserID).
			Err(err).
			Msg("failed to create article")
		contr.handleError(c, err)
		return
	}

	duration := time.Since(start)
	contr.logger.Info().
		Str("method", "createArticle").
		Int("user_id", input.UserID).
		Dur("latency", duration).
		Msg("article created successfully")

	c.JSON(http.StatusCreated, CommonResponse{"article created successfully"})
}

// @Summary Get all articles of the user
// @Description Returns articles of the authenticated user
// @Tags articles
// @Accept json
// @Produce json
// @Param id path int true "Article ID"
// @Success 200 {array} models.Article
// @Failure 500 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Router /articles/{id} [get]
// @Security BearerAuth
func (contr *ControllerStruct) getArticleByID(c *gin.Context) {
	start := time.Now()

	userID, exists := c.Get("user_id")
	if !exists {
		contr.logger.Warn().
			Str("method", "getArticleByID").
			Msg("user_id not found in context")
		contr.handleError(c, errs.ErrUserIDNotFoundInContext)
		return
	}
	id := userID.(int)

	articleIDStr := c.Param("id")
	articleID, err := strconv.Atoi(articleIDStr)
	if err != nil || articleID < 1 {
		contr.logger.Warn().
			Str("method", "getArticleByID").
			Str("article_id_param", articleIDStr).
			Err(err).
			Msg("invalid article ID parameter")
		contr.handleError(c, errs.ErrInvalidPathParam)
		return
	}

	article, err := contr.serv.GetArticleByID(id, articleID)
	duration := time.Since(start)

	if err != nil {
		contr.logger.Error().
			Str("method", "getArticleByID").
			Int("user_id", id).
			Int("article_id", articleID).
			Dur("latency", duration).
			Err(err).
			Msg("failed to get article by ID")
		contr.handleError(c, err)
		return
	}

	contr.logger.Info().
		Str("method", "getArticleByID").
		Int("user_id", id).
		Int("article_id", articleID).
		Dur("latency", duration).
		Msg("article retrieved successfully")

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
// @Security BearerAuth
func (contr *ControllerStruct) PatchArticle(c *gin.Context) {
	start := time.Now()

	userIDInterface, exists := c.Get("user_id")
	if !exists {
		contr.logger.Warn().
			Str("method", "PatchArticle").
			Msg("user_id not found in context")
		contr.handleError(c, errs.ErrUserIDNotFoundInContext)
		return
	}

	userID, ok := userIDInterface.(int)
	if !ok {
		contr.logger.Error().
			Str("method", "PatchArticle").
			Msg("invalid user_id type in context")
		contr.handleError(c, errs.ErrUserIDNotFoundInContext)
		return
	}

	articleID := c.Param("id")

	var updates map[string]interface{}
	if err := c.ShouldBindJSON(&updates); err != nil {
		contr.logger.Warn().
			Str("method", "PatchArticle").
			Str("article_id", articleID).
			Err(err).
			Msg("invalid request body")
		contr.handleError(c, errors.Join(errs.ErrInvalidRequestBody, err))
		return
	}

	if err := contr.serv.PatchArticle(userID, articleID, updates); err != nil {
		duration := time.Since(start)
		contr.logger.Error().
			Str("method", "PatchArticle").
			Int("user_id", userID).
			Str("article_id", articleID).
			Dur("latency", duration).
			Err(err).
			Msg("failed to update article")
		contr.handleError(c, err)
		return
	}

	duration := time.Since(start)
	contr.logger.Info().
		Str("method", "PatchArticle").
		Int("user_id", userID).
		Str("article_id", articleID).
		Dur("latency", duration).
		Msg("article updated successfully")

	c.JSON(http.StatusOK, gin.H{"message": "article updated"})
}

// DeleteArticle godoc
// @Summary Delete an article
// @Description Delete an article by ID for the authenticated user
// @Tags articles
// @Accept json
// @Produce json
// @Param id path int true "Article ID"
// @Success 204 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /articles/{id} [delete]
// @Security BearerAuth
func (contr *ControllerStruct) deleteArticle(c *gin.Context) {
	start := time.Now()

	userID, exists := c.Get("user_id")
	if !exists {
		contr.logger.Warn().
			Str("method", "deleteArticle").
			Msg("user_id not found in context")
		contr.handleError(c, errs.ErrUserIDNotFoundInContext)
		return
	}
	id := userID.(int)

	articleIDStr := c.Param("id")
	articleID, err := strconv.Atoi(articleIDStr)
	if err != nil || articleID < 1 {
		contr.logger.Warn().
			Str("method", "deleteArticle").
			Str("article_id_param", articleIDStr).
			Err(err).
			Msg("invalid article ID parameter")
		contr.handleError(c, errs.ErrInvalidPathParam)
		return
	}

	err = contr.serv.DeleteArticle(id, articleIDStr)
	duration := time.Since(start)

	if err != nil {
		contr.logger.Error().
			Str("method", "deleteArticle").
			Int("user_id", id).
			Int("article_id", articleID).
			Dur("latency", duration).
			Err(err).
			Msg("failed to delete article")
		contr.handleError(c, err)
		return
	}

	contr.logger.Info().
		Str("method", "deleteArticle").
		Int("user_id", id).
		Int("article_id", articleID).
		Dur("latency", duration).
		Msg("article deleted successfully")

	c.JSON(http.StatusNoContent, gin.H{
		"message": "successfull deleted",
	})
}

func (contr *ControllerStruct) ping(c *gin.Context) {
	start := time.Now()

	contr.logger.Info().
		Str("method", "ping").
		Str("endpoint", "/ping").
		Msg("ping endpoint accessed")

	c.JSON(http.StatusOK, gin.H{
		"message": "successfull",
	})

	duration := time.Since(start)
	contr.logger.Info().
		Str("method", "ping").
		Dur("latency", duration).
		Msg("ping response sent")
}
