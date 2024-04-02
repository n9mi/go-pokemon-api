package model

type Pokemon struct {
	ID             int     `json:"id"`
	Name           string  `json:"name"`
	Weight         int     `json:"weight"`
	Height         int     `json:"height"`
	Species        Species `json:"species"`
	Stats          []Stat  `json:"stats"`
	Types          []Type  `json:"types"`
	Forms          []Form  `json:"forms"`
	Cries          Cries   `json:"cries"`
	Sprites        Sprites `json:"sprites"`
	BaseExperience int     `json:"base_experience"`
}
