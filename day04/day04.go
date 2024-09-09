package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func main() {
	file, ferr := os.Open("data.txt")
	if ferr != nil {
		panic(ferr)
	}

	scanner := bufio.NewScanner(file)

	regCard := regexp.MustCompile(`[:]`)
	regPipe := regexp.MustCompile(`[|]`)

	gamePoint := 0
	numGame := 1

	for scanner.Scan() {

		points := 0

		line := scanner.Text()

		cardSlice := regCard.Split(line, -1)
		gameSlice := regPipe.Split(cardSlice[1], -1)

		winnings := gameSlice[0]
		chances := gameSlice[1]

		winningsSlice := strings.Fields(winnings)
		chancesSlice := strings.Fields(chances)

		for _, chance := range chancesSlice {
			for _, win := range winningsSlice {
				if chance == win && points == 0 {
					points = 1
				} else if chance == win {
					points = points * 2
				}
			}
		}
		numGame++
		gamePoint = gamePoint + points

	}

	fmt.Printf("Points are %d\n", gamePoint)
	part2()
}

func part2() {
	file, ferr := os.Open("data.txt")
	if ferr != nil {
		panic(ferr)
	}

	scanner := bufio.NewScanner(file)

	regCard := regexp.MustCompile(`[:]`)
	regPipe := regexp.MustCompile(`[|]`)

	numGame := 0
	totCards := 0

	var cards [213][]int

	for scanner.Scan() {

		points := 0

		line := scanner.Text()

		cardSlice := regCard.Split(line, -1)
		gameSlice := regPipe.Split(cardSlice[1], -1)

		winnings := gameSlice[0]
		chances := gameSlice[1]

		winningsSlice := strings.Fields(winnings)
		chancesSlice := strings.Fields(chances)

		cards[numGame] = append(cards[numGame], 1)

		for _, chance := range chancesSlice {
			for _, win := range winningsSlice {
				if chance == win {
					points = points + 1
				}
			}
		}

		for i := 0; i < len(cards[numGame]); i++ {
			for i := 1; i <= points; i++ {
				cards[numGame+i] = append(cards[numGame+i], 1)
			}
		}

		totCards = totCards + len(cards[numGame])
		numGame++
	}

	fmt.Printf("Total number of cards are %d\n", totCards)
}
