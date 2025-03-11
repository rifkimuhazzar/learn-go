package main

import (
	"bufio"
	"io"
	"os"
)

func createNewFile(name string, message string) error {
	file, err := os.OpenFile(name, os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		return err
	}
	defer file.Close()
	file.WriteString(message)
	return nil
}

func addToFile(name string, message string) error {
	file, err := os.OpenFile(name, os.O_RDWR|os.O_APPEND, 0666)
	if err != nil {
		return err
	}
	defer file.Close()
	file.WriteString(message)
	return nil
}


func readFile(name string) (string, error) {
	file, err := os.OpenFile(name, os.O_RDONLY, 0666)
	if err != nil {
		return "", err
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	var message string
	for {
		line, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}
		message += string(line) + "\n"
	}
	return message, nil
}

func main() {
	// createNewFile("18-sample.log", "this is sample log")

	// result, err := readFile("18-sample.log")
	// if err != nil {
	// 	fmt.Println("if")
	// 	fmt.Println(err)
	// } else {
	// 	fmt.Println("else")
	// 	fmt.Println(result)
	// }

	addToFile("18-sample.log", "\nthis is add message")
	addToFile("18-sample.log", "\nthis is add message")
	addToFile("18-sample.log", "\nthis is add message")
}