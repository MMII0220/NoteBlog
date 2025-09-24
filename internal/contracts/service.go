package contracts

import "myasd/internal/models"

type ServiceI interface {
	// User mewthods
	CreateUser(user models.User) error
	GetUser(login, password string) (models.TokenResponse, error)

	// JWT logic
	RefreshToken(refreshToken string) (string, error)
	GenerateTokens(userID int) (accessToken, refreshToken string, err error)
	ValidateAccessToken(tokenStr string) (userID int, err error)
	ValidateRefreshToken(tokenStr string) (userID int, err error)

	// Articles Methods
	GetAllArticles(userID int) ([]models.Article, error)
	CreateArticle(input models.Article) error
	GetArticleByID(userID int, articleID int) (models.Article, error)
	PatchArticle(userID int, articleID string, updates map[string]interface{}) error
	DeleteArticle(userID int, articleID string) error
}
