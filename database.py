import mysql.connector

def get_connection():
    return mysql.connector.connect(
        host="db-trab-api-rest.mysql.database.azure.com",
        user="admrest",
        password="Tr@b@lh0r3st",
        database="produtos"
    )
