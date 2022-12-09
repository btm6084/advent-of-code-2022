package day06

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
	line := lines[0]

	for i := 4; i < len(line); i++ {
		uniq := slice.Unique([]byte(line[i-4 : i]))
		if len(uniq) == 4 {
			return cast.ToString(i)
		}
	}

	return "error"
}

func Part2() string {
	lines := config.RawInput("./pkg/" + pkg + "/" + config.INPUTFILE)
	line := lines[0]

	for i := 14; i < len(line); i++ {
		uniq := slice.Unique([]byte(line[i-14 : i]))
		if len(uniq) == 14 {
			return cast.ToString(i)
		}
	}

	return "error"
}
