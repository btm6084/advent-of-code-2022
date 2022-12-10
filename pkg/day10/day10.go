package day10

import (
	"os"
	"reflect"
	"strings"

	"github.com/btm6084/advent-of-code-2022/pkg/config"
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

	cycles := 0
	x := 1
	signal := 0

	for _, line := range lines {
		if line == "noop" {
			cycles++
			signal += sum(x, cycles)
			continue
		}

		cycles++
		signal += sum(x, cycles)
		n := cast.ToInt(strings.Split(line, " ")[1])
		cycles++
		signal += sum(x, cycles)
		x += n
	}

	return cast.ToString(signal)
}

func Part2() string {
	lines := config.RawInput("./pkg/" + pkg + "/" + config.INPUTFILE)

	x := 1

	crt := make([]string, 240)

	for i := 0; i < 240; i++ {
		line := lines[0]
		lines = lines[1:]

		px := i % 40

		if px >= x-1 && px <= x+1 {
			// Draw
			crt[i] = "#"
		} else {
			crt[i] = " "
		}

		if line == "noop" {
			continue
		}

		n := cast.ToInt(strings.Split(line, " ")[1])

		i++ // 2nd cycle
		px = i % 40

		if px >= x-1 && px <= x+1 {
			// Draw
			crt[i] = "#"
		} else {
			crt[i] = " "
		}

		x += n
	}

	return printCRT(crt)
}

func printCRT(crt []string) string {
	out := "\n"
	for i := 0; i < len(crt); i++ {
		if i > 0 && i%40 == 0 {
			out += "\n"
		}

		out += crt[i]
	}

	return out
}

func sum(x, cycle int) int {
	if slice.ContainsInt([]int{20, 60, 100, 140, 180, 220}, cycle) {
		return x * cycle
	}

	return 0
}
