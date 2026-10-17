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
		PresenteDoIndicativo: model.VerbConjugations{
			Eu:            verbRow.PresenteDoIndicativoEu,
			Tu:            verbRow.PresenteDoIndicativoTu,
			VoceEleEla:    verbRow.PresenteDoIndicativoVoceEleEla,
			Nos:           verbRow.PresenteDoIndicativoNos,
			Vos:           verbRow.PresenteDoIndicativoVos,
			VocesElesElas: verbRow.PresenteDoIndicativoVocesElesElas,
		},
		PreteritoImperfeitoDoIndicativo: model.VerbConjugations{
			Eu:            verbRow.PreteritoImperfeitoDoIndicativoEu,
			Tu:            verbRow.PreteritoImperfeitoDoIndicativoTu,
			VoceEleEla:    verbRow.PreteritoImperfeitoDoIndicativoVoceEleEla,
			Nos:           verbRow.PreteritoImperfeitoDoIndicativoNos,
			Vos:           verbRow.PreteritoImperfeitoDoIndicativoVos,
			VocesElesElas: verbRow.PreteritoImperfeitoDoIndicativoVocesElesElas,
		},
		PreteritoPerfeitoDoIndicativo: model.VerbConjugations{
			Eu:            verbRow.PreteritoPerfeitoDoIndicativoEu,
			Tu:            verbRow.PreteritoPerfeitoDoIndicativoTu,
			VoceEleEla:    verbRow.PreteritoPerfeitoDoIndicativoVoceEleEla,
			Nos:           verbRow.PreteritoPerfeitoDoIndicativoNos,
			Vos:           verbRow.PreteritoPerfeitoDoIndicativoVos,
			VocesElesElas: verbRow.PreteritoPerfeitoDoIndicativoVocesElesElas,
		},
		PreteritoMaisQuePerfeitoDoIndicativo: model.VerbConjugations{
			Eu:            verbRow.PreteritoMaisQuePerfeitoDoIndicativoEu,
			Tu:            verbRow.PreteritoMaisQuePerfeitoDoIndicativoTu,
			VoceEleEla:    verbRow.PreteritoMaisQuePerfeitoDoIndicativoVoceEleEla,
			Nos:           verbRow.PreteritoMaisQuePerfeitoDoIndicativoNos,
			Vos:           verbRow.PreteritoMaisQuePerfeitoDoIndicativoVos,
			VocesElesElas: verbRow.PreteritoMaisQuePerfeitoDoIndicativoVocesElesElas,
		},
		FuturoDoPresenteDoIndicativo: model.VerbConjugations{
			Eu:            verbRow.FuturoDoPresenteDoIndicativoEu,
			Tu:            verbRow.FuturoDoPresenteDoIndicativoTu,
			VoceEleEla:    verbRow.FuturoDoPresenteDoIndicativoVoceEleEla,
			Nos:           verbRow.FuturoDoPresenteDoIndicativoNos,
			Vos:           verbRow.FuturoDoPresenteDoIndicativoVos,
			VocesElesElas: verbRow.FuturoDoPresenteDoIndicativoVocesElesElas,
		},
		FuturoDoPreteritoDoIndicativo: model.VerbConjugations{
			Eu:            verbRow.FuturoDoPreteritoDoIndicativoEu,
			Tu:            verbRow.FuturoDoPreteritoDoIndicativoTu,
			VoceEleEla:    verbRow.FuturoDoPreteritoDoIndicativoVoceEleEla,
			Nos:           verbRow.FuturoDoPreteritoDoIndicativoNos,
			Vos:           verbRow.FuturoDoPreteritoDoIndicativoVos,
			VocesElesElas: verbRow.FuturoDoPreteritoDoIndicativoVocesElesElas,
		},
	}
}
