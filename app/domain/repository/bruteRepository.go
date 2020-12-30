package repository

import (
	"github.com/bryutus/brute/app/domain/model"
)

type BruteRepository interface {
	FindBy() (*model.Brute, error)
}
