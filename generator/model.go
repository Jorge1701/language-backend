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

var validTenses = func() map[Tense]bool {
	valid := make(map[Tense]bool, len(AllTenses))
	for _, tense := range AllTenses {
		valid[tense] = true
	}
	return valid
}()

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

var AllPronouns = []Pronoun{
	Eu,
	VoceEleEla,
	Nos,
	VocesElesElas,
	Tu,
}

var validPronouns = func() map[Pronoun]bool {
	valid := make(map[Pronoun]bool, len(AllPronouns))
	for _, pronoun := range AllPronouns {
		valid[pronoun] = true
	}
	return valid
}()

func (p Pronoun) IsValid() bool {
	return validPronouns[p]
}
