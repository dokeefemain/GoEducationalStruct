package main

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
	"strconv"
	"strings"
)

type Education struct {
	Name        string
	School      string
	Grade       int
	Gpa         float32
	FinalGrades []string
	CreditHours int `18`
	Password1   string
}

func main() {
	devin := Education{
		Name:        "Devin",
		School:      "Georgia State University",
		Grade:       14,
		Gpa:         3.37,
		FinalGrades: []string{"A", "A", "B", "A"},
		CreditHours: 14,
		Password1:   "password",
	}
	stew := Education{
		Name:        "Stewart",
		School:      "Georgia Institute of Technology",
		Grade:       14,
		Gpa:         4.0,
		FinalGrades: []string{"A", "A", "A", "A", "A"},
		CreditHours: 18,
		Password1:   "password",
	}
	jonah := Education{
		Name:        "Jonah",
		School:      "Georgia Highlands College",
		Grade:       14,
		Gpa:         2.75,
		FinalGrades: []string{"B", "B"},
		CreditHours: 6,
		Password1:   "password",
	}
	erin := Education{
		Name:        "Erin",
		School:      "Harrison High School",
		Grade:       12,
		Gpa:         4.25,
		FinalGrades: []string{"A", "A", "A", "A"},
		CreditHours: 4,
		Password1:   "password",
	}
	sarah := Education{
		Name:        "Sarah",
		School:      "Kennesaw State University",
		Grade:       13,
		Gpa:         2.6,
		FinalGrades: []string{"A", "B", "A", "C"},
		CreditHours: 13,
		Password1:   "password",
	}
	person := Education{}
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter Username: ")
	username, _ := reader.ReadString('\n')
	username = strings.Replace(username, "\r\n", "", -1)

	reader1 := bufio.NewReader(os.Stdin)
	fmt.Print("Enter Password: ")
	password, _ := reader1.ReadString('\n')
	password = strings.Replace(password, "\r\n", "", -1)

	test := false
	if username == "dokeefe" {
		person = devin
		if password == devin.Password1 {
			test = true
		} else {
			fmt.Println("The Password you have entered is incorrect")
		}
	} else if username == "sfronk" {
		person = stew
		if password == person.Password1 {
			test = true
		} else {
			fmt.Println("The Password you have entered is incorrect")
		}
	} else if username == "jcoleman" {
		person = jonah
		if password == person.Password1 {
			test = true
		} else {
			fmt.Println("The Password you have entered is incorrect")
		}
	} else if username == "eokeefe" {
		person = erin
		if password == person.Password1 {
			test = true
		} else {
			fmt.Println("The Password you have entered is incorrect")
		}
	} else if username == "smiller" {
		person = sarah
		if password == person.Password1 {
			test = true
		} else {
			fmt.Println("The Password you have entered is incorrect")
		}
	} else {
		fmt.Println("The username you entered was not found")
	}
	if test == true {
		t := reflect.TypeOf(Education{})
		field, _ := t.FieldByName("CreditHours")
		var i string = string(field.Tag)
		j, err := strconv.Atoi(i)
		if err == nil {
			if j >= person.CreditHours {
				fmt.Printf("Name: %v Grade: %v GPA: %v Final Grades: %v Credit Hours: %v ", person.Name, person.Grade, person.Gpa, person.FinalGrades, person.CreditHours)
			} else {
				fmt.Printf("Name: %v Grade: %v GPA: %v Final Grades: %v %v is currently signed up for more than the max ammount of credit hours that can be taken in a semster", person.Name, person.Grade, person.Gpa, person.FinalGrades, person.Name)
			}
		}
	}

}
