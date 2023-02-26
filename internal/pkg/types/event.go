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
	T1            Team          `json:"t1"`
	T2            Team          `json:"t2"`
	HasVod        bool          `json:"has_vod"`
	MatchStrategy MatchStrategy `json:"match_stategy,omitempty"`
}

type Team struct {
	Code   string
	Result MatchTeamResult
}

type MatchTeamResult struct {
	Outcome  string `json:"outcome"`
	GameWins int    `json:"gameWins"`
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
