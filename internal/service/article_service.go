package service

import (
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
			return models.Article{}, errors.New("статья c таким ID не найдена")
		}
		return models.Article{}, err
	}
	return article, nil
}

// article, err := s.GetArticleByID(id, articleID)
// 	if err != nil {
// 		if errors.Is(err, sql.ErrNoRows) {
// 			return models.Article{}, errors.New("статья c таким ID не найдена")
// 		}
// 		return models.Article{}, err
// 	}
// 	return article, nil
// func (s *ServiceStruct) UpdateArticle(id, articleID int) error {
// 	return s.repo.UpdateArticle(id, articleID)
// }

func (s *ServiceStruct) PatchArticle(userID int, articleID string, updates map[string]interface{}) error {
	if len(updates) == 0 {
		return errors.New("no fields to update")
	}

	return s.repo.PatchArticle(userID, articleID, updates)
}

func (s *ServiceStruct) DeleteArticle(id int, articleID string) error {
	return s.repo.DeleteArticle(id, articleID)
}
