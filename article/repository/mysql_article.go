package repository

import (
	"database/sql"
	"github.com/mnuma/my-go-clean-arch/article"
	"github.com/mnuma/my-go-clean-arch/author"
)

type mysqlArticleRepository struct {
	Conn *sql.DB
}

func NewMysqlArticleRepository(Conn *sql.DB) ArticleRepository {
	return &mysqlArticleRepository{Conn}
}

func (m *mysqlArticleRepository) GetByID(id int64) (*article.Article, error) {
	query := `SELECT id,title,content, author_id, updated_at, created_at
  						FROM article WHERE ID = ?`

	list, err := m.fetch(query, id)
	if err != nil {
		return nil, err
	}

	a := &article.Article{}
	if len(list) > 0 {
		a = list[0]
	} else {
		return nil, article.NOT_FOUND_ERROR
	}

	return a, nil
}

func (m *mysqlArticleRepository) fetch(query string, args ...interface{}) ([]*article.Article, error) {

	rows, err := m.Conn.Query(query, args...)

	if err != nil {

		return nil, article.INTERNAL_SERVER_ERROR
	}
	defer rows.Close()
	result := make([]*article.Article, 0)
	for rows.Next() {
		t := new(article.Article)
		authorID := int64(0)
		err = rows.Scan(
			&t.ID,
			&t.Title,
			&t.Content,
			&authorID,
			&t.UpdatedAt,
			&t.CreatedAt,
		)

		if err != nil {

			return nil, article.INTERNAL_SERVER_ERROR
		}
		t.Author = author.Author{
			ID: authorID,
		}
		result = append(result, t)
	}

	return result, nil
}