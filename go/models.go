package main

type Produto struct {
	ID_PRODUTO int     `json:"id_produto"`
	DESCRICAO  string  `json:"descricao"`
	PRECO      float64 `json:"preco"`
	MARCA      string  `json:"marca"`
}
