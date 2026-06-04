package endpoints

import (
	"language-backend/database"
	"language-backend/model"
)

func mapVerbRow(verbRow database.VerbRow) model.VerbInfo {
	return model.VerbInfo{
		Type:             verbRow.Type,
		Infinitive:       verbRow.Infinitive,
		PresentPaticiple: verbRow.PresentParticiple,
		PastParticiple:   verbRow.PastParticiple,
		SimplePresent: model.VerbConjugations{
			FirstPersonSingular:  verbRow.SimplePresentFPS,
			SecondPersonSingular: verbRow.SimplePresentSPS,
			ThirdPersonSingular:  verbRow.SimplePresentTPS,
			FirstPersonPlural:    verbRow.SimplePresentFPP,
			SecondPersonPlural:   verbRow.SimplePresentSPP,
			ThirdPersonPlural:    verbRow.SimplePresentTPP,
		},
		ImperfectPast: model.VerbConjugations{
			FirstPersonSingular:  verbRow.ImperfectPastFPS,
			SecondPersonSingular: verbRow.ImperfectPastSPS,
			ThirdPersonSingular:  verbRow.ImperfectPastTPS,
			FirstPersonPlural:    verbRow.ImperfectPastFPP,
			SecondPersonPlural:   verbRow.ImperfectPastSPP,
			ThirdPersonPlural:    verbRow.ImperfectPastTPP,
		},
		SimplePast: model.VerbConjugations{
			FirstPersonSingular:  verbRow.SimplePastFPS,
			SecondPersonSingular: verbRow.SimplePastSPS,
			ThirdPersonSingular:  verbRow.SimplePastTPS,
			FirstPersonPlural:    verbRow.SimplePastFPP,
			SecondPersonPlural:   verbRow.SimplePastSPP,
			ThirdPersonPlural:    verbRow.SimplePastTPP,
		},
		PastPerfect: model.VerbConjugations{
			FirstPersonSingular:  verbRow.PastPerfectFPS,
			SecondPersonSingular: verbRow.PastPerfectSPS,
			ThirdPersonSingular:  verbRow.PastPerfectTPS,
			FirstPersonPlural:    verbRow.PastPerfectFPP,
			SecondPersonPlural:   verbRow.PastPerfectSPP,
			ThirdPersonPlural:    verbRow.PastPerfectTPP,
		},
		SimpleFuture: model.VerbConjugations{
			FirstPersonSingular:  verbRow.SimpleFutureFPS,
			SecondPersonSingular: verbRow.SimpleFutureSPS,
			ThirdPersonSingular:  verbRow.SimpleFutureTPS,
			FirstPersonPlural:    verbRow.SimpleFutureFPP,
			SecondPersonPlural:   verbRow.SimpleFutureSPP,
			ThirdPersonPlural:    verbRow.SimpleFutureTPP,
		},
		Conditional: model.VerbConjugations{
			FirstPersonSingular:  verbRow.ConditionalFPS,
			SecondPersonSingular: verbRow.ConditionalSPS,
			ThirdPersonSingular:  verbRow.ConditionalTPS,
			FirstPersonPlural:    verbRow.ConditionalFPP,
			SecondPersonPlural:   verbRow.ConditionalSPP,
			ThirdPersonPlural:    verbRow.ConditionalTPP,
		},
	}
}
