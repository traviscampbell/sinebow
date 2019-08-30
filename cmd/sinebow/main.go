package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"

	"github.com/traviscampbell/sinebow"
)

func main() {
	bg := flag.Bool("bg", false, "Invert the sexyness")
	flag.Parse()

	scanner := bufio.NewScanner(os.Stdin)
	
	for scanner.Scan() {
		if *bg {
			fmt.Println(sinebow.RainbroBG(scanner.Text()))
		} else {
			fmt.Println(sinebow.Rainbro(scanner.Text()))
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(sinebow.Rainbro(err.Error()))
		os.Exit(1)
	}
}