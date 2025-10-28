package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	ConnectDB()

	r := mux.NewRouter()

	r.HandleFunc("/produtos", GetProdutos).Methods("GET")
	r.HandleFunc("/produtos", CreateProduto).Methods("POST")
	r.HandleFunc("/produtos/{id}", DeleteProduto).Methods("DELETE")

	log.Println("🚀 Servidor rodando em http://localhost:8080")
	http.ListenAndServe(":8080", r)
}
