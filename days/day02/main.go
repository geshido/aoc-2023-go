package main

import (
	"fmt"
	"github.com/geshido/aoc-2023-go/util/file"
	"github.com/samber/lo"
	"strings"
)

type Game struct {
	ID   int64
	Sets []Set
}

func (g Game) MaxGreen() int64 {
	return lo.MaxBy(g.Sets, func(s1, s2 Set) bool {
		return s1.Green > s2.Green
	}).Green
}
func (g Game) MaxBlue() int64 {
	return lo.MaxBy(g.Sets, func(s1, s2 Set) bool {
		return s1.Blue > s2.Blue
	}).Blue
}
func (g Game) MaxRed() int64 {
	return lo.MaxBy(g.Sets, func(s1, s2 Set) bool {
		return s1.Red > s2.Red
	}).Red
}
func (g Game) Power() int64 {
	return g.MaxGreen() * g.MaxRed() * g.MaxBlue()
}

func (g Game) Valid(limit Set) bool {
	return g.MaxGreen() <= limit.Green && g.MaxRed() <= limit.Red && g.MaxBlue() <= limit.Blue
}

type Set struct {
	Red   int64
	Blue  int64
	Green int64
}

func prepareInput(lines []string) []Game {
	return lo.Map(lines, func(item string, _ int) Game {
		g := Game{}

		parts := strings.Split(item, ": ")

		_, err := fmt.Sscanf(parts[0], "Game %d", &g.ID)
		if err != nil {
			panic(err)
		}

		g.Sets = lo.Map(strings.Split(parts[1], "; "), func(s string, _ int) Set {
			set := Set{}

			for _, item := range strings.Split(s, ", ") {
				var color string
				var n int64

				_, err := fmt.Sscanf(item, "%d %s", &n, &color)
				if err != nil {
					panic(err)
				}

				switch color {
				case "red":
					set.Red = n
				case "blue":
					set.Blue = n
				case "green":
					set.Green = n
				}
			}

			return set
		})

		return g
	})
}

func main() {
	fmt.Println("part1 test:", part1(prepareInput(file.LoadStrings("test1.txt"))))
	fmt.Println("part1:", part1(prepareInput(file.LoadStrings("input01.txt"))))
	fmt.Println("part2 test:", part2(prepareInput(file.LoadStrings("test1.txt"))))
	fmt.Println("part2:", part2(prepareInput(file.LoadStrings("input01.txt"))))
}

func part1(games []Game) int64 {
	limit := Set{
		Red:   12,
		Blue:  14,
		Green: 13,
	}

	var sumIds int64
	for _, g := range games {
		if g.Valid(limit) {
			sumIds += g.ID
		}
	}

	return sumIds
}

func part2(games []Game) int64 {
	return lo.SumBy(games, func(g Game) int64 {
		return g.Power()
	})
}
