package repository

import (
	"github.com/bryutus/brute/app/domain/model"
)

type AphorismRepository interface {
	FindBy(code string) (*model.Aphorism, error)
	Save(code string, phrase string) (*model.Aphorism, error)
}
