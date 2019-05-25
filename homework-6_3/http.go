package main

/*
 Syntax Go. Homework 6.3 - Standard library part 2
Echopage fun experiment, displays on the page the path entered
 Nick Nikulin, dated Мay 17, 2019
 Teacher, Sergey Iryupin
*/

import (
	"fmt"
	"html"
	"net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Это %q", html.EscapeString(r.URL.Path))
}

func main() {
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/", fs)
	http.HandleFunc("/pageecho", hello)

	http.ListenAndServe(":8080", nil)
}
