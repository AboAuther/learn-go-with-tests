package main

import (
	"fmt"
	"io"
	"net/http"
)

func Greet(writer io.Writer, name string) {
	fmt.Fprintf(writer, "hello,%s", name)
}
func MyGreetHandler(w http.ResponseWriter, t *http.Request) {
	Greet(w, "world")
}
func main() {
	err := http.ListenAndServe(":5000", http.HandlerFunc(MyGreetHandler))
	if err != nil {
		fmt.Println(err)
	}

}
