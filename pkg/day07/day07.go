package day07

import (
	"fmt"
	"log"
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

const (
	FILE int = 1 + iota
	DIRECTORY
)

type Node struct {
	Name     string
	Type     int
	Size     int
	Parent   *Node
	Children []*Node
}

func (n *Node) mkdir(name string) {
	if n.Type == FILE {
		log.Fatal("mkdir called on file")
	}

	n.Children = append(n.Children, &Node{
		Name:   name,
		Type:   DIRECTORY,
		Parent: n,
	})
}

func (n *Node) touch(size int, name string) {
	if n.Type == FILE {
		log.Fatal("touch called on file")
	}

	n.Children = append(n.Children, &Node{
		Name:   name,
		Type:   FILE,
		Size:   size,
		Parent: n,
	})
}

func (n *Node) sum() int {
	if n.Type == FILE {
		return n.Size
	}

	sum := 0

	for _, c := range n.Children {
		if c.Type == FILE {
			sum += c.Size
		} else {
			sum += c.sum()
		}
	}

	n.Size = sum

	return sum
}

func (n *Node) sumP1() int {
	if n.Type == FILE {
		return 0
	}

	sum := 0
	if n.Size <= 100000 {
		sum = n.Size
	}

	for _, c := range n.Children {
		if c.Type == DIRECTORY {
			sum += c.sumP1()
		}
	}

	return sum
}

func (n *Node) sumP2(need int, have int) int {
	if n.Type == FILE {
		return 0
	}

	out := have
	if need <= n.Size {
		out = conv.MinInt(have, n.Size)
	}

	for _, c := range n.Children {
		if c.Type == DIRECTORY {
			out = conv.MinInt(c.sumP2(need, out), out)
		}
	}

	return out
}

func init() {
	pkgRaw := strings.Split(reflect.TypeOf(Empty{}).PkgPath(), string(os.PathSeparator))

	pkg = pkgRaw[len(pkgRaw)-1]
}

func Part1() string {
	lines := config.RawInput("./pkg/" + pkg + "/" + config.INPUTFILE)
	root := parseFS(lines)
	root.sum()
	return cast.ToString(root.sumP1())
}

func Part2() string {
	lines := config.RawInput("./pkg/" + pkg + "/" + config.INPUTFILE)
	root := parseFS(lines)
	root.sum()
	need := 30000000 - (70000000 - root.Size)
	return cast.ToString(root.sumP2(need, root.Size))
}

func parseFS(lines []string) *Node {
	lines = lines[1:]

	var cwd *Node
	var root = Node{
		Name: "/",
		Type: DIRECTORY,
	}

	cwd = &root

	for len(lines) > 0 {
		input := lines[0]
		lines = lines[1:]

		if input[0] != '$' {
			log.Fatal("Expected Instruction processing: ", input)
		}

		switch input[:4] {
		case "$ cd":
			dir := strings.Split(input, " ")[2]

			if dir == "/" {
				cwd = &root
				break
			}

			if dir == ".." {
				cwd = cwd.Parent
				break
			}

			cwd = CD(cwd, dir)
		case "$ ls":
			lines = LS(cwd, lines)
		}
	}

	return &root
}

func CD(cwd *Node, dir string) *Node {
	for _, n := range cwd.Children {
		if n.Name == dir {
			return n
		}
	}

	log.Fatal("no directory named " + dir + " found in " + cwd.Name)
	return nil
}

func LS(cwd *Node, lines []string) []string {

	for len(lines) > 0 {
		input := lines[0]

		if input[0] == '$' {
			return lines
		}

		lines = lines[1:]

		switch input[0] {
		case 'd':
			cwd.mkdir(strings.Split(input, " ")[1])
		default:
			s := strings.Split(input, " ")
			cwd.touch(cast.ToInt(s[0]), s[1])
		}
	}

	return []string{}
}

func printDir(cwd *Node, depth int) {
	for i := 0; i < depth; i++ {
		fmt.Print("\t")
	}

	if cwd.Type == FILE {
		fmt.Printf("- %s (file, size=%d)\n", cwd.Name, cwd.Size)
		return
	}

	fmt.Printf("- %s (dir, size=%d)\n", cwd.Name, cwd.Size)

	for _, n := range cwd.Children {
		if n.Type == DIRECTORY {
			printDir(n, depth+1)
		}
	}

	for _, n := range cwd.Children {
		if n.Type == FILE {
			for i := 0; i < depth+1; i++ {
				fmt.Print("\t")
			}
			fmt.Printf("- %s (file, size=%d)\n", n.Name, n.Size)
		}
	}
}
