package apifootball

import (
	"log"
	"testing"
)

func TestGetStandingsByLeagueID(t *testing.T) {
	setup()
	client := NewAPIClient(apikey).SetRequestsLimit(5, 5)

	res, err := client.GetStandings(2097)
	if err != nil {
		t.Error(err)
	}
	log.Println(res.Api.cRPD)
	for i := range res.Api.Standings {
		r := res.Api.Standings[i]
		if len(r) != 20 {
			t.Error()
		}
	}
}
