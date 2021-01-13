package usecase

import (
	"github.com/bryutus/brute/app/domain/model"
	"github.com/bryutus/brute/app/domain/repository"
)

type saveBruteUseCaseImplement struct {
	AphorismRepository repository.AphorismRepository
}

type saveBruteUseCase interface {
	Exec(code string, phrase string) (*model.Aphorism, error)
}

func NewSaveBruteUseCaseImplement(aphorismRepository repository.AphorismRepository) saveBruteUseCase {
	return saveBruteUseCaseImplement{
		AphorismRepository: aphorismRepository,
	}
}

func (usecase saveBruteUseCaseImplement) Exec(code string, phrase string) (*model.Aphorism, error) {
	brute, err := usecase.AphorismRepository.Save(code, phrase)
	if err != nil {
		return nil, err
	}
	return brute, nil
}
