package generator

type Tense string

const (
	PresenteDoIndicativo          Tense = "PRESENTE_DO_INDICATIVO"
	PreteritoPerfeitoDoIndicativo Tense = "PRETERITO_PERFEITO_DO_INDICATIVO"
)

var validTenses = map[Tense]bool{
	PresenteDoIndicativo: true,
}

func (t Tense) IsValid() bool {
	return validTenses[t]
}

type Pronoun string

const (
	Eu            Pronoun = "EU"
	VoceEleEla    Pronoun = "VOCE_ELE_ELA"
	Nos           Pronoun = "NOS"
	VocesElesElas Pronoun = "VOCES_ELES_ELAS"
	Tu            Pronoun = "TU"
)

var validPronouns = map[Pronoun]bool{
	Eu:            true,
	VoceEleEla:    true,
	Nos:           true,
	VocesElesElas: true,
	Tu:            true,
}

func (p Pronoun) IsValid() bool {
	return validPronouns[p]
}
