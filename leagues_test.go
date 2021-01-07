package apifootball

import (
	"testing"
)

func TestGetLeagueByID(t *testing.T) {
	setup()
	client := NewAPIClient(apikey).SetRequestsLimit(10, 10)
	res, err := client.GetLeagueByID(leagueID)
	if err != nil {
		t.Error(err)
	}

	leagues := res.Api.Leagues
	for i := range leagues {
		if leagues[i].LeagueID != 2790 {
			t.Error()
		}
	}
}
