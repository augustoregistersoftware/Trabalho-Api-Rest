# Projeto APi Rest - Produtos

Bem-vindo ao projeto **APi Rest - Produtos**. Este guia tem como objetivo ajudar na configuração inicial do ambiente de desenvolvimento para que você possa executar o projeto corretamente.

## Requisitos

- [XAMPP](https://www.apachefriends.org/index.html) instalado (contendo Apache, MySQL, PHP).
- Navegador de sua preferência.

## Passos para Configuração

### 1. Baixar e Instalar o XAMPP

1. Baixe a versão nao tao recente do XAMPP a partir do [site oficial](https://www.apachefriends.org/index.html).
2. Instale o XAMPP em seu sistema. Durante a instalação, certifique-se de que os módulos **Apache** e **MySQL** estão selecionados.

### 2. Iniciar o XAMPP

1. Abra o **XAMPP Control Panel**.
2. Inicie o **Apache** e o **MySQL** clicando nos botões "Start" ao lado de cada um.

### 3. Configurar o Projeto

1. Pegue o Backup do banco na pasta `\backup`.
2. Crie um Banco de dados no MySQL com o nome produtos
3. Acesse o seu banco de dados MySQL e faça o backup
## Python
```bash
pip install fastapi uvicorn mysql-connector-python
```

```bash
uvicorn main:app --reload
````
```
A API vai rodar em:

🌍 http://127.0.0.1:8000
```

### 4. Executar o Projeto

1. Abra seu PostMan ou Navegador
## POST 
```bash
curl -X POST -H "Content-Type: application/json" \
-d '{"descricao":"Teste","preco":5899.9,"marca":"Teste Marca"}' \
http://localhost:8080/produtos
```

## GET
```bash
curl http://localhost:8080/produtos
```

## DELETE
```bash
curl -X DELETE http://localhost:8080/produtos/id
```

## Problemas Comuns

- **Porta Ocupada**: Se o Apache não iniciar, pode ser que a porta 80 esteja em uso. Você pode alterar a porta no arquivo de configuração do Apache (httpd.conf) ou liberar a porta fechando o programa que está utilizando-a.

## Contato

Para dúvidas ou suporte, entre em contato pelo e-mail do desenvolvedor.

