# Description
This project is a web server that takes a sample of text and returns the ascii art form of that text.
It can handle the carriage return \r as well as the newline character \n and can be set to the following ascii art forms: standard, shadow or thinkertoy.

# Usage
- Step 1: Go to directory where program is stored in the console. A code editor such as Visual Studio Code can be used to do so.
- Step 2: Type `go run main.go` in the terminal.
- Step 3: Go to `127.0.0.1:8080` or `http://localhost:8080/` in your browser.
- Step 4: Select a banner and enter the text you want converted. You can also press enter to get multiline output.
- Step 5: Click submit and check out the result.

# Errors Handled
- 404: if file is not found.
- 400: bad request if banner is not selected or if there is no text to convert.
- 500: internal server error for unhandled errors.

# Authors
* Karolis Milius
* Coderbro
* Maddina