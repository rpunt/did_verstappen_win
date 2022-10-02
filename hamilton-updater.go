package main

import (
	"fmt"
	"strconv"

	"github.com/rpunt/f1apireader"
)

func main() {
	// branch = 'main'
	// g = Git.open(__dir__)
	// g.checkout(branch)
	// # g.pull

	race, err := f1apireader.RaceResults("https://api.formula1.com/v1/event-tracker")
	if err != nil {
		fmt.Println("err:", err)
	}

	// config = YAML.load_file("#{__dir__}/_config.yml")
	// config['description'] = "#{race_status['race']['meetingCountryName']}: #{raceStart.strftime('%d %B %Y')}"
	// File.write("#{__dir__}/_config.yml", config.to_yaml)

	for _, event := range race.SeasonContext.Timetables {
		if event.Description == "Race" {
			switch event.State {
			case "upcoming":
				race.CurrentRaceStatus = "The race hasn't started yet"
			case "started":
				race.CurrentRaceStatus = "They're racing now"
			case "completed":
				race.CurrentRaceStatus = "They're done racing"
				for _, result := range race.RaceResults {
					position, err := strconv.Atoi(result.PositionNumber)
					if err != nil {
						fmt.Println("err:", err)
					}
					if position == 1 {
						race.Winner = result.DriverTLA
					}
				}
			}
		}
	}

	fmt.Println("state:", race.CurrentRaceStatus)
	fmt.Println("winner:", race.Winner)
	if race.Winner == "HAM" {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}

	// File.write("#{__dir__}/index.md", "# #{answer}")
	// File.write("#{__dir__}/api_results/#{Time.now.strftime('%Y%m%d-%H%M%S')}.json", JSON.pretty_generate(race_status))

	// if g.status.changed.count.positive?
	// 	g.add(all: true)
	// 	g.commit_all("updating for #{race_status['race']['meetingCountryName']}: #{raceStart.strftime('%d %B %Y')}; race #{raceinfo['state']}")
	// 	g.push(g.remote('origin'), branch)
	// end
}
