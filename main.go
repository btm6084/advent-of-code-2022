package main

import (
	"flag"
	"fmt"

	"github.com/btm6084/advent-of-code-2022/pkg/config"
	"github.com/btm6084/advent-of-code-2022/pkg/day01"
	"github.com/btm6084/advent-of-code-2022/pkg/day02"
	"github.com/btm6084/advent-of-code-2022/pkg/day03"
	"github.com/btm6084/advent-of-code-2022/pkg/day04"
	"github.com/btm6084/advent-of-code-2022/pkg/day05"
	"github.com/btm6084/advent-of-code-2022/pkg/day06"
	"github.com/btm6084/advent-of-code-2022/pkg/day07"
	"github.com/btm6084/advent-of-code-2022/pkg/day08"
	"github.com/btm6084/advent-of-code-2022/pkg/day09"
	"github.com/btm6084/advent-of-code-2022/pkg/day10"
	"github.com/btm6084/advent-of-code-2022/pkg/day11"
	"github.com/btm6084/advent-of-code-2022/pkg/day12"
	"github.com/btm6084/advent-of-code-2022/pkg/day13"
	"github.com/btm6084/advent-of-code-2022/pkg/day14"
	"github.com/btm6084/advent-of-code-2022/pkg/day15"
	"github.com/btm6084/advent-of-code-2022/pkg/day16"
	"github.com/btm6084/advent-of-code-2022/pkg/day17"
	"github.com/btm6084/advent-of-code-2022/pkg/day18"
	"github.com/btm6084/advent-of-code-2022/pkg/day19"
	"github.com/btm6084/advent-of-code-2022/pkg/day20"
	"github.com/btm6084/advent-of-code-2022/pkg/day21"
	"github.com/btm6084/advent-of-code-2022/pkg/day22"
	"github.com/btm6084/advent-of-code-2022/pkg/day23"
	"github.com/btm6084/advent-of-code-2022/pkg/day24"
	"github.com/btm6084/advent-of-code-2022/pkg/day25"
)

type Runner func() string

var (
	runners = [][]Runner{
		{
			day01.Part1,
			day01.Part2,
		},
		{
			day02.Part1,
			day02.Part2,
		},
		{
			day03.Part1,
			day03.Part2,
		},
		{
			day04.Part1,
			day04.Part2,
		},
		{
			day05.Part1,
			day05.Part2,
		},
		{
			day06.Part1,
			day06.Part2,
		},
		{
			day07.Part1,
			day07.Part2,
		},
		{
			day08.Part1,
			day08.Part2,
		},
		{
			day09.Part1,
			day09.Part2,
		},
		{
			day10.Part1,
			day10.Part2,
		},
		{
			day11.Part1,
			day11.Part2,
		},
		{
			day12.Part1,
			day12.Part2,
		},
		{
			day13.Part1,
			day13.Part2,
		},
		{
			day14.Part1,
			day14.Part2,
		},
		{
			day15.Part1,
			day15.Part2,
		},
		{
			day16.Part1,
			day16.Part2,
		},
		{
			day17.Part1,
			day17.Part2,
		},
		{
			day18.Part1,
			day18.Part2,
		},
		{
			day19.Part1,
			day19.Part2,
		},
		{
			day20.Part1,
			day20.Part2,
		},
		{
			day21.Part1,
			day21.Part2,
		},
		{
			day22.Part1,
			day22.Part2,
		},
		{
			day23.Part1,
			day23.Part2,
		},
		{
			day24.Part1,
			day24.Part2,
		},
		{
			day25.Part1,
			day25.Part2,
		},
	}
)

func main() {

	var day = flag.Int("d", 0, "specific day to run")
	var test = flag.Bool("t", false, "use example data")
	flag.Parse()

	if *test {
		config.INPUTFILE = "example.txt"
	}

	if *day > 0 {
		fmt.Printf("Day %d Part 1: %s\n", *day-1, runners[*day-1][0]())
		fmt.Printf("Day %d Part 2: %s\n", *day-1, runners[*day-1][1]())

		return
	}

	for i := 0; i < len(runners); i++ {
		fmt.Printf("Day %d Part 1: %s\n", i+1, runners[i][0]())
		fmt.Printf("Day %d Part 2: %s\n", i+1, runners[i][1]())
		fmt.Println()
	}
}
