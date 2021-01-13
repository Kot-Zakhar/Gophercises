package main

import (
	"fmt"
	"os"
	"encoding/csv"
	"io"
	"strings"
	"flag"
	"time"
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

	csvReader := csv.NewReader(file)

	for {
		record, err := csvReader.Read()
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
	var (
		filename string
		limit time.Duration
	)

	flag.StringVar(&filename, "csv", "./problems.csv", "a csv file in the format of 'question,answer'")
	flag.DurationVar(&limit, "limit", 30 * time.Second, "the time limit for the quiz in seconds")

	flag.Parse()

	tasks, err := getTasks(filename)
	if err != nil {
		fmt.Print(err)
		return
	}

	var (
		correct = 0
		answer string
	)

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
