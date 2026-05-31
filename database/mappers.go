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
		&verbRow.SimplePresentFPS,
		&verbRow.SimplePresentSPS,
		&verbRow.SimplePresentTPS,
		&verbRow.SimplePresentFPP,
		&verbRow.SimplePresentSPP,
		&verbRow.SimplePresentTPP,
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
		simple_present_fps,
		simple_present_sps,
		simple_present_tps,
		simple_present_fpp,
		simple_present_spp,
		simple_present_tpp,
		imperfect_past_fps,
		imperfect_past_sps,
		imperfect_past_tps,
		imperfect_past_fpp,
		imperfect_past_spp,
		imperfect_past_tpp,
		simple_past_fps,
		simple_past_sps,
		simple_past_tps,
		simple_past_fpp,
		simple_past_spp,
		simple_past_tpp,
		past_perfect_fps,
		past_perfect_sps,
		past_perfect_tps,
		past_perfect_fpp,
		past_perfect_spp,
		past_perfect_tpp,
		simple_future_fps,
		simple_future_sps,
		simple_future_tps,
		simple_future_fpp,
		simple_future_spp,
		simple_future_tpp,
		conditional_fps,
		conditional_sps,
		conditional_tps,
		conditional_fpp,
		conditional_spp,
		conditional_tpp
	`
}
