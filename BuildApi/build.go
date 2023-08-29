package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type Courses struct {
	CourseId    string  `json:"courseid"`
	CourseName  string  `json:"coursename"`
	CoursePrice int     `json:"price"`
	Author      *Author `json:"author"`
}

type Author struct {
	FullName string `json:"fullname"`
	Website  string `json:"website"`
}

// fake DB
var courses []Courses

//middleware,helper-file

func (c *Courses) IsEmpty() bool {
	return c.CourseId == "" && c.CourseName == ""
}

func main() {

}

func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to API learncodeOnline</h1>"))
}

func getAllCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get all courses")
	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(courses)
}
func getOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get one courses")
	w.Header().Set("content-type", "application/json")

	//grab id from request

	params := mux.Vars(r)

	//loop through the coures,find matching id and return the respone

	for _,course:=range courses{
		if course.CourseId==params["id"]{
			json.NewEncoder(w).Encode(course)
			return
		}

	}

	json.NewEncoder(w).Encode("no course found in thid id")
	return
}

func createOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("create one courses")
	w.Header().Set("content-type", "application/json")
}
