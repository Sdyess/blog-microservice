package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"os"
)

func getAllBlogPosts(w http.ResponseWriter, r *http.Request) {

}

func getBlogPost(w http.ResponseWriter, r *http.Request) {

}

func main() {

	router := mux.NewRouter()
	router.HandleFunc("/blog/posts", getAllBlogPosts).Methods("GET")
	router.HandleFunc("/blog/posts/{id}", getBlogPost).Methods("GET")

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	err := http.ListenAndServe(":"+port, router)
	if err != nil {
		fmt.Print(err)
	}
}
