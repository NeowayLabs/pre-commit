package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"

	"github.com/AlecAivazis/survey/v2"
	"github.com/spf13/cobra"
)

const (
	fixType      = "[fix]: A code/bug fix;"
	featType     = "[feat]: A new feature;"
	docsType     = "[docs]: Documentation only changes;"
	testType     = "[test]: Adding missing tests or correcting existing tests;"
	refactorType = "[refactor]: A code change that neither fixes a bug nor adds a feature;"
	styleType    = "[style]: Changes that do not affect the meaning of the code (white-space, formatting, missing semi-colons, etc);"
	perfType     = "[perf]: A code change that improves performance;"
	buildType    = "[build]: Changes that affect the build system or external dependencies (example scopes: gulp, broccoli, npm);"
	ciType       = "[ci]: Changes to our CI configuration files and scripts (example scopes: Travis, Circle, BrowserStack, SauceLabs);"
	breakingType = "[breaking change]: Change that will require other changes in dependant applications;"

	colorYellow = "\033[33m"
	colorGreen  = "\033[32m"
	colorReset  = "\033[0m"
	textBold    = "\033[1m"
	colorRed    = "\033[31m"
)

var (
	mapCommitTypes = map[string]string{
		fixType:      "fix",
		featType:     "feat",
		docsType:     "docs",
		testType:     "test",
		refactorType: "refactor",
		styleType:    "style",
		perfType:     "perf",
		buildType:    "build",
		ciType:       "ci",
		breakingType: "breaking change",
	}
)

func main() {
	var devToolCmd = &cobra.Command{
		Use:   ".",
		Short: "A commit message CLI helper",
		Run:   runCli,
	}

	var rootCmd = &cobra.Command{Use: "root"}
	rootCmd.AddCommand(devToolCmd)
	rootCmd.Execute()
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

func getCommitType(selectedValue string) string {
	return mapCommitTypes[selectedValue]
}

func runCli(cmd *cobra.Command, args []string) {
	options := []string{
		fixType,
		featType,
		docsType,
		testType,
		refactorType,
		styleType,
		perfType,
		buildType,
		ciType,
		breakingType,
	}

	var commitType string
	prompt := &survey.Select{
		Message: "Select the commit type:",
		Options: options,
	}
	survey.AskOne(prompt, &commitType)

	commitMessage := stringPrompt("What is the commit message?")

	commitType = getCommitType(commitType)
	fullMessage := fmt.Sprintf("type: [%s], message: %s", commitType, commitMessage)
	fmt.Printf("Your commit message is: %s.\n", fullMessage)

	yesNoOptions := []string{"yes", "no"}
	var commitNow string
	prompt = &survey.Select{
		Message: "Do you want to commit now?",
		Options: yesNoOptions,
	}
	survey.AskOne(prompt, &commitNow)

	if commitNow == "yes" {
		commit(fullMessage)
	} else {
		fmt.Println("Bye!")
	}
}

func stringPrompt(label string) string {
	var s string
	r := bufio.NewReader(os.Stdin)
	for {
		fmt.Fprint(os.Stderr, label+" ")
		s, _ = r.ReadString('\n')
		if s != "" {
			break
		}
	}
	return strings.TrimSpace(s)
}

func commit(message string) {
	fmt.Println("commiting...")
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
