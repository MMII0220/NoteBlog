package service

import (
	"myasd/internal/models"
	"myasd/internal/repository"
)

func GetAllArticles(userID int) ([]models.Article, error) {
	return repository.GetAllArticles(userID)
}

func CreateArticle(input models.Article) error {
	return repository.CreateArticle(input)
}
