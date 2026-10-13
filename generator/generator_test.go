package generator

import (
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
				"tu":              "estiviste",
			},
		},
		{
			"ter",
			map[string]string{
				"eu":              "tive",
				"você/ele/ela":    "teve",
				"nós":             "tivemos",
				"vocês/eles/elas": "tiveram",
				"tu":              "tiviste",
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
