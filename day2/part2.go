package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func part2(file string) position {
	var part2Position position

	data, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}
	defer data.Close()
	scanner := bufio.NewScanner(data)
	for scanner.Scan() {
		command := strings.Fields(scanner.Text())
		distance, _ := strconv.Atoi(command[1])
		if command[0] == "up" {
			part2Position.aim -= distance
		} else if command[0] == "down" {
			part2Position.aim += distance
		} else if command[0] == "forward" {
			part2Position.horizontal += distance
			part2Position.depth += part2Position.aim * distance
		}

	}

	return part2Position
}
