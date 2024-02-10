package sort

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"sort"
)

type Line struct {
	Index int
	Value string
}

var kFlag *string
var nFlag bool
var rFlag bool
var uFlag bool
var MFlag bool
var bFlag bool
var cFlag bool
var hFlag bool

func init() {
	kFlag = flag.String("k", "", "column number for sorting")
	flag.BoolVar(&nFlag, "n", false, "sort numerically")
	flag.BoolVar(&rFlag, "r", false, "reverse the result of comparisons")
	flag.BoolVar(&uFlag, "u", false, "output only the first of an equal run")
	flag.BoolVar(&MFlag, "M", false, "sort by month names")
	flag.BoolVar(&bFlag, "b", false, "ignore leading blanks")
	flag.BoolVar(&cFlag, "c", false, "check for sorted input")
	flag.BoolVar(&hFlag, "h", false, "compare human readable numbers")
}

func readLines(filename string) ([]string, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	var lines []string
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func compareLines(a, b *Line) bool {
	// Реализация сравнения в зависимости от флагов
}

func sortLines(lines []*Line) {
	sort.SliceStable(lines, func(i, j int) bool {
		return compareLines(lines[i], lines[j])
	})
}

func removeDuplicates(lines []*Line) []*Line {
	seen := make(map[string]struct{})
	j := 0
	for _, v := range lines {
		if _, ok := seen[v.Value]; ok {
			continue
		}
		seen[v.Value] = struct{}{}
		lines[j] = v
		j++
	}
	return lines[:j]
}

func writeLines(filename string, lines []*Line) error {
	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	w := bufio.NewWriter(f)
	for _, line := range lines {
		fmt.Fprintln(w, line.Value)
	}
	return w.Flush()
}

func main() {
	flag.Parse()

	lines, err := readLines(flag.Arg(0))
	if err != nil {
		log.Fatal(err)
	}

	// Преобразование строк в структуры Line
	// Применение функции сортировки и удаление дубликатов
	// Запись результата в файл

	if err := writeLines("output.txt", sortedLines); err != nil {
		log.Fatal(err)
	}
}
