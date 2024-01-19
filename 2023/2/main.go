package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func findMinSet(input string) int {
	var total int
	rounds := strings.Split(input, "\n")
	fmt.Println("rounds len: ", len(rounds))

	for _, round := range rounds {
		// cur_game := i + 1
		multiply_res := 0
		red := 0
		green := 0
		blue := 0

		max_red := 0
		max_green := 0
		max_blue := 0 
		
		sets := strings.Split(round, ":")
		plays := strings.Split(sets[1], ";")

		for _, play := range plays {
			cubes := strings.Split(play, ",")
			for _, cube := range cubes {
				gem := strings.Split(cube, " ")
				switch gem[2] {
				case "red":
					red, _ = strconv.Atoi(gem[1])
				case "green":
					green, _ = strconv.Atoi(gem[1])
				case "blue":
					blue, _ = strconv.Atoi(gem[1])
				}
			}

			if max_red < red {
				max_red = red
			}

			if max_green < green {
				max_green = green
			}

			if max_blue < blue {
				max_blue = blue
			}
		}
		
		multiply_res = max_red * max_green * max_blue
		total += multiply_res
	}
	fmt.Println("result totall: ", total)
	return 0
}

func process(input string) int {
	var total int
	rounds := strings.Split(input, "\n")
	fmt.Println("rounds len: ", len(rounds))

	for i, round := range rounds {
		cur_game := i + 1
		red := 0
		green := 0
		blue := 0
		
		sets := strings.Split(round, ":")
		plays := strings.Split(sets[1], ";")

		for j, play := range plays {
			cubes := strings.Split(play, ",")
			for _, cube := range cubes {
				gem := strings.Split(cube, " ")
				switch gem[2] {
				case "red":
					red, _ = strconv.Atoi(gem[1])
				case "green":
					green, _ = strconv.Atoi(gem[1])
				case "blue":
					blue, _ = strconv.Atoi(gem[1])
				}
			}
			if (red > 12 || green > 13 || blue > 14) {
				break
			}
			if j == len(plays) - 1 {
				total += cur_game
			}
		}
	}
	fmt.Println("totalll : ", total)
	return 0
}

func main() {
	// input, err := os.ReadFile("example.txt")
	input, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(input))

	// res := process(string(input))
	res := findMinSet(string(input))

	fmt.Println(res)
}