package main

/*
 Syntax Go. Homework 5.3 - Standard library
A simple program that extracts data from people.csv and outputs it to the console.
 Nick Nikulin, dated Мay 11, 2019
 Teacher, Sergey Iryupin
*/

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
)

type Person struct {
	Firstname string
	Lastname  string
	Address   string
}

func main() {
	csvFile, _ := os.Open("people.csv")
	reader := csv.NewReader(bufio.NewReader(csvFile))
	var people []Person
	for {
		line, error := reader.Read()
		if error == io.EOF {
			break
		} else if error != nil {
			log.Fatal(error)
		}
		people = append(people, Person{
			Firstname: line[0],
			Lastname:  line[1],
			Address:   line[2]})
	}

	fmt.Println(people)
}
