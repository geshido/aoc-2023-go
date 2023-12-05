package main

import (
	"fmt"
	"github.com/geshido/aoc-2023-go/util/conv"
	"github.com/geshido/aoc-2023-go/util/field"
	"github.com/geshido/aoc-2023-go/util/file"
	"github.com/samber/lo"
)

func main() {
	fmt.Println("part1 test:", part1(prepareInput(file.LoadStrings("test.txt"))))
	fmt.Println("part1:", part1(prepareInput(file.LoadStrings("input.txt"))))
	fmt.Println("part2 test:", part2(prepareInput(file.LoadStrings("test.txt"))))
	fmt.Println("part2:", part2(prepareInput(file.LoadStrings("input.txt"))))
}

func prepareInput(lines []string) field.Board {
	return field.New(lines)
}

func part1(board field.Board) int64 {
	var foundParts []int64

	possibleParts := findPossibleParts(board)
	for _, part := range possibleParts {
		if lo.SomeBy(part.neighbours(board), func(item byte) bool {
			return item != '.' && (item < '0' || item > '9')
		}) {
			foundParts = append(foundParts, conv.Int(string(board.HorizontalSlice(part.Start(), part.End()))))
		}
	}

	return lo.Sum(foundParts)
}

func part2(board field.Board) int64 {
	var result int64

	parts := lo.GroupBy(findPossibleParts(board), func(item Part) int {
		return item.Start().Row()
	})
	for _, gear := range board.All('*') {
		possibleNeighbours := lo.Flatten([][]Part{
			parts[gear.Row()-1],
			parts[gear.Row()],
			parts[gear.Row()+1],
		})

		neighbours := lo.Filter(possibleNeighbours, func(item Part, _ int) bool {
			if item.Start().Col() > gear.Col()+1 {
				return false
			}

			if item.End().Col() < gear.Col()-1 {
				return false
			}

			return true
		})

		if len(neighbours) == 2 {
			result += conv.Int(string(board.HorizontalSlice(neighbours[0].Start(), neighbours[0].End()))) *
				conv.Int(string(board.HorizontalSlice(neighbours[1].Start(), neighbours[1].End())))
		}

	}

	return result
}

type Part [2]field.Point

func (p Part) Start() field.Point {
	return p[0]
}
func (p Part) End() field.Point {
	return p[1]
}
func (p Part) neighbours(board field.Board) []byte {
	var res []byte
	for row := p.Start().Row() - 1; row <= p.Start().Row()+1; row++ {
		for col := p.Start().Col() - 1; col <= p.End().Col()+1; col++ {
			// пропускаем сами себя
			if row == p.Start().Row() && col >= p.Start().Col() && col <= p.End().Col() {
				continue
			}

			char := board.At(field.Point{row, col})
			if char == nil {
				continue
			}

			res = append(res, *char)
		}
	}
	return res
}

func findPossibleParts(board field.Board) []Part {
	var parts []Part
	for _, line := range board.Rows() {
		var (
			start, end   field.Point
			insideNumber bool
		)

		for _, point := range line {
			curChar := *board.At(point)
			if curChar < '0' || curChar > '9' {
				if insideNumber {
					insideNumber = false
					parts = append(parts, Part{start, end})
				}

				continue
			}

			if !insideNumber {
				start = point
				insideNumber = true
			}
			end = point
		}

		if insideNumber {
			parts = append(parts, Part{start, end})
		}
	}
	return parts
}
