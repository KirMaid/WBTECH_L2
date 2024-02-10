package cut

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var fieldsFlag string
var delimiterFlag string
var separatedFlag bool

func init() {
	flag.StringVar(&fieldsFlag, "f", "", "выбрать поля (колонки)")
	flag.StringVar(&delimiterFlag, "d", "\t", "использовать другой разделитель")
	flag.BoolVar(&separatedFlag, "s", false, "только строки с разделителем")
}

func main() {
	flag.Parse()

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		if !separatedFlag || strings.Contains(line, delimiterFlag) {
			cols := strings.Split(line, delimiterFlag)
			selectedCols := selectColumns(cols, fieldsFlag)
			fmt.Println(strings.Join(selectedCols, delimiterFlag))
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "ошибка чтения ввода:", err)
	}
}

func selectColumns(cols []string, fieldsStr string) []string {
	var selected []string
	if fieldsStr != "" {
		fieldIndices := parseFieldIndices(fieldsStr)
		for _, idx := range fieldIndices {
			if idx < len(cols) {
				selected = append(selected, cols[idx])
			}
		}
	} else {
		selected = cols
	}
	return selected
}

func parseFieldIndices(fieldsStr string) []int {
	var indices []int
	for _, f := range strings.Split(fieldsStr, ",") {
		if idx, err := strconv.Atoi(f); err == nil {
			indices = append(indices, idx)
		}
	}
	return indices
}
