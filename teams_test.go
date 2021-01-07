package apifootball

import (
	"testing"
)

var (
	expectedNumOfteams = 20
)

func TestTeamsByLeague(t *testing.T) {
	setup()
	client := NewAPIClient(apikey).SetRequestsLimit(5, 10)
	res, err := client.GetTeamsByLeague(leagueID)

	if err != nil {
		t.Error()
	}

	if res.Api.Results != expectedNumOfteams {
		t.Error()
	}
}

func TestTeamsByLeagueFailure(t *testing.T) {
	setup()
	client := NewAPIClient(apikey).SetRequestsLimit(5, 10)
	teamsByLeague = "test"
	_, err := client.GetTeamsByLeague(leagueID)

	if err == nil {
		t.Error()
	}
}
