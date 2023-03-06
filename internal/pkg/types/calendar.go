package types

import "strings"

type CalendarRequest struct {
	Leagues  []string `schema:"leagues"`
	FavTeams []string `schema:"teams"`
	Spoiler  bool     `schema:"spoiler"`
}

func (r *CalendarRequest) Normalize() {
	if len(r.Leagues) > 0 && strings.Contains(r.Leagues[0], ",") {
		r.Leagues = strings.Split(r.Leagues[0], ",")
	}
	for i, l := range r.Leagues {
		r.Leagues[i] = strings.ToLower(l)
	}

	if len(r.FavTeams) > 0 && strings.Contains(r.FavTeams[0], ",") {
		r.FavTeams = strings.Split(r.FavTeams[0], ",")
	}
	for i, t := range r.FavTeams {
		r.FavTeams[i] = strings.ToUpper(t)
	}
}
