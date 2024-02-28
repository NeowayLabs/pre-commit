package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"log"
)

const (
	colorYellow = "\033[33m"
	colorGreen  = "\033[32m"
	colorReset  = "\033[0m"
	textBold    = "\033[1m"
	colorRed    = "\033[31m"

	TypeFix            = "fix"
	TypeDocumentation  = "docs"
	TypeFeature        = "feat"
	TypeRefactory      = "refactor"
	TypeStyle          = "style"
	TypeTest           = "test"
	TypePerformance    = "perf"
	TypeBreakingChange = "breaking change"
	TypeBuild          = "build"
	TypeCi             = "ci"
	TypeExit           = "exit"
)

var (
	deployTypeMap = map[string]string{
		"1":  TypeFix,
		"2":  TypeDocumentation,
		"3":  TypeFeature,
		"4":  TypeRefactory,
		"5":  TypeStyle,
		"6":  TypeTest,
		"7":  TypePerformance,
		"8":  TypeBreakingChange,
		"9":  TypeBuild,
		"10": TypeCi,
		"0":  TypeExit}
)

func main() {
	printRed("Welcome to Data Platform pre commit!")

	commitType, err := getCommitType()
	if err != nil {
		log.Fatal(err)
	}

	commitMessage := getCommitMessage(0)

	commitFullMessage := fmt.Sprintf("type: [%s], message: %s", commitType, commitMessage)

	printGreenYellow("[Your commit message is] ", commitFullMessage)

	commit(commitFullMessage)
}

func printGreenYellow(tag, text string) {
	fmt.Println(textBold + colorGreen + tag + colorReset + colorYellow + text + colorReset)
}

func printRed(text string) {
	fmt.Println(colorRed + text + colorReset)
}

func printGreen(text string) {
	fmt.Println(colorRed + text + colorReset)
}

func getCommitType() (string, error) {
	printGreenYellow("[COMMIT TYPE]", ` Which one is your commit type? *insert the option number*

		1.              [fix]: A code/bug fix;
		2.             [docs]: Documentation only changes;
		3.             [feat]: A new feature;
		4.         [refactor]: A code change that neither fixes a bug nor adds a feature;
		5.            [style]: Changes that do not affect the meaning of the code (white-space, formatting, missing semi-colons, etc);
		6.             [test]: Adding missing tests or correcting existing tests;
		7.             [perf]: A code change that improves performance;
		8.  [breaking change]: Change that will require other changes in dependant applications;
		9.            [build]: Changes that affect the build system or external dependencies (example scopes: gulp, broccoli, npm);
		10.              [ci]: Changes to our CI configuration files and scripts (example scopes: Travis, Circle, BrowserStack, SauceLabs);
		0.                     Exit.`)

	deployType := getUserInput()

	result := deployTypeMap[deployType]

	if result == "" {
		printRed("Invalid option!")
		os.Exit(1)
	}

	if result == TypeExit {
		return "", fmt.Errorf("Bye!")
	}

	printGreenYellow("\nSelected commit type: ", result+"\n")

	return result, nil
}

func getCommitMessage(count int) string {
	printGreenYellow("\n[Message] ", "What is the commit message?")
	commitMessage := getUserInput()
	if commitMessage == "" {
		if count == 2 {
			printRed("Commit message cannot be empty!")
			os.Exit(1)
		}
		count++
		getCommitMessage(count)
	}
	return commitMessage
}

func getUserInput() string {
	reader := bufio.NewReader(os.Stdin)
	return readAndPurgeBreakLine(reader)
}

func readAndPurgeBreakLine(reader *bufio.Reader) string {
	result, _ := reader.ReadString('\n')
	return strings.TrimSuffix(result, "\n")
}

func commit(message string) {
	path, err := exec.Command("pwd").Output()
	if err != nil {
		printRed(fmt.Sprintf("Oops! Error to get current path: %s", err.Error()))
		log.Fatal(err)
	}

	location := string(path[:])

	command := fmt.Sprintf("git commit -m \"%s\"", message)

	printGreenYellow("[Running] ", fmt.Sprintf("%s from %s", command, location))

	out, errRun := exec.Command("bash", "-c", command).Output()

	if errRun != nil {
		printRed(fmt.Sprintf("Oops! Error to commit: %s", errRun.Error()))
		if strings.Contains(errRun.Error(), "exit status 128") {
			printRed("It looks like you are not logged in git, try: [git config --global user.email \"user@domain.com\"]")
		}
		log.Fatal(errRun)
	}

	output := string(out[:])
	printGreen(output)
}
