package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func Part1(file string) counter {
	var counters counter

	data, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}
	defer data.Close()
	var current, previous int

	scanner := bufio.NewScanner(data)
	for scanner.Scan() {
		current, err = strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}
		if previous != 0 {
			if previous < current {
				counters.increases++
			} else if previous > current {
				counters.decreases++
			}
		}
		previous = current
	}
	return counters
}
