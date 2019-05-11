package main

/*
 Syntax Go. Homework 5.2 - Standard library
 Reading from file
 Nick Nikulin, dated Мay 11, 2019
 Teacher, Sergey Iryupin
*/

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("main.go")
	if err != nil {
		return
	}
	defer file.Close()

	// getting size of file
	stat, err := file.Stat()
	if err != nil {
		return
	}

	// reading file
	bs := make([]byte, stat.Size())
	_, err = file.Read(bs)
	if err != nil {
		return
	}
	// инициализируем переменную show, передадим ей значения переменной bs, строкой
	show := string(bs)
	// Расспечатаем show
	fmt.Println(show)
}
