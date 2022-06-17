package main

import (
	"fmt"
	"os"
)

type Question struct {
	Problem   string
	Solution  int64
	Answer    int64
	IsCorrect bool
}

func(q *Question) Ask() {
	reader := os.Stdin
	fmt.Printf("What %s, sir?", q.Problem)
	
	var result int64
	fmt.Fscanln(reader, &result)
	q.Answer = result

	if result == q.Solution {
		q.IsCorrect = true
	} else {
		q.IsCorrect = false
	}
}