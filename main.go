package main

import (
	"fmt"

	"github.com/btm6084/advent-of-code-2022/pkg/day1"
	"github.com/btm6084/advent-of-code-2022/pkg/day2"
)

type Runner func() string

var (
	runners = []Runner{
		day1.Part1,
		day1.Part2,
		day2.Part1,
		day2.Part2,
	}

	labels = []string{
		"Day 1 Part 1",
		"Day 1 Part 2",
		"Day 2 Part 1",
		"Day 2 Part 2",
	}
)

func main() {
	for i := range runners {
		fmt.Printf("%s: %s\n", labels[i], runners[i]())
	}

	fmt.Println()
}
