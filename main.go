package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/blog-microservice/models"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"net/http"
	"os"
	"time"
)

var client *mongo.Client

func getAllBlogPosts(w http.ResponseWriter, r *http.Request) {
	var posts []models.Post
	collection := client.Database(os.Getenv("MongoDatabase")).Collection(os.Getenv("MongoCollection"))
	fmt.Println("Database: " + os.Getenv("MongoDatabase"))
	fmt.Println("Collection: " + os.Getenv("MongoCollection"))

	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var post models.Post
		cursor.Decode(&post)
		posts = append(posts, post)
	}
	if err := cursor.Err(); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}

	postResp := models.Posts{Data: posts}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(postResp)
}

func getBlogPost(w http.ResponseWriter, r *http.Request) {
	var post models.Post
	collection := client.Database(os.Getenv("MongoDatabase")).Collection(os.Getenv("MongoCollection"))
	fmt.Println("Database: " + os.Getenv("MongoDatabase"))
	fmt.Println("Collection: " + os.Getenv("MongoCollection"))

	params := mux.Vars(r)
	postId := params["id"]

	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	err := collection.FindOne(ctx, bson.M{"id": postId}).Decode(&post)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(post)
}

func main() {

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	var err error
	client, err = mongo.Connect(ctx, options.Client().ApplyURI(os.Getenv("MongoUri")))
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("MongoUri: " + os.Getenv("MongoUri"))
	router := mux.NewRouter()
	router.HandleFunc("/blog/posts", getAllBlogPosts).Methods("GET")
	router.HandleFunc("/blog/posts/{id}", getBlogPost).Methods("GET")

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	log.Fatal(http.ListenAndServe(":"+port, handlers.CORS(handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}), handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"}), handlers.AllowedOrigins([]string{"*"}))(router)))
}
