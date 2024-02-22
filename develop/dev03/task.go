package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
	"unicode"
)

/*

	Отсортировать строки в файле по аналогии с консольной утилитой sort (man sort — смотрим описание и основные параметры): на входе подается файл из несортированными строками, на выходе — файл с отсортированными.

	Реализовать поддержку утилитой следующих ключей:

	-k — указание колонки для сортировки (слова в строке могут выступать в качестве колонок, по умолчанию разделитель — пробел)
	-n — сортировать по числовому значению
	-r — сортировать в обратном порядке
	-u — не выводить повторяющиеся строки

*/

func main() {
	os.Exit(cli(os.Args[1:]))
}

// CLI runs the go-sort command line app and returns its exit status.
func cli(args []string) int {
	var app appEnv
	err := app.fromArgs(args)
	if err != nil {
		return 2
	}
	if err = app.run(); err != nil {
		fmt.Fprintf(os.Stderr, "Runtime error: %v\n", err)
		return 1
	}
	return 0
}

type appEnv struct {
	isNumeric       bool
	isReverse       bool
	deleteDuplicate bool
	column          int
	reader          io.ReadCloser
}

func (app *appEnv) fromArgs(args []string) error {
	fl := flag.NewFlagSet("sortfile", flag.ContinueOnError)
	fl.IntVar(&app.column, "k", 1, "sort via a column")
	fl.BoolVar(&app.isNumeric, "n", false, "compare according to string numerical value")
	fl.BoolVar(&app.isReverse, "r", false, "reverse the result of comparisons")
	fl.BoolVar(&app.deleteDuplicate, "u", false, "delete duplicate strings")

	if err := fl.Parse(args); err != nil {
		fl.Usage()
		return err
	}

	stat, _ := os.Stdin.Stat()
	if (stat.Mode() & os.ModeCharDevice) == 0 {
		app.reader = os.Stdin
		return nil
	}

	file, err := os.Open(fl.Arg(0))
	if err != nil {
		fmt.Fprintf(os.Stderr, "can't open file %s: %v\n", fl.Arg(0), err)
		return err
	}
	app.reader = file

	return nil
}

func (app *appEnv) run() error {
	defer app.reader.Close()
	data := make([]string, 0)

	scanner := bufio.NewScanner(app.reader)
	for scanner.Scan() {
		data = append(data, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return err
	}

	if app.column == 1 && !app.isNumeric {
		data = app.sort(data)
		writeToOutput(data)
		return nil
	}

	data = app.sortColumns(data)
	writeToOutput(data)

	return nil
}

func (app *appEnv) sort(data []string) []string {
	if app.isReverse {
		sort.Sort(sort.Reverse(sort.StringSlice(data)))
	} else {
		sort.Strings(data)
	}

	if app.deleteDuplicate {
		data = delDuplicate(data)
	}

	return data
}

func (app *appEnv) sortColumns(data []string) []string {
	t := stringTable{
		data:      make([][]string, 0, len(data)),
		column:    app.column - 1,
		isNumeric: app.isNumeric,
	}

	for _, v := range data {
		t.data = append(t.data, strings.Fields(v))
	}

	if app.isReverse {
		sort.Sort(sort.Reverse(t))
	} else {
		sort.Sort(t)
	}

	for i, v := range t.data {
		data[i] = strings.Join(v, " ")
	}

	if app.deleteDuplicate {
		data = delDuplicate(data)
	}

	return data
}

func delDuplicate(data []string) []string {
	exists := make(map[string]struct{}, len(data))
	res := make([]string, 0, len(data))
	for _, v := range data {
		if _, ok := exists[v]; ok {
			continue
		}
		res = append(res, v)
		exists[v] = struct{}{}
	}

	return res
}

func writeToOutput(data []string) {
	for _, v := range data {
		fmt.Fprintf(os.Stdout, "%s\n", v)
	}
}

// trimNonNumber deletes non number runes from the end of the string
func trimNonNumber(str string) string {
	return strings.TrimRightFunc(str, func(r rune) bool {
		return !unicode.IsNumber(r)
	})
}

type stringTable struct {
	data      [][]string
	column    int
	isNumeric bool
}

func (t stringTable) Len() int {
	return len(t.data)
}

func (t stringTable) Less(i, j int) bool {
	col := t.column
	if col > len(t.data[i])-1 || col > len(t.data[j]) {
		col = 0
	}

	if t.isNumeric {
		n1 := trimNonNumber(t.data[i][col])
		n2 := trimNonNumber(t.data[j][col])

		i1, err := strconv.Atoi(n1)
		if err != nil {
			return (t.data[i][col] < t.data[j][col])
		}
		j1, err := strconv.Atoi(n2)
		if err != nil {
			return (t.data[i][col] < t.data[j][col])
		}

		return i1 < j1
	}
	return (t.data[i][col] < t.data[j][col])
}

func (t stringTable) Swap(i, j int) {
	t.data[i], t.data[j] = t.data[j], t.data[i]
}
