package clients

import (
	"encoding/json"
	"net/http"
	"net/url"
	"strings"

	"github.com/minmax1996/lolesports-calendar/internal/pkg/common"
	"github.com/minmax1996/lolesports-calendar/internal/pkg/types"
)

type LolEsportsClient struct {
	httpClient *http.Client
	token      string
}

func NewLolEsportsClient(token string) *LolEsportsClient {
	return &LolEsportsClient{
		httpClient: http.DefaultClient,
		token:      token,
	}
}

var leagueMap = map[string]string{
	"lec":    "98767991302996019",
	"lcl":    "98767991355908944",
	"lcs":    "98767991299243165",
	"lck":    "98767991310872058",
	"lpl":    "98767991314006698",
	"msi":    "98767991325878492",
	"worlds": "98767975604431411",
}

func (c *LolEsportsClient) buildGetScheduleRequest(leagues []string, page string) (*http.Request, error) {
	reqURL, err := url.Parse("https://esports-api.lolesports.com/persisted/gw/getSchedule")
	if err != nil {
		return nil, err
	}

	query := reqURL.Query()
	query.Add("hl", "en-GB")
	var leagueIds []string
	for _, l := range leagues {
		leagueIds = append(leagueIds, leagueMap[l])
	}
	query.Add("leagueId", strings.Join(leagueIds, ","))
	if page != "" {
		query.Add("pageToken", page)
	}

	reqURL.RawQuery = query.Encode()

	req, err := http.NewRequest(http.MethodGet, reqURL.String(), nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("authority", "esports-api.lolesports.com")
	req.Header.Add("sec-ch-ua", `" Not;A Brand";v="99", "Google Chrome";v="97", "Chromium";v="97"`)
	req.Header.Add("sec-ch-ua-mobile", "?0")
	req.Header.Add("user-agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/97.0.4692.71 Safari/537.36")
	req.Header.Add("x-api-key", c.token)
	req.Header.Add("sec-ch-ua-platform", `"macOS"`)
	req.Header.Add("accept", "*/*")
	req.Header.Add("origin", "https://lolesports.com")
	req.Header.Add("sec-fetch-site", "same-site")
	req.Header.Add("sec-fetch-mode", "cors")
	req.Header.Add("sec-fetch-dest", "empty")
	req.Header.Add("referer", "https://lolesports.com/")
	req.Header.Add("accept-language", "ru-RU,ru;q=0.9,en-US;q=0.8,en;q=0.7")
	return req, nil
}

func (c *LolEsportsClient) GetSchedule(leagues []string) ([]types.Event, error) {
	var result []types.Event
	var page string
	for {
		req, err := c.buildGetScheduleRequest(leagues, page)
		if err != nil {
			return nil, err
		}

		resp, err := c.httpClient.Do(req)
		if err != nil {
			return nil, err
		}

		defer func() {
			_ = resp.Body.Close()
		}()

		var res getScheduleResponse
		if err = json.NewDecoder(resp.Body).Decode(&res); err != nil {
			return nil, err
		}

		for _, e := range res.Data.Schedule.Events {
			if e.Type == "match" {
				result = append(result, types.Event{
					ID:         e.Match.ID,
					LeagueName: e.League.Name,
					WeekName:   e.BlockName,
					StartTime:  e.StartTime,
					T1:         e.Match.Teams[0].Code,
					T2:         e.Match.Teams[1].Code,
					HasVod:     common.Contains(e.Match.Flags, "hasVod"),
					MatchStrategy: types.MatchStrategy{
						Type:  types.MatchStrategyType(e.Match.Strategy.Type),
						Count: e.Match.Strategy.Count,
					},
				})
			}
		}

		if res.Data.Schedule.Pages.Newer == nil {
			break
		}
		page = *res.Data.Schedule.Pages.Newer
	}

	return result, nil
}
