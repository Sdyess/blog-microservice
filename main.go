package main

import (
	"encoding/json"
	"github.com/blog-microservice/models"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
)

func getAllBlogPosts(w http.ResponseWriter, r *http.Request) {
	const postContent = "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Suspendisse mauris ligula, bibendum sit amet ex eu, pharetra dignissim dui. Nunc sed imperdiet ligula, vitae ultricies justo. Donec vehicula pellentesque ligula. Integer sit amet diam finibus, blandit ligula a, pretium enim. Suspendisse semper neque rhoncus mattis lobortis. Ut commodo mollis nunc. Donec vehicula, nibh vel eleifend facilisis, purus odio efficitur nunc, nec facilisis enim dui ut enim. Phasellus et dolor eu massa aliquet vehicula varius et risus."

	post1 := models.Post{Id: 1, Title: "Blog Post 1", Content: postContent}
	post2 := models.Post{Id: 2, Title: "Blog Post 2", Content: postContent}
	post3 := models.Post{Id: 3, Title: "Blog Post 3", Content: postContent}

	posts := []models.Post{post1, post2, post3}
	postResp := models.Posts{Data: posts}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(postResp)
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

	log.Fatal(http.ListenAndServe(":"+port, handlers.CORS(handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}), handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"}), handlers.AllowedOrigins([]string{"*"}))(router)))
}
