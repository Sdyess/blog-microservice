package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func getAllBlogPosts(w http.ResponseWriter, r *http.Request) {

}

func getBlogPost(w http.ResponseWriter, r *http.Request) {

}

func main() {

	router := mux.NewRouter()
	router.HandleFunc("/blog/posts", getAllBlogPosts).Methods("GET")
	router.HandleFunc("/blog/posts/{id}", getBlogPost).Methods("GET")

	err := http.ListenAndServe(":8000", router)
	if err != nil {
		fmt.Print(err)
	}
}
