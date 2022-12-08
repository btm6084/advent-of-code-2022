package day04

import (
	"os"
	"reflect"
	"strings"

	"github.com/btm6084/advent-of-code-2022/pkg/config"
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

	overlaps := 0
	for _, line := range lines {
		a := strings.Split(line, ",")
		b := strings.Split(a[0], "-")
		c := strings.Split(a[1], "-")

		aStart, aEnd := cast.ToInt(b[0]), cast.ToInt(b[1])
		bStart, bEnd := cast.ToInt(c[0]), cast.ToInt(c[1])

		if aStart >= bStart && aEnd <= bEnd {
			overlaps++
			continue
		}

		if bStart >= aStart && bEnd <= aEnd {
			overlaps++
			continue
		}
	}

	return cast.ToString(overlaps)
}

func Part2() string {
	lines := config.RawInput("./pkg/" + pkg + "/" + config.INPUTFILE)

	overlaps := 0
	for _, line := range lines {
		a := strings.Split(line, ",")
		b := strings.Split(a[0], "-")
		c := strings.Split(a[1], "-")

		aStart, aEnd := cast.ToInt(b[0]), cast.ToInt(b[1])
		bStart, bEnd := cast.ToInt(c[0]), cast.ToInt(c[1])

		if aStart >= bStart && aStart <= bEnd {
			overlaps++
			continue
		}

		if bStart >= aStart && bStart <= aEnd {
			overlaps++
			continue
		}
	}

	return cast.ToString(overlaps)
}
