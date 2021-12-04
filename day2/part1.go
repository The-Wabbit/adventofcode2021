package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func part1(file string) position {
	var part1Position position
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
			part1Position.depth -= distance
		} else if command[0] == "down" {
			part1Position.depth += distance
		} else if command[0] == "forward" {
			part1Position.horizontal += distance
		}

	}
	return part1Position
}
