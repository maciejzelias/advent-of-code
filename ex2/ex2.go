package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {

	file, err := os.Open("../ex2/data.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	games := readDataToArray(file)

	max_red := 12
	max_green := 13
	max_blue := 14

	impossibleGames := 0

	possibleGamesFewestSum := 0

	for _, game := range games {
		impossibleCount := 0
		red := 0
		blue := 0
		green := 0
		for _, drop := range game.drops {
			for _, stepDrop := range drop {
				switch stepDrop.color {
				case "red":
					if stepDrop.count > max_red {
						impossibleCount++
					}
					if stepDrop.count > red {
						red = stepDrop.count
					}
				case "green":
					if stepDrop.count > max_green {
						impossibleCount++
					}
					if stepDrop.count > green {
						green = stepDrop.count
					}
				case "blue":
					if stepDrop.count > max_blue {
						impossibleCount++
					}
					if stepDrop.count > blue {
						blue = stepDrop.count
					}
				}
			}
		}
		possibleGamesFewestSum += red * green * blue
		if impossibleCount == 0 {
			impossibleGames += game.id
		}
	}
	println(impossibleGames)
	println(possibleGamesFewestSum)

}

type Drop struct {
	count int
	color string
}

type Game struct {
	id    int
	drops [][]Drop
}

func readDataToArray(file *os.File) []Game {
	var games []Game
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, ":")

		id, err := strconv.Atoi(strings.Split(parts[0], " ")[1])

		if err != nil {
			log.Fatal(err)
		}

		dropsStr := strings.Split(parts[1], ";")

		var drops [][]Drop

		for _, drop := range dropsStr {
			dropParts := strings.Split(drop, ",")
			var stepDrops []Drop
			for _, dropPart := range dropParts {
				trimmedDropPart := strings.Split(strings.TrimSpace(dropPart), " ")
				count, err := strconv.Atoi(trimmedDropPart[0])
				if err != nil {
					log.Fatal(err)
				}
				stepDrops = append(stepDrops, Drop{
					count: count,
					color: trimmedDropPart[1],
				})
			}
			drops = append(drops, stepDrops)
		}

		games = append(games, Game{
			id:    id,
			drops: drops,
		})
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return games
}
