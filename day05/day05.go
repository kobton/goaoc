package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

var seedSlice []string
var seedToSoilSlice [][]string
var soilToFertilizerSlice [][]string
var fertilizerToWaterSlice [][]string
var waterToLightSlice [][]string
var lightToTemperatureSlice [][]string
var temperatureToHumiditySlice [][]string
var humidityToLocationSlice [][]string

func main() {
	createMaps()
	part1()
	// part2()

}

func part1(){
	numSeeds, err := convertStringSeeds(seedSlice)
	if err != nil {
		fmt.Printf("error",err)
		return
	}
	var locations []int
	for _, num := range numSeeds{
		locationNum := passNumber(num)
		locations = append(locations, locationNum)
	}

	minLocation , err := MinIntSlice(locations)
	if err != nil{
		return
	}

	fmt.Printf("The min location part1 is: %d\n", minLocation)
}

func part2(){
	numSeeds, err := convertStringSeeds(seedSlice)
	if err != nil {
		fmt.Printf("error",err)
		return
	}

	var allSeeds []int
	for i := 0 ; i < len(numSeeds)-1; {
		for j := 0 ; j < numSeeds[i+1] ; j++ {
			allSeeds = append(allSeeds, numSeeds[i]+j)
		}
		i = i+2
 	}


	var locations []int
	for _, num := range allSeeds{
		locationNum := passNumber(num)
		locations = append(locations, locationNum)
	}

	minLocation , err := MinIntSlice(locations)
	if err != nil{
		return
	}

	fmt.Printf("The min location part2 is: %d\n", minLocation)
}

func MinIntSlice(numbers []int) (int, error) {
    if len(numbers) == 0 {
        return 0, fmt.Errorf("slice is empty")
    }

    min := math.MaxInt

    for _, num := range numbers {
        if num < min {
            min = num
        }
    }

    return min, nil
}

func passNumber(num int) (int) {
	maps := [][][]string{seedToSoilSlice, soilToFertilizerSlice, fertilizerToWaterSlice, waterToLightSlice, lightToTemperatureSlice, temperatureToHumiditySlice, humidityToLocationSlice}

	tempLocation := num

	for _,stringSlice := range maps{
		numSlice, err := convertStringSlice(stringSlice)
		if err != nil{
			return 0
		}

		tempLocation = findLocation(tempLocation, numSlice)
	}
	return tempLocation
}

func createMaps(){
	file, ferr := os.Open("data.txt")
	if ferr != nil {
		panic(ferr)
	}

	scanner := bufio.NewScanner(file)

	lineIndex := 0
	currentMap := ""

	for scanner.Scan() {

		line := scanner.Text()

		if(lineIndex == 0){
			seedString := line[6:]
			seedSlice = strings.Fields(seedString)
		}

		if(line == ""){
			currentMap = ""
		}

		switch currentMap {
		case "seedToSoil":
			seedToSoilSlice = append(seedToSoilSlice, strings.Fields(line))
		case "soilToFertilizer":
			soilToFertilizerSlice = append(soilToFertilizerSlice, strings.Fields(line))
		case "fertilizerToWater":
			fertilizerToWaterSlice = append(fertilizerToWaterSlice, strings.Fields(line))
		case "waterToLight":
			waterToLightSlice = append(waterToLightSlice, strings.Fields(line))
		case "lightToTemperature":
			lightToTemperatureSlice = append(lightToTemperatureSlice, strings.Fields(line))
		case "temperatureToHumidity":
			temperatureToHumiditySlice = append(temperatureToHumiditySlice, strings.Fields(line))
		case "humidityToLocation":
			humidityToLocationSlice = append(humidityToLocationSlice, strings.Fields(line))
		default:
		}


		if(strings.Contains(line, "seed-to-soil")){
			currentMap = "seedToSoil"
		}
		if(strings.Contains(line, "soil-to-fertilizer")){
			currentMap = "soilToFertilizer"
		}
		if(strings.Contains(line, "fertilizer-to-water")){
			currentMap = "fertilizerToWater"
		}
		if(strings.Contains(line, "water-to-light")){
			currentMap = "waterToLight"
		}
		if(strings.Contains(line, "light-to-temperature")){
			currentMap = "lightToTemperature"
		}
		if(strings.Contains(line, "temperature-to-humidity")){
			currentMap = "temperatureToHumidity"
		}
		if(strings.Contains(line, "humidity-to-location")){
			currentMap = "humidityToLocation"
		}
		lineIndex++
	}

}

func findLocation(seed int, sliceMap[][]int) (int)  {
	newNum := 0
	for _, mapRow := range sliceMap {
		if (seed >= mapRow[1] && seed < mapRow[1]+mapRow[2]){
			newNum = seed + mapRow[0] - mapRow[1]
		} 
	}
	if(newNum == 0){
		newNum = seed
	}
	return newNum
 }

 func convertStringSlice(slice [][]string) ([][]int, error){

	intSlice := make([][]int, len(slice))

	for i, row := range slice{
		intSlice[i] = make([]int, len(row))
		for j, element := range row{
			numElement, err := strconv.Atoi(element)
			if err != nil{
				return intSlice, err
			}
			intSlice[i][j] = numElement
		}
	}
	return intSlice, nil
 }

 func convertStringSeeds(slice []string) ([]int, error){
	var intSlice []int
	for _, element := range slice{
			numElement, err := strconv.Atoi(element)
			if err != nil{
				return nil, err
			}
			intSlice = append(intSlice, numElement)
	}
	return intSlice, nil
 }

