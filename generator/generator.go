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
	ErrTenseNotImplemented    = errors.New("tense not pronoun")
	ErrNotImplemented         = errors.New("conjugation not implemented")
	ErrConjugationUnavailable = errors.New("conjugation unavailable")
)

//go:embed data/*.yml
var verbFiles embed.FS

type Generator struct {
	tenseData map[model.Tense]TenseData
}

func New() (*Generator, error) {
	simplePresentTenseData, err := LoadTenseData("data/simple_present.yml")
	if err != nil {
		return &Generator{}, err
	}

	return &Generator{
		tenseData: map[model.Tense]TenseData{
			model.SimplePresent: *simplePresentTenseData,
		},
	}, nil
}

func (g *Generator) Conjugate(inputVerb string, tense model.Tense, pronoun model.Pronoun) (string, error) {
	verb := strings.ToLower(strings.TrimSpace(inputVerb))

	if !validateInfinitiveForm(verb) {
		return "", ErrInvalidVerb
	}

	pronounText, err := PronounText(pronoun)
	if err != nil {
		return "", ErrInvalidPronoun
	}

	tenseData, hasTenseData := g.tenseData[tense]
	if !hasTenseData {
		return "", ErrTenseNotImplemented
	}

	irregularData, isIrregular := tenseData.Irregular[verb]
	if isIrregular {
		return irregularData[pronounText], nil
	} else {
		base, infinitiveEnding := extractBase(verb)

		conjugationEnding, found := tenseData.RegularEndings[infinitiveEnding][pronounText]
		if !found {
			return "", ErrConjugationUnavailable
		}

		return base + conjugationEnding, nil
	}
}

func validateInfinitiveForm(verb string) bool {
	return strings.HasSuffix(verb, "ar") || strings.HasSuffix(verb, "er") || strings.HasSuffix(verb, "ir")
}

func extractBase(infinitiveVerb string) (string, string) {
	letters := []rune(infinitiveVerb)
	if len(letters) <= 2 {
		return "", infinitiveVerb
	}
	return string(letters[:len(letters)-2]), string(letters[len(letters)-2:])
}
