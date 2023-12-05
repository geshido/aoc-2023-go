package field

type Point [2]int

func (p Point) Row() int {
	return p[0]
}
func (p Point) Col() int {
	return p[1]
}

func (p Point) Up() Point {
	return Point{p.Row() - 1, p.Col()}
}
func (p Point) Down() Point {
	return Point{p.Row() + 1, p.Col()}
}
func (p Point) Left() Point {
	return Point{p.Row(), p.Col() - 1}
}
func (p Point) Right() Point {
	return Point{p.Row(), p.Col() + 1}
}
func (p Point) AllNeighbours() []Point {
	return []Point{
		p.Up().Left(),
		p.Up(),
		p.Up().Right(),
		p.Right(),
		p.Down().Right(),
		p.Down(),
		p.Down().Left(),
		p.Left(),
	}
}
