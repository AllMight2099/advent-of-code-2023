package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func check_error(e error) {
    if e != nil {
        log.Fatal(e)
    }
}

func main() {
	part_one()
	part_two()

}

func part_one() {
	file, err := os.Open("input.txt")
	check_error(err)
	defer file.Close()
	var sum int = 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()
		re := regexp.MustCompile("[0-9]+")
		extracted_number := strings.Join(re.FindAllString(text, -1), "")
		first_digit := extracted_number[0:1]
		last_digit := extracted_number[len(extracted_number)-1:] 
		number := first_digit+last_digit
		int_extracted_number, err := strconv.Atoi(number)
		check_error(err)
		sum+=int_extracted_number
	}
	fmt.Println("Part One:")
	fmt.Println("Final Callibration Number: ", sum)
	fmt.Println()

	check_error(scanner.Err())
}

func part_two() {
	file, err := os.Open("input.txt")
	check_error(err)
	defer file.Close()
	numbers := map[string]int{"nine":9, "eight":8, "seven":7, "six":6, "five":5,  "four":4,"three":3, "two":2, "one": 1}
	extra_numbers := map[string]string{"nine": "nninee","eight":"eeightt","seven" :"ssevenn","six" :"ssixx", "five":"ffivee","four" : "ffourr","three": "tthreee","two" :"ttwoo", "one": "oonee"}
	ordered_numbers := []string{"nine", "eight", "seven", "six", "five", "four", "three", "two", "one"}
	var sum int = 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()
		for _, key:= range ordered_numbers {
			text = strings.Replace(text, key, extra_numbers[key], -1)
			text = strings.Replace(text, key, strconv.Itoa(numbers[key]), -1)
		}
		re := regexp.MustCompile("[0-9]+")
		extracted_number := strings.Join(re.FindAllString(text, -1), "")
		first_digit := extracted_number[0:1]
		last_digit := extracted_number[len(extracted_number)-1:] 
		number := first_digit+last_digit
		int_extracted_number, err := strconv.Atoi(number)
		check_error(err)
		sum+=int_extracted_number
	}

	fmt.Println("Part Two:")
	fmt.Println("Final Callibration Number: ", sum)
	fmt.Println()
}