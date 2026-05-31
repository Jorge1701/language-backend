package endpoints

import "language-backend/database"

func mapVerbRow(verbRow database.VerbRow) VerbInfo {
	return VerbInfo{
		Type:             verbRow.Type,
		Infinitive:       verbRow.Infinitive,
		PresentPaticiple: verbRow.PresentParticiple,
		PastParticiple:   verbRow.PastParticiple,
		SimplePresent: VerbConjugations{
			FirstPersonSingular:  verbRow.SimplePresentFPS,
			SecondPersonSingular: verbRow.SimplePresentSPS,
			ThirdPersonSingular:  verbRow.SimplePresentTPS,
			FirstPersonPlural:    verbRow.SimplePresentFPP,
			SecondPersonPlural:   verbRow.SimplePresentSPP,
			ThirdPersonPlural:    verbRow.SimplePresentTPP,
		},
		ImperfectPast: VerbConjugations{
			FirstPersonSingular:  verbRow.ImperfectPastFPS,
			SecondPersonSingular: verbRow.ImperfectPastSPS,
			ThirdPersonSingular:  verbRow.ImperfectPastTPS,
			FirstPersonPlural:    verbRow.ImperfectPastFPP,
			SecondPersonPlural:   verbRow.ImperfectPastSPP,
			ThirdPersonPlural:    verbRow.ImperfectPastTPP,
		},
		SimplePast: VerbConjugations{
			FirstPersonSingular:  verbRow.SimplePastFPS,
			SecondPersonSingular: verbRow.SimplePastSPS,
			ThirdPersonSingular:  verbRow.SimplePastTPS,
			FirstPersonPlural:    verbRow.SimplePastFPP,
			SecondPersonPlural:   verbRow.SimplePastSPP,
			ThirdPersonPlural:    verbRow.SimplePastTPP,
		},
		PastPerfect: VerbConjugations{
			FirstPersonSingular:  verbRow.PastPerfectFPS,
			SecondPersonSingular: verbRow.PastPerfectSPS,
			ThirdPersonSingular:  verbRow.PastPerfectTPS,
			FirstPersonPlural:    verbRow.PastPerfectFPP,
			SecondPersonPlural:   verbRow.PastPerfectSPP,
			ThirdPersonPlural:    verbRow.PastPerfectTPP,
		},
		SimpleFuture: VerbConjugations{
			FirstPersonSingular:  verbRow.SimpleFutureFPS,
			SecondPersonSingular: verbRow.SimpleFutureSPS,
			ThirdPersonSingular:  verbRow.SimpleFutureTPS,
			FirstPersonPlural:    verbRow.SimpleFutureFPP,
			SecondPersonPlural:   verbRow.SimpleFutureSPP,
			ThirdPersonPlural:    verbRow.SimpleFutureTPP,
		},
		Conditional: VerbConjugations{
			FirstPersonSingular:  verbRow.ConditionalFPS,
			SecondPersonSingular: verbRow.ConditionalSPS,
			ThirdPersonSingular:  verbRow.ConditionalTPS,
			FirstPersonPlural:    verbRow.ConditionalFPP,
			SecondPersonPlural:   verbRow.ConditionalSPP,
			ThirdPersonPlural:    verbRow.ConditionalTPP,
		},
	}
}
