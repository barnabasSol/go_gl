package gamelogic

import "fmt"

type Level uint8

type ScoreTrack struct {
	Points         int
	KillCount      int
	PassedMaxLevel int
}

func NewScore(killcount, passedmaxlevel int) *ScoreTrack {
	return &ScoreTrack{
		KillCount:      killcount,
		PassedMaxLevel: passedmaxlevel,
	}

}

func (score *ScoreTrack) UpdateScore() {

}

func PrintGameStats(scoreTrack *ScoreTrack) {
	fmt.Println("\n┌───────────── Game Stats ─────────────┐")
	fmt.Printf("│ Kill Count:       %-20d │\n", scoreTrack.KillCount)
	fmt.Printf("│ Points:           %-20d │\n", scoreTrack.Points)
	fmt.Printf("│ Max Level Passed: %-20d │\n", scoreTrack.PassedMaxLevel)
	fmt.Println("└─────────────────────────────────────┘")
}
