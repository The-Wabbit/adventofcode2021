package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

type measurements struct {
	first  []int
	second []int
}
type counter struct {
	increases int
	decreases int
}

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

func main() {
	var file = "puzzleInput"
	var part1Counter, part2Counter counter

	part1Counter = Part1(file)
	part2Counter = Part2(file)

	fmt.Println("Part 1 Increases count: ", part1Counter.increases)
	fmt.Println("Part 1 Decreases count: ", part1Counter.decreases)
	fmt.Println("Part 2 Increases count: ", part2Counter.increases)
	fmt.Println("Part 2 Decreases count: ", part2Counter.decreases)
}
