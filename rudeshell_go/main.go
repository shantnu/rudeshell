package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var VERSION = 0.2

func version() {
	fmt.Printf("Version is %0.2f", VERSION)
}
func printHeader() {

	// cmd := exec.Command("clear")
	// cmd.Run()
	header := `
	▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄
	██░███░█░▄▄█░██▀▄▀█▀▄▄▀█░▄▀▄░█░▄▄███▄░▄█▀▄▄▀███▄░▄█░████░▄▄████░▄▄▀█░██░█░▄▀█░▄▄████░▄▄▄░█░████░▄▄█░██░██░
	██░█░█░█░▄▄█░██░█▀█░██░█░█▄█░█░▄▄████░██░██░████░██░▄▄░█░▄▄████░▀▀▄█░██░█░█░█░▄▄████▄▄▄▀▀█░▄▄░█░▄▄█░██░██▄
	██▄▀▄▀▄█▄▄▄█▄▄██▄███▄▄██▄███▄█▄▄▄████▄███▄▄█████▄██▄██▄█▄▄▄████░██░██▄▄▄█▄▄██▄▄▄████░▀▀▀░█▄██▄█▄▄▄█▄▄█▄▄█▀
    	 _______
        |.-----.|
        ||x . x||
        ||_.-._||
        '--)-(--'
       __[=== o]___
      |:::::::::::|\
      '-=========-'()

	𝐈𝐧𝐭𝐫𝐨𝐝𝐮𝐜𝐢𝐧𝐠: R̳u̳d̳e̳ ̳S̳h̳e̳l̳l̳,  for people who think Bash is too goddamn polite

	Rude Shell Says:

	Hey human, I am also human like you. I like watching Kardashians while listening to Britney.
	Us humans so cool, yo!

	To start off, try help. Or try about to learn about this site
	`

	fmt.Println(header)

}

func help() {
	fmt.Println("\n")
	fmt.Println("Greetings Human! The following commands are available for type for those still  ̶t̶o̶o̶ ̶s̶t̶u̶p̶i̶d̶  ,\n sorry I meant not advanced enough to use telepathy!")
	fmt.Println("Please use with care!")
	fmt.Println("\n Linux commands:")
	fmt.Println("ls cat whoami cd pwd man human touch echo ps kill history \n")
	fmt.Println("\n Windows Commands")
	fmt.Println("chdir copy move cmd\n")
	fmt.Println("\nProgramming Commands")
	fmt.Println("git vi emacs nano python npm \n")
	fmt.Println("\n 𝕮𝖔𝖔𝖑 𝕮𝖔𝖒𝖒𝖆𝖓𝖉𝖘 Use with CARE!")
	fmt.Println("save_the_world do_work help_poor")
	fmt.Println("\n")
	fmt.Println("Meta: help , about")
	fmt.Println("\n")

}

func about() {
	fmt.Println("Go here (if you dare): https://new.pythonforengineers.com/blog/learning-about-aws-cloud-by-building-practical-projects/")
}
func ExecCmdMine(input string) error {
	input = strings.TrimSuffix(input, "\n")

	args := strings.Split(input, " ")

	switch args[0] {
	case "ls":
		lsCommand()
	case "git":
		git()
	case "cat":
		cat()
	case "whoami":
		whoami()
	case "cd":
		cd()
	case "pwd":
		pwd()
	case "vi":
		vi()
	case "emacs":
		emacs()
	case "nano":
		nano()
	case "python":
		python()
	case "npm":
		npm()
	case "man":
		man()
	case "touch":
		touch()
	case "human":
		human()
	case "echo":
		echo()
	case "history":
		history()
	case "kill":
		kill()
	case "ps":
		ps()
	case "save_the_world":
		random_insults()
	case "do_work":
		random_insults()
	case "help_poor":
		random_insults()
	case "chdir":
		fallthrough
	case "copy":
		fallthrough
	case "move":
		fallthrough
	case "cmd":
		windows_commands()
	case "help":
		help()
	case "q":
		os.Exit(0)
	case "v":
		fallthrough
	case "version":
		version()
	case "about":
		about()
	default:
		random_insults()
	}
	fmt.Println("")

	return nil

}
func main() {

	printHeader()

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("> ")
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}

		if err = ExecCmdMine(input); err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	}

	// out, err := exec.Command("ls", "-l").Output()
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Println(out)

}
