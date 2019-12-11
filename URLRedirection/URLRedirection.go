package main

import(
	"net/http"
	"os"
	"fmt"
	"gopkg.in/yaml.v2"
	"log"
	"io/ioutil"
)

type AppYAML struct {
	Path string `yaml:"path"`
	Url  string `yaml:"url"`
}

func check(e error) {
	if e != nil {
	  panic(e)
	}
  }

func setArrayStruct() []AppYAML{
	if _, err := os.Stat("settings.yaml"); os.IsNotExist(err) {
		fmt.Println("You need to have \"settings.yaml\" file in your current directory, to setup the redirection config")
		os.Exit(-1) }

	file, err := ioutil.ReadFile("./settings.yaml")
	check(err)

	var config []AppYAML
	if err = yaml.Unmarshal(file, &config); err != nil {
        log.Fatalf("Cannot Unmarshal Data: %v", err)
	}
	return config
}

func urlRedirection(rw http.ResponseWriter, req *http.Request){
	req.ParseForm()
	if (req.URL.Path == "/favicon.ico"){
		return }
	config := setArrayStruct()
	for i, _:= range config {
		if (config[i].Path == req.URL.Path) {
			fmt.Println("Redirection made to :", config[i].Url)
			http.Redirect(rw, req, config[i].Url, 301)
		}
	}
}

func main(){
	fmt.Println("Listening to localhost:9090 ...")
	go http.HandleFunc("/", urlRedirection)
	err := http.ListenAndServe(":9090", nil)
    if err != nil {
        log.Fatal("ListenAndServe: ", err)
	}
}
