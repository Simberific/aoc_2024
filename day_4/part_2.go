package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Println("Part 2 LFG")
	file, err := os.Open("day_4/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	topMap := inputToMap(scanner)
	fmt.Println("TopMap length: ", len(topMap))

	count := countXChristmases(topMap)
	fmt.Println(count)
}

func countXChristmases(topMap map[int]map[int]byte) int {
	count := 0
	for columnIndex := 0; columnIndex < len(topMap[0])-3; columnIndex++ {
		// Look below me to grab the 4 vertically stacking bytes until I can't anymore.
		for rowIndex := 0; rowIndex < len(topMap)-3; rowIndex++ {
			arr := []byte{
				topMap[rowIndex][columnIndex],
				topMap[rowIndex+1][columnIndex+1],
				topMap[rowIndex+2][columnIndex+2],
				topMap[rowIndex+3][columnIndex+3],
			}
			if isXmas(arr) {
				fmt.Println("Found a forwards Diagonal Xmas starting in Column", columnIndex)
				count++
			}
		}
	}
	return count
}

func equalsXmas(arr []byte) bool {
	xmas := []byte("XMAS")
	return bytes.Equal(xmas, arr)
}

func equalsSamx(arr []byte) bool {
	samx := []byte("SAMX")
	return bytes.Equal(samx, arr)
}
