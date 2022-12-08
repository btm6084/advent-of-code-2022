package day05

import (
	"fmt"
	"log"
	"os"
	"reflect"
	"strings"

	"github.com/btm6084/advent-of-code-2022/pkg/config"
	"github.com/btm6084/utilities/conv"
	"github.com/btm6084/utilities/slice"
	"github.com/spf13/cast"
)

type Empty struct{}

var pkg string

var ()

func init() {
	pkgRaw := strings.Split(reflect.TypeOf(Empty{}).PkgPath(), string(os.PathSeparator))

	pkg = pkgRaw[len(pkgRaw)-1]
}

func Part1() string {
	lines := config.RawInput("./pkg/" + pkg + "/" + config.INPUTFILE)
	stacks, instructions := parseInput(lines)

	return process(stacks, instructions, move)
}

func Part2() string {
	lines := config.RawInput("./pkg/" + pkg + "/" + config.INPUTFILE)
	stacks, instructions := parseInput(lines)

	return process(stacks, instructions, moveStack)
}

func parseInput(lines []string) ([]slice.Stack[string], []string) {
	var boxesRaw []string
	var instructions []string

	i := 0
	maxWidth := 0
	for ; i < len(lines); i++ {
		if lines[i] == "" {
			break
		}

		maxWidth = conv.MaxInt(maxWidth, len(lines[i]))
	}

	maxWidth = (maxWidth / 4) + 1
	boxesRaw = lines[:i-1]
	instructions = lines[i+1:]

	stacks := make([]slice.Stack[string], maxWidth)

	for i = len(boxesRaw) - 1; i >= 0; i-- {
		row := boxesRaw[i]
		pos := 0
		for {
			if len(row) < 3 {
				break
			}

			box := row[:3]
			row = row[3:]

			// Consume the spacer
			if len(row) > 2 {
				row = row[1:]
			}

			if box == "   " {
				pos++
				continue
			}

			stacks[pos].Push(string(box[1]))
			pos++
		}
	}

	return stacks, instructions
}

func move(stacks []slice.Stack[string], from, to, num int) {
	for i := 0; i < num; i++ {
		itm, err := stacks[from].Pop()
		if err != nil {
			log.Fatal(err)
			return
		}

		stacks[to].Push(itm)
	}
}

func moveStack(stacks []slice.Stack[string], from, to, num int) {
	var arm slice.Stack[string]
	for i := 0; i < num; i++ {
		itm, err := stacks[from].Pop()
		if err != nil {
			log.Fatal(err)
			return
		}

		arm.Push(itm)
	}

	for i := 0; i < num; i++ {
		itm, err := arm.Pop()
		if err != nil {
			log.Fatal(err)
			return
		}

		stacks[to].Push(itm)
	}
}

func process(stacks []slice.Stack[string], instructions []string, armFn func([]slice.Stack[string], int, int, int)) string {
	for _, line := range instructions {
		inst := strings.Split(line, " ")
		from := cast.ToInt(inst[3]) - 1
		to := cast.ToInt(inst[5]) - 1
		num := cast.ToInt(inst[1])

		armFn(stacks, from, to, num)
	}

	out := ""
	for i := 0; i < len(stacks); i++ {
		v, err := stacks[i].Peek()
		if err != nil {
			log.Fatal(err)
		}

		out += v
	}

	return out
}

func printStack(s []slice.Stack[string]) {
	stacks := make([]slice.Stack[string], len(s))
	height := 0
	for i := 0; i < len(s); i++ {
		stacks[i] = s[i].Copy()
		height = conv.MaxInt(height, s[i].Len())
	}

	for i := height; i > 0; i-- {
		for b := 0; b < len(stacks); b++ {
			if stacks[b].Len() >= i {
				s, _ := stacks[b].Pop()
				fmt.Printf("[%s].", s)
			} else {
				fmt.Print("....")
			}
		}
		fmt.Println()
	}

	for i := 0; i < len(stacks); i++ {
		fmt.Print(" ", i+1, "  ")
	}

	fmt.Println()
}
