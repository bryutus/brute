package model

// Aphorism Domain Model
type Aphorism struct {
	Phrase       string
	LanguageCode string `form:"language_code" binding:"omitempty,len=2"`
}
