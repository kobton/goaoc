package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	file, ferr := os.Open("data.txt")
	if ferr != nil {
		panic(ferr)
	}

	scanner := bufio.NewScanner(file)

	re := regexp.MustCompile(`[;:]`)
	regex := regexp.MustCompile(`[^0-9.]`) 
	regSet := regexp.MustCompile(`[,]`)


	gamesAmount:= 0
	gameTotalPower := 0


	for scanner.Scan() {

		line := scanner.Text()
		noWhite := strings.ReplaceAll(line, " ", "")

		splitLine := re.Split(noWhite, -1)
		numGameString := string(splitLine[0])

		numnumGameString := regex.ReplaceAllString(numGameString, "")


		numGame, err := strconv.Atoi(numnumGameString)
		if err != nil {
			panic(err)
		}

		allSets := splitLine[1:]
		fmt.Printf("%q\n", allSets)

		countGame := true

		topGreenAmount := 0
		topRedAmount := 0
		topBlueAmount := 0


		for _, set := range allSets {

			var greenAmount int
			var redAmount int
			var blueAmount int	

			


			splitSet := regSet.Split(set, -1)

			for _, element := range splitSet {

			amountString := regex.ReplaceAllString(element, "")

			amountNum, err := strconv.Atoi(amountString)
			if err != nil {
				panic(err)
			}

			color := string(element[len(element)-1])

			if(color == "n"){
				greenAmount = greenAmount+amountNum
			}
			if(color == "d"){
				redAmount = redAmount + amountNum
			}
			if(color == "e"){
				blueAmount = blueAmount + amountNum
			}
			}
			fmt.Printf("The index: %d\n", numGame)
			fmt.Printf("The colors: %d, %d, %d\n", greenAmount, blueAmount, redAmount)
	
			if(blueAmount > 14 || greenAmount > 13 || redAmount > 12){
				countGame = false
			}

			if(blueAmount > topBlueAmount){
				topBlueAmount = blueAmount
			}
			if(redAmount > topRedAmount){
				topRedAmount = redAmount
			}
			if(greenAmount > topGreenAmount){
				topGreenAmount = greenAmount
			}
		}
		gamePower := topBlueAmount*topRedAmount*topGreenAmount

		gameTotalPower = gameTotalPower + gamePower


		if(countGame){
		gamesAmount = gamesAmount + numGame
		// fmt.Printf("New gamesAmount: %d\n", gamesAmount)
		}

	}
		

	fmt.Printf("Gamesamount part 1 : %d\n",gamesAmount)
	fmt.Printf("Total game power part 2 : %d\n",gameTotalPower)


}




