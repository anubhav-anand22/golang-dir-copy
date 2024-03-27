package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/mattn/go-tty"
)

func main() {
	ttya, err := tty.Open()
	if err != nil {
		log.Fatal(err)
	}
	defer ttya.Close()

	a := make([]string, 0)
	a = append(a, "kunal")
	a = append(a, "kumar")
	a = append(a, "anu")
	a = append(a, "bhav")
	a = append(a, "anand")

	dropdown(a, ttya)

}

func dropdownHelper(arr []string, selected int, title string) {
	fmt.Println(title)
	for index, element := range arr {
		if index == selected {
			fmt.Println(">>> " + element)
		} else {
			fmt.Println("    " + element)
		}
	}
	fmt.Println("Press 's' to to go up, 'd' to go down and 'e' to select")
}

func dropdown(list []string, ttya *tty.TTY) int {

	selected := 0
	for {
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
		dropdownHelper(list, selected, "ABC")
		r, err := ttya.ReadRune()
		if err != nil {
			log.Fatal(err)
		}
		key := string(r)
		if key == "e" {
			break
		} else if key == "s" {
			if selected <= 0 {
				selected = len(list) - 1
			} else {
				selected -= 1
			}
		} else if key == "d" {
			if selected >= (len(list) - 1) {
				selected = 0
			} else {
				selected += 1
			}
		}
	}

	return selected
}
