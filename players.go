package apifootball

import (
	"fmt"
	"log"
)

var (
	playersSquadByTeam = "players/squad"
)

// GetSquadByTeam returns list of players in required team by season
func (c *APIClient) GetSquadByTeam(teamID, year int) (*TeamSquad, error) {
	var squad TeamSquad
	uri := fmt.Sprintf("%s/%d/%d", playersSquadByTeam, teamID, year)

	res, err := c.DoRequests("GET", uri, nil)
	if err != nil {
		log.Println(fmt.Errorf(err.Error()))
		return nil, err
	}

	err = res.Decode(&squad)
	if err != nil {
		log.Println(fmt.Errorf(err.Error()))
		return nil, err
	}

	return &squad, nil
}
