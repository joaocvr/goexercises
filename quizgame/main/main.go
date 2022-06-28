package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	filePath, timeLimitSeconds := getParams()

	questions := readProblemsFile(filePath)

	askQuestions(questions, timeLimitSeconds)

	printResult(questions)
}

func getParams() (string, int64) {
	filePath := ""
	timeLimitSeconds := int64(5)

	args := os.Args
	for i, arg := range args {
		if arg == "-f" && (i+1) < len(args) {
			filePath = args[i+1]
		} else if arg == "-t" && (i+1) < len(args) {
			timeLimitSeconds, _ = strconv.ParseInt(args[i+1], 10, 64)
		}
	}

	if filePath == "" {
		filePath = "../../quizgame/main/problems.csv"
	}

	return filePath, timeLimitSeconds
}

func readProblemsFile(filePath string) []Question {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		if err = file.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	fileBytes, err := ioutil.ReadAll(file)

	lines := strings.Split(string(fileBytes), "\n")

	var questions []Question

	for _, line := range lines {
		lineData := strings.Split(line, ",")
		problem := lineData[0]
		solution, _ := strconv.ParseInt(lineData[1], 10, 64)
		questions = append(questions, Question{Problem: problem, Solution: solution})
	}

	return questions
}

func askQuestions(questions []Question, timeLimitSeconds int64) {

	keyPressed := make(chan []byte)
	go func(keyPressed chan []byte) {
		var b []byte = make([]byte, 1)
		fmt.Printf("Press ENTER to start.")
		for {
			os.Stdin.Read(b)
			keyPressed <- b
		}
	}(keyPressed)

	for {
		stdin := <-keyPressed
		enter := append([]byte{}, 10)
		if stdin[0] == enter[0] {
			break
		}
	}

	fmt.Printf("You have %d seconds to answer the questions.\n", timeLimitSeconds)
	deadline := time.Second * time.Duration(timeLimitSeconds)
	ctx, cancel := context.WithTimeout(context.TODO(), deadline)
	defer cancel()
	ask := func() {
		for index := range questions {
			questions[index].Ask()
		}
	}
	go ask()

	select {
	case <-ctx.Done():
		fmt.Println("\nTime is out!")
	case <-time.After(deadline):
		fmt.Println("\nTime is out!")
	}
}

func printResult(questions []Question) {
	correct := 0
	answered := 0
	for _, question := range questions {
		if question.IsAnswered {
			answered++
			if question.IsCorrect {
				correct++
			}
		}
	}

	fmt.Printf("%d answered questions from %d, and %d of them is correct.\n", answered, len(questions), correct)
}
