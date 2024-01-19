package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func build2DArray(input string, lineNumber int) [][]string {
	res := make([][]string, lineNumber)
	for i := range res {
			res[i] = make([]string, lineNumber)
	}
	lines := strings.Split(input, "\n")
	for i, line := range lines {
		for j, li := range line {
			res[i][j] = string(li)
		}
	}
	return res
}

func isNumeric(str string) bool {
	_, err := strconv.Atoi(str)
	return err == nil
}

func process(input string) int {
	cur_num := 0
	total := 0
	isAdjacent := false
	lines := strings.Split(input, "\n")
	arr := build2DArray(input, len(lines))
	for i, line := range lines {
		for j, li := range line {
			c := rune(li)
			if unicode.IsDigit(c) {
				// loop and check around current position
				for _, cr := range []int{-1,0,1} {
					for _, cc := range []int{-1,0,1} {
						if (i+cr)>=0 && (i+cr)<len(lines) && (j+cc)>=0 && (j+cc)<len(lines) {
							cur_check := arr[i+cr][j+cc]
							if !isNumeric(cur_check) && string(cur_check) != "."  {
								isAdjacent = true
							}
						}
					}
				}
				n, _ := strconv.Atoi(string(li))
				cur_num = cur_num*10 + n
				// handle last char is number
				if j == len(line)-1 {
					if cur_num != 0 {
						if isAdjacent {
							total += cur_num
						}
						cur_num = 0
						isAdjacent = false
					}
				}
			} else {
				if cur_num != 0 {
					if isAdjacent {
						fmt.Println("RESULT: ", cur_num, total)
						total += cur_num
					}
					cur_num = 0
					isAdjacent = false
				}
			}
		}
		cur_num = 0
		isAdjacent = false
		// fmt.Println()
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

	fmt.Println(res)
}