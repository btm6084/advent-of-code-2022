package day09

import (
	"fmt"
	"os"
	"reflect"
	"strings"

	"github.com/btm6084/advent-of-code-2022/pkg/config"
	"github.com/spf13/cast"
)

type Empty struct{}

var pkg string

var ()

type Knot struct {
	x int
	y int
}

func init() {
	pkgRaw := strings.Split(reflect.TypeOf(Empty{}).PkgPath(), string(os.PathSeparator))

	pkg = pkgRaw[len(pkgRaw)-1]
}

func Part1() string {
	lines := config.RawInput("./pkg/" + pkg + "/" + config.INPUTFILE)

	knots := []Knot{
		{0, 0},
		{0, 0},
	}

	sum := playSnake(lines, knots)

	return cast.ToString(sum)
}

func Part2() string {
	lines := config.RawInput("./pkg/" + pkg + "/" + config.INPUTFILE)

	knots := []Knot{
		{0, 0},
		{0, 0},
		{0, 0},
		{0, 0},
		{0, 0},
		{0, 0},
		{0, 0},
		{0, 0},
		{0, 0},
		{0, 0},
	}

	sum := playSnake(lines, knots)

	return cast.ToString(sum)
}

func playSnake(lines []string, knots []Knot) int {
	visited := make(map[int]map[int]int)
	visited[0] = map[int]int{}
	visited[0][0]++

	// visualize(knots)

	for _, line := range lines {
		s := strings.Split(line, " ")
		dir := s[0]
		num := cast.ToInt(s[1])

		// fmt.Println()
		// fmt.Println("==", line, "==")
		// fmt.Println()

		for i := 0; i < num; i++ {
			switch dir {
			case "R":
				knots[0].x++
			case "L":
				knots[0].x--
			case "D":
				knots[0].y--
			case "U":
				knots[0].y++
			}

			for k := 1; k < len(knots); k++ {
				knots[k].x, knots[k].y = delta(knots[k-1].x, knots[k-1].y, knots[k].x, knots[k].y)
			}

			t := len(knots) - 1
			if visited[knots[t].x] == nil {
				visited[knots[t].x] = make(map[int]int)
			}
			visited[knots[t].x][knots[t].y]++
		}
		// visualize(knots)
	}

	sum := 0
	for _, m := range visited {
		sum += len(m)
	}

	return sum
}

func delta(hX, hY, tX, tY int) (int, int) {

	dX := hX - tX
	dY := hY - tY

	if hX != tX && hY != tY && (abs(dX)+abs(dY) > 2) {
		x := 1
		if dX < 0 {
			x = -1
		}

		y := 1
		if dY < 0 {
			y = -1
		}
		// diagonal
		return tX + x, tY + y
	}

	if dX > 0 {
		dX--
	} else if dX < 0 {
		dX++
	}

	if dY > 0 {
		dY--
	} else if dY < 0 {
		dY++
	}

	return tX + dX, tY + dY
}

func abs(x int) int {
	if x < 0 {
		return -x
	}

	return x
}

func visualize(knots []Knot) {
	for row := 25; row >= -25; row-- {
		for col := -25; col < 25; col++ {
			chr := "."
			if 0 == row && 0 == col {
				chr = "s"
			}
			for i := len(knots) - 1; i >= 0; i-- {
				if knots[i].x == col && knots[i].y == row {
					chr = cast.ToString(i)

					if i == len(knots)-1 {
						chr = "T"
					}

					if i == 0 {
						chr = "H"
					}
				}
			}

			fmt.Print(chr)
		}
		fmt.Println()
	}
	fmt.Println()
}
