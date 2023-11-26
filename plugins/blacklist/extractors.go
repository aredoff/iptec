package blacklist

import (
	"strings"
)

func simpleExtracotor(data string) []string {
	lines := []string{}
	for _, line := range strings.Split(data, "\n") {
		if len(line) < 2 || len(line) > 45 || line[0] == '#' {
			continue
		}
		lines = append(lines, line)
	}
	return lines
}

func firstAddressSpaсeExtracotor(data string) []string {
	lines := []string{}
	for _, line := range strings.Split(data, "\n") {
		if len(line) < 2 || line[0] == '#' {
			continue
		}
		words := strings.Fields(line)

		if len(words) > 0 {
			lines = append(lines, words[0])
		}
	}
	return lines
}
