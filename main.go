package main

import (
	"fmt"
	flag "github.com/ogier/pflag"
	"os"
)

var (
	options struct {
		fixtures bool
		team     string
	}
)

func main() {

	flag.Parse()
	if flag.NFlag() == 0 {
		printUsage()
	}

	if options.fixtures {
		result := getFixtures()
		for _, f := range result.Fixtures {
			fmt.Println(`Matchday: `, f.Matchday)
			fmt.Println(`Date: `, f.Date)

			if f.Status == "SCHEDULED" {
				fmt.Println(`TBD vs TBD`)
				fmt.Println("/", "-", "/")
			} else if f.Status == "TIMED" {
				fmt.Printf(`%s vs %s`, f.HomeTeamName, f.AwayTeamName)
				fmt.Println("")
				fmt.Println("/", "-", "/")
			} else {
				fmt.Printf(`%s vs %s`, f.HomeTeamName, f.AwayTeamName)
				fmt.Println("")
				fmt.Println(f.Result.GoalsHomeTeam, "-", f.Result.GoalsAwayTeam)
			}

			fmt.Println("")
		}
	}

	if options.team != "" {
		result := getSquad(getTeamCode(options.team))
		fmt.Printf("%-3v| %-25v| %-20v\n", "NUM", "NAME", "POSITION")
		for _, player := range result.Players {
			fmt.Printf("%-3v| %-25v| %-20v\n", player.JerseyNumber, player.Name, player.Position)
		}
	}

}

func init() {
	flag.BoolVarP(&options.fixtures, "fixtures", "f", false, "Get fixtures")
	flag.StringVarP(&options.team, "squad", "s", "", "Get squad by team name")
}

func printUsage() {
	fmt.Printf("Usage %s [options]", os.Args[0])
	fmt.Println("Options:")
	flag.PrintDefaults()
	os.Exit(1)
}
