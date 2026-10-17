package generator

import (
	"errors"
	"testing"
)

type conjugationTestCase struct {
	verb                string
	correctConjugations map[Pronoun]string
}

func newTestGenerator(t *testing.T) *Generator {
	t.Helper()
	g, err := New()
	if err != nil {
		t.Fatal(err)
	}
	return g
}

func TestNewWithDifferentWorkingDirectory(t *testing.T) {
	t.Chdir(t.TempDir())
	g := newTestGenerator(t)

	cases := []struct {
		tense Tense
		verb  string
		want  string
	}{
		{PresenteDoIndicativo, "falar", "falo"},
		{PresenteDoIndicativo, "ir", "vou"},
		{PreteritoPerfeitoDoIndicativo, "falar", "falei"},
		{PreteritoPerfeitoDoIndicativo, "ir", "fui"},
	}

	for _, tt := range cases {
		t.Run(string(tt.tense)+"/"+tt.verb, func(t *testing.T) {
			got, err := g.Conjugate(tt.verb, tt.tense, Eu)
			if err != nil || got != tt.want {
				t.Fatalf("got (%q, %v), want %q", got, err, tt.want)
			}
		})
	}
}

func TestConjugateIrregularOverrides(t *testing.T) {
	cases := []struct {
		name      string
		overrides map[string]string
		pronoun   Pronoun
		want      string
		wantErr   error
	}{
		{
			name:      "present override",
			overrides: map[string]string{"eu": "conheço"},
			pronoun:   Eu,
			want:      "conheço",
		},
		{
			name:      "missing pronoun uses regular rule",
			overrides: map[string]string{"eu": "conheço"},
			pronoun:   Tu,
			want:      "conheces",
		},
		{
			name:      "empty form is unavailable",
			overrides: map[string]string{"eu": ""},
			pronoun:   Eu,
			wantErr:   ErrConjugationUnavailable,
		},
		{
			name:    "nil override map uses regular rule",
			pronoun: Tu,
			want:    "conheces",
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			g := newTestGenerator(t)
			g.tenseData[PresenteDoIndicativo].Irregular["conhecer"] = tt.overrides

			got, err := g.Conjugate("conhecer", PresenteDoIndicativo, tt.pronoun)
			if got != tt.want || !errors.Is(err, tt.wantErr) {
				t.Fatalf("got (%q, %v), want (%q, %v)", got, err, tt.want, tt.wantErr)
			}
		})
	}
}

func TestConjugateMissingIrregularAndRegularForm(t *testing.T) {
	g := newTestGenerator(t)
	g.tenseData[PresenteDoIndicativo].Irregular["conhecer"] = map[string]string{"eu": "conheço"}
	delete(g.tenseData[PresenteDoIndicativo].RegularEndings["er"], "tu")

	got, err := g.Conjugate("conhecer", PresenteDoIndicativo, Tu)
	if got != "" || !errors.Is(err, ErrConjugationUnavailable) {
		t.Fatalf("got (%q, %v), want (%q, %v)", got, err, "", ErrConjugationUnavailable)
	}
}

func TestPresenteDoIndicativoConjugations(t *testing.T) {
	g := newTestGenerator(t)

	cases := []conjugationTestCase{
		// Regular examples
		{
			"gostar",
			map[Pronoun]string{
				Eu:            "gosto",
				VoceEleEla:    "gosta",
				Nos:           "gostamos",
				VocesElesElas: "gostam",
				Tu:            "gostas",
			},
		},
		{
			"comer",
			map[Pronoun]string{
				Eu:            "como",
				VoceEleEla:    "come",
				Nos:           "comemos",
				VocesElesElas: "comem",
				Tu:            "comes",
			},
		},
		{
			"assistir",
			map[Pronoun]string{
				Eu:            "assisto",
				VoceEleEla:    "assiste",
				Nos:           "assistimos",
				VocesElesElas: "assistem",
				Tu:            "assistes",
			},
		},
		// Irregular examples
		{
			"ser",
			map[Pronoun]string{
				Eu:            "sou",
				VoceEleEla:    "é",
				Nos:           "somos",
				VocesElesElas: "são",
				Tu:            "és",
			},
		},
		{
			"estar",
			map[Pronoun]string{
				Eu:            "estou",
				VoceEleEla:    "está",
				Nos:           "estamos",
				VocesElesElas: "estão",
				Tu:            "estás",
			},
		},
		{
			"ter",
			map[Pronoun]string{
				Eu:            "tenho",
				VoceEleEla:    "tem",
				Nos:           "temos",
				VocesElesElas: "têm",
				Tu:            "tens",
			},
		},
		{
			"ir",
			map[Pronoun]string{
				Eu:            "vou",
				VoceEleEla:    "vai",
				Nos:           "vamos",
				VocesElesElas: "vão",
				Tu:            "vais",
			},
		},
		// Mixed regular and irregular
		{
			"conhecer",
			map[Pronoun]string{
				Eu:            "conheço",
				VoceEleEla:    "conhece",
				Nos:           "conhecemos",
				VocesElesElas: "conhecem",
				Tu:            "conheces",
			},
		},
	}

	execute(t, g, PresenteDoIndicativo, cases)
}

func TestPreteritoPerfeitoDoIndicativoConjugations(t *testing.T) {
	g := newTestGenerator(t)

	cases := []conjugationTestCase{
		// Regular examples
		{
			"andar",
			map[Pronoun]string{
				Eu:            "andei",
				VoceEleEla:    "andou",
				Nos:           "andamos",
				VocesElesElas: "andaram",
				Tu:            "andaste",
			},
		},
		{
			"conhecer",
			map[Pronoun]string{
				Eu:            "conheci",
				VoceEleEla:    "conheceu",
				Nos:           "conhecemos",
				VocesElesElas: "conheceram",
				Tu:            "conheceste",
			},
		},
		{
			"assistir",
			map[Pronoun]string{
				Eu:            "assisti",
				VoceEleEla:    "assistiu",
				Nos:           "assistimos",
				VocesElesElas: "assistiram",
				Tu:            "assististe",
			},
		},
		// Irregular examples
		{
			"ser",
			map[Pronoun]string{
				Eu:            "fui",
				VoceEleEla:    "foi",
				Nos:           "fomos",
				VocesElesElas: "foram",
				Tu:            "foste",
			},
		},
		{
			"estar",
			map[Pronoun]string{
				Eu:            "estive",
				VoceEleEla:    "esteve",
				Nos:           "estivemos",
				VocesElesElas: "estiveram",
				Tu:            "estiveste",
			},
		},
		{
			"ter",
			map[Pronoun]string{
				Eu:            "tive",
				VoceEleEla:    "teve",
				Nos:           "tivemos",
				VocesElesElas: "tiveram",
				Tu:            "tiveste",
			},
		},
		{
			"ir",
			map[Pronoun]string{
				Eu:            "fui",
				VoceEleEla:    "foi",
				Nos:           "fomos",
				VocesElesElas: "foram",
				Tu:            "foste",
			},
		},
		// Mixed regular and irregular
		{
			"ficar",
			map[Pronoun]string{
				Eu:            "fiquei",
				VoceEleEla:    "ficou",
				Nos:           "ficamos",
				VocesElesElas: "ficaram",
				Tu:            "ficaste",
			},
		},
	}

	execute(t, g, PreteritoPerfeitoDoIndicativo, cases)
}

func execute(t *testing.T, g *Generator, tense Tense, cases []conjugationTestCase) {
	for _, tt := range cases {
		for pronoun, correctConjugation := range tt.correctConjugations {
			t.Run(string(tense)+"/"+string(pronoun)+"/"+tt.verb, func(t *testing.T) {
				got, err := g.Conjugate(tt.verb, tense, pronoun)

				if err != nil || got != correctConjugation {
					t.Fatalf("got (%q, %v), want %q", got, err, correctConjugation)
				}
			})
		}
	}
}
