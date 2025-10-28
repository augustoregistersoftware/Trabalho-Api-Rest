from pydantic import BaseModel

class Produto(BaseModel):
    id: int | None = None
    descricao: str
    preco: float
    marca: str
