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

	count := countHorizontalXmases(topMap)
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

func countChristmases() {}

func countHorizontalXmases(topMap map[int]map[int]byte) int {
	count := 0
	for i, line := range topMap {
		for j := 0; j < len(line)-3; j++ {
			arr := []byte{line[j], line[j+1], line[j+2], line[j+3]}
			if isXmas(arr) {
				fmt.Println("Found an Xmas in Line ", i)
				count++
			}
		}
	}
	return count
}

func isXmas(arr []byte) bool {
	xmas := []byte("XMAS")
	return bytes.Equal(xmas, arr)
}
