package file

import (
	"os"
	"strings"
)

func LoadStrings(filename string) []string {
	b, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	return strings.Split(string(b), "\n")
}

func LoadString(filename string) string {
	b, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	return string(b)
}
