package main

import (
	"fmt"
	"github.com/geshido/aoc-2023-go/util/conv"
	"github.com/geshido/aoc-2023-go/util/file"
	"regexp"
	"strings"
)

func main() {
	fmt.Println("Part01 test:", part01(file.LoadStrings("test.txt")))
	fmt.Println("Part01:", part01(file.LoadStrings("part01.txt")))
	fmt.Println("Part02 test:", part02(file.LoadStrings("test2")))
	fmt.Println("Part02:", part02(file.LoadStrings("part01.txt")))
}

func part01(list []string) int64 {
	var res int64
	re := regexp.MustCompile("[0-9]")

	for _, s := range list {
		digits := re.FindAllString(s, -1)
		cur := conv.Int(digits[0] + digits[len(digits)-1])
		res += cur
	}

	return res
}

func part02(list []string) int64 {
	var res int64

	for _, s := range list {
		var digits []string
		for idx, c := range s {
			switch c {
			case '1', '2', '3', '4', '5', '6', '7', '8', '9':
				digits = append(digits, string(c))
			default:
				switch 0 {
				case strings.Index(s[idx:], "one"):
					digits = append(digits, "1")
				case strings.Index(s[idx:], "two"):
					digits = append(digits, "2")
				case strings.Index(s[idx:], "three"):
					digits = append(digits, "3")
				case strings.Index(s[idx:], "four"):
					digits = append(digits, "4")
				case strings.Index(s[idx:], "five"):
					digits = append(digits, "5")
				case strings.Index(s[idx:], "six"):
					digits = append(digits, "6")
				case strings.Index(s[idx:], "seven"):
					digits = append(digits, "7")
				case strings.Index(s[idx:], "eight"):
					digits = append(digits, "8")
				case strings.Index(s[idx:], "nine"):
					digits = append(digits, "9")
				}
			}
		}

		res += conv.Int(digits[0] + digits[len(digits)-1])
	}

	return res
}
