package main

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// GET /produtos
func GetProdutos(w http.ResponseWriter, r *http.Request) {
	rows, err := db.Query("SELECT id_produto,descricao,preco,marca FROM produto")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var produtos []Produto
	for rows.Next() {
		var u Produto
		rows.Scan(&u.ID_PRODUTO, &u.DESCRICAO, &u.PRECO, &u.MARCA)
		produtos = append(produtos, u)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(produtos)
}

// POST /produtos
func CreateProduto(w http.ResponseWriter, r *http.Request) {
	var u Produto
	json.NewDecoder(r.Body).Decode(&u)

	stmt, err := db.Prepare("INSERT INTO produto(descricao, preco, marca) VALUES(?, ?, ?)")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec(u.DESCRICAO, u.PRECO, u.MARCA)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(u)
}

// DELETE /produto/{id}
func DeleteProduto(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	stmt, err := db.Prepare("DELETE FROM produto WHERE id_produto = ?")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
