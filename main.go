package main

import (
	"fmt"
)

type Position struct {
	x int
	y int
}

type Transform struct {
	pos   Position
	angle int
}

type StackTranforms struct {
	items []Transform
}

func (s *StackTranforms) push(t Transform) {
	s.items = append(s.items, t)
}

func (s *StackTranforms) pop() {
	if len(s.items) == 0 {
		return
	}
	s.items = s.items[:len(s.items)-1]
}

func (pos *Position) posToIndex() (x int, y int) {
	x = pos.x
	y = (n - (pos.y + 1))
	return
}

func printTree(treeArray [n][n]string) {
	for _, level := range treeArray {
		for _, c := range level {
			fmt.Printf(" %s", c)
		}
		fmt.Println()

	}
}

// L - system fractal plant

func fractalPlan(n int, i int, input string) string {
	if i >= n {
		return input
	}
	out := ""
	for _, c := range input {
		switch c {
		case 'X':
			out = out + "F+[[X]-X]-F[-FX]+X"
		case 'F':
			out = out + "FF"
		default:
			out = out + string(c)
		}
	}
	i = i + 1
	return fractalPlan(n, i, out)
}

const n = 100

var anglecharMap = map[int]string{
	0:  "|",
	1:  "\\",
	-1: "/",
}

func renderFractcalPlant(input string, treeArray *[n][n]string, startPos Position) {
	currTransform := Transform{pos: startPos}
	transformStack := StackTranforms{items: make([]Transform, 0)}
	for _, c := range input {
		x, y := currTransform.pos.posToIndex()
		switch c {
		case '-':
			if currTransform.angle == -1 {
				currTransform.angle = -1
				continue
			}
			currTransform.angle -= 1
		case '+':
			if currTransform.angle == 1 {
				currTransform.angle = 1
				continue
			}
			currTransform.angle += 1
		case 'F':
			treeArray[y][x] = anglecharMap[currTransform.angle]
		case '[':
			transformStack.push(currTransform)
		case ']':
			currTransform = transformStack.items[len(transformStack.items)-1]
			transformStack.pop()
		}
		//moving the cursor
		if c == 'F' {
			switch currTransform.angle {
			case -1:
				currTransform.pos.x += 1
			case 1:
				currTransform.pos.x -= 1
			}
			currTransform.pos.y += 1
			if currTransform.pos.x < 0 || currTransform.pos.y < 0 || currTransform.pos.x >= n || currTransform.pos.y >= n {
				fmt.Println("Out of bounds", currTransform.pos)
				return
			}
			//x, y := currTransform.pos.posToIndex()
			//fmt.Println("New pos : ", x, y)
		}
	}

}

func main() {
	treeArray := [n][n]string{}
	printTree(treeArray)

	out := fractalPlan(4, 0, "-X")

	fmt.Println(out)
	renderFractcalPlant(out, &treeArray, Position{0, 0})
	printTree(treeArray)

}
