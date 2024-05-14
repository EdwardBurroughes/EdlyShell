package main

import (
	// "fmt"

	"github.com/edwardBurroughes/edlyshell/internal/shell"
	// "bufio"
	// "bytes"
	// "fmt"
	// "os"
	// "os/exec"
	// "os/signal"
	// "strings"
	// "syscall"
)

// // TODO - handle signals i.e. up arrow and down arrow + handle control C
// // TODO - don't exit on control C but kill the command
// // TODO - create history command

// type Shell struct {
// 	workingDir string
// 	signalChan chan os.Signal
// }

// func getRelWorkingDir() string  {
// 	// fail for anything other then Mac
// 	homeDir := os.Getenv("HOME")
// 	wd, err := os.Getwd()
// 	if err != nil {
// 		panic(err)
// 	}
// 	return strings.Replace(wd, homeDir, "~", -1)
// }

// func buildShellTitle() string {
// 	relDir := getRelWorkingDir()
// 	titleStr := fmt.Sprintf("%s 🚀", relDir)
// 	colorString := Cyan(titleStr)
// 	return Bold(colorString)
// }

// func handleCDCommand(command string) (string,error) {
// 	cmd := cleanCmdBase(command)
// 	if len(cmd.args) > 1 {
// 		return "", fmt.Errorf("more then one path provides %v", cmd.args)
// 	}
// 	path := replaceTildaHome(cmd.args[0])
// 	err := os.Chdir(path)
// 	if err != nil {
// 		return "", err
// 	}
// 	return "", nil
// }

// func replaceTildaHome(path string) string {
// 	if strings.HasPrefix(path, "~") {
// 		// don't need to fail here if not Mac OS
// 		homeDir := os.Getenv("HOME")
// 		return strings.Replace(path, "~", homeDir, -1)
// 	}
// 	return path
// }

// func cleanCmdBase(command string ) Command {
// 	splitCmd := strings.Split(command, " ")
// 	return Command{splitCmd[0], splitCmd[1:]}
// }

// type Command struct {
// 	name string
// 	args []string
// }

// func cleanCmd(command string) []*exec.Cmd {
// 	splitPipeCmd := strings.Split(command, "|")

// 	cmds := make([]*exec.Cmd, 0)
// 	for _, cmd := range splitPipeCmd {
// 		newCmd := cleanCmdBase(strings.TrimSpace(cmd))
// 		cmds = append(cmds, exec.Command(newCmd.name, newCmd.args...))
// 	}
// 	return cmds
// }

// func executeCommands(cmdStack []*exec.Cmd, pipeStack []*os.File) error {
// 	// can I improve this hmm !!
// 	for _, cmd := range cmdStack {
// 		if err := cmd.Start(); err != nil {
// 			return err
// 		}
// 	}

// 	for _, pipe := range pipeStack {
// 		pipe.Close()
// 	}

// 	for _, cmd := range cmdStack {
// 		if err := cmd.Wait(); err != nil {
// 			return err
// 		}
// 	}
// 	return nil
// }

// func handleUnixCommand(command string) (string, error) {
// 	cmdStack := cleanCmd(command)

// 	var stderr bytes.Buffer
// 	var output bytes.Buffer
// 	pipeStack := make([]*os.File, len(cmdStack)-1)

// 	for i := 0; i < len(cmdStack) -1; i++ {
// 		pipeIn, pipeOut, _ := os.Pipe()
// 		cmdStack[i].Stdout = pipeOut
// 		cmdStack[i].Stderr = &stderr
// 		cmdStack[i + 1].Stdin = pipeIn
// 		pipeStack[i] = pipeOut
// 	}
// 	cmdStack[len(cmdStack) - 1].Stdout = &output
// 	if err := executeCommands(cmdStack, pipeStack); err != nil {
// 		return "", err
// 	}
// 	trimmedOutput := strings.TrimSpace(output.String())
// 	return trimmedOutput, nil
// }

// func handleCommand(command string) (string, error){
// 	// handle signals
// 	trimmedCmd := strings.TrimSpace(command)
// 	if strings.HasPrefix(trimmedCmd, "exit") {
// 		fmt.Println("Goodbye 👋")
// 		os.Exit(0)
// 	}
// 	if strings.HasPrefix(trimmedCmd, "cd") {
// 		return handleCDCommand(trimmedCmd)
// 	} else {
// 		return handleUnixCommand(trimmedCmd)
// 	}
// }

// func main(){
// 	// bug when hitting the command should start a new line
// 	fmt.Println("\n")
// 	sigs := make(chan os.Signal, 1)
// 	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
// 	interrupt := make(chan bool, 1)

//     go func() {
//         for {
//             sig := <-sigs
//             if sig == syscall.SIGINT {
//                 // Signal received, send to interrupt channel
//                 interrupt <- true
//             }
//         }
//     }()

// 	for {
// 		fmt.Println(Cyan(buildShellTitle()))
// 		fmt.Print(Red("> "))
// 		reader := bufio.NewReader(os.Stdin)
// 		input, _ := reader.ReadString('\n')
// 		select {
// 		case <-interrupt:
//             // If interrupt flag is set, continue to the next iteration
//             continue
//         default:
//             // If no interrupt, handle the command as usual
//             output, err := handleCommand(input)
//             if err != nil {
//                 fmt.Println(err)
//             } else {
//                 fmt.Println(output)
//             }
//         }
// 	}
// }
func main() {
	sh := shell.ShellInit()
	for {
		sh.ExecuteCommand()
	}
}

