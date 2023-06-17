package main

import (
	"fmt"
	"net/http"
)

type Course struct {
	Name        string
	Description string
	Price       int
}

func (c Course) GetFullInfo() string {
	return fmt.Sprintf("Name: %s, Description: %s, Price: %d", c.Name, c.Description, c.Price)
}

func main() {
	course := Course{
		Name:        "Golang",
		Description: "Golang Course",
		Price:       100,
	}
	fmt.Println(course.GetFullInfo())

	http.HandleFunc("/", home)
	http.ListenAndServe(":8000", nil)
}

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Você está na home"))
}
