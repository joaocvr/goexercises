package main

import (
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	//GET FILE PATH
	filePath := getFilePath()

	//READ CSV FILE
	questions := readProblemsFile(filePath)

	//ASK QUESTIONS
	askQuestions(questions)

}

func getFilePath() string {
	filePath := ""
	if len(os.Args) > 1 {
		filePath = os.Args[1]
	} else {
		filePath = "../../quizgame/main/problems.csv"
	}

	return filePath
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
