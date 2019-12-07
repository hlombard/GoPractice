package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
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

func getInput() string {
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	text = strings.TrimRight(text, "\r\n") // \r for windows
	return text
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
	ptr := csv.NewReader(file)
	i := 1
	for {
		record, err := ptr.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Question %d: %s : ", i, record[0])
		input := getInput()
		if checkAnswer(input, record[1]) == false {
			fmt.Println("False... Correct answer was", record[1])
			os.Exit(-2)
		}
		i++
	}
	fmt.Printf("\nCongratz all answers correct : %d/%d", i, i)
}
