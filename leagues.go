package apifootball

import (
	"fmt"
	"log"
)

var (
	leaguesPrefix          = "leagues"
	leaguePrefix           = "leagues/league"
	leagueSeasonsAvailable = "leagues/seasonsAvailable"
)

// GetLeagues returns list of available leagues
func (c *APIClient) GetLeagues() (*League, error) {
	var league League
	res, err := c.DoRequests("GET", leaguesPrefix, nil)

	if err != nil {
		log.Println(fmt.Errorf(err.Error()))
		return nil, err
	}

	err = res.Decode(&league)
	if err != nil {
		log.Println(fmt.Errorf(err.Error()))
		return nil, err
	}
	return &league, nil
}

// GetLeagueByID returns league by ID
func (c *APIClient) GetLeagueByID(leagueID int) (*League, error) {
	var league League
	uri := fmt.Sprintf("%s/%d", leaguePrefix, leagueID)
	res, err := c.DoRequests("GET", uri, nil)

	if err != nil {
		log.Println(err)
		return nil, err
	}

	err = res.Decode(&league)
	if err != nil {
		log.Panicln(fmt.Errorf(err.Error()))
		return nil, err
	}
	return &league, nil
}

// GetLeagueAvailableSeasons returns list of available seasons for required league
func (c *APIClient) GetLeagueAvailableSeasons(leagueID int) (*League, error) {
	var league League
	uri := fmt.Sprintf("%s/%d", leagueSeasonsAvailable, leagueID)

	res, err := c.DoRequests("GET", uri, nil)
	if err != nil {
		log.Println(fmt.Sprintf(err.Error()))
		return nil, err
	}

	err = res.Decode(&league)
	if err != nil {
		log.Println(fmt.Sprintf(err.Error()))
		return nil, err
	}

	return &league, nil
}

// GetLeagueInfoBySeason returns required season for specified league
func (c *APIClient) GetLeagueInfoBySeason(leagueID, season int) (*League, error) {
	var league League
	uri := fmt.Sprintf("%s/%d/%d", leagueSeasonsAvailable, leagueID, season)

	res, err := c.DoRequests("GET", uri, nil)
	if err != nil {
		log.Println(fmt.Errorf(err.Error()))
		return nil, err
	}

	err = res.Decode(&league)
	if err != nil {
		log.Println(fmt.Errorf(err.Error()))
		return nil, err
	}

	return &league, nil
}
