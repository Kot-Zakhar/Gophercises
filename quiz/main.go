package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
	"time"
)

type Task struct {
	Question string
	Answer   string
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
		} else if err != nil {
			return nil, err
		}
		tasks = append(tasks, Task{record[0], record[1]})
	}

	return tasks, nil
}

func main() {
	var (
		filename string
		limit    time.Duration
	)

	flag.StringVar(&filename, "csv", "./problems.csv", "a csv file in the format of 'question,answer'")
	flag.DurationVar(&limit, "limit", 30*time.Second, "the time limit for the quiz in seconds")

	flag.Parse()

	tasks, err := getTasks(filename)
	if err != nil {
		fmt.Print(err)
		return
	}

	var (
		correct      = 0
		questionsAmt = 0
	)

	timer := time.NewTimer(limit)

	taskChannel := make(chan bool)

	questionFunc := func(i int, t *Task) {
		var answer string
		fmt.Printf("%d) %s\n", i+1, t.Question)
		fmt.Scanln(&answer)
		answer = strings.TrimSpace(answer)
		taskChannel <- answer == t.Answer
	}

	func() {
		for i, t := range tasks {
			go questionFunc(i, &t)
			select {
			case <-timer.C:
				fmt.Println("Time is up")
				return
			case answerRight := <-taskChannel:
				if answerRight {
					correct++
				}
				questionsAmt++
			}
		}
	}()

	fmt.Printf("\n%d correct answers from %d\n", correct, questionsAmt)
}
