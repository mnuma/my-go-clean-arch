package usecase

import (
	"github.com/mnuma/my-go-clean-arch/article"
	"github.com/mnuma/my-go-clean-arch/article/repository"

	authorRepository "github.com/mnuma/my-go-clean-arch/author/repository"
)

type ArticleUsecase interface {
	GetByID(id int64) (*article.Article, error)
}

type articleUsecase struct {
	articleRepos repository.ArticleRepository
	authorRepo   authorRepository.AuthorRepository
}

func NewArticleUsecase(a repository.ArticleRepository, ar authorRepository.AuthorRepository) ArticleUsecase {
	return &articleUsecase{
		articleRepos: a,
		authorRepo:   ar,
	}
}

func (a *articleUsecase) GetByID(id int64) (*article.Article, error) {

	res, err := a.articleRepos.GetByID(id)
	if err != nil {
		return nil, err
	}

	resAuthor, err := a.authorRepo.GetByID(res.Author.ID)
	if err != nil {
		return nil, err
	}
	res.Author = *resAuthor
	return res, nil
}