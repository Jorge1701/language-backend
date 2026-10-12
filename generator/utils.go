package generator

import (
	"errors"
	"language-backend/model"
	"strings"
)

// PronounText returns the lowercase Portuguese display form of a pronoun.
func PronounText(pronoun model.Pronoun) (string, error) {
	switch pronoun {
	case model.FirstPersonalSingular:
		return "eu", nil
	case model.SecondPersonalSingular:
		return "tu", nil
	case model.ThirdPersonalSingular:
		return "você/ele/ela", nil
	case model.FirstPersonalPlural:
		return "nós", nil
	case model.SecondPersonalPlural:
		return "vós", nil
	case model.ThirdPersonalPlural:
		return "vocês/eles/elas", nil
	default:
		return "", errors.New("Pronoun not implemented")
	}
}

// PronounFromText parses a Portuguese display form, ignoring case and surrounding whitespace.
func PronounFromText(text string) (model.Pronoun, error) {
	switch strings.ToLower(strings.TrimSpace(text)) {
	case "eu":
		return model.FirstPersonalSingular, nil
	case "tu":
		return model.SecondPersonalSingular, nil
	case "você/ele/ela":
		return model.ThirdPersonalSingular, nil
	case "nós":
		return model.FirstPersonalPlural, nil
	case "vós":
		return model.SecondPersonalPlural, nil
	case "vocês/eles/elas":
		return model.ThirdPersonalPlural, nil
	default:
		return "", errors.New("Pronoun not implemented")
	}
}

// TenseText returns the Portuguese display name of a tense.
func TenseText(tense model.Tense) (string, error) {
	switch tense {
	case model.SimplePresent:
		return "Presente do Indicativo", nil
	case model.ImperfectPast:
		return "Pretérito Imperfeito", nil
	case model.SimplePast:
		return "Pretérito Perfeito", nil
	case model.PastPerfect:
		return "Pretérito Mais-que-perfeito", nil
	case model.SimpleFuture:
		return "Futuro do Presente", nil
	case model.Conditional:
		return "Futuro do Pretérito", nil
	default:
		return "", errors.New("Tense not implemented")
	}
}

// TenseFromText parses a Portuguese display name, ignoring case and surrounding whitespace.
func TenseFromText(text string) (model.Tense, error) {
	switch strings.ToLower(strings.TrimSpace(text)) {
	case "presente":
		return model.SimplePresent, nil
	case "pretérito imperfeito":
		return model.ImperfectPast, nil
	case "pretérito perfeito":
		return model.SimplePast, nil
	case "pretérito mais-que-perfeito":
		return model.PastPerfect, nil
	case "futuro do presente":
		return model.SimpleFuture, nil
	case "futuro do pretérito":
		return model.Conditional, nil
	default:
		return "", errors.New("Tense not implemented")
	}
}
