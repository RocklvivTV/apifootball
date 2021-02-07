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
	log.Println(res.API.cRPD)
	for i := range res.API.Standings {
		r := res.API.Standings[i]
		if len(r) != 20 {
			t.Error()
		}
	}
}
