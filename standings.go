package apifootball

import (
	"fmt"
	"log"
)

var (
	standingsPrefix = "leagueTable"
)

// GetStandings returns standings of league by ID
func (c *APIClient) GetStandings(leagueID int) (*Standings, error) {
	var standings Standings
	res, err := c.DoRequests("GET", standingsPrefix, nil)
	if err != nil {
		log.Println(fmt.Errorf(err.Error()))
		return nil, err
	}

	err = res.Decode(&standings)
	if err != nil {
		log.Println(fmt.Errorf(err.Error()))
		return nil, err
	}
	return &standings, nil
}
