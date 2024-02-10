package grep

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
)

type GrepOptions struct {
	Pattern       string
	FilePath      string
	PrintLineNum  bool
	IgnoreCase    bool
	InvertMatch   bool
	FixedStrings  bool
	BeforeContext int
	AfterContext  int
	CountOnly     bool
}

func parseFlags() (*GrepOptions, error) {
	opts := &GrepOptions{}

	flag.StringVar(&opts.Pattern, "pattern", "", "Pattern to search for")
	flag.IntVar(&opts.BeforeContext, "B", 0, "Print N lines before match")
	flag.IntVar(&opts.AfterContext, "A", 0, "Print N lines after match")
	flag.BoolVar(&opts.CountOnly, "c", false, "Show only count of matching lines")
	flag.BoolVar(&opts.IgnoreCase, "i", false, "Ignore case distinctions")
	flag.BoolVar(&opts.InvertMatch, "v", false, "Select non-matching lines")
	flag.BoolVar(&opts.FixedStrings, "F", false, "Interpret pattern as a fixed string")
	flag.BoolVar(&opts.PrintLineNum, "n", false, "Print line numbers with output lines")

	flag.Parse()

	if len(flag.Args()) > 0 {
		opts.FilePath = flag.Args()[0]
	} else {
		return nil, errors.New("no file path provided")
	}

	return opts, nil
}
func grep(opts *GrepOptions) error {
	file, err := os.Open(opts.FilePath)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lineNum := 1
	matchCount := 0

	for scanner.Scan() {
		line := scanner.Text()
		if opts.IgnoreCase {
			line = strings.ToLower(line)
		}

		var matched bool
		if opts.FixedStrings {
			matched = strings.Contains(line, opts.Pattern)
		} else {
			matched, err = regexp.MatchString(opts.Pattern, line)
			if err != nil {
				return err
			}
		}

		if (matched && !opts.InvertMatch) || (!matched && opts.InvertMatch) {
			if opts.CountOnly {
				matchCount++
			} else {
				if opts.PrintLineNum {
					fmt.Printf("%d: ", lineNum)
				}
				fmt.Println(line)
			}
		}

		lineNum++
	}

	if opts.CountOnly {
		fmt.Println(matchCount)
	}

	return scanner.Err()
}
func main() {
	opts, err := parseFlags()
	if err != nil {
		log.Fatal(err)
	}

	err = grep(opts)
	if err != nil {
		log.Fatal(err)
	}
}
