package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("\nenter your first name : ")

	var Fname string
	fmt.Scanln(&Fname)

	fmt.Println("enter your last name : ")

	var Lname string
	fmt.Scanln(&Lname)

	fmt.Println("your fullname is : \n" + Fname + " " + Lname + "\n")

	IsAdmin := false
	if strings.ToLower(Fname) == "pouya" && strings.ToLower(Lname) == "tirdad" {
		fmt.Println("Welcome \n ")
		IsAdmin = true
	}

	if IsAdmin == false && strings.Contains(Fname, "pouya") && strings.Contains(Lname, "tirdad") {
		fmt.Println("ohh , sorry i could not recognize you \nWelcome \n ")
		IsAdmin = true
	}
}
