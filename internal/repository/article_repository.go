package repository

import (
	"fmt"

	_ "github.com/lib/pq"

	// "myasd/config"
	"myasd/internal/models"
	"time"
)

func (r *RepositoryStruct) GetAllArticles(userID int) ([]models.Article, error) {
	start := time.Now()
	var input []models.Article
	query := `select id, name, content, created_at from articles where user_id=$1 and deleted_at IS NULL`

	duration := time.Since(start)
	err := r.db.Select(&input, query, userID)
	if err != nil {
		r.logger.Error().Str("method ", "GetAllArticles").Int("user_id ", userID).Dur("latency ", duration).Err(err).Msg("failed to fecth articles")
		return []models.Article{}, r.translateError(err)
	}

	r.logger.Info().Str("method ", "GetAllArticles").Int("user_id ", userID).Int("articles_count ", len(input)).Dur("duration ", duration).Msg("fethced all articles successfully")
	return input, nil
}

func (r *RepositoryStruct) CreateArticle(input models.Article) error {
	start := time.Now()
	duration := time.Since(start)
	r.logger.Info().Str("method ", "CreateArticle").Str("article_name ", input.Name).Int("user_id ", input.UserID).Msg("attempt to create article")

	query := `insert into articles (name, content, user_id) values($1, $2, $3)`
	_, err := r.db.Exec(query, input.Name, input.Content, input.UserID)
	if err != nil {
		r.logger.Error().Str("method ", "CreatedArticle").Err(err).Dur("duration ", duration).Msg("failed to create article")

		return r.translateError(err)
	}

	r.logger.Info().Str("method ", "CreateArticle").Str("article-name ", input.Name).Int("user_id ", input.UserID).Dur("duration ", duration).Msg("article created successfully")
	return nil
}

func (r *RepositoryStruct) GetArticleByID(userID int, articleID int) (models.Article, error) {
	start := time.Now()
	var article models.Article
	const query = `SELECT id, 
       name,
       content,
       user_id,
       created_at
					FROM articles 
					WHERE deleted_at IS NULL AND id=$1 and user_id=$2`

	err := r.db.Get(&article, query, articleID, userID)
	duration := time.Since(start)

	if err != nil {
		r.logger.Error().
			Str("method", "GetArticleByID").
			Int("article_id", articleID).
			Dur("latency", duration).
			Err(err).
			Msg("failed to fetch article by ID")
		return models.Article{}, r.translateError(err)
	}

	r.logger.Info().
		Str("method", "GetArticleByID").
		Int("article_id", articleID).
		Dur("latency", duration).
		Msg("article fetched successfully")
	return article, nil
}

// func (r *RepositoryStruct) UpdateArticle(id int, articleID string) error {
// 	return r.UpdateArticle(id, articleID)
// }

func (r *RepositoryStruct) PatchArticle(id int, articleID string, updates map[string]interface{}) error {
	start := time.Now()
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
	duration := time.Since(start)
	if err != nil {
		r.logger.Error().
			Str("method", "PatchArticle").
			Int("id", id).
			Interface("updates", updates).
			Dur("latency", duration).
			Err(err).
			Msg("failed to patch article")
		return r.translateError(err)
	}

	r.logger.Info().
		Str("method", "PatchArticle").
		Int("id", id).
		Interface("updates", updates).
		Dur("latency", duration).
		Msg("article patched successfully")
	return nil
}

func (r *RepositoryStruct) DeleteArticle(userID int, articleID string) error {
	start := time.Now()
	query := "UPDATE articles SET deleted_at = NOW() WHERE id = $1 AND user_id = $2"

	_, err := r.db.Exec(query, articleID, userID)
	duration := time.Since(start)
	if err != nil {
		r.logger.Error().
			Str("method", "DeleteArticle").
			Str("articleID", articleID).
			Dur("latency", duration).
			Err(err).
			Msg("failed to delete article")
		return r.translateError(err)
	}

	r.logger.Info().
		Str("method", "DeleteArticle").
		Str("articleID", articleID).
		Dur("latency", duration).
		Msg("article deleted successfully")

	// rows, err := res.RowsAffected()
	// if err != nil {
	// 	return err
	// }
	// if rows == 0 {
	// 	return r.tranlsateError(err)
	// }

	return nil
}
