package apifootball

import (
	"fmt"
	"os"
	"testing"
)

var (
	apikey   string
	leagueID = 2790
)

func TestNewClient(t *testing.T) {
	setup()
	client := NewAPIClient(apikey)

	if client.apikey != apikey {
		t.Error()
	} else if client.apikey == "" {
		t.Error()
	}
}

func TestNewClientWithMaxRPM(t *testing.T) {
	setup()
	client := NewAPIClient(apikey).SetRequestsLimit(20, 200)

	if client.rpm != 20 {
		t.Error()
	}

	if client.reqDailyLimit != 200 {
		t.Error()
	}
}

func TestNewClientWithTimeouts(t *testing.T) {
	setup()
	client := NewAPIClient(apikey).SetTimeoutes(50, 10)

	if client.timeout != 10 {
		t.Error()
	}

	if client.dailyTimeout != 50 {
		t.Error()
	}
}

func TestRequest(t *testing.T) {
	var v League
	setup()
	client := NewAPIClient(apikey)

	res, err := client.DoRequests("GET", fmt.Sprintf("leagues/league/%d", leagueID), nil)
	if err != nil {
		t.Error(err)
	}

	if res == nil {
		t.Error()
	}
	res.Decode(&v)

	for k := range v.API.Leagues {
		if v.API.Leagues[k].LeagueID != 2790 {
			t.Error()
		}
	}
}

func TestNewClientWithoutApiToken(t *testing.T) {
	apikey = ""
	NewAPIClient(apikey)
}

// Setup function
func setup() {
	apikey = os.Getenv("FOOTBALL_API_KEY")
}
