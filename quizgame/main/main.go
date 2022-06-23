package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	//GET FILE PATH AND TIME LIMIT PER QUESTION
	filePath, timeLimite := getParams()
	fmt.Println(filePath, timeLimite)

	//READ CSV FILE
	questions := readProblemsFile(filePath)

	//ASK QUESTIONS
	askQuestions(questions)

	//SHOW RESULT
	printResult(questions)
}

func getParams() (string, int64) {
	filePath := ""
	timeLimit := int64(30)

	args := os.Args
	for i, arg := range args {
		if arg == "-f" && (i+1) < len(args) {
			filePath = args[i+1]
		} else if arg == "-t" && (i+1) < len(args) {
			timeLimit, _ = strconv.ParseInt(args[i+1], 10, 64)
		}
	}

	if filePath == "" {
		filePath = "../../quizgame/main/problems.csv"
	}

	return filePath, timeLimit
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

func askQuestions(questions []Question) {
	for index := range questions {
		questions[index].Ask()
	}
}

func printResult(questions []Question) {
	correct := 0
	for _, question := range questions {
		if question.IsCorrect {
			correct++
		}
	}

	fmt.Printf("Total of %d correct quesitons of %d", correct, len(questions))
}
