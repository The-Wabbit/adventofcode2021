package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func Part2(file string) counter {
	var counters counter
	var measures measurements
	var current int

	data, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}
	defer data.Close()
	scanner := bufio.NewScanner(data)
	for scanner.Scan() {
		current, err = strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}
		// initializing the stacks
		if len(measures.second) < 3 {
			if len(measures.first) < 1 {
				measures.first = append(measures.first, current)
			} else if len(measures.first) == 3 {
				measures.second = append(measures.second, current)
			} else {
				measures.first = append(measures.first, current)
				measures.second = append(measures.second, current)
			}
			// after both stacks are full, we can begin counting
		} else {
			if measures.first[0]+measures.first[1]+measures.first[2] < measures.second[0]+measures.second[1]+measures.second[2] {
				counters.increases++
			} else if measures.first[0]+measures.first[1]+measures.first[2] > measures.second[0]+measures.second[1]+measures.second[2] {
				counters.decreases++
			}
			measures.first = measures.first[1:]
			measures.first = append(measures.first, measures.second[2])
			measures.second = measures.second[1:]
			measures.second = append(measures.second, current)
		}
	}

	//final check after pushing the last line in the file to the second stack
	if measures.first[0]+measures.first[1]+measures.first[2] < measures.second[0]+measures.second[1]+measures.second[2] {
		counters.increases++
	} else if measures.first[0]+measures.first[1]+measures.first[2] > measures.second[0]+measures.second[1]+measures.second[2] {
		counters.decreases++
	}
	return counters
}
