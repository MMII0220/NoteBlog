package repository

import (
	"fmt"

	_ "github.com/lib/pq"

	// "myasd/config"
	"myasd/internal/models"
)

func (r *RepositoryStruct) GetAllArticles(userID int) ([]models.Article, error) {
	var input []models.Article
	query := `select id, name, content, created_at from articles where user_id=$1 and deleted_at IS NULL`
	err := r.db.Select(&input, query, userID)
	if err != nil {
		return []models.Article{}, r.translateError(err)
	}
	return input, nil
}

func (r *RepositoryStruct) CreateArticle(input models.Article) error {
	query := `insert into articles (name, content, user_id) values($1, $2, $3)`
	_, err := r.db.Exec(query, input.Name, input.Content, input.UserID)
	if err != nil {
		return r.translateError(err)
	}
	return nil
}

func (r *RepositoryStruct) GetArticleByID(userID int, articleID int) (models.Article, error) {
	var article models.Article
	const query = `SELECT id, 
       name,
       content,
       user_id,
       created_at
					FROM articles 
					WHERE deleted_at IS NULL AND id=$1 and user_id=$2`

	err := r.db.Get(&article, query, articleID, userID)
	if err != nil {
		return models.Article{}, r.translateError(err)
	}
	return article, nil
}

// func (r *RepositoryStruct) UpdateArticle(id int, articleID string) error {
// 	return r.UpdateArticle(id, articleID)
// }

func (r *RepositoryStruct) PatchArticle(id int, articleID string, updates map[string]interface{}) error {
	if len(updates) == 0 {
		return nil
	}

	query := "UPDATE articles SET "
	args := []interface{}{}
	i := 1

	for field, value := range updates {
		if i > 1 {
			query += ", "
		}
		query += field + " = $" + fmt.Sprint(i) // для PostgreSQL, для MySQL будет ?
		args = append(args, value)
		i++
	}

	// обновляем updated_at автоматически
	query += ", updated_at = NOW()"

	// условие WHERE
	query += " WHERE id = $" + fmt.Sprint(i)
	args = append(args, articleID)

	// выполняем
	_, err := r.db.Exec(query, args...)
	if err != nil {
		return r.translateError(err)
	}
	return nil
}

func (r *RepositoryStruct) DeleteArticle(userID int, articleID string) error {
	query := "UPDATE articles SET deleted_at = NOW() WHERE id = $1 AND user_id = $2"

	_, err := r.db.Exec(query, articleID, userID)
	if err != nil {
		return r.translateError(err)
	}

	// rows, err := res.RowsAffected()
	// if err != nil {
	// 	return err
	// }
	// if rows == 0 {
	// 	return r.tranlsateError(err)
	// }

	return nil
}
