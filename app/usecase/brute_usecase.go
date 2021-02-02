package usecase

import (
	"fmt"

	"github.com/bryutus/brute/app/domain/model"
	"github.com/bryutus/brute/app/domain/repository"
)

type bruteUseCaseImplement struct {
	AphorismRepository repository.AphorismRepository
}

type bruteUseCase interface {
	Exec(code string) (*model.Aphorism, error)
}

func NewBruteUseCaseImplement(aphorismRepository repository.AphorismRepository) bruteUseCase {
	return bruteUseCaseImplement{
		AphorismRepository: aphorismRepository,
	}
}

func (usecase bruteUseCaseImplement) Exec(code string) (*model.Aphorism, error) {
	brute, err := usecase.AphorismRepository.FindBy(code)
	if err != nil {
		return nil, err
	}

	if brute == nil {
		return nil, fmt.Errorf("record not found: language_code=%s", code)
	}

	return brute, nil
}
