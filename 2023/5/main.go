package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func process(input string) int {
	res_location := 9999999999999
	// index of first line and last line of each map
	line_map := map[int][]int{
		0: {0, 0},
		1: {0, 0},
		2: {0, 0},
		3: {0, 0},
		4: {0, 0},
		5: {0 ,0},
		6: {0, 0},
	}

	lines := strings.Split(input, "\n")
	firstLine := strings.Split(lines[0], ":")
	seedsString := firstLine[1]
	fmt.Println("seeds: ", seedsString)

	// build slice of seeds
	seedz := strings.Split(strings.Trim(seedsString, " "), " ")
	// fmt.Println("seedzzz: ", seedz, len(seedz))
	// fmt.Println("linezz LAST: ", len(lines)-1, lines[len(lines)-1])

	for i, line := range lines {
		switch line {
			case "seed-to-soil map:":
				line_map[0][0] = i+1
			case "soil-to-fertilizer map:":
				line_map[0][1] = i-1
				line_map[1][0] = i+1
			case "fertilizer-to-water map:":
				line_map[1][1] = i-1
				line_map[2][0] = i+1
			case "water-to-light map:":
				line_map[2][1] = i-1
				line_map[3][0] = i+1
			case "light-to-temperature map:":
				line_map[3][1] = i-1
				line_map[4][0] = i+1
			case "temperature-to-humidity map:":
				line_map[4][1] = i-1
				line_map[5][0] = i+1
			case "humidity-to-location map:":
				line_map[5][1] = i-1
				line_map[6][0] = i+1
				line_map[6][1] = len(lines)
		}
	}

	for _, seed := range seedz {
		// fmt.Println("CURRENT SEED: ", seed)
		test_seed, _ := strconv.Atoi(seed)
		test_soil := 0
		test_fertilizer := 0
		test_water := 0
		test_light := 0
		test_temp := 0
		test_hum := 0
		test_loc := 0

		for i := 0; i < len(line_map); i++ {
			for k := line_map[i][0]; k < line_map[i][1]; k++ {
				nums_str := strings.Split(lines[k], " ")
				cur_num := [3]int{}
				for ii := 0; ii < len(nums_str); ii++ {
					cur_num[ii], _ = strconv.Atoi(nums_str[ii])
				}
				// fmt.Println("cur_nummmm: ", cur_num[0], cur_num[1], cur_num[2])
				if i == 0 { // seed to soil
					if test_soil > 0 {
						continue
					} else if cur_num[1] <= test_seed && (cur_num[1] + cur_num[2]-1) >= test_seed {
						test_soil = cur_num[0] + (test_seed - cur_num[1])
					} else {
						if test_soil == 0 && k == line_map[i][1]-1 {
							test_soil = test_seed
						}
					}
					// fmt.Println("SOIL: ", test_soil)
				} else if i == 1 { // soil to fertilizer
					if test_fertilizer > 0 {
						continue
					} else if cur_num[1] <= test_soil && (cur_num[1] + cur_num[2]-1) >= test_soil {
						test_fertilizer = cur_num[0] + (test_soil - cur_num[1])
					} else {
						if test_fertilizer == 0 && k == line_map[i][1]-1 {
							test_fertilizer = test_soil
						}
					}
					// fmt.Println("FERTILIZER: ", test_fertilizer)
				} else if i == 2 { // fertilizer to water
					if test_water > 0 {
						continue
					} else if cur_num[1] <= test_fertilizer && (cur_num[1] + cur_num[2]-1) >= test_fertilizer {
						test_water = cur_num[0] + (test_fertilizer - cur_num[1])
					} else {
						if test_water == 0 && k == line_map[i][1]-1 {
							test_water = test_fertilizer
						}
					}
					// fmt.Println("WATERRR: ", test_water)
				} else if i == 3 { // water to light
					if test_light > 0 {
						continue
					} else if cur_num[1] <= test_water && (cur_num[1] + cur_num[2]-1) >= test_water {
						test_light = cur_num[0] + (test_water - cur_num[1])
					} else {
						if test_light == 0 && k == line_map[i][1]-1 {
							test_light = test_water
						}
					}
					// fmt.Println("LIGHTTT: ", test_light)
				} else if i == 4 { // light to temperature
					if test_temp > 0 {
						continue
					} else if cur_num[1] <= test_light && (cur_num[1] + cur_num[2]-1) >= test_light {
						test_temp = cur_num[0] + (test_light - cur_num[1])
					} else {
						if test_temp == 0 && k == line_map[i][1]-1 {
							test_temp = test_light
						}
					}
					// fmt.Println("TEMPERATUREE: ", test_temp)
				} else if i == 5 { // temperature to humidity
					if test_hum > 0 {
						continue
					} else if cur_num[1] <= test_temp && (cur_num[1] + cur_num[2]-1) >= test_temp {
						test_hum = cur_num[0] + (test_temp - cur_num[1])
					} else {
						if test_hum == 0 && k == line_map[i][1]-1 {
							test_hum = test_temp
						}
					}
					// fmt.Println("HUMIDITY: ", test_hum)
				} else if i == 6 { // humidity to location
					if test_loc > 0 {
						continue
					} else if cur_num[1] <= test_hum && (cur_num[1] + cur_num[2]-1) >= test_hum {
						test_loc = cur_num[0] + (test_hum - cur_num[1])
					} else {
						if test_loc == 0 && k == line_map[i][1]-1 {
							test_loc = test_hum
						}
					}
		
					if test_loc > 0 && res_location > test_loc {
						// fmt.Println("LOCATIONNN: ", res_location, test_loc)
						// fmt.Println("HEREHERE !!!!!!! ", test_loc)
						res_location = test_loc
						// fmt.Println("ress loca now: ", res_location)
					}
				}
			}
		}
		// fmt.Println()
	}

	// res := inputSeed(0)
	return res_location
}

func main() {
	// input, err := os.ReadFile("example.txt")
	input, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	// fmt.Println(string(input))
	res := process(string(input))
	fmt.Println("\nresult: ", res)
}