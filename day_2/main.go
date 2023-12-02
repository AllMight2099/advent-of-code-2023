package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
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
	scanner:
	for scanner.Scan() {
		text := scanner.Text()
		game_text := strings.Split(text, ":")
		// fmt.Println(game_text[0])

		game_number, err := strconv.Atoi(strings.Split(game_text[0], " ")[1])
		check_error(err)
		// fmt.Println(game_number)
		game_list := strings.Split(game_text[1], ";")
		// temp_sum := 0;
		for _, game := range game_list {
			// fmt.Println(game)
			balls := strings.Split(game, ",")
			// ball_count = map[string]int{"red":0, "green":0, "blue":0}
			for _, ball := range balls {
				// fmt.Println("'here'",strings.Split(strings.TrimSpace(ball), " ")[0])
				no, err := strconv.Atoi(strings.Split(strings.TrimSpace(ball), " ")[0])
				check_error(err)
				color := strings.Split(strings.TrimSpace(ball), " ")[1]
				if (color == "red" && no>12) {
					continue scanner
				} else if (color == "green" && no>13) {
					continue scanner
				} else if  (color == "blue" && no >14) {
					continue scanner
				}
			}
		}
		sum+=game_number
		// fmt.Println(games)
	}
	fmt.Println("Part One:")
	fmt.Println("Sum of game numbers", sum)
	fmt.Println()
	check_error(scanner.Err())
}

func part_two() {
	file, err := os.Open("input.txt")
	check_error(err)
	defer file.Close()
	var sum int = 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()
		game_text := strings.Split(text, ":")
		// fmt.Println(game_text[0])

		// game_number, err := strconv.Atoi(strings.Split(game_text[0], " ")[1])
		check_error(err)
		// fmt.Println(game_number)
		game_list := strings.Split(game_text[1], ";")
		// temp_sum := 0;
		ball_count := map[string]int{"red":0, "green":0, "blue":0}
		for _, game := range game_list {
			// fmt.Println(game)
			balls := strings.Split(game, ",")
			for _, ball := range balls {
				no, err := strconv.Atoi(strings.Split(strings.TrimSpace(ball), " ")[0])
				check_error(err)
				color := strings.Split(strings.TrimSpace(ball), " ")[1]
				if no>ball_count[color]{
					ball_count[color]=no
				}
			}
			
		}
		power:=ball_count["red"]*ball_count["green"]*ball_count["blue"]
		sum+=power
		// fmt.Println(games)
	}
	fmt.Println("Part Two")
	fmt.Println("Sum of powers: " ,sum)
	fmt.Println()
	check_error(scanner.Err())

}