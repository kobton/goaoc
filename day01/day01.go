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
}