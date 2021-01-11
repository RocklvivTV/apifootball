package apifootball

// League represents a league(s) api call response
type League struct {
	Api struct {
		Results int       `json:"results"`
		Leagues []Leagues `json:"leagues"`
	}
}

// Leagues represents a array of leagues returned by api call
type Leagues struct {
	LeagueID    int            `json:"league_id"`
	Name        string         `json:"name"`
	LeagueType  string         `json:"type"`
	Country     string         `json:"country"`
	CountryCode string         `json:"country_code"`
	Season      int            `json:"season"`
	SeasonStart string         `json:"season_start"`
	SeasonEnd   string         `json:"season_end"`
	Logo        string         `json:"logo"`
	Flag        string         `json:"flag"`
	Standings   int            `json:"standings"`
	IsCurrent   int            `json:"is_current"`
	Coverage    LeagueCoverage `json:"coverage"`
}

// LeagueCoverage represents a status of each league and what is covered by RapidAPI APIFOOTBALL provider
type LeagueCoverage struct {
	Stangings   bool                   `json:"standings"`
	Fixtures    LeagueCoverageFixtures `json:"fixtures"`
	Players     bool                   `json:"players"`
	TopScorers  bool                   `json:"topScorers"`
	Predictions bool                   `json:"predictions"`
	Odds        bool                   `json:"odds"`
}

// LeagueCoverageFixtures represents a status of each league and what is covered by each fixture
type LeagueCoverageFixtures struct {
	Events            bool `json:"events"`
	Lineups           bool `json:"lineups"`
	Statistics        bool `json:"statistics"`
	PlayersStatistics bool `json:"players_statistics"`
}

// Teams represents teams in required league
type Teams struct {
	Api struct {
		Results int           `json:"results"`
		Teams   []LeagueTeams `json:"teams"`
	}
}

// LeagueTeams represents a information about team in required league
type LeagueTeams struct {
	TeamID        int    `json:"team_id"`
	Name          string `json:"name"`
	Code          string `json:"code"`
	Logo          string `json:"logo"`
	Country       string `json:"country"`
	IsNational    bool   `json:"is_national"`
	Founded       int    `json:"founded"`
	VenueName     string `json:"venue_name"`
	VenueSurface  string `json:"venue_surface"`
	VenueAddress  string `json:"venue_address"`
	VenueCity     string `json:"venue_city"`
	VenueCapacity int    `json:"venue_capacity"`
}

// TeamSquad represents a Api output for squad of the team
type TeamSquad struct {
	Api struct {
		Results int      `json:"results"`
		Players []Player `json:"players"`
	}
}

// Player represents an personal information about football player
type Player struct {
	PlayerID     int    `json:"player_id"`
	PlayerName   string `json:"player_name"`
	FirstName    string `json:"firstname"`
	LastName     string `json:"lastname"`
	Number       int    `json:"number,omitempty"`
	Position     string `json:"position"`
	Age          int    `json:"age"`
	BirthDay     string `json:"birth_day"`
	BirthPlace   string `json:"birth_place"`
	BirthCountry string `json:"birth_country"`
	Nationality  string `json:"nationality"`
	Height       string `json:"height"`
	Weight       string `json:"weight"`
}

// Standings represents a league table
type Standings struct {
	Api struct {
		Results   int        `json:"results"`
		Standings []Standing `json:"standings"`
		cRPM      int
		cRPD      int
	}
}

// Standing represents standing in league
type Standing []struct {
	Rank        int         `json:"rank"`
	TeamID      int         `json:"team_id"`
	TeamName    string      `json:"teamName"`
	Logo        string      `json:"logo"`
	Group       string      `json:"group"`
	Forme       string      `json:"forme"`
	Description string      `json:"description"`
	All         FormeStruct `json:"all"`
	Home        FormeStruct `json:"home"`
	Away        FormeStruct `json:"away"`
	GoalsDiff   int         `json:"goalsDiff"`
	Points      int         `json:"points"`
	LastUpdate  string      `json:"lastUpdate"`
}

// FormeStruct represents a team forme like: games played, GA, GF
type FormeStruct struct {
	MatchesPlayed int `json:"matchesPlayed"`
	Win           int `json:"win"`
	Draw          int `json:"draw"`
	Lose          int `json:"lose"`
	GoalsFor      int `json:"goalsFor"`
	GoalsAgainst  int `json:"goalsAgainst"`
}

// LeagueFixtures represents all fixtures in required league
type LeagueFixtures struct {
	Api struct {
		Results  int        `json:"results"`
		Fixtures []Fixtures `json:"fixtures"`
		cRPM     int
		cRPD     int
	}
}

// Fixtures represents data about each fixture in a league
type Fixtures struct {
	FixtureID int `json:"fixture_id"`
	LeagueID  int `json:"league_id"`
	League    struct {
		Name    string `json:"name"`
		Country string `json:"country"`
		Logo    string `json:"logo"`
		Flag    string `json:"flag"`
	} `json:"league"`
	EventDate       string `json:"event_date"`
	EventTimestamp  int    `json:"event_timestamp"`
	FirstHalfStart  int    `json:"firstHalfStart"`
	SecondHalfStart int    `json:"secondHalfStart"`
	Round           string `json:"round"`
	Status          string `json:"status"`
	StatusShort     string `json:"status_short"`
	Elapsed         int    `json:"elapsed"`
	Venue           string `json:"venue"`
	Referee         string `json:"referee"`
	HomeTeam        struct {
		TeamID   int    `json:"team_id"`
		TeamName string `json:"team_name"`
		Logo     string `json:"logo"`
	} `json:"home_team"`
	AwayTeam struct {
		TeamID   int    `json:"team_id"`
		TeamName string `json:"team_name"`
		Logo     string `json:"logo"`
	} `json:"away_team"`
	GoalsHomeTeam int `json:"goalsHomeTeam"`
	GoalsAwayTeam int `json:"goalsAwayTeam"`
	Score         struct {
		Halftime  string `json:"halftime"`
		Fulltime  string `json:"fulltime"`
		Extratime string `json:"extratime"`
		Penalty   string `json:"penalty"`
	} `json:"score"`
}
