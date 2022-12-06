package day06

import (
	"os"
	"reflect"
	"strings"

	"github.com/btm6084/advent-of-code-2022/pkg/config"
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
	_ = lines

	return "@todo"
}

func Part2() string {
	lines := config.RawInput("./pkg/" + pkg + "/" + config.INPUTFILE)
	_ = lines

	return "@todo"
}
