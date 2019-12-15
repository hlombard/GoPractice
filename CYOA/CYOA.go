package main

import(
	"fmt"
	"encoding/json"
	"os"
	"io/ioutil"
	"log"
	"bufio"
	"strconv"
	"strings"
)

type OptionStruct struct{
	Txt string `json:"text"`
	Chapter string `json:"arc"`
}

type ChapterStruct struct{
	Title string `json:"title"`
	Story []string `json:"story"`
	Option[] OptionStruct `json:"options"`
}

type Story map[string]ChapterStruct

func printStory(ptr ChapterStruct){

	fmt.Printf("\n[%s]\n\n", ptr.Title)
	for _, j := range ptr.Story{
		fmt.Printf("%s\n", j)
	}
}

func parseFile() Story{

	var FullStory Story

	content, err := ioutil.ReadFile(os.Args[1])
	if err != nil{
		fmt.Println("Error reading the file")
		log.Fatal(err)
	}

	err = json.Unmarshal(content, &FullStory)
	if err != nil {
		fmt.Println("error:", err)
		os.Exit(-2)
	}

	return FullStory
}

func readChapter(ptr ChapterStruct) string{

	printStory(ptr)

	if (len(ptr.Option) <= 0){ //END OF THE STORY
		return "" }

	fmt.Printf("\nPress Enter to check options... ")
	reader := bufio.NewReader(os.Stdin)
	reader.ReadString('\n')
	fmt.Printf("\n")

	nb := 0
	for j, _ := range ptr.Option{
		fmt.Printf("%d : %s\n", nb, ptr.Option[j].Txt)
		nb++
	}

	fmt.Printf("\nEnter option Number and continue your adventure ! ... ")

	for{
		reader = bufio.NewReader(os.Stdin)
		str, _ := reader.ReadString('\n')
		str = strings.TrimRight(str, "\r\n")
		if value, err := strconv.Atoi(str); err == nil && value >= 0 && value < len(ptr.Option){
			return ptr.Option[value].Chapter
		} else{
			fmt.Printf("Nope... Try again\n")
			continue
		}
	}

	return ""
}

func startReading(FullStory Story){
	if _, set := FullStory["intro"]; set == false{
		fmt.Printf("Your story needs a start, set an 'intro' chapter :)\n")
		return;
	} else {
		nextChapter := readChapter(FullStory["intro"])
		for
		{
			nextChapter = readChapter(FullStory[nextChapter])
			if _, val := FullStory[nextChapter]; !val{
				break
			}
		}
	}
	fmt.Printf("\n\n\nEnd of the story.\n")
}

func main(){

	if (len(os.Args) != 2){
		fmt.Printf("Usage: go run CYAO.go [story.json]\n")
		return }

	if _, err := os.Stat(os.Args[1]); err != nil{
			fmt.Printf("Cannot find file \"%s\"\n", os.Args[1])
		}

	FullStory := parseFile()
	startReading(FullStory)

}
