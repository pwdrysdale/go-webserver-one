// https://www.youtube.com/watch?v=YS4e4q9oBaU
// my first http server in go
// thanks github copilot
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler)

	http.HandleFunc("/hello", helloHandler)

	http.HandleFunc("/todos", handleTodos)

	// post request
	http.HandleFunc("/test", test)

	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s!", "Go")
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World")
}

func handleTodos(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		getTodos(w, r)
	//case "POST":
	//	addTodo(w, r)
	default:
		w.WriteHeader(http.StatusNotImplemented)
		w.Write([]byte(http.StatusText(http.StatusNotImplemented)))
	}
}

func getTodos(w http.ResponseWriter, r *http.Request) {
	// get todos from json placeholder
	resp, err := http.Get("https://jsonplaceholder.typicode.com/todos")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(http.StatusText(http.StatusInternalServerError)))
		return
	}

	//send the todos back to the client
	type Todo struct {
		UserId int    `json:"userId"`
		Id     int    `json:"id"`
		Title  string `json:"title"`
		Done   bool   `json:"completed"`
	}

	// get body in desired format for use in go
	todos := make([]Todo, 0, 10)
	err = json.NewDecoder(resp.Body).Decode(&todos)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	// get ready to send as json and send it
	marshalled, ok := json.Marshal(todos)
	if ok != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(ok.Error()))
		return
	} else {
		fmt.Fprintf(w, "%v", string(marshalled))
	}

}

type test_struct struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
	Test string `json:"test"`
}

func test(rw http.ResponseWriter, req *http.Request) {
	decoder := json.NewDecoder(req.Body)
	var t test_struct
	err := decoder.Decode(&t)
	if err != nil {
		panic(err)
	}
	log.Println(t)
}
