package types

import (
	"fmt"
	"time"
)

type Event struct {
	ID            string        `json:"id"`
	LeagueName    string        `json:"league_name"`
	WeekName      string        `json:"week_name"`
	StartTime     time.Time     `json:"start_time"`
	T1            string        `json:"t1"`
	T2            string        `json:"t2"`
	HasVod        bool          `json:"has_vod"`
	MatchStrategy MatchStrategy `json:"match_stategy,omitempty"`
}

type MatchStrategyType string

const (
	MatchStrategyTypeBestOf MatchStrategyType = "bestOf"
)

type MatchStrategy struct {
	Type  MatchStrategyType `json:"type"`
	Count int               `json:"count"`
}

func (ms MatchStrategy) ToString() string {
	switch ms.Type {
	case MatchStrategyTypeBestOf:
		return fmt.Sprintf("Bo%d", ms.Count)
	default:
		return fmt.Sprintf("%s%d", ms.Type, ms.Count)
	}
}
