package field

import (
	"github.com/samber/lo"
	"strings"
)

type Board struct {
	data [][]byte
}

func New(lines []string) Board {
	f := Board{
		data: lo.Map(lines, func(line string, _ int) []byte {
			return []byte(line)
		}),
	}

	return f
}

func (b Board) At(p Point) *byte {
	if p.Row() < 0 || p.Row() > len(b.data)-1 {
		return nil
	}

	if p.Col() < 0 || p.Col() > len(b.data[0])-1 {
		return nil
	}

	return &b.data[p.Row()][p.Col()]
}

func (b Board) Rows() [][]Point {
	res := make([][]Point, len(b.data))
	for row := range b.data {
		for col := range b.data[row] {
			res[row] = append(res[row], Point{row, col})
		}
	}
	return res
}

func (b Board) All(char byte) []Point {
	var res []Point
	for row := range b.data {
		for col := range b.data[row] {
			if b.data[row][col] == char {
				res = append(res, Point{row, col})
			}
		}
	}
	return res
}

func (b Board) HorizontalSlice(start, end Point) []byte {
	if start.Row() != end.Row() {
		return nil
	}

	return b.data[start.Row()][start.Col() : end.Col()+1]
}

func (b Board) String() string {
	sb := strings.Builder{}
	for _, row := range b.data {
		sb.Write(row)
		sb.Write([]byte("\r\n"))
	}

	return sb.String()
}
