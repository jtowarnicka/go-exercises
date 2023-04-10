package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
)

func printFile(scanner *bufio.Scanner, option bool, status string) {
	for scanner.Scan() {
		checkbox := scanner.Text()[0:3]

		switch status {
		case "todo":
			if checkbox != "[ ]" {
				continue
			}
		case "done":
			if checkbox != "[x]" {
				continue
			}
		case "nope":
			if checkbox != "[-]" {
				continue
			}
		case "ongoing":
			if checkbox != "[+]" {
				continue
			}
		// default:
			// do nothing
		}

		if option {
			description := scanner.Text()[4:]

			switch checkbox {
			case "[ ]":
				// not done
				fmt.Printf("\033[34;1m%s\033[0m", "[ ]")
				fmt.Printf("\033[37m %s\033[0m\n", description)
			case "[x]":
				// done
				fmt.Printf("\033[32m%s\033[0m", "[x]")
				fmt.Printf("\033[37m %s\033[0m\n", description)
			case "[-]":
				// nope
				fmt.Printf("\033[31;1m%s\033[0m", "[-]")
				fmt.Printf("\033[37m %s\033[0m\n", description)
			case "[+]":
				// ongoing
				fmt.Printf("\033[35m%s\033[0m", "[+]")
				fmt.Printf("\033[37m %s\033[0m\n", description)
			default:
				// comment
				fmt.Printf("\033[30;1m%s\033[0m\n", scanner.Text())
			}
		} else {
			fmt.Printf("\033[30;1m%s\033[0m\n", scanner.Text())
		}
	}

}

func loadFromFile(filename string, option bool, status string) {
	// fmt.Println("Loading from file: " + filename)
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	printFile(scanner, option, status)
}

type arrayFlags []string

func (i *arrayFlags) String() string {
    return fmt.Sprintf("%v", *i)
}

func (i *arrayFlags) Set(value string) error {
    *i = append(*i, value)
    return nil
}

func main() {
	var filename arrayFlags
    flag.Var(&filename, "filename", "files to load")
	option := flag.Bool("option", true, "disable color if true color is supported")
	status := flag.String("status", "", "only show tasks with this tag")
	flag.Parse()

	for run := false; !run; {
		fmt.Print("Enter command (list/exit): ")
		var command string
		fmt.Scanln(&command)

		switch command {
		case "list":
			for _, name := range filename {
				loadFromFile(name, *option, *status)
			}
		case "exit":
			run = true
		default:
			fmt.Println("Unknown command")
		}
	}

	// go run todo.go -filename=todo.txt -option=false
	// go run todo.go -filename=todo.txt -status=ongoing
	// go run todo.go -filename todo.txt -filename todo2.txt
}
