package main

/*
 Syntax Go. Homework 4.2 - Methods and interfaces
 Simple implementation of sorting for data structure
 Nick Nikulin, dated Мay 05, 2019
 Teacher, Sergey Iryupin
*/

import (
	"fmt"
	"sort"
)

type Book struct {
	Name    string
	Surname string
	Phone   int
}

type Numb []Book

func (numb Numb) Len() int {
	return len(numb)
}

func (numb Numb) Less(i, j int) bool {
	return numb[i].Surname < numb[j].Surname
}

func (numb Numb) Swap(i, j int) {
	numb[i], numb[j] = numb[j], numb[i]
}

func main() {
	contact := []Book{
		{"Николай", "Никулин", 89179001111},
		{"Егор", "Aрхиров", 89271317288},
		{"Евгений", "Косолапов", 89504330200},
		{"Ильдар", "Мугинов", 89504330200},
		{"Евгений", "Королев", 89504330200},
	}
	fmt.Println(contact)

	sort.Sort(Numb(contact))
	fmt.Println(contact)
}
