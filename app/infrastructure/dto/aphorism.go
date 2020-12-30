package dto

import "github.com/bryutus/brute/app/domain/model"

type Aphorism struct {
	ID           uint
	Phrase       string
	LanguageCode string
}

func (a *Aphorism) ConvertToModel() *model.Brute {
	return &model.Brute{
		Phrase:       a.Phrase,
		LanguageCode: a.LanguageCode,
	}
}
