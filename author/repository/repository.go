package repository

import "github.com/mnuma/my-go-clean-arch/author"

type AuthorRepository interface {
	GetByID(id int64) (*author.Author, error)
}