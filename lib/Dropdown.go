package lib

import (
	"fmt"
	"log"
	"os"
	"os/exec"

	// "os"
	// "os/exec"

	"github.com/mattn/go-tty"
)

func dropdownHelper(arr []string, selected int, title string, botTxt string) {
	fmt.Println(title)
	fmt.Println("")
	for index, element := range arr {
		if index == selected {
			fmt.Println(">>> " + element)
		} else {
			fmt.Println("    " + element)
		}
	}
	fmt.Println("")
	fmt.Println("Press 'a' to to go up, 'b' to go down" + botTxt)
}

type CallbackFunc func(int, string) int

func Dropdown(list []string, title string, ttya *tty.TTY, botTxt string, cb CallbackFunc) int {
	if len(botTxt) == 0 {
		botTxt = " and press 'SPACE' to select"
	}

	selected := 0
	for {
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
		dropdownHelper(list, selected, title, botTxt)
		r, err := ttya.ReadRune()
		if err != nil {
			log.Fatal(err)
		}
		key := string(r)
		if key == "e" {
			os.Exit(1)
			break
		} else if key == " " {
			break
		} else if key == "a" || key == "A" {
			if selected <= 0 {
				selected = len(list) - 1
			} else {
				selected -= 1
			}
		} else if key == "b" || key == "B" {
			if selected >= (len(list) - 1) {
				selected = 0
			} else {
				selected += 1
			}
		} else {
			cbR := cb(selected, key)
			if cbR == 1 {
				break
			}
		}
	}

	return selected
}
