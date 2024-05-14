// create a struct for the simplest case i.e. print and execuet command
package shell

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"unicode"
)

type Shell struct {
	currentDir string
	Input string
	// workingDir string
}

func isValidChar(b byte) bool {
	if b == '\n' {
		return true
	}
	if b == '[' {
		return false
	}
	r := rune(b)
	return unicode.IsSpace(r) || unicode.IsDigit(r) || unicode.IsLetter(r) || unicode.IsPunct(r) || unicode.IsSymbol(r)
}


func (s *Shell) prompt() {
	titleStr := fmt.Sprintf("%s 🚀", s.currentDir)
	colorString := Cyan(titleStr)
	boldTitle := Bold(colorString)
	fmt.Println(boldTitle)
	fmt.Print(Red("> "))
}


func (s *Shell) ReadInput() (string) {
	s.prompt()
	scanner := bufio.NewReader(os.Stdin)
	s.Input = "" 
	for {
		b, err := scanner.ReadByte()
		if err != nil {
			break
		}

		if isValidChar(b) {
			s.Input += string(b)
		}
		if b == '\n' {
			break
		}
	}

	trimmedInput := strings.TrimSpace(s.Input)
	return trimmedInput
}

func handleUnixCommand(name string, args []string){
	cmd := exec.Command(name, args...)
	cmd.Dir = s.currentDir

	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		fmt.Println(err)
	}
}

func commandHandler(name string, args []string){
	switch name {
	case "exit":
		fmt.Println("Goodbye 👋")
		os.Exit(0)
	default:
		
	}
}

func (s *Shell) ExecuteCommand() {
	input := s.ReadInput()
	fields := strings.Fields(input)
	name := fields[0]
	args := fields[1:]

	cmd := exec.Command(name, args...)

	// handle different commands
	

	
}


func ShellInit() (*Shell) {
	pwd, _ := os.Getwd()
	return &Shell{currentDir: pwd}
}





