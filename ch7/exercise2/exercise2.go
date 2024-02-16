package main

import (
	"fmt"
	"sort"
)

type Team struct {
	Name    string
	Players []string
}

type League struct {
	Teams map[string]Team
	Wins  map[string]int
}

func (l *League) MatchResult(team1 string, team1_score int, team2 string, team2_score int) {
	if _, ok := l.Teams[team1]; !ok {
		return
	}

	if _, ok := l.Teams[team2]; !ok {
		return
	}

	if team1_score > team2_score {
		l.Wins[team1]++
	} else if team1_score < team2_score {
		l.Wins[team2]++
	} else if team1_score == team2_score {
		return
	}
}

func (l League) Ranking() []string {
	names := make([]string, 0, len(l.Teams))
	for k := range l.Teams {
		names = append(names, k)
	}

	sort.Slice(names, func(i, j int) bool {
		return l.Wins[names[i]] > l.Wins[names[j]]
	})

	return names
}

func main() {
	demoLeague := League{
		Teams: map[string]Team{
			"Team 1": {
				Name: "Team 1",
			},
			"Team 2": {
				Name: "Team 2",
			},
			"Team 3": {
				Name: "Team 3",
			},
			"Team 4": {
				Name: "Team 4",
			},
		},
		Wins: map[string]int{},
	}

	// Game 1
	demoLeague.MatchResult("Team 1", 5, "Team 2", 2)

	// Game 2
	demoLeague.MatchResult("Team 4", 1, "Team 3", 7)

	// Game 3
	demoLeague.MatchResult("Team 1", 9, "Team 4", 5)

	// Game 4
	demoLeague.MatchResult("Team 2", 4, "Team 3", 5)

	// Game 5
	demoLeague.MatchResult("Team 1", 9, "Team 3", 6)

	// Results
	results := demoLeague.Ranking()
	fmt.Println(results)
}
