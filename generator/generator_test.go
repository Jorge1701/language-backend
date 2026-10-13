package generator

import (
	"errors"
	"testing"
)

type conjugationTestCase struct {
	verb                string
	correctConjugations map[string]string
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

func TestSimplePresentConjugations(t *testing.T) {
	g := newTestGenerator(t)

	cases := []conjugationTestCase{
		// Regular examples
		{
			"gostar",
			map[string]string{
				"eu":              "gosto",
				"você/ele/ela":    "gosta",
				"nós":             "gostamos",
				"vocês/eles/elas": "gostam",
				"tu":              "gostas",
			},
		},
		{
			"comer",
			map[string]string{
				"eu":              "como",
				"você/ele/ela":    "come",
				"nós":             "comemos",
				"vocês/eles/elas": "comem",
				"tu":              "comes",
			},
		},
		{
			"assistir",
			map[string]string{
				"eu":              "assisto",
				"você/ele/ela":    "assiste",
				"nós":             "assistimos",
				"vocês/eles/elas": "assistem",
				"tu":              "assistes",
			},
		},
		// Irregular examples
		{
			"ser",
			map[string]string{
				"eu":              "sou",
				"você/ele/ela":    "é",
				"nós":             "somos",
				"vocês/eles/elas": "são",
				"tu":              "és",
			},
		},
		{
			"estar",
			map[string]string{
				"eu":              "estou",
				"você/ele/ela":    "está",
				"nós":             "estamos",
				"vocês/eles/elas": "estão",
				"tu":              "estás",
			},
		},
		{
			"ter",
			map[string]string{
				"eu":              "tenho",
				"você/ele/ela":    "tem",
				"nós":             "temos",
				"vocês/eles/elas": "têm",
				"tu":              "tens",
			},
		},
		{
			"ir",
			map[string]string{
				"eu":              "vou",
				"você/ele/ela":    "vai",
				"nós":             "vamos",
				"vocês/eles/elas": "vão",
				"tu":              "vais",
			},
		},
	}

	execute(t, g, PresenteDoIndicativo, cases)
}

func TestSimplePastConjugations(t *testing.T) {
	g := newTestGenerator(t)

	cases := []conjugationTestCase{
		// Regular examples
		{
			"andar",
			map[string]string{
				"eu":              "andei",
				"você/ele/ela":    "andou",
				"nós":             "andamos",
				"vocês/eles/elas": "andaram",
				"tu":              "andaste",
			},
		},
		{
			"conhecer",
			map[string]string{
				"eu":              "conheci",
				"você/ele/ela":    "conheceu",
				"nós":             "conhecemos",
				"vocês/eles/elas": "conheceram",
				"tu":              "conheceste",
			},
		},
		{
			"assistir",
			map[string]string{
				"eu":              "assisti",
				"você/ele/ela":    "assistiu",
				"nós":             "assistimos",
				"vocês/eles/elas": "assistiram",
				"tu":              "assististe",
			},
		},
		// Irregular examples
		{
			"ser",
			map[string]string{
				"eu":              "fui",
				"você/ele/ela":    "foi",
				"nós":             "fomos",
				"vocês/eles/elas": "foram",
				"tu":              "foste",
			},
		},
		{
			"estar",
			map[string]string{
				"eu":              "estive",
				"você/ele/ela":    "esteve",
				"nós":             "estivemos",
				"vocês/eles/elas": "estiveram",
				"tu":              "estiveste",
			},
		},
		{
			"ter",
			map[string]string{
				"eu":              "tive",
				"você/ele/ela":    "teve",
				"nós":             "tivemos",
				"vocês/eles/elas": "tiveram",
				"tu":              "tiveste",
			},
		},
		{
			"ir",
			map[string]string{
				"eu":              "fui",
				"você/ele/ela":    "foi",
				"nós":             "fomos",
				"vocês/eles/elas": "foram",
				"tu":              "foste",
			},
		},
	}

	execute(t, g, PreteritoPerfeitoDoIndicativo, cases)
}

func execute(t *testing.T, g *Generator, tense Tense, cases []conjugationTestCase) {
	tenseText, err := TenseText(tense)
	if err != nil {
		t.Fatal(err)
	}

	for _, tt := range cases {
		for pronounText, correctConjugation := range tt.correctConjugations {
			pronoun, err := PronounFromText(pronounText)
			if err != nil {
				t.Fatal(err)
			}

			t.Run(tenseText+"/"+pronounText+"/"+tt.verb, func(t *testing.T) {
				got, err := g.Conjugate(tt.verb, tense, pronoun)

				if err != nil || got != correctConjugation {
					t.Fatalf("got (%q, %v), want %q", got, err, correctConjugation)
				}
			})
		}
	}
}
