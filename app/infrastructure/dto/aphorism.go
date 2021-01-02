package dto

import "github.com/bryutus/brute/app/domain/model"

type Aphorism struct {
	ID           uint
	Phrase       string
	LanguageCode string
}

func (a *Aphorism) ConvertToModel() *model.Aphorism {
	return &model.Aphorism{
		Phrase:       a.Phrase,
		LanguageCode: a.LanguageCode,
	}
}
