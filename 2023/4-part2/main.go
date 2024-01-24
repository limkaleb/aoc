package main

import (
	"fmt"
	"log"
	"os"
	"reflect"
	"slices"
	"sort"
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
	copyMap := make(map[int]int)

	lines := strings.Split(input, "\n")
	fmt.Println("lines: ", len(lines))
	
	for i := 0 ; i < len(lines); i++ {
		cards := strings.Split(lines[i], ":")
		separate := strings.Split(cards[1], "|")

		winMap[i] = buildMap(separate[0])
		haveMap[i] = buildMap(separate[1])

		// fmt.Println("current: ", i, winMap[i], haveMap[i])
	}

	keys := reflect.ValueOf(haveMap).MapKeys()
	keysOrder := func(i, j int) bool { return keys[i].Interface().(int) < keys[j].Interface().(int) }
	sort.Slice(keys, keysOrder)

	// process map in key-sorted order
	for i, key := range keys {
		value := haveMap[key.Interface().(int)]
		// fmt.Println("keyval: ", key, value)
		copyMap[i] += 1
		for _, currentValue := range value {
			if slices.Contains(winMap[i], currentValue) {
				score += 1
			}
		}

		for j := i+1; j < i+score+1; j++ {
			copyMap[j] += copyMap[i]
			// fmt.Println("containnn: ", j, score, copyMap[j])
		}
	
		// fmt.Println("current status after: ", i, copyMap)
		score = 0
	}

	// calculate total
	for _, val := range copyMap {
		total += val
	}

	return total
}

func main() {
	// input, err := os.ReadFile("example.txt")
	input, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	// fmt.Println(string(input))
	res := process(string(input))
	fmt.Println("result: ", res)
}