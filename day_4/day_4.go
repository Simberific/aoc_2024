package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Println("LFG")
	file, err := os.Open("day_4/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	topMap := inputToMap(scanner)
	fmt.Println("TopMap length: ", len(topMap))

	count := countChristmases(topMap)
	fmt.Println(count)
}

func inputToMap(scanner *bufio.Scanner) map[int]map[int]byte {
	topMap := make(map[int]map[int]byte)
	i := 0
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println("Line:", line)
		rowLength := len(line)
		lineMap := make(map[int]byte)
		for j := 0; j < rowLength; j++ {
			lineMap[j] = line[j]
		}
		topMap[i] = lineMap
		i++
	}
	return topMap
}

func countChristmases(topMap map[int]map[int]byte) int {
	return countHorizontalXmases(topMap) + countVerticalXmases(topMap)
}

func countHorizontalXmases(topMap map[int]map[int]byte) int {
	count := 0
	for i, line := range topMap {
		for j := 0; j < len(line)-3; j++ {
			arr := []byte{line[j], line[j+1], line[j+2], line[j+3]}
			if isXmas(arr) {
				fmt.Println("Found an Xmas in Row ", i)
				count++
			}
		}
	}
	return count
}

func countVerticalXmases(topMap map[int]map[int]byte) int {
	count := 0
	for columnIndex := 0; columnIndex < len(topMap[0]); columnIndex++ {
		// Look below me to grab the 4 vertically stacking bytes until I can't anymore.
		for rowIndex := 0; rowIndex < len(topMap)-3; rowIndex++ {
			arr := []byte{
				topMap[rowIndex][columnIndex],
				topMap[rowIndex+1][columnIndex],
				topMap[rowIndex+2][columnIndex],
				topMap[rowIndex+3][columnIndex],
			}
			if isXmas(arr) {
				fmt.Println("Found an Xmas in Column ", columnIndex)
				count++
			}
		}
	}

	return count
}

func isXmas(arr []byte) bool {
	xmas := []byte("XMAS")
	samx := []byte("SAMX")
	return bytes.Equal(xmas, arr) || bytes.Equal(samx, arr)
}
