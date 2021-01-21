package apifootball

import (
	"fmt"
	"log"
)

var (
	fixturesLeaguePrefix      = "fixtures/league"
	timezone                  = "Europe/London"
	currentFixtureRoundPrefix = "fixtures/rounds"
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

// GetCurrentFixtureRound returns current fixture round
func (c *APIClient) GetCurrentFixtureRound(leagueID int) (*CurrentFixtureRound, error) {
	var currentFixtureRound CurrentFixtureRound

	res, err := c.DoRequests("GET", fmt.Sprintf("%s/%d/current", currentFixtureRoundPrefix, leagueID), nil)
	if err != nil {
		log.Println(fmt.Errorf(err.Error()))
		return nil, err
	}
	err = res.Decode(&currentFixtureRound)
	if err != nil {
		log.Println(fmt.Errorf(err.Error()))
		return nil, err
	}

	return &currentFixtureRound, nil
}
