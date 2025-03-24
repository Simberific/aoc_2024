package main

import (
	"bufio"
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
	fmt.Println(topMap)
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
	}
	return topMap
}
