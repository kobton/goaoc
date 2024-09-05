package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	file, ferr := os.Open("data.txt")
	if ferr != nil {
		panic(ferr)
	}
	scanner := bufio.NewScanner(file)

	symReg := regexp.MustCompile(`[^0-9.]`)
	numReg := regexp.MustCompile(`\d+`)

	var numbers [][][]int
	var symbols [][][]int
	var lines []string

	i := 0

	for scanner.Scan() {

		line := scanner.Text()
		lines = append(lines, line)

		numberIndexSets := numReg.FindAllStringIndex(line, -1)

		numbers = append(numbers, numberIndexSets)

		symbolIndexSets := symReg.FindAllStringIndex(line, -1)

		symbols = append(symbols, symbolIndexSets)

		i++
	}
	// fmt.Printf("The numbers %v\n", numbers)
	// fmt.Printf("The symbols %v\n", symbols)

	// fmt.Printf("Tot rowindex %d\n", i)

	row := 0
	count := 0
	sum := 0

	for _, numRow := range numbers {
		//check same row
		for _, symPair := range symbols[row] {
			for _, numPair := range numRow {
				if symPair[0] == numPair[0]-1 || numPair[1] == symPair[0] {
					num, err := strconv.Atoi(lines[row][numPair[0]:numPair[1]])
					if err != nil {
						fmt.Println("Error:", err)
						return
					}
					sum = sum + num
					count++
				}
			}
		}
		//check row above
		if row > 0 {
			for _, symPair := range symbols[row-1] {
				for _, numPair := range numRow {
					if symPair[0] >= numPair[0]-1 && numPair[1] >= symPair[0] {
						num, err := strconv.Atoi(lines[row][numPair[0]:numPair[1]])
						if err != nil {
							fmt.Println("Error:", err)
							return
						}
						sum = sum + num
						count++
					}
				}
			}
		}
		//check row below
		if row < len(lines)-1 {
			for _, symPair := range symbols[row+1] {
				for _, numPair := range numRow {
					if symPair[0] >= numPair[0]-1 && numPair[1] >= symPair[0] {
						num, err := strconv.Atoi(lines[row][numPair[0]:numPair[1]])
						if err != nil {
							fmt.Println("Error:", err)
							return
						}
						sum = sum + num
						count++
					}
				}
			}
		}
		row++
	}
	fmt.Printf("Count: %d\n", count)
	fmt.Printf("Sum: %d\n", sum)
	part2()

}

func part2() {
	file, ferr := os.Open("data.txt")
	if ferr != nil {
		panic(ferr)
	}
	scanner := bufio.NewScanner(file)

	symReg := regexp.MustCompile(`\*`)
	numReg := regexp.MustCompile(`\d+`)

	var numbers [][][]int
	var symbols [][][]int
	var lines []string

	i := 0

	for scanner.Scan() {

		line := scanner.Text()
		lines = append(lines, line)

		numberIndexSets := numReg.FindAllStringIndex(line, -1)

		numbers = append(numbers, numberIndexSets)

		symbolIndexSets := symReg.FindAllStringIndex(line, -1)

		symbols = append(symbols, symbolIndexSets)

		i++
	}
	// fmt.Printf("The numbers %v\n", numbers)
	// fmt.Printf("The symbols %v\n", symbols)

	// fmt.Printf("Tot rowindex %d\n", i)

	row := 0
	count := 0
	sum := 0

	for _, _symRow := range symbols {
		for _, symPair := range _symRow {
			one := 1
			two := 1
			three := 1
			four := 1

			var above []int
			var below []int


			if row < len(lines)-1 {
				for _, numBelowPair := range numbers[row+1] {
					if symPair[0] >= numBelowPair[0]-1 && numBelowPair[1] >= symPair[0] {
						num1, err := strconv.Atoi(lines[row+1][numBelowPair[0]:numBelowPair[1]])
						if err != nil {
							fmt.Println("Error:", err)
							return
						}
						below = append(below, num1)
					}
				}
				num1 := 1
				for _, hit := range below {
					num1 = num1*hit
					count++
				}

				one =  num1
			}
			if row > 0 {
				for _, numAbovePair := range numbers[row-1] {
					if symPair[0] >= numAbovePair[0]-1 && numAbovePair[1] >= symPair[0] {
						num1, err := strconv.Atoi(lines[row-1][numAbovePair[0]:numAbovePair[1]])
						if err != nil {
							fmt.Println("Error:", err)
							return
						}
						above = append(above, num1)
					}
				}
				num1 := 1
				for _, hit := range above {
					num1 = num1*hit
					count++
				}

				two =  num1
			}
			for _, numSamePair := range numbers[row] {
				if symPair[0] == numSamePair[0]-1 {
					num1, err := strconv.Atoi(lines[row][numSamePair[0]:numSamePair[1]])
					if err != nil {
						fmt.Println("Error:", err)
						return
					}
					count++
					three = num1
				}
			}
			for _, numSamePair2 := range numbers[row] {
				if numSamePair2[1] == symPair[0] {
					num1, err := strconv.Atoi(lines[row][numSamePair2[0]:numSamePair2[1]])
					if err != nil {
						fmt.Println("Error:", err)
						return
					}
					count++
					four = num1
				}
			}
			if count == 2 {
				sum = sum + one*two*three*four
			}
			count = 0
		}
		row++
	}

	fmt.Printf("Sum2: %d\n", sum)

}
