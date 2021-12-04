package main

import (
	"fmt"
)

type measurements struct {
	first  []int
	second []int
}
type counter struct {
	increases int
	decreases int
}

func main() {
	var file = "day1/puzzleInput"
	var part1Counter, part2Counter counter

	part1Counter = Part1(file)
	part2Counter = Part2(file)

	fmt.Println("Part 1 Increases count: ", part1Counter.increases)
	fmt.Println("Part 1 Decreases count: ", part1Counter.decreases)
	fmt.Println("Part 2 Increases count: ", part2Counter.increases)
	fmt.Println("Part 2 Decreases count: ", part2Counter.decreases)
}
