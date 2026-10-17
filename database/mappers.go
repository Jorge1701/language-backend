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
		&verbRow.ImperfectPastFPS,
		&verbRow.ImperfectPastSPS,
		&verbRow.ImperfectPastTPS,
		&verbRow.ImperfectPastFPP,
		&verbRow.ImperfectPastSPP,
		&verbRow.ImperfectPastTPP,
		&verbRow.SimplePastFPS,
		&verbRow.SimplePastSPS,
		&verbRow.SimplePastTPS,
		&verbRow.SimplePastFPP,
		&verbRow.SimplePastSPP,
		&verbRow.SimplePastTPP,
		&verbRow.PastPerfectFPS,
		&verbRow.PastPerfectSPS,
		&verbRow.PastPerfectTPS,
		&verbRow.PastPerfectFPP,
		&verbRow.PastPerfectSPP,
		&verbRow.PastPerfectTPP,
		&verbRow.SimpleFutureFPS,
		&verbRow.SimpleFutureSPS,
		&verbRow.SimpleFutureTPS,
		&verbRow.SimpleFutureFPP,
		&verbRow.SimpleFutureSPP,
		&verbRow.SimpleFutureTPP,
		&verbRow.ConditionalFPS,
		&verbRow.ConditionalSPS,
		&verbRow.ConditionalTPS,
		&verbRow.ConditionalFPP,
		&verbRow.ConditionalSPP,
		&verbRow.ConditionalTPP,
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
