package model

type Ability struct {
	IsHidden   bool       `json:"is_hidden"`
	Slot       int        `json:"slot"`
	SubAbility SubAbility `json:"ability"`
}
