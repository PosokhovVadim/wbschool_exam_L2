package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"regexp"
	"strings"
)

/*
=== Утилита grep ===

Реализовать утилиту фильтрации (man grep)

Поддержать флаги:
-A - "after" печатать +N строк после совпадения
-B - "before" печатать +N строк до совпадения
-C - "context" (A+B) печатать ±N строк вокруг совпадения
-c - "count" (количество строк)
-i - "ignore-case" (игнорировать регистр)
-v - "invert" (вместо совпадения, исключать)
-F - "fixed", точное совпадение со строкой, не паттерн
-n - "line num", печатать номер строки

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

var (
	folderPath = "develop/dev05/" // to run from root
)

var (
	a     int  // DONE
	b     int  // DONE
	ctx   int  // DONE
	count bool // DONE
	i     bool // DONE
	v     bool // DONE
	f     bool // DONE
	n     bool // DONE
)

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

func isMatch(line string, pattern string, re *regexp.Regexp) bool {
	if f {
		if i {
			return strings.EqualFold(line, pattern)
		}
		return strings.Contains(line, pattern)
	} else {
		return re.MatchString(line)
	}
}

func compileRegex(pattern string) *regexp.Regexp {
	if i {
		pattern = "(?i)" + pattern
	}
	return regexp.MustCompile(pattern)
}

func grepProcess(lines []string, pattern string) []string {
	resultMap := make(map[int]string)
	var resultLines []string
	re := compileRegex(pattern)
	left := max(b, ctx)
	right := max(a, ctx)

	for idx, line := range lines {
		matched := isMatch(line, pattern, re)

		if matched == v {
			continue
		}

		start := max(0, idx-left)
		end := min(len(lines)-1, idx+right)

		for i := start; i <= end; i++ {
			if _, ok := resultMap[i]; !ok {
				resultMap[i] = lines[i]
			}

		}
	}

	for key, v := range resultMap {
		if n {
			resultLines = append(resultLines, fmt.Sprintf("%d: %s", key+1, v))
		} else {
			resultLines = append(resultLines, v)
		}
	}

	return resultLines
}

func printLines(lines []string) {

	if count {
		fmt.Println(len(lines))
		return
	}

	fmt.Fprintf(os.Stdout, "%s\n", strings.Join(lines, "\n"))

}
func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

func run() error {
	flag.IntVar(&a, "A", 0, "print n strings after match")
	flag.IntVar(&b, "B", 0, "print n strings before match")
	flag.IntVar(&ctx, "C", 0, "print ±n strings around match")
	flag.BoolVar(&count, "c", false, "count matches")
	flag.BoolVar(&i, "i", false, "ignore case")
	flag.BoolVar(&v, "v", false, "invert match")
	flag.BoolVar(&f, "F", false, "find fixed string")
	flag.BoolVar(&n, "n", false, "print line number")

	flag.Parse()
	args := flag.Args()

	if len(args) != 2 {
		return fmt.Errorf("expected pattern and file path, got %v", args)
	}

	lines, err := readFile(args[1])
	if err != nil {
		return err
	}

	pattern := args[0]

	result := grepProcess(lines, pattern)

	printLines(result)
	return nil
}

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}
