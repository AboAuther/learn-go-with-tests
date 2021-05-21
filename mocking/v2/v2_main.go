package main

import (
	"fmt"
	"io"
	"os"
)

const finalWord = "Go!!"
const countdownstart = 3

func CountDown(out io.Writer) {
	for i := countdownstart; i > 0; i-- {
		fmt.Fprintln(out, i)
	}
	fmt.Fprint(out, finalWord)
}
func main() {
	CountDown(os.Stdout)
}
