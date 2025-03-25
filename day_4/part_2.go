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
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	topMap := inputToMap2(scanner)
	fmt.Println("TopMap length: ", len(topMap))

	count := countXChristmases(topMap)
	fmt.Println(count)
}

func inputToMap2(scanner *bufio.Scanner) map[int]map[int]byte {
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

func countXChristmases(topMap map[int]map[int]byte) int {
	count := 0
	for columnIndex := 0; columnIndex < len(topMap[0])-2; columnIndex++ {
		// Look below me to grab the X shaped bytes until I can't anymore.
		for rowIndex := 0; rowIndex < len(topMap)-2; rowIndex++ {
			arr1 := []byte{
				topMap[rowIndex][columnIndex],
				topMap[rowIndex+1][columnIndex+1],
				topMap[rowIndex+2][columnIndex+2],
			}
			arr2 := []byte{
				topMap[rowIndex][columnIndex+2],
				topMap[rowIndex+1][columnIndex+1],
				topMap[rowIndex+2][columnIndex],
			}

			if isSamOrMas(arr1) && isSamOrMas(arr2) {
				fmt.Println("Found an X Christmas starting in Column ", columnIndex)
				count++
			}
		}
	}
	return count
}

func isSamOrMas(arr []byte) bool {
	mas := []byte("MAS")
	sam := []byte("SAM")
	return bytes.Equal(sam, arr) || bytes.Equal(mas, arr)
}
