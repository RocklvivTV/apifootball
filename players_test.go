package apifootball

import (
	"testing"
)

var (
	teamID int = 85
	season int = 2019
)

func TestGetSquadByTeam(t *testing.T) {
	setup()
	client := NewAPIClient(apikey).SetRequestsLimit(20, 10)

	res, err := client.GetSquadByTeam(teamID, season)
	if err != nil {
		t.Error()
	}
	if res.Api.Results < 1 {
		t.Error()
	}
}

func TestGetSquadByTeamFailure(t *testing.T) {
	setup()
	client := NewAPIClient(apikey).SetRequestsLimit(20, 10)
	playersSquadByTeam = "test"
	_, err := client.GetSquadByTeam(0, 50000)
	if err == nil {
		t.Error()
	}
}
