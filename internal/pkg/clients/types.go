package clients

import "time"

type event struct {
	StartTime time.Time `json:"startTime"`
	State     string    `json:"state"`
	Type      string    `json:"type"`
	BlockName string    `json:"blockName"`
	League    struct {
		Name string `json:"name"`
		Slug string `json:"slug"`
	} `json:"league"`
	Match struct {
		ID    string   `json:"id"`
		Flags []string `json:"flags"`
		Teams []struct {
			Name   string `json:"name"`
			Code   string `json:"code"`
			Image  string `json:"image"`
			Result struct {
				Outcome  string `json:"outcome"`
				GameWins int    `json:"gameWins"`
			} `json:"result"`
			Record struct {
				Wins   int `json:"wins"`
				Losses int `json:"losses"`
			} `json:"record"`
		} `json:"teams"`
		Strategy MatchStrategy `json:"strategy"`
	} `json:"match"`
}

type MatchStrategy struct {
	Type  string `json:"type"`
	Count int    `json:"count"`
}

type getScheduleResponse struct {
	Data struct {
		Schedule struct {
			Pages struct {
				Older string  `json:"older"`
				Newer *string `json:"newer"`
			} `json:"pages"`
			Events []event `json:"events"`
		} `json:"schedule"`
	} `json:"data"`
}
