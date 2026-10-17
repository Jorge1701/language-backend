package database

import "database/sql"

func mapError(err error) error {
	if err == nil {
		return err
	}
	switch err.Error() {
	case "sql: no rows in result set":
		return ErrNotFoundInDB
	default:
		return err
	}
}

func extractVerbRow(row *sql.Row) (VerbRow, error) {
	var verbRow VerbRow
	err := row.Scan(
		&verbRow.Id,
		&verbRow.Type,
		&verbRow.Infinitive,
		&verbRow.PresentParticiple,
		&verbRow.PastParticiple,
		&verbRow.PresenteDoIndicativoEu,
		&verbRow.PresenteDoIndicativoTu,
		&verbRow.PresenteDoIndicativoVoceEleEla,
		&verbRow.PresenteDoIndicativoNos,
		&verbRow.PresenteDoIndicativoVos,
		&verbRow.PresenteDoIndicativoVocesElesElas,
		&verbRow.PreteritoImperfeitoDoIndicativoEu,
		&verbRow.PreteritoImperfeitoDoIndicativoTu,
		&verbRow.PreteritoImperfeitoDoIndicativoVoceEleEla,
		&verbRow.PreteritoImperfeitoDoIndicativoNos,
		&verbRow.PreteritoImperfeitoDoIndicativoVos,
		&verbRow.PreteritoImperfeitoDoIndicativoVocesElesElas,
		&verbRow.PreteritoPerfeitoDoIndicativoEu,
		&verbRow.PreteritoPerfeitoDoIndicativoTu,
		&verbRow.PreteritoPerfeitoDoIndicativoVoceEleEla,
		&verbRow.PreteritoPerfeitoDoIndicativoNos,
		&verbRow.PreteritoPerfeitoDoIndicativoVos,
		&verbRow.PreteritoPerfeitoDoIndicativoVocesElesElas,
		&verbRow.PreteritoMaisQuePerfeitoDoIndicativoEu,
		&verbRow.PreteritoMaisQuePerfeitoDoIndicativoTu,
		&verbRow.PreteritoMaisQuePerfeitoDoIndicativoVoceEleEla,
		&verbRow.PreteritoMaisQuePerfeitoDoIndicativoNos,
		&verbRow.PreteritoMaisQuePerfeitoDoIndicativoVos,
		&verbRow.PreteritoMaisQuePerfeitoDoIndicativoVocesElesElas,
		&verbRow.FuturoDoPresenteDoIndicativoEu,
		&verbRow.FuturoDoPresenteDoIndicativoTu,
		&verbRow.FuturoDoPresenteDoIndicativoVoceEleEla,
		&verbRow.FuturoDoPresenteDoIndicativoNos,
		&verbRow.FuturoDoPresenteDoIndicativoVos,
		&verbRow.FuturoDoPresenteDoIndicativoVocesElesElas,
		&verbRow.FuturoDoPreteritoDoIndicativoEu,
		&verbRow.FuturoDoPreteritoDoIndicativoTu,
		&verbRow.FuturoDoPreteritoDoIndicativoVoceEleEla,
		&verbRow.FuturoDoPreteritoDoIndicativoNos,
		&verbRow.FuturoDoPreteritoDoIndicativoVos,
		&verbRow.FuturoDoPreteritoDoIndicativoVocesElesElas,
	)

	return verbRow, mapError(err)
}

func verbsProperties() string {
	return `
		id,
		type,
		infinitive,
		present_participle,
		past_participle,
		presente_do_indicativo_eu,
		presente_do_indicativo_tu,
		presente_do_indicativo_voce_ele_ela,
		presente_do_indicativo_nos,
		presente_do_indicativo_vos,
		presente_do_indicativo_voces_eles_elas,
		preterito_imperfeito_do_indicativo_eu,
		preterito_imperfeito_do_indicativo_tu,
		preterito_imperfeito_do_indicativo_voce_ele_ela,
		preterito_imperfeito_do_indicativo_nos,
		preterito_imperfeito_do_indicativo_vos,
		preterito_imperfeito_do_indicativo_voces_eles_elas,
		preterito_perfeito_do_indicativo_eu,
		preterito_perfeito_do_indicativo_tu,
		preterito_perfeito_do_indicativo_voce_ele_ela,
		preterito_perfeito_do_indicativo_nos,
		preterito_perfeito_do_indicativo_vos,
		preterito_perfeito_do_indicativo_voces_eles_elas,
		preterito_mais_que_perfeito_do_indicativo_eu,
		preterito_mais_que_perfeito_do_indicativo_tu,
		preterito_mais_que_perfeito_do_indicativo_voce_ele_ela,
		preterito_mais_que_perfeito_do_indicativo_nos,
		preterito_mais_que_perfeito_do_indicativo_vos,
		preterito_mais_que_perfeito_do_indicativo_voces_eles_elas,
		futuro_do_presente_do_indicativo_eu,
		futuro_do_presente_do_indicativo_tu,
		futuro_do_presente_do_indicativo_voce_ele_ela,
		futuro_do_presente_do_indicativo_nos,
		futuro_do_presente_do_indicativo_vos,
		futuro_do_presente_do_indicativo_voces_eles_elas,
		futuro_do_preterito_do_indicativo_eu,
		futuro_do_preterito_do_indicativo_tu,
		futuro_do_preterito_do_indicativo_voce_ele_ela,
		futuro_do_preterito_do_indicativo_nos,
		futuro_do_preterito_do_indicativo_vos,
		futuro_do_preterito_do_indicativo_voces_eles_elas
	`
}
