package models

type GameRequest struct {
	ID       int     `json:"id,omitempty"`
	Niveau   int     `json:"niveau"`
	Location string  `json:"location"`
	Time     string  `json:"time"`
	Gender   string  `json:"gender"`
	Amount   int     `json:"amount"`
	Price    float64 `json:"price"`
}
