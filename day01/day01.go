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

	regex := regexp.MustCompile(`\D+`) 

	var acc int

	for scanner.Scan() {
		line := scanner.Text()
		reg := regex.ReplaceAllString(line, "")
		snum := string(reg[0])+string(reg[len(reg)-1])

		   num, err := strconv.Atoi(snum)
		   if err != nil {
			   panic(err)
		   }
		acc = acc+num
	}
	fmt.Println(acc)	

	part2()
}

func part2(){
	file, ferr := os.Open("data.txt")

	if ferr != nil {
		panic(ferr)
	}

	scanner := bufio.NewScanner(file)

	numbers := map[string]int{
		"one": 1,
		"two":   2,
		"three": 3,
		"four": 4,
		"five": 5,
		"six": 6,
		"seven": 7,
		"eight": 8,
		"nine": 9,
	}

	var acc int

	for scanner.Scan() {
		line := scanner.Text()

		firstIndex := 100
		firstNumber := 0

		lastIndex := 0
		lastNumber := 0

		for stringNum, numNum := range numbers {
			resString := strings.Index(line, stringNum)
			resLastString := strings.LastIndex(line, stringNum)
			resNum := strings.Index(line, strconv.Itoa(numNum))
			resLastNum := strings.LastIndex(line, strconv.Itoa(numNum))


			if ( -1 != resString && resString <= firstIndex ){
				firstIndex = resString
				firstNumber = numNum
			}
			if (-1 != resNum && resNum <= firstIndex ){
				firstIndex = resNum
				firstNumber = numNum
			}
			if ( -1 != resLastString && resLastString >= lastIndex ){
				lastIndex = resLastString
				lastNumber = numNum
			}
			if (-1 != resLastNum && resLastNum >= lastIndex ){
				lastIndex = resLastNum
				lastNumber = numNum
			}
		}


		together := strconv.Itoa(firstNumber)+strconv.Itoa(lastNumber)
		togetherNum, err := strconv.Atoi(together)
		if err != nil {
			panic(err)
		}

		acc = acc + togetherNum
	}
	fmt.Println(acc)

}
