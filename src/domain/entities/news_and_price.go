package domain

type NewsAndPrice struct {
	Price Price  `json:"price"`
	News  []News `json:"news"`
}
