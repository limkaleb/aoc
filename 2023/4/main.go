package main

import (
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

func buildMap(inputString string) []int {
	haveTrimmed := strings.Trim(inputString, " ")
		nums := strings.Split(haveTrimmed, " ")
		var sliceInts []int
		for k := 0; k < len(nums); k++ {
			if nums[k] == "" {
				continue
			} else {
				con, _ := strconv.Atoi(nums[k])
				sliceInts = append(sliceInts, con)
			}
		}
		return sliceInts
}

func process(input string) int {
	total := 0
	winMap := make(map[int][]int)
	haveMap := make(map[int][]int)
	score := 0
	lines := strings.Split(input, "\n")
	// fmt.Println("lines: ", len(lines))
	
	for i := 0 ; i < len(lines); i++ {
		cards := strings.Split(lines[i], ":")
		separate := strings.Split(cards[1], "|")

		winMap[i] = buildMap(separate[0])
		haveMap[i] = buildMap(separate[1])

		// fmt.Println("current: ", i, winMap[i], haveMap[i])
	}

	for k, v := range haveMap {
		fmt.Println("keyval: ", k, v)
		for _, currentValue := range v {
			contain := slices.Contains(winMap[k], currentValue)
			if contain {
				if score == 0 {
					score = 1
				} else if score > 0 {
					score *= 2
				}
			} 
			// fmt.Println("containnn: ", currentValue, contain, score)
		}
		fmt.Println()
		total += score
		score = 0
	}

	return total
}

func main() {
	input, err := os.ReadFile("example.txt")
	// input, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(input))
	res := process(string(input))
	fmt.Println("result: ", res)
}