package day03

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

	sum := 0

	for _, line := range lines {
		seen := make(map[byte]bool)

		for i := 0; i < (len(line) / 2); i++ {
			seen[line[i]] = true
		}

		for i := (len(line) / 2); i < len(line); i++ {
			if seen[line[i]] {
				asInt := cast.ToInt(line[i])
				priority := 0
				if asInt <= 122 && asInt >= 97 {
					// Lowercase
					priority = asInt - 96
				} else if asInt <= 90 && asInt >= 65 {
					// Uppercase
					priority = asInt - 38
				}

				sum += priority
				break
			}
		}
	}

	return cast.ToString(sum)
}

func Part2() string {
	lines := config.RawInput("./pkg/" + pkg + "/" + config.INPUTFILE)
	sum := 0

	for len(lines) >= 3 {

		var uniq []string

		for _, line := range lines[:3] {
			uniq = append(uniq, string(slice.Unique([]byte(line))))
		}

		seen := make(map[byte]int)

		for _, line := range uniq {
			for i := range line {
				seen[line[i]]++
			}
		}

		for c, count := range seen {
			if count > 2 {
				asInt := cast.ToInt(c)
				priority := 0
				if asInt <= 122 && asInt >= 97 {
					// Lowercase
					priority = asInt - 96
				} else if asInt <= 90 && asInt >= 65 {
					// Uppercase
					priority = asInt - 38
				}

				sum += priority
				break
			}
		}

		lines = lines[3:]
	}

	return cast.ToString(sum)
}
