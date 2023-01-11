package day13

import (
	"fmt"
	"os"
	"reflect"
	"strings"

	"github.com/btm6084/advent-of-code-2022/pkg/config"
	"github.com/btm6084/gojson"
	"github.com/davecgh/go-spew/spew"
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

	for len(lines) >= 2 {
		l := lines[:2]
		lines = lines[2:]

		var a []int
		gojson.Unmarshal([]byte(l[0]), &a)

		var b []int
		gojson.Unmarshal([]byte(l[0]), &b)

		spew.Dump(a, b)
		fmt.Println()
		fmt.Println()

		if len(lines) > 2 {
			lines = lines[1:]
		}
	}

	return "@todo"
}

func Part2() string {
	lines := config.RawInput("./pkg/" + pkg + "/" + config.INPUTFILE)
	_ = lines

	return "@todo"
}
