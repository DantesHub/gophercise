package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

//main package for executing the quiz program
var correctAnswers []string
var totalQuestions []string
var usersAnswer []string
var numCorrectAnswers = 0;
func main() {
	reader := bufio.NewReader(os.Stdin)
	data, err := ioutil.ReadFile("problems.csv")
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v", err)
	}
	for _, line := range strings.Split(string(data), "\n") {
		totalQuestions = append(totalQuestions,(line[0:3]))
		correctAnswers = append(correctAnswers, (line[4:]))

	}

	fmt.Println("The quiz begins now! You may only exit when all questions are answered")
	for i :=0; i < len(totalQuestions); i++ {
		fmt.Println(totalQuestions[i]);
		fmt.Println(correctAnswers[i]);
		text, _ := reader.ReadString('\n' )
		fmt.Println(text);
		//didnt realize there was extra white space, took me a long time to figure out
		if(strings.TrimRight(text, "\n") == correctAnswers[i]) {
			numCorrectAnswers++
		}
	}

	fmt.Printf("Your total score is %d/%d",numCorrectAnswers,len(totalQuestions));


}
