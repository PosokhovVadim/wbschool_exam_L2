package main

import (
	"bufio"
	"flag"
	"os"
	"testing"
)

func TestSortUtility(t *testing.T) {
	createFiles()
	tests := []struct {
		name      string
		args      []string
		inputFile string
		wantFile  string
	}{
		{
			name:      "sort by default",
			args:      []string{"input1.txt"},
			inputFile: "input1.txt",
			wantFile:  "expected1.txt",
		},
		{
			name:      "sort numerically",
			args:      []string{"-n", "input2.txt"},
			inputFile: "input2.txt",
			wantFile:  "expected2.txt",
		},
		{
			name:      "sort reverse",
			args:      []string{"-r", "input3.txt"},
			inputFile: "input3.txt",
			wantFile:  "expected3.txt",
		},
		{
			name:      "sort by month name",
			args:      []string{"-M", "input4.txt"},
			inputFile: "input4.txt",
			wantFile:  "expected4.txt",
		},
		{
			name:      "sort by second column",
			args:      []string{"-k", "2", "input5.txt"},
			inputFile: "input5.txt",
			wantFile:  "expected5.txt",
		},
		{
			name:      "sort by second column in reverse order",
			args:      []string{"-k", "2", "-r", "input6.txt"},
			inputFile: "input6.txt",
			wantFile:  "expected6.txt",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			inFile := tt.args[len(tt.args)-1]

			oldArgs := os.Args
			defer func() { os.Args = oldArgs }()

			flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError)
			os.Args = append([]string{"cmd"}, tt.args...)

			if err := run(); err != nil {
				t.Fatalf("run() error = %v", err)
			}

			got, err := readFile("sorted-" + inFile)
			if err != nil {
				t.Fatalf("failed to read output file: %v", err)
			}

			want, err := readFile(tt.wantFile)
			if err != nil {
				t.Fatalf("failed to read expected file: %v", err)
			}

			if !compareSlices(got, want) {
				t.Errorf("unexpected output for %s:\ngot:\n%v\nwant:\n%v", tt.name, got, want)
			}
		})
	}
}

func compareSlices(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func createFiles() {
	createFile("input1.txt", []string{
		"apple",
		"banana",
		"cherry",
		"apple",
		"banana",
	})

	createFile("expected1.txt", []string{
		"apple",
		"apple",
		"banana",
		"banana",
		"cherry",
	})

	createFile("input2.txt", []string{
		"3",
		"1",
		"4",
		"1",
		"5",
		"9",
		"2",
		"6",
		"5",
		"3",
		"5",
	})

	createFile("expected2.txt", []string{
		"1",
		"1",
		"2",
		"3",
		"3",
		"4",
		"5",
		"5",
		"5",
		"6",
		"9",
	})

	createFile("input3.txt", []string{
		"delta",
		"charlie",
		"bravo",
		"alpha",
		"echo",
		"foxtrot",
		"golf",
		"hotel",
		"india",
		"juliet",
	})

	createFile("expected3.txt", []string{
		"juliet",
		"india",
		"hotel",
		"golf",
		"foxtrot",
		"echo",
		"delta",
		"charlie",
		"bravo",
		"alpha",
	})

	createFile("input4.txt", []string{
		"jan",
		"feb",
		"mar",
		"apr",
		"may",
		"jun",
		"jul",
		"aug",
		"sep",
		"oct",
		"nov",
		"dec",
	})

	createFile("expected4.txt", []string{
		"jan",
		"feb",
		"mar",
		"apr",
		"may",
		"jun",
		"jul",
		"aug",
		"sep",
		"oct",
		"nov",
		"dec",
	})

	createFile("input5.txt", []string{
		"apple 3",
		"banana 1",
		"cherry 2",
		"apple 2",
		"banana 3",
	})

	createFile("expected5.txt", []string{
		"banana 1",
		"cherry 2",
		"apple 2",
		"apple 3",
		"banana 3",
	})

	createFile("input6.txt", []string{
		"apple 3",
		"banana 1",
		"cherry 2",
		"apple 2",
		"banana 3",
	})

	createFile("expected6.txt", []string{
		"apple 3",
		"banana 3",
		"cherry 2",
		"apple 2",
		"banana 1",
	})

}

func createFile(filename string, lines []string) {
	file, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	for _, line := range lines {
		_, err := writer.WriteString(line + "\n")
		if err != nil {
			panic(err)
		}
	}
	writer.Flush()

}
