package endpoints

type VerbInfo struct {
	Type             string           `json:"type"`
	Infinitive       string           `json:"infinitive,omitempty"`
	PresentPaticiple string           `json:"present_participle,omitempty"`
	PastParticiple   string           `json:"past_participle,omitempty"`
	SimplePresent    VerbConjugations `json:"simple_present"`
	ImperfectPast    VerbConjugations `json:"imperfect_past"`
	SimplePast       VerbConjugations `json:"simple_past"`
	PastPerfect      VerbConjugations `json:"past_perfect"`
	SimpleFuture     VerbConjugations `json:"simple_future"`
	Conditional      VerbConjugations `json:"conditional"`
}

type VerbConjugations struct {
	FirstPersonSingular  string `json:"fps"`
	SecondPersonSingular string `json:"sps"`
	ThirdPersonSingular  string `json:"tps"`
	FirstPersonPlural    string `json:"fpp"`
	SecondPersonPlural   string `json:"spp"`
	ThirdPersonPlural    string `json:"tpp"`
}
