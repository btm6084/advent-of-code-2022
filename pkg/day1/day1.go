package day1

import (
	"os"
	"reflect"
	"sort"
	"strings"

	"github.com/btm6084/advent-of-code-2022/pkg/config"
	"github.com/btm6084/utilities/slice"
	"github.com/spf13/cast"
)

type Empty struct{}

var pkg string

func init() {
	pkgRaw := strings.Split(reflect.TypeOf(Empty{}).PkgPath(), string(os.PathSeparator))

	pkg = pkgRaw[len(pkgRaw)-1]
}

func sumElves(lines []string) []int {
	var sums []int

	i := 0
	sums = append(sums, 0)
	for _, line := range lines {
		if line == "" {
			sums = append(sums, 0)
			i++
			continue
		}

		sums[i] += cast.ToInt(line)
	}

	return sums
}

func Part1() string {
	lines := config.RawInput("./pkg/" + pkg + "/" + config.INPUTFILE)
	sums := sumElves(lines)

	return cast.ToString(slice.Max(sums))
}

func Part2() string {
	lines := config.RawInput("./pkg/" + pkg + "/" + config.INPUTFILE)
	sums := sumElves(lines)

	sort.Ints(sums)

	return cast.ToString(slice.Sum(sums[len(sums)-3:]))
}
