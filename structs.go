package main

import (
	"time"
)

type Fixtures struct {
	Links struct {
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
		Competition struct {
			Href string `json:"href"`
		} `json:"competition"`
	} `json:"_links"`
	Count    int `json:"count"`
	Fixtures []struct {
		Links struct {
			Self struct {
				Href string `json:"href"`
			} `json:"self"`
			Competition struct {
				Href string `json:"href"`
			} `json:"competition"`
			HomeTeam struct {
				Href string `json:"href"`
			} `json:"homeTeam"`
			AwayTeam struct {
				Href string `json:"href"`
			} `json:"awayTeam"`
		} `json:"_links"`
		Date         time.Time `json:"date"`
		Status       string    `json:"status"`
		Matchday     int       `json:"matchday"`
		HomeTeamName string    `json:"homeTeamName"`
		AwayTeamName string    `json:"awayTeamName"`
		Result       struct {
			GoalsHomeTeam interface{} `json:"goalsHomeTeam"`
			GoalsAwayTeam interface{} `json:"goalsAwayTeam"`
		} `json:"result"`
		Odds interface{} `json:"odds"`
	} `json:"fixtures"`
}

type Team struct {
	Links struct {
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
	} `json:"_links"`
	SearchQuery string `json:"searchQuery"`
	Count       int    `json:"count"`
	Teams       []struct {
		ID                 int         `json:"id"`
		Name               string      `json:"name"`
		CurrentCompetition interface{} `json:"currentCompetition"`
		CurrentLeague      interface{} `json:"currentLeague"`
	} `json:"teams"`
}

type Squad struct {
	Links struct {
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
		Team struct {
			Href string `json:"href"`
		} `json:"team"`
	} `json:"_links"`
	Count   int `json:"count"`
	Players []struct {
		Name          string      `json:"name"`
		Position      string      `json:"position"`
		JerseyNumber  int         `json:"jerseyNumber"`
		DateOfBirth   string      `json:"dateOfBirth"`
		Nationality   string      `json:"nationality"`
		ContractUntil string      `json:"contractUntil"`
		MarketValue   interface{} `json:"marketValue"`
	} `json:"players"`
}
