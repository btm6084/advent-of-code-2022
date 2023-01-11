package day12

import (
	"fmt"
	"os"
	"reflect"
	"sort"
	"strings"

	"github.com/btm6084/advent-of-code-2022/pkg/config"
	"github.com/spf13/cast"
)

type Empty struct{}

var pkg string

var (
	alpha = map[byte]int{
		'a': 0, 'b': 1, 'c': 2, 'd': 3, 'e': 4, 'f': 5, 'g': 6, 'h': 7, 'i': 8, 'j': 9,
		'k': 10, 'l': 11, 'm': 12, 'n': 13, 'o': 14, 'p': 15, 'q': 16, 'r': 17, 's': 18, 't': 19,
		'u': 20, 'v': 21, 'w': 22, 'x': 23, 'y': 24, 'z': 25,
		'S': 0, 'E': 25,
	}
)

func init() {
	pkgRaw := strings.Split(reflect.TypeOf(Empty{}).PkgPath(), string(os.PathSeparator))

	pkg = pkgRaw[len(pkgRaw)-1]
}

type Node struct {
	ID      string
	Cost    int
	Visited bool
	Height  int
	IsGoal  bool
	IsStart bool
	Row     int
	Col     int
	Parent  *Node
}

func Part1() string {
	lines := config.RawInput("./pkg/" + pkg + "/" + config.INPUTFILE)

	graph := make([][]*Node, len(lines))
	for i := 0; i < len(lines); i++ {
		graph[i] = make([]*Node, len(lines[i]))
	}

	var toVisit []*Node
	var goal *Node

	for row, line := range lines {
		for col := 0; col < len(line); col++ {
			graph[row][col] = &Node{
				ID:      fmt.Sprintf("[%d,%d]", row, col),
				Height:  alpha[line[col]],
				Cost:    999,
				Visited: false,
				Row:     row,
				Col:     col,
			}

			if line[col] == 'S' {
				graph[row][col].IsStart = true
				graph[row][col].Parent = &Node{Cost: -1}
				toVisit = append(toVisit, graph[row][col])
			}

			if line[col] == 'E' {
				graph[row][col].IsGoal = true
				goal = graph[row][col]
			}
		}
	}

	i := 0
	for len(toVisit) > 0 {
		i++
		toVisit = visit(graph, toVisit, i)
	}

	return cast.ToString(goal.Cost)
}

func Part2() string {
	lines := config.RawInput("./pkg/" + pkg + "/" + config.INPUTFILE)

	type pt struct {
		Row int
		Col int
	}

	var starts []pt
	for row, line := range lines {
		for col := 0; col < len(line); col++ {
			if line[col] == 'a' || line[col] == 'S' {
				starts = append(starts, pt{Row: row, Col: col})
			}
		}
	}

	var costs []int
	for _, p := range starts {
		costs = append(costs, run(lines, p.Row, p.Col))
	}

	sort.Ints(costs)

	return cast.ToString(costs[0])
}

func run(lines []string, startX, startY int) int {
	graph := make([][]*Node, len(lines))
	for i := 0; i < len(lines); i++ {
		graph[i] = make([]*Node, len(lines[i]))
	}

	var toVisit []*Node
	var goal *Node

	for row, line := range lines {
		for col := 0; col < len(line); col++ {
			graph[row][col] = &Node{
				ID:      fmt.Sprintf("[%d,%d]", row, col),
				Height:  alpha[line[col]],
				Cost:    999,
				Visited: false,
				Row:     row,
				Col:     col,
			}

			if row == startX && col == startY {
				graph[row][col].IsStart = true
				graph[row][col].Parent = &Node{Cost: -1}
				toVisit = append(toVisit, graph[row][col])
			}

			if line[col] == 'E' {
				graph[row][col].IsGoal = true
				goal = graph[row][col]
			}
		}
	}

	i := 0
	for len(toVisit) > 0 {
		i++
		toVisit = visit(graph, toVisit, i)
	}

	return goal.Cost
}

func visit(graph [][]*Node, toVisit []*Node, i int) []*Node {
	node := toVisit[0]
	toVisit = toVisit[1:]

	if node.Visited {
		return toVisit
	}

	node.Cost = node.Parent.Cost + 1
	node.Visited = true

	if node.Row-1 >= 0 {
		nn := graph[node.Row-1][node.Col]
		if !nn.Visited && node.Height-nn.Height >= -1 {
			nn.Parent = node
			toVisit = append(toVisit, nn)
		}
	}

	if node.Row+1 < len(graph) {
		nn := graph[node.Row+1][node.Col]
		if !nn.Visited && node.Height-nn.Height >= -1 {
			nn.Parent = node
			toVisit = append(toVisit, nn)
		}
	}

	if node.Col-1 >= 0 {
		nn := graph[node.Row][node.Col-1]
		if !nn.Visited && node.Height-nn.Height >= -1 {
			nn.Parent = node
			toVisit = append(toVisit, nn)
		}
	}

	if node.Col+1 < len(graph[0]) {
		nn := graph[node.Row][node.Col+1]
		if !nn.Visited && node.Height-nn.Height >= -1 {
			nn.Parent = node
			toVisit = append(toVisit, nn)
		}
	}

	return toVisit
}
