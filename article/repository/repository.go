package repository

import "github.com/mnuma/my-go-clean-arch/article"

type ArticleRepository interface {
	GetByID(id int64) (*article.Article, error)
}