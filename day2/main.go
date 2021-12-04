package main

import "fmt"

func main() {
	part1Position := part1(file)
	fmt.Println("Depth x Horisontal is: ", part1Position.depth*part1Position.horizontal)
	part2Position := part2(file)
	fmt.Println("Depth x Horisontal is: ", part2Position.depth*part2Position.horizontal)
}
