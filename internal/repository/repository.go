package repository

import (
	"fmt"
	_ "github.com/lib/pq"
	"myasd/config"
	"myasd/internal/models"
)

func GetAllArticles(userID int) ([]models.Article, error) {
	var input []models.Article
	query := `select id, name, content, created_at from articles where user_id=$1`
	err := config.GetDBConnection().Select(&input, query, userID)
	if err != nil {
		return []models.Article{}, fmt.Errorf("error in select query for database %v", err)
	}
	return input, err
}

func CreateArticle(input models.Article) error {
	query := `insert into articles (name, content, user_id) values($1, $2, $3)`
	_, err := config.GetDBConnection().Exec(query, input.Name, input.Content, input.UserID)
	if err != nil {
		return fmt.Errorf("error in query request insert into %v: ", err)
	}
	return nil
}
