// Package generator conjugates Portuguese verbs using YAML overrides and regular rules.
package generator

import (
	"embed"
	"errors"
	"language-backend/model"
	"strings"
)

var (
	ErrInvalidVerb            = errors.New("invalid verb infinitive")
	ErrInvalidTense           = errors.New("invalid tense")
	ErrInvalidPronoun         = errors.New("invalid pronoun")
	ErrNotImplemented         = errors.New("conjugation not implemented")
	ErrConjugationUnavailable = errors.New("conjugation unavailable")
)

//go:embed data/*.yml
var verbFiles embed.FS

type Generator struct{}

func New() (*Generator, error) {
	return &Generator{}, nil
}

func (g *Generator) Conjugate(verb string, tense model.Tense, pronoun model.Pronoun) (string, error) {
	return "TODO", nil
}

func validateInfinitiveForm(verb string) bool {
	return strings.HasSuffix(verb, "ar") || strings.HasSuffix(verb, "er") || strings.HasSuffix(verb, "ir")
}
