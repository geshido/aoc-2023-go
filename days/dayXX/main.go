package main

import (
	"fmt"
	"github.com/geshido/aoc-2023-go/util/file"
)

func main() {
	fmt.Println("part1 test:", part1(prepareInput(file.LoadStrings("test.txt"))))
	fmt.Println("part1:", part1(prepareInput(file.LoadStrings("input.txt"))))
	fmt.Println("part2 test:", part2(prepareInput(file.LoadStrings("test.txt"))))
	fmt.Println("part2:", part2(prepareInput(file.LoadStrings("input.txt"))))
}

func prepareInput(lines []string) []string {
	return lines
}

func part1(lines []string) int64 {
	return 0
}

func part2(lines []string) int64 {
	return 0
}
