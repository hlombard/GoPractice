package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"time"
)

func showUsage() {
	fmt.Println("Usage: go run CSVQuiz.go [file.csv]")
}

func checkArgs() bool {
	if len(os.Args) != 2 {
		showUsage()
		return false
	} else {
		if _, err := os.Stat(os.Args[1]); os.IsNotExist(err) {
			return false
		} else {
			return true
		}
	}
}

func checkAnswer(eval string, answer string) bool {
	if eval == answer {
		return true
	} else {
		return false
	}
}

func getInput(input chan string) {
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	text = strings.TrimRight(text, "\r\n")
	input <- text
}

func iswellFormated(str []string) bool {
	if len(str) == 2 && strings.ContainsAny(str[0], "+-*/") {
		return true
	} else {
		return false
	}
}

func main() {

	if ok := checkArgs(); ok == false {
		if len(os.Args) == 2 {
			fmt.Printf("Error: file \"%s\" doesn't exists\n", os.Args[1])
		}
		os.Exit(-1)
	}
	file, error := os.Open(os.Args[1])
	if error != nil {
		log.Fatalln("Couldn't open the csv file", error)
	}

	fmt.Println("You have 2 seconds to answer each questions")
	fmt.Println("Press ENTER when you're ready !")
	for{
		tmp := bufio.NewReader(os.Stdin)
		tmp.ReadString('\n')
		break
	}

	ptr := csv.NewReader(file)
	total := 0
	correct := 0
	for loop := true; loop == true;{
		record, err := ptr.Read()
		if err == io.EOF {
			break
		}
		if iswellFormated(record) == false {
			continue
		}
		fmt.Printf("Question %d: %s : ", total, record[0])
		total++
		input := make(chan string, 1)
		go getInput(input)

		select {
		case i:= <-input:
			if checkAnswer(i, record[1]) == false{
			fmt.Println("False... Correct answer was", record[1])
			} else{
				correct++
			}
		case <-time.After(2000 * time.Millisecond):
			fmt.Println("You need to be quicker...")
			loop = false
		}
	}
	fmt.Printf("\nTest finished : %d/%d Correct Answers !", correct, total)
}
