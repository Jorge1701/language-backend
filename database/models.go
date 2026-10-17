package database

type VerbRow struct {
	Id                                                int    `json:"id"`
	Type                                              string `json:"type"`
	Infinitive                                        string `json:"infinitive"`
	PresentParticiple                                 string `json:"present_participle"`
	PastParticiple                                    string `json:"past_participle"`
	PresenteDoIndicativoEu                            string `json:"presente_do_indicativo_eu"`
	PresenteDoIndicativoTu                            string `json:"presente_do_indicativo_tu"`
	PresenteDoIndicativoVoceEleEla                    string `json:"presente_do_indicativo_voce_ele_ela"`
	PresenteDoIndicativoNos                           string `json:"presente_do_indicativo_nos"`
	PresenteDoIndicativoVos                           string `json:"presente_do_indicativo_vos"`
	PresenteDoIndicativoVocesElesElas                 string `json:"presente_do_indicativo_voces_eles_elas"`
	PreteritoImperfeitoDoIndicativoEu                 string `json:"preterito_imperfeito_do_indicativo_eu"`
	PreteritoImperfeitoDoIndicativoTu                 string `json:"preterito_imperfeito_do_indicativo_tu"`
	PreteritoImperfeitoDoIndicativoVoceEleEla         string `json:"preterito_imperfeito_do_indicativo_voce_ele_ela"`
	PreteritoImperfeitoDoIndicativoNos                string `json:"preterito_imperfeito_do_indicativo_nos"`
	PreteritoImperfeitoDoIndicativoVos                string `json:"preterito_imperfeito_do_indicativo_vos"`
	PreteritoImperfeitoDoIndicativoVocesElesElas      string `json:"preterito_imperfeito_do_indicativo_voces_eles_elas"`
	PreteritoPerfeitoDoIndicativoEu                   string `json:"preterito_perfeito_do_indicativo_eu"`
	PreteritoPerfeitoDoIndicativoTu                   string `json:"preterito_perfeito_do_indicativo_tu"`
	PreteritoPerfeitoDoIndicativoVoceEleEla           string `json:"preterito_perfeito_do_indicativo_voce_ele_ela"`
	PreteritoPerfeitoDoIndicativoNos                  string `json:"preterito_perfeito_do_indicativo_nos"`
	PreteritoPerfeitoDoIndicativoVos                  string `json:"preterito_perfeito_do_indicativo_vos"`
	PreteritoPerfeitoDoIndicativoVocesElesElas        string `json:"preterito_perfeito_do_indicativo_voces_eles_elas"`
	PreteritoMaisQuePerfeitoDoIndicativoEu            string `json:"preterito_mais_que_perfeito_do_indicativo_eu"`
	PreteritoMaisQuePerfeitoDoIndicativoTu            string `json:"preterito_mais_que_perfeito_do_indicativo_tu"`
	PreteritoMaisQuePerfeitoDoIndicativoVoceEleEla    string `json:"preterito_mais_que_perfeito_do_indicativo_voce_ele_ela"`
	PreteritoMaisQuePerfeitoDoIndicativoNos           string `json:"preterito_mais_que_perfeito_do_indicativo_nos"`
	PreteritoMaisQuePerfeitoDoIndicativoVos           string `json:"preterito_mais_que_perfeito_do_indicativo_vos"`
	PreteritoMaisQuePerfeitoDoIndicativoVocesElesElas string `json:"preterito_mais_que_perfeito_do_indicativo_voces_eles_elas"`
	FuturoDoPresenteDoIndicativoEu                    string `json:"futuro_do_presente_do_indicativo_eu"`
	FuturoDoPresenteDoIndicativoTu                    string `json:"futuro_do_presente_do_indicativo_tu"`
	FuturoDoPresenteDoIndicativoVoceEleEla            string `json:"futuro_do_presente_do_indicativo_voce_ele_ela"`
	FuturoDoPresenteDoIndicativoNos                   string `json:"futuro_do_presente_do_indicativo_nos"`
	FuturoDoPresenteDoIndicativoVos                   string `json:"futuro_do_presente_do_indicativo_vos"`
	FuturoDoPresenteDoIndicativoVocesElesElas         string `json:"futuro_do_presente_do_indicativo_voces_eles_elas"`
	FuturoDoPreteritoDoIndicativoEu                   string `json:"futuro_do_preterito_do_indicativo_eu"`
	FuturoDoPreteritoDoIndicativoTu                   string `json:"futuro_do_preterito_do_indicativo_tu"`
	FuturoDoPreteritoDoIndicativoVoceEleEla           string `json:"futuro_do_preterito_do_indicativo_voce_ele_ela"`
	FuturoDoPreteritoDoIndicativoNos                  string `json:"futuro_do_preterito_do_indicativo_nos"`
	FuturoDoPreteritoDoIndicativoVos                  string `json:"futuro_do_preterito_do_indicativo_vos"`
	FuturoDoPreteritoDoIndicativoVocesElesElas        string `json:"futuro_do_preterito_do_indicativo_voces_eles_elas"`
}

type ExampleRow struct {
	Pt string
	Es string
}

type User struct {
	Id        int
	Username  string
	Password  string
	CreatedAt string
	UpdatedAt string
}
