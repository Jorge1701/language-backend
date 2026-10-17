package model

import "fmt"

type Tense string

const (
	PresenteDoIndicativo                 Tense = "PresenteDoIndicativo"
	PreteritoImperfeitoDoIndicativo      Tense = "PreteritoImperfeitoDoIndicativo"
	PreteritoPerfeitoDoIndicativo        Tense = "PreteritoPerfeitoDoIndicativo"
	PreteritoMaisQuePerfeitoDoIndicativo Tense = "PreteritoMaisQuePerfeitoDoIndicativo"
	FuturoDoPresenteDoIndicativo         Tense = "FuturoDoPresenteDoIndicativo"
	FuturoDoPreteritoDoIndicativo        Tense = "FuturoDoPreteritoDoIndicativo"
)

var validTenses = map[Tense]bool{
	PresenteDoIndicativo:                 true,
	PreteritoImperfeitoDoIndicativo:      true,
	PreteritoPerfeitoDoIndicativo:        true,
	PreteritoMaisQuePerfeitoDoIndicativo: true,
	FuturoDoPresenteDoIndicativo:         true,
	FuturoDoPreteritoDoIndicativo:        true,
}

func (t Tense) IsValid() bool {
	return validTenses[t]
}

type Pronoun string

const (
	Eu            Pronoun = "Eu"
	Tu            Pronoun = "Tu"
	VoceEleEla    Pronoun = "VoceEleEla"
	Nos           Pronoun = "Nos"
	Vos           Pronoun = "Vos"
	VocesElesElas Pronoun = "VocesElesElas"
)

var validPronouns = map[Pronoun]bool{
	Eu:            true,
	Tu:            true,
	VoceEleEla:    true,
	Nos:           true,
	Vos:           true,
	VocesElesElas: true,
}

func (p Pronoun) IsValid() bool {
	return validPronouns[p]
}

type ExerciseExample struct {
	Pt string `json:"pt,omitempty"`
	Es string `json:"es,omitempty"`
}

func (e *ExerciseExample) String() string {
	return fmt.Sprintf("{%s %s}", e.Pt, e.Es)
}

type VerbExerciseInfo struct {
	Tense       Tense            `json:"tense"`
	Pronoun     string           `json:"pronoun"`
	Verb        string           `json:"verb"`
	Conjugation string           `json:"conjugation"`
	Example     *ExerciseExample `json:"example,omitempty"`
}

type VerbInfo struct {
	Type             string           `json:"type"`
	Infinitive       string           `json:"infinitive,omitempty"`
	PresentPaticiple string           `json:"present_participle,omitempty"`
	PastParticiple   string           `json:"past_participle,omitempty"`
	SimplePresent    VerbConjugations `json:"simple_present"`
	ImperfectPast    VerbConjugations `json:"imperfect_past"`
	SimplePast       VerbConjugations `json:"simple_past"`
	PastPerfect      VerbConjugations `json:"past_perfect"`
	SimpleFuture     VerbConjugations `json:"simple_future"`
	Conditional      VerbConjugations `json:"conditional"`
}

type VerbConjugations struct {
	FirstPersonSingular  string `json:"fps"`
	SecondPersonSingular string `json:"sps"`
	ThirdPersonSingular  string `json:"tps"`
	FirstPersonPlural    string `json:"fpp"`
	SecondPersonPlural   string `json:"spp"`
	ThirdPersonPlural    string `json:"tpp"`
}
