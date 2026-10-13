package generator

import (
	"errors"
	"strings"
)

// PronounText returns the lowercase Portuguese display form of a pronoun.
func PronounText(pronoun Pronoun) (string, error) {
	switch pronoun {
	case Eu:
		return "eu", nil
	case Tu:
		return "tu", nil
	case VoceEleEla:
		return "você/ele/ela", nil
	case Nos:
		return "nós", nil
	case VocesElesElas:
		return "vocês/eles/elas", nil
	default:
		return "", errors.New("Pronoun not implemented")
	}
}

// PronounFromText parses a Portuguese display form, ignoring case and surrounding whitespace.
func PronounFromText(text string) (Pronoun, error) {
	switch strings.ToLower(strings.TrimSpace(text)) {
	case "eu":
		return Eu, nil
	case "tu":
		return Tu, nil
	case "você/ele/ela":
		return VoceEleEla, nil
	case "nós":
		return Nos, nil
	case "vocês/eles/elas":
		return VocesElesElas, nil
	default:
		return "", errors.New("Pronoun not implemented")
	}
}

// TenseText returns the Portuguese display name of a tense.
func TenseText(tense Tense) (string, error) {
	switch tense {
	case PresenteDoIndicativo:
		return "Presente do Indicativo", nil
	case PreteritoPerfeitoDoIndicativo:
		return "Pretérito Perfeito", nil
	default:
		return "", errors.New("Tense not implemented")
	}
}

// TenseFromText parses a Portuguese display name, ignoring case and surrounding whitespace.
func TenseFromText(text string) (Tense, error) {
	switch strings.ToLower(strings.TrimSpace(text)) {
	case "presente":
		return PresenteDoIndicativo, nil
	case "pretérito perfeito":
		return PreteritoPerfeitoDoIndicativo, nil
	default:
		return "", errors.New("Tense not implemented")
	}
}
