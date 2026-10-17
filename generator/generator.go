// Package generator conjugates Portuguese verbs using YAML overrides and regular rules.
package generator

import (
	"errors"
	"strings"
)

var (
	ErrInvalidVerb            = errors.New("invalid verb infinitive")
	ErrInvalidPronoun         = errors.New("invalid pronoun")
	ErrTenseNotImplemented    = errors.New("tense not pronoun")
	ErrConjugationUnavailable = errors.New("conjugation unavailable")
)

type Generator struct {
	tenseData map[Tense]TenseData
}

func New() (*Generator, error) {
	tenseData := make(map[Tense]TenseData, len(AllTenses))

	for _, tense := range AllTenses {
		tenseFile := strings.ToLower(string(tense))
		td, err := LoadTenseData("data/" + tenseFile + ".yml")
		if err != nil {
			return &Generator{}, err
		}

		tenseData[tense] = *td
	}

	return &Generator{
		tenseData: tenseData,
	}, nil
}

func (g *Generator) Conjugate(inputVerb string, tense Tense, pronoun Pronoun) (string, error) {
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

	if conjugation, found := tenseData.Irregular[verb][pronounText]; found {
		if conjugation == "" {
			return "", ErrConjugationUnavailable
		}
		return conjugation, nil
	}

	base, infinitiveEnding := extractBase(verb)
	conjugationEnding, found := tenseData.RegularEndings[infinitiveEnding][pronounText]
	if !found {
		return "", ErrConjugationUnavailable
	}

	return base + conjugationEnding, nil
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
