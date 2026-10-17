package verbs

import (
	"errors"
	"fmt"
	"language-backend/database"
	"language-backend/model"
	"strings"
)

var (
	ErrInvalidInfinitiveVerb = errors.New("VerbConjugator: Invalid infinitive verb")
)

type VerbConjugator struct {
	db *database.Database
}

func New(db *database.Database) (*VerbConjugator, error) {
	return &VerbConjugator{
		db: db,
	}, nil
}

func (vc *VerbConjugator) Conjugate(inputVerb string, tense model.Tense, pronoun model.Pronoun) (string, error) {
	verb := strings.ToLower(strings.TrimSpace(inputVerb))

	if !validateInfinitiveForm(verb) {
		return "", ErrInvalidInfinitiveVerb
	}

	conjugation, err := vc.db.FindConjugation(tense, pronoun, verb)

	if err != nil {
		return "", fmt.Errorf("VerbConjugator: Error finding conjugation: %w", err)
	}

	return conjugation, nil
}

func validateInfinitiveForm(verb string) bool {
	return strings.HasSuffix(verb, "ar") || strings.HasSuffix(verb, "er") || strings.HasSuffix(verb, "ir")
}
