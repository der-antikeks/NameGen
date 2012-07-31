package main


import (
	"fmt"
	"strings"

	"github.com/der-antikeks/namegen"
)

func main() {
	// [a-zA-Z]+
	s := strings.Split("Jacob Mason William Jayden Noah Michael Ethan Alexander Aiden Daniel Anthony Matthew Elijah Joshua Liam Andrew James David Benjamin Logan Christopher Joseph Jackson Gabriel Ryan Samuel John Nathan Lucas Christian Jonathan Caleb Dylan Landon Isaac Gavin Brayden Tyler Luke Evan Carter Nicholas Isaiah Owen Jack Jordan Brandon Wyatt Julian Aaron Jeremiah Angel Cameron Connor Hunter Adrian Henry Eli Justin Austin Robert Charles Thomas Zachary Jose Levi Kevin Sebastian Chase Ayden Jason Ian Blake Colton Bentley Dominic Xavier Oliver Parker Josiah Adam Cooper Brody Nathaniel Carson Jaxon Tristan Luis Juan Hayden Carlos Jesus Nolan Cole Alex Max Grayson Bryson Diego Jaden", " ")

	n := namegen.NewNameGen(s)

	for _, name := range n.GenerateMultiple(10) {
		fmt.Println(name)
	}
}
