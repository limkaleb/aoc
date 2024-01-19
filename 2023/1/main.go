package main

import (
	"fmt"
	"log"
	"os"
	"reflect"
	"strconv"
	"strings"
	"unicode"
)

func process(input string) int {
	map_number := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
		"six":   6,
		"seven": 7,
		"eight": 8,
		"nine":  9,
	}

	keys := reflect.ValueOf(map_number).MapKeys()

	var result int = 0
	fmt.Println(input)
	fmt.Println()

	items := strings.Fields(input)
	first := 0
	second := 0
	total := 0
	key_found_index_first := -1
	key_found_index_last := -1

	for i, item := range items {
		fmt.Println(i, item)
		chars := []rune(item)

		for _, key := range keys {
			key_found_index_1 := strings.Index(item, key.Interface().(string))
			key_found_index_2 := strings.LastIndex(item, key.Interface().(string))
			current_number := map_number[key.String()]
			if key_found_index_first == -1 && key_found_index_1 >= 0 {
				key_found_index_first = key_found_index_1
				first = 10 * current_number
			}
			if key_found_index_last == -1 && key_found_index_2 >= 0 {
				key_found_index_last = key_found_index_2
				second = current_number
			}
			if key_found_index_1 >= 0 && key_found_index_first > key_found_index_1 {
				key_found_index_first = key_found_index_1
				first = 10 * current_number
			}
			if key_found_index_last >= 0 && key_found_index_last < key_found_index_2 {
				key_found_index_last = key_found_index_2
				second = current_number
			}
		}

		for j := 0; j < len(item); j++ {
			if (unicode.IsDigit(chars[j]) && j < key_found_index_first) || (unicode.IsDigit(chars[j]) && key_found_index_first == -1) {
				first, _ = strconv.Atoi(string(item[j]))
				first = 10 * first
				break
			}
		}

		for k := len(item) - 1; k >= 0; k-- {
			if (unicode.IsDigit(chars[k]) && k > key_found_index_last) || (unicode.IsDigit(chars[k]) && key_found_index_last == -1) {
				second, _ = strconv.Atoi(string(item[k]))
				break
			}
		}

		fmt.Println("last first second: ", first, second)
		key_found_index_first = -1
		key_found_index_last = -1
		total = total + first + second
		fmt.Println()
	}

	fmt.Println("TOTALL: ", total)
	return result
}

func main() {
	// input, err := os.ReadFile("example2.txt")
	input, err := os.ReadFile("input.txt")
	// input, err := os.ReadFile("test.txt")
	if err != nil {
		log.Fatal(err)
	}

	res := process(string(input))
	fmt.Println(res)
}
