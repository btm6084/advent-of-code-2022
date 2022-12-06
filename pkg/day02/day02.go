package day02

import (
	"os"
	"reflect"
	"strings"

	"github.com/btm6084/advent-of-code-2022/pkg/config"
	"github.com/spf13/cast"
)

type Empty struct{}

var pkg string

var (
	score = map[string]int{
		"A X": 4,
		"A Y": 8,
		"A Z": 3,
		"B X": 1,
		"B Y": 5,
		"B Z": 9,
		"C X": 7,
		"C Y": 2,
		"C Z": 6,
	}

	strategy = map[string]int{
		"A X": 3, // L
		"A Y": 4, // D
		"A Z": 8, // W
		"B X": 1, // L
		"B Y": 5, // D
		"B Z": 9, // W
		"C X": 2, // L
		"C Y": 6, // D
		"C Z": 7, // W
	}
)

func init() {
	pkgRaw := strings.Split(reflect.TypeOf(Empty{}).PkgPath(), string(os.PathSeparator))

	pkg = pkgRaw[len(pkgRaw)-1]
}

func Part1() string {
	lines := config.RawInput("./pkg/" + pkg + "/" + config.INPUTFILE)

	sum := 0
	for _, line := range lines {
		sum += score[line]
	}
	return cast.ToString(sum)
}

func Part2() string {
	lines := config.RawInput("./pkg/" + pkg + "/" + config.INPUTFILE)

	sum := 0
	for _, line := range lines {
		sum += strategy[line]
	}
	return cast.ToString(sum)
}
