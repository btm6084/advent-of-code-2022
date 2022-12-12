package day11

import (
	"os"
	"reflect"
	"sort"
	"strings"

	"github.com/btm6084/advent-of-code-2022/pkg/config"
	"github.com/btm6084/utilities/queue"
	"github.com/spf13/cast"
)

type Empty struct{}

var pkg string

var ()

type Monkey struct {
	Inventory   queue.Queue[int]
	Operand     string
	Modifier    int
	Test        int
	OnTrue      int
	OnFalse     int
	Inspections int
}

func init() {
	pkgRaw := strings.Split(reflect.TypeOf(Empty{}).PkgPath(), string(os.PathSeparator))

	pkg = pkgRaw[len(pkgRaw)-1]
}

func Part1() string {
	lines := config.RawInput("./pkg/" + pkg + "/" + config.INPUTFILE)
	first, second := p1(lines, 20)
	return cast.ToString(first * second)
}

func Part2() string {
	lines := config.RawInput("./pkg/" + pkg + "/" + config.INPUTFILE)
	first, second := p2(lines, 10000, false)
	return cast.ToString(first * second)
}

func p1(lines []string, iterations int) (int, int) {
	monkeys := parseInput(lines)

	for i := 0; i < iterations; i++ {
		for m := range monkeys {
			for {
				// Monkey Inspects
				item, err := monkeys[m].Inventory.Dequeue()
				if err != nil {
					break
				}

				monkeys[m].Inspections++

				// Worry levels rising.
				switch monkeys[m].Operand {
				case "+":
					item += monkeys[m].Modifier
				case "*":
					item *= monkeys[m].Modifier
				case "sqr":
					item *= item
				}

				item /= 3

				// Monkey throws it
				recip := monkeys[m].OnFalse
				if item%monkeys[m].Test == 0 {
					recip = monkeys[m].OnTrue
				}

				monkeys[recip].Inventory.Enqueue(item)
			}
		}
	}

	inspections := make([]int, len(monkeys))
	for k, m := range monkeys {
		inspections[k] = m.Inspections
	}

	sort.Ints(inspections)
	n := len(inspections)

	return inspections[n-1], inspections[n-2]
}

func p2(lines []string, iterations int, div bool) (int, int) {
	monkeys := parseInput(lines)

	divisor := monkeys[0].Test

	for m := 1; m < len(monkeys); m++ {
		divisor *= monkeys[m].Test
	}

	for i := 0; i < iterations; i++ {
		for m := range monkeys {
			for {
				// Monkey Inspects
				item, err := monkeys[m].Inventory.Dequeue()
				if err != nil {
					break
				}

				item = item % divisor

				monkeys[m].Inspections++

				// Worry levels rising.
				switch monkeys[m].Operand {
				case "+":
					item += monkeys[m].Modifier
				case "*":
					item *= monkeys[m].Modifier
				case "sqr":
					item *= item
				}

				// Monkey throws it
				recip := monkeys[m].OnFalse
				if item%monkeys[m].Test == 0 {
					recip = monkeys[m].OnTrue
				}

				monkeys[recip].Inventory.Enqueue(item)
			}
		}
	}

	inspections := make([]int, len(monkeys))
	for k, m := range monkeys {
		inspections[k] = m.Inspections
	}

	sort.Ints(inspections)
	n := len(inspections)

	return inspections[n-1], inspections[n-2]
}

func parseInput(lines []string) []Monkey {
	var monkeys []Monkey

	for len(lines) > 0 {
		monkey := Monkey{}

		l := lines[:6]
		lines = lines[6:]

		// Consume line break
		if len(lines) > 6 {
			lines = lines[1:]
		}

		// Skip Monkey designator
		l = l[1:]

		inv := strings.Split(strings.Split(l[0], ": ")[1], ", ")
		for i := 0; i < len(inv); i++ {
			monkey.Inventory.Enqueue((cast.ToInt(inv[i])))
		}

		// Move past inventory
		l = l[1:]

		mod := strings.Split(strings.Split(l[0], "old ")[1], " ")
		monkey.Operand = mod[0]
		monkey.Modifier = cast.ToInt(mod[1])

		if mod[1] == "old" {
			monkey.Operand = "sqr"
		}

		// Move past Operation
		l = l[1:]
		monkey.Test = cast.ToInt(strings.Split(strings.Split(l[0], "by ")[1], " ")[0])

		// Move past Test
		l = l[1:]
		monkey.OnTrue = cast.ToInt(strings.Split(strings.Split(l[0], "monkey ")[1], " ")[0])

		// Move past True
		l = l[1:]
		monkey.OnFalse = cast.ToInt(strings.Split(strings.Split(l[0], "monkey ")[1], " ")[0])

		monkeys = append(monkeys, monkey)
	}

	return monkeys
}
