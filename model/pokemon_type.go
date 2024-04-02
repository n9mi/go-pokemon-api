package model

type Type struct {
	Slot    int     `json:"slot"`
	SubType SubType `json:"type"`
}
