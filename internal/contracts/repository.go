package contracts

import "myasd/internal/models"

type RepositoryI interface {
	//User Methods
	CreateUser(user models.User) error
	GetUserByLogin(login string) (models.User, error)

	// Articles Methods
	GetAllArticles(userID int) ([]models.Article, error)
	CreateArticle(input models.Article) error
	GetArticleByID(id int) (models.Article, error)
	PatchArticle(id int, articleID string, updates map[string]interface{}) error
	DeleteArticle(id int, articleID string) error
}
