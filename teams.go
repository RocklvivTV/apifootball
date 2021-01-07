package apifootball

import (
	"fmt"
	"log"
)

var (
	teamsByLeague = "teams/league"
)

// GetTeamsByLeague returns list of teams in required league
func (c *APIClient) GetTeamsByLeague(leagueID int) (*Teams, error) {
	var teams Teams
	uri := fmt.Sprintf("%s/%d", teamsByLeague, leagueID)

	res, err := c.DoRequests("GET", uri, nil)
	if err != nil {
		log.Println(fmt.Errorf(err.Error()))
		return nil, err
	}

	err = res.Decode(&teams)
	if err != nil {
		log.Println(fmt.Errorf(err.Error()))
		return nil, err
	}

	return &teams, nil
}
