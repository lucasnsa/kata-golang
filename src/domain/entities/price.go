package domain

type Price struct {
	Code   string `json:"code"`
	Codein string `json:"codein"`
	Name   string `json:"name"`
	High   string `json:"high"`
	Low    string `json:"low"`
}
