package main

import (
	"fmt"

	"net/http"

	"github.com/gorilla/mux"
)

func htppHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	user := vars["user"]
	password := vars["password"]
	fmt.Fprintf(w, "Your user is %s and password is %s\n", user, password)
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/userRegister/{user}/{password}", htppHandler).Methods("POST")
	http.ListenAndServe(":8080", r)
	fmt.Printf("ola mundo")
	err := http.ListenAndServe(":8080", r)
	if err != nil {
		fmt.Printf("Erro ao iniciar o servidor: %s", err)
	}
}
