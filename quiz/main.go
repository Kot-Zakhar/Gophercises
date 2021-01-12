package main

import (
	"fmt"
	"os"
	"encoding/csv"
	"io"
	"strings"
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
	var tasks, _ = getTasks("./problems.csv")
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
