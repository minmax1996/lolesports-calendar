package types

import "time"

type Event struct {
	ID         string    `json:"id"`
	LeagueName string    `json:"league_name"`
	WeekName   string    `json:"week_name"`
	StartTime  time.Time `json:"start_time"`
	T1         string    `json:"t1"`
	T2         string    `json:"t2"`
	HasVod     bool      `json:"has_vod"`
}
