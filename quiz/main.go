package main

import (
	"fmt"
	"os"
	"encoding/csv"
	"io"
	"strings"
	"flag"
)

type Task struct {
	Question string
	Answer string
}

func getTasks(filename string) (tasks []Task, err error) {

	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	csv_reader := csv.NewReader(file)

	for {
		record, err := csv_reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}

		tasks = append(tasks, Task{record[0], record[1]})
	}

	return tasks, nil
}

func main() {
	var filename string
	var limit int

	flag.StringVar(&filename, "csv", "./problems.csv", "a csv file in the format of 'question,answer' (default \"problems.csv\"")
	flag.IntVar(&limit, "limit", 30, "the time limit for the quiz in seconds (default 30)")

	flag.Parse()

	var tasks, _ = getTasks(filename)
	// TODO: add error processing

	var correct = 0
	var answer string

	for i, v := range tasks {
		fmt.Printf("%d) %s\n", i, v.Question)
		fmt.Scanln(&answer)
		answer = strings.TrimSpace(answer)
		if answer == v.Answer {
			correct += 1
		}
	}

	fmt.Printf("\n%d correct answers from %d\n", correct, len(tasks))
}
