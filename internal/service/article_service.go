package service

import (
	"myasd/internal/errs"
	"myasd/internal/models"

	// "myasd/internal/repository"
	"database/sql"
	"errors"
	// "log"
	"strconv"
	"time"
)

func (s *ServiceStruct) GetAllArticles(userID int) ([]models.Article, error) {
	start := time.Now()
	articles, err := s.repo.GetAllArticles(userID)
	duration := time.Since(start)

	if err != nil {
		s.logger.Error().
			Str("method", "GetAllArticles").
			Int("user_id", userID).
			Dur("latency", duration).
			Err(err).
			Msg("failed to fetch all articles in service")
		return nil, err
	}

	s.logger.Info().
		Str("method", "GetAllArticles").
		Int("user_id", userID).
		Int("articles_count", len(articles)).
		Dur("latency", duration).
		Msg("fetched all articles successfully in service")

	return articles, nil
}

func (s *ServiceStruct) CreateArticle(input models.Article) error {
	start := time.Now()
	s.logger.Info().
		Str("method", "CreateArticle").
		Str("article_name", input.Name).
		Int("user_id", input.UserID).
		Msg("attempting to create article in service")

	err := s.repo.CreateArticle(input)
	duration := time.Since(start)

	if err != nil {
		s.logger.Error().
			Str("method", "CreateArticle").
			Str("article_name", input.Name).
			Int("user_id", input.UserID).
			Dur("latency", duration).
			Err(err).
			Msg("failed to create article in service")
		return err
	}

	s.logger.Info().
		Str("method", "CreateArticle").
		Str("article_name", input.Name).
		Int("user_id", input.UserID).
		Dur("latency", duration).
		Msg("article created successfully in service")

	return nil
}

func (s *ServiceStruct) GetArticleByID(userID int, articleID int) (models.Article, error) {
	start := time.Now()
	article, err := s.repo.GetArticleByID(userID, articleID)
	duration := time.Since(start)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			s.logger.Warn().
				Str("method", "GetArticleByID").
				Int("user_id", userID).
				Int("article_id", articleID).
				Dur("latency", duration).
				Msg("article not found")
			return models.Article{}, errs.ErrArticleNotFound
		}

		s.logger.Error().
			Str("method", "GetArticleByID").
			Int("user_id", userID).
			Int("article_id", articleID).
			Dur("latency", duration).
			Err(err).
			Msg("failed to fetch article by ID")
		return models.Article{}, err
	}

	s.logger.Info().
		Str("method", "GetArticleByID").
		Int("user_id", userID).
		Int("article_id", articleID).
		Dur("latency", duration).
		Msg("article fetched successfully")

	return article, nil
}

func (s *ServiceStruct) PatchArticle(userID int, articleID string, updates map[string]interface{}) error {
	start := time.Now()

	if len(updates) == 0 {
		s.logger.Warn().
			Str("method", "PatchArticle").
			Int("user_id", userID).
			Str("article_id", articleID).
			Msg("no fields to update")
		return errors.New("no fields to update")
	}

	articleIDInt, err := strconv.Atoi(articleID)
	if err != nil {
		s.logger.Error().
			Str("method", "PatchArticle").
			Int("user_id", userID).
			Str("article_id", articleID).
			Err(err).
			Msg("invalid article ID")
		return errs.ErrArticleNotFound
	}

	_, err = s.repo.GetArticleByID(userID, articleIDInt)
	if err != nil {
		s.logger.Error().
			Str("method", "PatchArticle").
			Int("user_id", userID).
			Str("article_id", articleID).
			Err(err).
			Msg("article not found in service")
		return errs.ErrArticleNotFound
	}

	err = s.repo.PatchArticle(userID, articleID, updates)
	duration := time.Since(start)

	if err != nil {
		s.logger.Error().
			Str("method", "PatchArticle").
			Int("user_id", userID).
			Str("article_id", articleID).
			Interface("updates", updates).
			Dur("latency", duration).
			Err(err).
			Msg("failed to patch article in service")
		return err
	}

	s.logger.Info().
		Str("method", "PatchArticle").
		Int("user_id", userID).
		Str("article_id", articleID).
		Interface("updates", updates).
		Dur("latency", duration).
		Msg("article patched successfully in service")

	return nil
}

func (s *ServiceStruct) DeleteArticle(userID int, articleID string) error {
	articleIDInt, err := strconv.Atoi(articleID)
	if err != nil {
		return errs.ErrArticleNotFound
	}

	_, err = s.repo.GetArticleByID(userID, articleIDInt)
	if err != nil {
		return errs.ErrArticleNotFound
	}

	return s.repo.DeleteArticle(userID, articleID)
}
