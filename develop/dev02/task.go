package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

/*
=== Задача на распаковку ===

Создать Go функцию, осуществляющую примитивную распаковку строки, содержащую повторяющиеся символы / руны, например:
  - "a4bc2d5e" => "aaaabccddddde"
  - "abcd" => "abcd"
  - "45" => "" (некорректная строка)
  - "" => ""

Дополнительное задание: поддержка escape - последовательностей
  - qwe\4\5 => qwe45 (*)
  - qwe\45 => qwe44444 (*)
  - qwe\\5 => qwe\\\\\ (*)

В случае если была передана некорректная строка функция должна возвращать ошибку. Написать unit-тесты.

Функция должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/
var errInvalidString = errors.New("invalid string")

// TODO: добавить поддержку escape
func unpack(s string) (string, error) {
	var res strings.Builder
	var prev rune

	for i, r := range s {
		if i == 0 && unicode.IsDigit(r) {
			return "", errInvalidString
		}

		if !unicode.IsDigit(r) {
			res.WriteRune(r)
			prev = r
			continue
		} else {
			if unicode.IsDigit(prev) {
				return "", errInvalidString
			}

			n, err := strconv.Atoi(string(r))

			if err != nil {
				return "", err
			}
			res.WriteString(strings.Repeat(string(prev), n-1))
			prev = r
		}
	}

	return res.String(), nil
}

func main() {
	res, err := unpack("")
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("result: %v\n", res)
}
