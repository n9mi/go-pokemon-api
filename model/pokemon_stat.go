package model

type Stat struct {
	BaseStat int     `json:"base_stat"`
	Effort   int     `json:"effort"`
	SubStat  SubStat `json:"stat"`
}
