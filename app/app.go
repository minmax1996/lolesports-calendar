package app

import (
	"fmt"
	"strings"
	"time"

	ics "github.com/arran4/golang-ical"
	"github.com/minmax1996/lolesports-calendar/internal/pkg/clients"
	"github.com/minmax1996/lolesports-calendar/internal/pkg/common"
	"github.com/minmax1996/lolesports-calendar/internal/pkg/types"
)

type App struct {
	lolesportsClient *clients.LolEsportsClient
}

func NewApp(c *clients.LolEsportsClient) *App {
	return &App{
		lolesportsClient: c,
	}
}

func (a *App) CalendarHandler(req types.CalendarRequest) ([]byte, error) {
	events, err := a.lolesportsClient.GetSchedule(req.Leagues)
	if err != nil {
		return nil, err
	}

	cal := ics.NewCalendarFor("name")
	for _, e := range events {
		if len(req.FavTeams) > 0 &&
			!common.Contains(req.FavTeams, e.T1) &&
			!common.Contains(req.FavTeams, e.T2) {
			continue
		}

		event := ics.NewEvent(e.ID)
		event.SetStartAt(e.StartTime)
		event.SetSummary(fmt.Sprintf("%s %s: %s vs %s (%s)", e.LeagueName, e.WeekName, e.T1, e.T2, e.MatchStrategy.ToString()))
		event.SetEndAt(e.StartTime.Add(time.Duration(e.MatchStrategy.Count) * time.Hour))
		event.SetDtStampTime(time.Now())
		if e.HasVod {
			event.SetDescription(fmt.Sprintf("game completed: find vod here: https://lolesports.com/vod/%s/1/", e.ID))
		} else {
			event.SetDescription(fmt.Sprintf("join live here: https://lolesports.com/live/%[1]s/%[1]s/", strings.ToLower(e.LeagueName)))
		}
		event.SetProperty(ics.ComponentProperty(ics.PropertyPriority), "5")
		cal.AddVEvent(event)
	}

	cal.Serialize()
	return []byte(cal.Serialize()), nil
}
