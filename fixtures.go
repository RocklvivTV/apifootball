package apifootball

import (
	"fmt"
	"log"
)

var (
	fixturesLeaguePrefix = "fixtures/league"
	timezone             = "Europe/London"
)

// GetFixturesByLeagueID returns fixtures in required league by ID
func (c *APIClient) GetFixturesByLeagueID(leagueID int) (*LeagueFixtures, error) {
	var fixtureLeagues LeagueFixtures

	res, err := c.DoRequests("GET", fmt.Sprintf("%s/%d", fixturesLeaguePrefix, leagueID), nil)
	if err != nil {
		log.Println(fmt.Errorf(err.Error()))
		return nil, err
	}

	err = res.Decode(&fixtureLeagues)
	if err != nil {
		log.Println(fmt.Errorf(err.Error()))
		return nil, err
	}

	return &fixtureLeagues, nil
}
