package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"flag"
	"time"
)

//main package for executing the quiz program
var correctAnswers []string
var totalQuestions []string
var usersAnswer []string
var numCorrectAnswers = 0
var answer string;

func init() {

}
func main() {
	csvFile := flag.String("csv", "problems.csv", "csv files that contain quiz question and answers'")
	timeLimit := flag.Int("limit", 30, "Time limit you have for answering all the questions in the quiz (in seconds)")
	flag.Parse()



	data, err := ioutil.ReadFile(*csvFile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v", err)
	}
	for _, line := range strings.Split(string(data), "\n") {
		totalQuestions = append(totalQuestions,(line[0:3]))
		correctAnswers = append(correctAnswers, (line[4:]))

	}


	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)

		fmt.Println("Press enter to begin quiz")
		buf := bufio.NewReader(os.Stdin)
		 buf.ReadBytes('\n')

		 //In order to break from both the case and for statements in a single call
		label:
		for i :=0; i < len(totalQuestions); i++ {
			askQuestion(i)

			answerChannel := make(chan string)

			go func() {
				var answer string
				reader := bufio.NewReader(os.Stdin)
				answer, _ = reader.ReadString('\n' )
				//send to this channel once answer comes through
				answerChannel<-answer
			} ()
			select {
			case <-timer.C:
				break label
			case answer := <-answerChannel:
				if(strings.TrimRight(answer, "\n") == correctAnswers[i]) {
					numCorrectAnswers++
				}
			}

		}
	fmt.Printf("Your total score is %d/%d",numCorrectAnswers,len(totalQuestions));









}

func checkAnswer(i int) {
	//didnt realize there was extra white space, took me a long time to figure out
	if(strings.TrimRight(answer, "\n") == correctAnswers[i]) {
		numCorrectAnswers++
	}
}

func askQuestion(i int) {
	fmt.Println(totalQuestions[i]);
}

func askAnswer(i int) {
	reader := bufio.NewReader(os.Stdin)
	answer, _ = reader.ReadString('\n' )
	checkAnswer(i)
}
