package unpack

import (
	"errors"
	"regexp"
	"strconv"
	"strings"
)

var unpackPattern = regexp.MustCompile(`(\\*)(\d*)([^\\\d]?)`)

func Unpack(input string) (string, error) {
	result := ""
	lastIndex := 0

	matches := unpackPattern.FindAllStringSubmatchIndex(input, -1)
	for _, match := range matches {
		text := input[match[2]:match[3]]
		numStr := input[match[4]:match[5]]
		char := input[match[6]:match[7]]

		count := 1
		if numStr != "" {
			var err error
			count, err = strconv.Atoi(numStr)
			if err != nil {
				return "", errors.New("Некорректное число после символа")
			}
		}

		if char != "" && char != "\\" {
			result += strings.Repeat(char, count)
		} else {
			result += text
		}

		lastIndex = match[7]
	}

	if lastIndex < len(input) {
		result += input[lastIndex:]
	}

	return result, nil
}
