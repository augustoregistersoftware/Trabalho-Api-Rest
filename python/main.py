from fastapi import FastAPI, HTTPException
from models import Produto
from database import get_connection

app = FastAPI(title="API de Produtos")

@app.get("/produtos")
def listar_produtos():
    conn = get_connection()
    cursor = conn.cursor(dictionary=True)
    cursor.execute("SELECT id_produto, descricao, preco, marca FROM produto")
    produtos = cursor.fetchall()
    cursor.close()
    conn.close()
    return produtos

@app.post("/produtos")
def criar_produto(produto: Produto):
    conn = get_connection()
    cursor = conn.cursor()
    sql = "INSERT INTO produto (descricao, preco, marca) VALUES (%s, %s, %s)"
    cursor.execute(sql, (produto.descricao, produto.preco, produto.marca))
    conn.commit()
    produto.id = cursor.lastrowid
    cursor.close()
    conn.close()
    return produto

@app.delete("/produtos/{id}")
def deletar_produto(id: int):
    conn = get_connection()
    cursor = conn.cursor()
    cursor.execute("DELETE FROM produto WHERE id_produto = %s", (id,))
    conn.commit()
    cursor.close()
    conn.close()
    return {"message": f"Produto {id} deletado com sucesso"}
