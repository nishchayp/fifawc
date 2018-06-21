package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

const (
	apiURL         = "http://api.football-data.org"
	apiVersion     = "/v1"
	apiKey         = "fd40d0d8c60842e9bc81ba025a33a719"
	competitionID  = "/467"
	compEndPoint   = "/competitions"
	teamsEndPoint  = "/teams"
	fixtureMethod  = "/fixtures"
	playersMethod  = "/players"
	nameFilter     = "?name="
	matchdayFilter = "?matchday="
)

func getFixtures(matchday int) Fixtures {

	client := &http.Client{}

	url := apiURL + apiVersion + compEndPoint + competitionID + fixtureMethod

	if matchday != -1 {
		url += matchdayFilter + strconv.Itoa(matchday)
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatalf("Error forming request: %s\n", err)
	}

	req.Header.Set("X-Auth-Token", apiKey)

	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Error retrieving data: %s\n", err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Error reading data: %s\n", err)
	}

	var fixtures Fixtures
	json.Unmarshal(body, &fixtures)
	return fixtures
}

func getTeamCode(teamName string) int {

	client := &http.Client{}

	url := apiURL + apiVersion + teamsEndPoint + nameFilter + teamName

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatalf("Error forming request: %s\n", err)
	}

	req.Header.Set("X-Auth-Token", apiKey)

	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Error retrieving data: %s\n", err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Error reading data: %s\n", err)
	}

	var team Team
	json.Unmarshal(body, &team)

	teamId := team.Teams[0].ID

	return teamId
}

func getSquad(teamCode int) Squad {

	client := &http.Client{}

	url := apiURL + apiVersion + teamsEndPoint + "/" + strconv.Itoa(teamCode) + playersMethod

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatalf("Error forming request: %s\n", err)
	}

	req.Header.Set("X-Auth-Token", apiKey)

	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Error retrieving data: %s\n", err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Error reading data: %s\n", err)
	}

	var squad Squad
	json.Unmarshal(body, &squad)
	return squad
}
