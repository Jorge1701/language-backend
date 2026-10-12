# language-backend

REST API backend supporting the [language-frontend](https://github.com/Jorge1701/language-frontend) Portuguese language learning app.

Conjugation data generated using [gconjugue](https://github.com/jalvesaq/gconjugue) by jalvesaq.

## Generator

The `generator` package uses the existing `model.Tense` and `model.Pronoun` types:

```go
g, err := generator.New()
if err != nil {
    return err
}
form, err := g.Conjugate("ir", model.SimplePresent, model.FirstPersonalSingular)
// form == "vou"
```

`New` loads the embedded `{tense}-irregulars.yml` files once. `Conjugate` first
looks up the verb and pronoun in the requested tense's map, then falls back to
that tense's regular rule. All six model tenses are implemented for regular
`-ar`, `-er`, and `-ir` verbs. `PastPerfect` produces the simple
pretérito mais-que-perfeito (`falara`).

The initial YAML files contain common irregular verbs and spelling exceptions,
seeded from the existing conjugation dataset. They are not an exhaustive list
of Portuguese irregular verbs: verbs without an override are assumed regular.
To add an exception, edit the corresponding file, using the model's pronoun names:

```yaml
ir:
  FIRST_PERSON_SINGULAR: vou
  SECOND_PERSON_SINGULAR: vais
```

Unlisted pronouns use regular rules. An explicit empty string (`""`) marks an
unavailable form and returns `generator.ErrConjugationUnavailable`. Invalid
inputs and unsupported rules return errors; missing tense functions return
`generator.ErrNotImplemented`. YAML is validated during initialization, and
file changes require rebuilding the application. The package does not require
the database at runtime.

## License
Licensed under the GNU AGPL v3. See LICENSE for details.
