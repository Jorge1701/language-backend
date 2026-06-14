package database

type VerbRow struct {
	Id                int    `json:"id"`
	Type              string `json:"type"`
	Infinitive        string `json:"infinitive"`
	PresentParticiple string `json:"present_participle"`
	PastParticiple    string `json:"past_participle"`
	SimplePresentFPS  string `json:"simple_present_fps"`
	SimplePresentSPS  string `json:"simple_present_sps"`
	SimplePresentTPS  string `json:"simple_present_tps"`
	SimplePresentFPP  string `json:"simple_present_fpp"`
	SimplePresentSPP  string `json:"simple_present_spp"`
	SimplePresentTPP  string `json:"simple_present_tpp"`
	ImperfectPastFPS  string `json:"imperfect_past_fps"`
	ImperfectPastSPS  string `json:"imperfect_past_sps"`
	ImperfectPastTPS  string `json:"imperfect_past_tps"`
	ImperfectPastFPP  string `json:"imperfect_past_fpp"`
	ImperfectPastSPP  string `json:"imperfect_past_spp"`
	ImperfectPastTPP  string `json:"imperfect_past_tpp"`
	SimplePastFPS     string `json:"simple_past_fps"`
	SimplePastSPS     string `json:"simple_past_sps"`
	SimplePastTPS     string `json:"simple_past_tps"`
	SimplePastFPP     string `json:"simple_past_fpp"`
	SimplePastSPP     string `json:"simple_past_spp"`
	SimplePastTPP     string `json:"simple_past_tpp"`
	PastPerfectFPS    string `json:"past_perfect_fps"`
	PastPerfectSPS    string `json:"past_perfect_sps"`
	PastPerfectTPS    string `json:"past_perfect_tps"`
	PastPerfectFPP    string `json:"past_perfect_fpp"`
	PastPerfectSPP    string `json:"past_perfect_spp"`
	PastPerfectTPP    string `json:"past_perfect_tpp"`
	SimpleFutureFPS   string `json:"simple_future_fps"`
	SimpleFutureSPS   string `json:"simple_future_sps"`
	SimpleFutureTPS   string `json:"simple_future_tps"`
	SimpleFutureFPP   string `json:"simple_future_fpp"`
	SimpleFutureSPP   string `json:"simple_future_spp"`
	SimpleFutureTPP   string `json:"simple_future_tpp"`
	ConditionalFPS    string `json:"conditional_fps"`
	ConditionalSPS    string `json:"conditional_sps"`
	ConditionalTPS    string `json:"conditional_tps"`
	ConditionalFPP    string `json:"conditional_fpp"`
	ConditionalSPP    string `json:"conditional_spp"`
	ConditionalTPP    string `json:"conditional_tpp"`
}

type ExampleRow struct {
	Pt string
	Es string
}

type User struct {
	Id        int
	Username  string
	Password  string
	CreatedAt string
	UpdatedAt string
}
