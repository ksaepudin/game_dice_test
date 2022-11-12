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
	setGame(3, 4)
}

func toCharStr(i int) string {
	return string('A' - 1 + i)
}

func setGame(pemain, dadu int) {
	var (
		rounde int
		score  int
		// player []listPlayer
	)

	fmt.Println("Pemain = ", pemain, ", Dadu = ", dadu)
	player := setDice(pemain, dadu)
	for {
		rounde++
		// start Gamae

		fmt.Println("================================================")
		fmt.Println("Giliran ", rounde, " lempar dadu:")
		for _, v := range player {
			fmt.Print("       Pemain #", v.Player, " (", score, "): ")
			for _, z := range v.Dice {
				fmt.Print(" ", z)
			}
			fmt.Println("")

		}
		resEvalu := evaluasi(player, rounde)
		fmt.Println("Setelah evaluasi rounde : ", rounde)
		for _, v := range resEvalu {
			fmt.Print("       Pemain #", v.Player, " (", v.Score, "): ")
			for _, z := range v.Dice {
				fmt.Print(" ", z)
			}
			fmt.Println("")

		}
		player = resEvalu
		if rounde > 20 {
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

func evaluasi(in []listPlayer, score int) []listPlayer {
	var (
		Dadu6 []int
		out   []listPlayer
		count int
	)
	for _, v := range in {
		player := listPlayer{}
		player.Player = v.Player
		for _, z := range v.Dice {
			if z == 6 {
				Dadu6 = append(Dadu6, z)
			} else {
				player.Dice = append(player.Dice, z)
			}
		}
		if count != 0 {
			if len(Dadu6) != 0 {
				for _, _ = range Dadu6 {
					player.Dice = append(player.Dice, 1)
				}
				player.Score = v.Score + 1
				Dadu6 = []int{}
			}
		}

		if count == (len(in) - 1) {
			if len(Dadu6) != 0 {
				for _, _ = range Dadu6 {
					out[0].Dice = append(out[0].Dice, 1)
				}
				out[0].Score = out[0].Score + 1
				Dadu6 = []int{}
			}
		}
		count++
		out = append(out, player)
	}
	return out
}
