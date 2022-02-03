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

func ReadFile(banner string) map[int][]string { // We want to return a map for simplicity
	fileInput := make(map[int][]string) // Alternative syntax: var fileInput map[int][]string
	file, err := os.Open(banner)        // file, err := os.Open("os.Args[1]") => e.g. go run main.go standard.txt; standard.txt would be 1st argument in this case
	if err != nil {
		fmt.Println("Invalid input. The named file does not exist.")
		os.Exit(1) // Exits program with "exit status 1" if file is not found; alternative: fmt.Println(err.Error()); "open standard.txt: no such file or directory"
	}
	defer file.Close()                // To close file is good practice; defer allows for all operations to be carried out before closing file
	scanner := bufio.NewScanner(file) // Scans the file and calls it "scanner"
	count := 31                       // Ignoring first line which is blank; in ascii manual (man ascii in terminal) integer 32 (decimal value) represents space (line 2 to 9 in standard.txt file); the first character is rune 32 (decimal value 32 - space); we could use fileInput := make(map[rune][]string) with slight changes to code elsewhere; there are 95 characters to print (dec value 32 - 126 in ascii manual)
	for scanner.Scan() {              // Scans the variable named scanner
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
	inputSlice := strings.Split(input, "\\n") // Takes input e.g. "Hello\nWorld" and converts to slice. In this case, we have two elements (words) in slice (splitting by new line)
	for _, input = range inputSlice {         // Range over inputSlice (Hello World) word by word
		for i := 0; i < 8; i++ { // Iterates through 8 lines as that's how many lines there are for each ascii character
			for _, inputChar := range input { // Range over input (H e l l o W o r l d) character by character
				output += (asciiMap[int(inputChar)][i]) // Casting inputChar (rune) to int as our map key is int; [i] print by index
			}
			output += "\n" // Alternative: fmt.Println(""); print new line is required so that characters are printed properly as opposed to on one line e.g. go run . "a" =>  / _` | | (_| |  \__,_|
		}
	}
	return output
}

func asciiGen(input string, banner string) string {
	output := ""

	output = outputAscii(ReadFile(banner), input) // ReadFile() is passed into asciiMap map[int][]string and os.Args[1] is passed into input string in parameter of outputAscii function

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
