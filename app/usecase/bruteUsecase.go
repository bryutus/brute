package usecase

import (
	"github.com/bryutus/brute/app/domain/model"
	"github.com/bryutus/brute/app/domain/repository"
)

type bruteUseCaseImplement struct {
	BruteRepository repository.BruteRepository
}

type bruteUseCase interface {
	Exec(code string) (*model.Brute, error)
}

func NewBruteUseCaseImplement(bruteRepository repository.BruteRepository) bruteUseCase {
	return bruteUseCaseImplement{
		BruteRepository: bruteRepository,
	}
}

func (usecase bruteUseCaseImplement) Exec(code string) (*model.Brute, error) {
	brute, err := usecase.BruteRepository.FindBy(code)
	if err != nil {
		return nil, err
	}

	return brute, nil
}
