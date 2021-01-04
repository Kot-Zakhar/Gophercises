package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type Task struct {
	Question string
	Answer string
}

func getTasks(filename string) (tasks []Task, err error) {
	content_bytes, err := ioutil.ReadFile(filename)
	content := string(content_bytes)
	if err != nil {
		return nil, err
	}

	lines := strings.Split(content, "\n")

	for _, line := range lines {
		if line == "" {
			continue
		}
		var values = strings.Split(line, ",")
		tasks = append(tasks, Task{values[0], values[1]})
	}

	return tasks, nil
}

func main() {
	var tasks, _ = getTasks("./problems.csv")

	for i, v := range tasks {
		fmt.Printf("%d) %s %s\n", i, v.Question, v.Answer)
	}
}
