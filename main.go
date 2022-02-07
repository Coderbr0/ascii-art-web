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
	inputSlice := strings.Split(input, "\r\n")
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
		tmpl, err := template.ParseFiles("templates/index.html")
		if err != nil {
			fmt.Fprintf(w, "404 file not found: %s", err.Error())
		} else {
			tmpl.Execute(w, nil)
		}
	} else {
		r.ParseForm()
		asciiInput := r.Form["asciiInput"]
		asciiBanner := r.Form["banner"]
		asciiBanner = append(asciiBanner, "") // Append as with if statement we would get index out of range error
		asciiInput = append(asciiInput, "")   // Append as with if statement we would get index out of range error
		if asciiBanner[0] == "" || asciiInput[0] == "" {
			tmpl, err := template.ParseFiles("templates/index.html")
			if err != nil {
				fmt.Fprintf(w, "404 file not found: %s", err.Error())
			} else {
				tmpl.Execute(w, fmt.Sprint("400 bad request"))
			}
		} else {
			asciiRaw := asciiGen(asciiInput[0], asciiBanner[0])
			if asciiRaw != "" {
				tmpl, err := template.ParseFiles("templates/index.html")
				if err != nil {
					fmt.Fprintf(w, "404 file not found: %s", err.Error())
				} else {
					tmpl.Execute(w, asciiRaw)
					fmt.Println("200 OK") // tmpl.Execute(w, asciiRaw) works, so we have fmt.Println("200 OK") here
				}
			} else {
				tmpl, err := template.ParseFiles("templates/index.html")
				if err != nil {
					fmt.Fprintf(w, "404 file not found: %s", err.Error())
				} else {
					tmpl.Execute(w, fmt.Sprintf("500 internal server error : %s", err.Error()))
				}
			}
		}
	}
}

func main() {
	http.HandleFunc("/", ascii)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalf("500 internal server error : %s", err.Error())
	}
}
