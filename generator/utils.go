package generator

import "errors"

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
