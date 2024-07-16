package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

/*
=== Утилита sort ===

Отсортировать строки (man sort)
Основное

Поддержать ключи

-k — указание колонки для сортировки
-n — сортировать по числовому значению
-r — сортировать в обратном порядке
-u — не выводить повторяющиеся строки

Дополнительное

Поддержать ключи

-M — сортировать по названию месяца
-b — игнорировать хвостовые пробелы
-c — проверять отсортированы ли данные
-h — сортировать по числовому значению с учётом суффиксов

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

var (
	folderPath = "develop/dev03/" // to run from root
)

var (
	k int  // DONE
	n bool // DONE
	r bool // DONE
	u bool // DONE
	m bool // DONE
	b bool // DONE
	c bool // DONE
	h bool // DONE
)

var (
	months = map[string]int{
		"jan": 1, "feb": 2, "mar": 3,
		"apr": 4, "may": 5, "jun": 6,
		"jul": 7, "aug": 8, "sep": 9,
		"oct": 10, "nov": 11, "dec": 12,
	}
)

var suffixes = map[string]float64{
	"k": 1e3,
	"M": 1e6,
	"g": 1e9,
	"t": 1e12,
	"p": 1e15,
	"e": 1e18,
}

func readFile(inFile string) ([]string, error) {
	file, err := os.Open(folderPath + inFile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}

	defer file.Close()

	var lines []string
	sc := bufio.NewScanner(file)

	for sc.Scan() {
		lines = append(lines, sc.Text())
	}

	return lines, sc.Err()
}

func writeFile(lines []string, inFile string) error {

	file, err := os.Create(folderPath + "sorted-" + inFile)

	if err != nil {
		return err
	}
	defer file.Close()

	for _, line := range lines {
		_, err := file.WriteString(line + "\n")
		if err != nil {
			return err
		}
	}

	return nil
}

func compareLines(a, b string) bool {

	if k <= 0 {
		if r {
			return a > b
		}
		return a < b
	}

	aFields := strings.Split(a, " ")
	bFields := strings.Split(b, " ")

	var akey, bkey string

	if k <= len(aFields) {
		akey = aFields[k-1]
	}
	if k <= len(bFields) {
		bkey = bFields[k-1]
	}

	if n {
		aNum, err := strconv.ParseFloat(akey, 64)
		if err != nil {
			return false
		}
		bNum, err := strconv.ParseFloat(bkey, 64)
		if err != nil {
			return false
		}
		if r {
			return aNum > bNum
		}
		return aNum < bNum
	}

	if m {
		aMonth, aOk := months[akey]
		bMonth, bOk := months[bkey]

		if aOk && bOk {
			if r {
				return aMonth > bMonth
			}
			return aMonth < bMonth
		}
	}

	if h {
		aNum := getNumWithSuffix(akey)
		bNum := getNumWithSuffix(bkey)

		if r {
			return aNum > bNum
		}
		return aNum < bNum
	}

	if r {
		return akey > bkey
	}

	return akey < bkey
}

func getNumWithSuffix(s string) float64 {

	for suf, val := range suffixes {
		if strings.HasSuffix(s, suf) {
			num, err := strconv.ParseFloat(s[:len(s)-len(suf)], 64)
			if err != nil {
				return 0
			}
			return num * val
		}
	}

	return 0
}
func unique(lines []string) []string {
	uMap := make(map[string]bool)

	res := make([]string, 0, len(lines))
	for _, line := range lines {
		if !uMap[line] {
			uMap[line] = true
			res = append(res, line)
		}
	}
	return res
}

func isSorted(lines []string) bool {
	for i := 1; i < len(lines); i++ {
		if compareLines(lines[i-1], lines[i]) {
			return false
		}
	}
	return true
}

func run() error {
	flag.IntVar(&k, "k", 1, "column number to sort by")
	flag.BoolVar(&n, "n", false, "sort by numerical value")
	flag.BoolVar(&r, "r", false, "reverse sort")
	flag.BoolVar(&u, "u", false, "skip duplicate lines")

	flag.BoolVar(&m, "M", false, "sort by month name")
	flag.BoolVar(&b, "b", false, "ignore trailing spaces")
	flag.BoolVar(&c, "c", false, "check for sorted input")
	flag.BoolVar(&h, "h", false, "compare human readable numbers")
	flag.Parse()

	if flag.NArg() != 1 {
		return fmt.Errorf("expected file name, got %v", flag.Arg(0))
	}

	inFile := flag.Arg(0)

	lines, err := readFile(inFile)

	if err != nil {
		return err
	}

	if c {
		if isSorted(lines) {
			fmt.Println("input file is sorted")
			return nil
		}
		return fmt.Errorf("input file is not sorted")
	}

	if b {
		for i := range lines {
			lines[i] = strings.TrimRight(lines[i], " ")
		}
	}
	sort.SliceStable(lines, func(i, j int) bool {

		return compareLines(lines[i], lines[j])
	})

	if u {
		lines = unique(lines)
	}

	if err := writeFile(lines, inFile); err != nil {
		return err
	}

	return nil

}

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}
