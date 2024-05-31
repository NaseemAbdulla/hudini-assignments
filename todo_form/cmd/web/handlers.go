package main

import (
	"html/template"
	"log"
	"net/http"
	"strconv"
)

// createing a structure of Todo
type Todo struct {
	Text string
	Id   int
}

// intialising a slice of allTods as empty
var allTodos []Todo
var id int

func getTask(w http.ResponseWriter, r *http.Request) {
	//  retrieving the home template
	files := []string{
		"./ui/html/home.page.tmpl",
	}
	// parsing the template to ts
	ts, err := template.ParseFiles(files...)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "interal server error", 500)
	}
	// Execute the template and show the slice of allTodo
	err = ts.Execute(w, allTodos)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "interal server error", 500)
	}
}
func addTask(w http.ResponseWriter, r *http.Request) {
	// creating an instance of Todo
	task := Todo{
		Text: r.FormValue("task"),
		Id:   id,
	}
	// incrmenting the id
	id++
	// appending the list
	allTodos = append(allTodos, task)
	http.Redirect(w, r, "/", http.StatusSeeOther)

}
func deleteTask(w http.ResponseWriter, r *http.Request) {
	// fetch the id from the form
	id, _ := strconv.Atoi(r.FormValue("id"))
	// iterate through the slice of Alltods and delete the item using id
	for i, value := range allTodos {
		if value.Id == id {
			allTodos = append(allTodos[:i], allTodos[i+1:]...)
			break
		}
	}
	// redirect the to the home
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
