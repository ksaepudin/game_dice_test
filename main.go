package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type listPlayer struct {
	Player string
	Dice   []int
	Score  int
}

var onlyOnce sync.Once

// prepare the dice
var dice = []int{1, 2, 3, 4, 5, 6}

func rollDice() int {
	onlyOnce.Do(func() {
		rand.Seed(time.Now().UnixNano()) // only run once
	})

	return dice[rand.Intn(len(dice))]
}

func main() {
	setGame(4, 10)
}

func toCharStr(i int) string {
	return string('A' - 1 + i)
}

func setGame(pemain, dadu int) {
	var (
		rounde int
	)

	fmt.Println("Pemain = ", pemain, ", Dadu = ", dadu)
	player := setDice(pemain, dadu)
	for {
		// start Gamae
		if rounde != 0 {
			player = setDiceNext(player)
		}
		rounde++

		fmt.Println("================================================")
		fmt.Println("Giliran ", rounde, " lempar dadu:")
		for _, v := range player {
			fmt.Print("       Pemain #", v.Player, " (", v.Score, "): ")
			if len(v.Dice) != 0 {
				for _, z := range v.Dice {
					fmt.Print(" ", z)
				}
			} else {
				fmt.Print("_ (Berhenti bermain karena tidak memiliki dadu)")
			}
			fmt.Println("")
		}

		resEvalu := evaluasi(player, rounde)
		fmt.Println("Setelah evaluasi rounde : ", rounde)
		for _, v := range resEvalu {
			fmt.Print("       Pemain #", v.Player, " (", v.Score, "): ")
			if len(v.Dice) != 0 {
				for _, z := range v.Dice {
					fmt.Print(" ", z)
				}
			} else {
				fmt.Print("_ (Berhenti bermain karena tidak memiliki dadu)")
			}
			fmt.Println("")
		}
		player = resEvalu
		if findWinner(resEvalu) {
			player = nil
			fmt.Println("================================================")
			break
		}
	}
}

func setDice(pemain, dadu int) []listPlayer {
	var player []listPlayer
	for i := 0; i < pemain; i++ {
		p := toCharStr(i + 1)
		var dice []int
		for z := 0; z < dadu; z++ {
			dice1 := rollDice()
			dice = append(dice, dice1)
		}
		player = append(player, listPlayer{
			Player: p,
			Dice:   dice,
		})
	}
	return player
}

func setDiceNext(player []listPlayer) []listPlayer {
	var playerRES []listPlayer
	for i := 0; i < len(player); i++ {
		p := toCharStr(i + 1)
		var dice []int
		for z := 0; z < len(player[i].Dice); z++ {
			dice1 := rollDice()
			dice = append(dice, dice1)
		}
		playerRES = append(playerRES, listPlayer{
			Player: p,
			Dice:   dice,
			Score:  player[i].Score,
		})
	}
	return playerRES
}

func evaluasi(in []listPlayer, score int) []listPlayer {
	var (
		Dadu6     []int
		DaduState []int
		out       []listPlayer
		count     int
	)
	for _, v := range in {
		player := listPlayer{}
		player.Player = v.Player
		player.Score = v.Score
		for _, z := range v.Dice {
			if z == 6 {
				Dadu6 = append(Dadu6, z)

			} else {
				player.Dice = append(player.Dice, z)
			}
		}
		if count != 0 {
			if len(DaduState) != 0 {
				for _, _ = range DaduState {
					player.Dice = append(player.Dice, 1)
				}
				player.Score = player.Score + 1
				DaduState = []int{}
				if len(Dadu6) != 0 {
					DaduState = Dadu6
					Dadu6 = []int{}
				}
				a := len(in) - 1
				if count == (a) {
					for _, _ = range DaduState {
						out[0].Dice = append(out[0].Dice, 1)
					}
					out[0].Score = out[0].Score + 1
					Dadu6 = []int{}
					DaduState = Dadu6
				}
			}
		} else {
			DaduState = Dadu6
			Dadu6 = []int{}
		}
		count++
		out = append(out, player)
	}
	return out
}

func findWinner(in []listPlayer) bool {
	var (
		player []string
		score  int
		pemain string
	)

	for _, v := range in {
		if len(v.Dice) != 0 {
			player = append(player, v.Player)
		}

		if score < v.Score {
			score = v.Score
			pemain = v.Player
		}
	}
	if len(player) == 1 {
		fmt.Println("================================================")
		fmt.Println("Winner = Player #", pemain)
		return true
	}
	return false
}
