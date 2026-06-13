package model

import "fmt"

type Tense string

const (
	SimplePresent Tense = "SIMPLE_PRESENT"
	ImperfectPast Tense = "IMPERFECT_PAST"
	SimplePast    Tense = "SIMPLE_PAST"
	PastPerfect   Tense = "PAST_PERFECT"
	SimpleFuture  Tense = "SIMPLE_FUTURE"
	Conditional   Tense = "CONDITIONAL"
)

var validTenses = map[Tense]bool{
	SimplePresent: true,
	ImperfectPast: true,
	SimplePast:    true,
	PastPerfect:   true,
	SimpleFuture:  true,
	Conditional:   true,
}

func (t Tense) IsValid() bool {
	return validTenses[t]
}

type Pronoun string

const (
	FirstPersonalSingular  Pronoun = "FIRST_PERSON_SINGULAR"
	SecondPersonalSingular Pronoun = "SECOND_PERSON_SINGULAR"
	ThirdPersonalSingular  Pronoun = "THIRD_PERSON_SINGULAR"
	FirstPersonalPlural    Pronoun = "FIRST_PERSON_PLULAR"
	SecondPersonalPlural   Pronoun = "SECOND_PERSON_PLULAR"
	ThirdPersonalPlural    Pronoun = "THIRD_PERSON_PLULAR"
)

var validPronouns = map[Pronoun]bool{
	FirstPersonalSingular:  true,
	SecondPersonalSingular: true,
	ThirdPersonalSingular:  true,
	FirstPersonalPlural:    true,
	SecondPersonalPlural:   true,
	ThirdPersonalPlural:    true,
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
