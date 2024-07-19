package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/*
=== Утилита cut ===

Принимает STDIN, разбивает по разделителю (TAB) на колонки, выводит запрошенные

Поддержать флаги:
-f - "fields" - выбрать поля (колонки)
-d - "delimiter" - использовать другой разделитель
-s - "separated" - только строки с разделителем

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

var (
	f []int
	d string
	s bool
)

func getFields(fList string) error {

	splitFields := strings.Split(fList, ",")
	for _, v := range splitFields {
		num, err := strconv.Atoi(v)
		if err != nil {
			return err
		}
		f = append(f, num)

	}
	return nil
}

func readData() ([]string, error) {
	scanner := bufio.NewScanner(os.Stdin)
	var data []string
	for scanner.Scan() {
		data = append(data, scanner.Text())
	}
	return data, nil
}

func cut(data []string) []string {
	var result []string
	var cutLine strings.Builder
	for _, line := range data {
		splitData := strings.Split(line, d)

		fmt.Printf("splitData: %v, len: %v\n", splitData, len(splitData))
		if s && len(splitData) == 1 {
			continue
		}

		for _, v := range f {
			if v <= len(splitData) {
				cutLine.WriteString(splitData[v-1])
				cutLine.WriteString(d)
			}

		}

		result = append(result, strings.TrimRight(cutLine.String(), d))
		cutLine.Reset()
	}

	return result
}

func run() error {
	var fList string
	flag.StringVar(&fList, "f", "", "fields")
	flag.StringVar(&d, "d", "\t", "delimiter")
	flag.BoolVar(&s, "s", false, "separated")

	flag.Parse()

	if fList == "" {
		return fmt.Errorf("fields is empty")
	}

	if err := getFields(fList); err != nil {
		return err
	}

	data, err := readData()
	if err != nil {
		return err
	}
	result := cut(data)

	fmt.Println(strings.Join(result, "\n"))
	return nil
}

func main() {

	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}
