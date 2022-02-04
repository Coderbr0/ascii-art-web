package main

import (
	"bufio"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"strings"
)

func ReadFile(banner string) map[int][]string {
	fileInput := make(map[int][]string)
	file, err := os.Open(banner)
	if err != nil {
		fmt.Println("Invalid input. The named file does not exist.")
		os.Exit(1)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	count := 31
	for scanner.Scan() {
		if scanner.Text() == "" {
			count++
		} else {
			fileInput[count] = append(fileInput[count], scanner.Text())
		}
	}
	return fileInput
}

func outputAscii(asciiMap map[int][]string, input string) string {
	var output string
	inputSlice := strings.Split(input, "\\n")
	for _, input = range inputSlice {
		for i := 0; i < 8; i++ {
			for _, inputChar := range input {
				output += (asciiMap[int(inputChar)][i])
			}
			output += "\n"
		}
	}
	return output
}

func asciiGen(input string, banner string) string {
	output := ""

	output = outputAscii(ReadFile(banner), input)

	return output
}

func ascii(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		tmpl, _ := template.ParseFiles("templates/index.html")
		tmpl.Execute(w, nil)
	} else {
		r.ParseForm()
		asciiInput := r.Form["asciiInput"]
		asciiBanner := r.Form["banner"]
		asciiRaw := asciiGen(asciiInput[0], asciiBanner[0])
		tmpl, _ := template.ParseFiles("templates/index.html")
		tmpl.Execute(w, asciiRaw)
	}
}

func main() {
	http.HandleFunc("/", ascii)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalf("500 internal server error : %s", err.Error())
	}
}
/*When writing new line in text box index is out of range"*/