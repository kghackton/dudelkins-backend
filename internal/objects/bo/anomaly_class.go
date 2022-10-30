package bo

type AnomalyClass struct {
	Verdict     bool   `json:"verdict"`
	Description string `json:"description,omitempty"`
}
