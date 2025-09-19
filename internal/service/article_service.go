package service

import (
	"myasd/internal/errs"
	"myasd/internal/models"
	// "myasd/internal/repository"
	"database/sql"
	"errors"
)

func (s *ServiceStruct) GetAllArticles(userID int) ([]models.Article, error) {
	return s.repo.GetAllArticles(userID)
}

func (s *ServiceStruct) CreateArticle(input models.Article) error {
	return s.repo.CreateArticle(input)
}

func (s *ServiceStruct) GetArticleByID(id int, articleID string) (models.Article, error) {
	article, err := s.repo.GetArticleByID(id)
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

	_, err := s.repo.GetArticleByID(userID)
	if err != nil {
		return errs.ErrArticleNotFound
	}

	return s.repo.PatchArticle(userID, articleID, updates)
}

func (s *ServiceStruct) DeleteArticle(id int, articleID string) error {
	_, err := s.repo.GetArticleByID(id)
	if err != nil {
		return errs.ErrArticleNotFound
	}

	return s.repo.DeleteArticle(id, articleID)
}
