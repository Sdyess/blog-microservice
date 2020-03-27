package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {

	router := mux.NewRouter()


	err := http.ListenAndServe(":8000", router)
	if err != nil {
		fmt.Print(err)
	}
}
