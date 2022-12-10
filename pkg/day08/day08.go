package day08

import (
	"os"
	"reflect"
	"strings"

	"github.com/btm6084/advent-of-code-2022/pkg/config"
	"github.com/btm6084/utilities/conv"
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
	grid := parseGrid(lines)

	visible := len(grid) * 2
	visible += (len(grid[0]) - 2) * 2

	for row := 1; row < len(grid)-1; row++ {
		for col := 1; col < len(grid)-1; col++ {
			n, s, e, w := true, true, true, true
			// n
			for i := row - 1; i >= 0; i-- {
				if grid[i][col] >= grid[row][col] {
					n = false
					break
				}
			}
			if n {
				visible++
				continue
			}

			// s
			for i := row + 1; i < len(grid); i++ {
				if grid[i][col] >= grid[row][col] {
					s = false
					break
				}
			}
			if s {
				visible++
				continue
			}

			// e
			for i := col + 1; i < len(grid[row]); i++ {
				if grid[row][i] >= grid[row][col] {
					e = false
					break
				}
			}
			if e {
				visible++
				continue
			}

			// w
			for i := col - 1; i >= 0; i-- {
				if grid[row][i] >= grid[row][col] {
					w = false
					break
				}
			}
			if w {
				visible++
				continue
			}
		}
	}

	return cast.ToString(visible)
}

func Part2() string {
	lines := config.RawInput("./pkg/" + pkg + "/" + config.INPUTFILE)
	grid := parseGrid(lines)

	score := make([][]int, len(grid))
	for row := 0; row < len(grid); row++ {
		score[row] = make([]int, len(grid[0]))
	}

	max := 0
	for row := 0; row < len(grid); row++ {
		for col := 0; col < len(grid); col++ {
			max = conv.MaxInt(max, calcScore(grid, row, col))
		}
	}

	return cast.ToString(max)
}

func calcScore(grid [][]int, row, col int) int {
	nv := 0
	for i := row - 1; i >= 0; i-- {
		nv++
		if grid[i][col] >= grid[row][col] {
			break
		}
	}

	sv := 0
	for i := row + 1; i < len(grid); i++ {
		sv++
		if grid[i][col] >= grid[row][col] {
			break
		}
	}

	ev := 0
	for i := col + 1; i < len(grid[row]); i++ {
		ev++
		if grid[row][i] >= grid[row][col] {
			break
		}
	}

	wv := 0
	for i := col - 1; i >= 0; i-- {
		wv++
		if grid[row][i] >= grid[row][col] {
			break
		}
	}

	return nv * sv * ev * wv
}

func parseGrid(lines []string) [][]int {
	if len(lines) < 1 {
		return [][]int{}
	}

	out := make([][]int, len(lines))
	for i := 0; i < len(out); i++ {
		out[i] = make([]int, len(lines[0]))
	}

	for i, l := range lines {
		for k, c := range l {
			out[i][k] = cast.ToInt(string(c))
		}
	}

	return out
}
