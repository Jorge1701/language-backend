package generator

type Tense string

const (
	PresenteDoIndicativo          Tense = "PRESENTE_DO_INDICATIVO"
	PreteritoPerfeitoDoIndicativo Tense = "PRETERITO_PERFEITO_DO_INDICATIVO"
)

var AllTenses = []Tense{
	PresenteDoIndicativo,
	PreteritoPerfeitoDoIndicativo,
}

type Pronoun string

const (
	Eu            Pronoun = "EU"
	VoceEleEla    Pronoun = "VOCE_ELE_ELA"
	Nos           Pronoun = "NOS"
	VocesElesElas Pronoun = "VOCES_ELES_ELAS"
	Tu            Pronoun = "TU"
)
