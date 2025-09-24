package service

import (
	"myasd/internal/errs"
	"myasd/internal/models"

	// "myasd/internal/repository"
	"database/sql"
	"errors"
	"strconv"
)

func (s *ServiceStruct) GetAllArticles(userID int) ([]models.Article, error) {
	return s.repo.GetAllArticles(userID)
}

func (s *ServiceStruct) CreateArticle(input models.Article) error {
	return s.repo.CreateArticle(input)
}

func (s *ServiceStruct) GetArticleByID(userID int, articleID int) (models.Article, error) {
	article, err := s.repo.GetArticleByID(userID, articleID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			// return models.Article{}, s.repo.translateError(err)
			return models.Article{}, errs.ErrArticleNotFound
		}
		return models.Article{}, err
	}
	return article, nil
}

func (s *ServiceStruct) PatchArticle(userID int, articleID string, updates map[string]interface{}) error {
	if len(updates) == 0 {
		return errors.New("no fields to update")
	}

	articleIDInt, err := strconv.Atoi(articleID)
	if err != nil {
		return errs.ErrArticleNotFound
	}

	_, err = s.repo.GetArticleByID(userID, articleIDInt)
	if err != nil {
		return errs.ErrArticleNotFound
	}

	return s.repo.PatchArticle(userID, articleID, updates)
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
