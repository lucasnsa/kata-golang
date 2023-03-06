package domain

type NewsAndPrice struct {
	Price Price        `json:"price"`
	News  []NewsEntity `json:"news"`
}
