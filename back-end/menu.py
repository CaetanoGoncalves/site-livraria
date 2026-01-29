import httpx
from tabulate import tabulate
import time
import subprocess
import os
from dotenv import load_dotenv

load_dotenv("config/.env")

homeUrl = os.getenv("BASE_LINK")
def clearTeminal():
    command = "cls"
    subprocess.run(command, shell=True, check=True)
    return
class DB:
    def __init__(self):
        self.listUrl = homeUrl+"list"
        self.deleteUrl = homeUrl+"delete"
        self.updateUrl = homeUrl+"update"
        self.insertUrl = homeUrl+"insert"
        self.searchUrl = homeUrl+"search"
        self.searchAuthorUrl = homeUrl+"author"
        self.searchNameUrl = homeUrl+"name"
    def searchByAuthor(self, author):
        payload = {"author": author}
        
        try:
            response = httpx.post(self.searchAuthorUrl, json=payload)
            returnResponse = {
                "status": response.status_code,
                "outArray": response.json()
            }
            return returnResponse
        except:
            return {"status": 500, "outArray": []}
    def searchByName(self, name):
        payload = {"name": name}
        
        try:
            response = httpx.post(self.searchNameUrl, json=payload)
            returnResponse = {
                "status": response.status_code,
                "outArray": response.json()
            }
            return returnResponse
        except:
            return {"status": 500, "outArray": []}
    def search(self, id):
        payload = {"id": id}
        
        try:
            response = httpx.post(url=self.searchUrl, json=payload)
            if response.status_code != 200:
                return {"status": response.status_code}
            data_list = response.json() 
            
            if not data_list:
                return {"status": 404}

            book = data_list[0]
            
            returnResponse = {
                "status": response.status_code,
                "id": book["id"],
                "name": book["name"],
                "author": book["author"],
                "price": book["price"]
            }
            return returnResponse

        except httpx.RequestError:
            return {"status": 500}
        except Exception:
            return {"status": 500}

    def listDB(self):
        try:
            response = httpx.get(self.listUrl)
            returnResponse = {
                "status": response.status_code,
                "outArray": response.json()
            }
            return returnResponse
        except:
            return {"status": 500, "outArray": []}

    def deleteFromDB(self, id):
        payload = {"id": id}
        try:
            response = httpx.post(url=self.deleteUrl, json=payload)
            return response.status_code
        except:
            return 500

    def insertIntoDB(self, name, author, price):
        payload = {
            "name": name,
            "author": author,
            "price": price
        }
        try:
            response = httpx.post(url=self.insertUrl, json=payload)
            return response.status_code
        except:
            return 500
            
    def updateElementFromDB(self, id="", newName="", newAuthor="", newPrice=""):
        payload = {
            "id":id,
            "newName": newName,
            "newAuthor": newAuthor,
            "newPrice": newPrice
        }
        try:
            response = httpx.patch(url=self.updateUrl, json=payload)
            return response.status_code
        except:
            return 500


class Menu:
    def __init__(self, DB_obj):
        self.DB_obj = DB_obj

    def start(self):
        print("\n--- Gerenciador de Livros ---")
        while True:
            
            print("\nO que você gostaria de fazer?")
            print("1. Listar todos livros")
            print("2. Deletar um livro especifico")
            print("3. Mudar as informações de um livro")
            print("4. Inserir um livro na tabela")
            print("5. Encontrar por ID")
            print("6. Encontrar por autor")
            print("7. Encontrar por nome")
            print("0. Sair")
            
            try:
                option = int(input("Escolha: "))
            except ValueError:
                clearTeminal()
                print("Por favor, digite um número.")
                continue

            if option == 0:
                print("Saindo...")
                break
            elif option == 1:
                self.listTable()
            elif option == 2:
                self.deleteItem()
            elif option == 3:
                self.updateItem()
            elif option == 4:
                self.insertItem()
            elif option == 5:
                self.searchItem()
            elif option == 6:
                self.searchAuthor()
            elif option == 7:
                self.searchName()
            else:
                print("Opção inválida.")
            input("Aperte qualquer tecla para continuar...")
            clearTeminal()
    def listTable(self):
        print("Listando os livros...")
        data = self.DB_obj.listDB()
        if data["status"] == 200:
            print(tabulate(data["outArray"], headers="keys", tablefmt="fancy_grid"))
        else:
            print("Erro ao conectar com o servidor.")

    def deleteItem(self):
        id_livro = input("Diga o ID do livro: ")
        print(f"Tem certeza de que gostaria de apagar o livro ID {id_livro}?")
        choosen = input("Y: Sim, Qualquer outra tecla: Não | ").upper()
        
        if choosen == "Y":
            response = self.DB_obj.deleteFromDB(id_livro)
            if response == 404:
                print("Nenhum livro encontrado com esse ID.")
            elif response == 200:
                print("Deletado com sucesso!")
            else:
                print(f"Erro no servidor: {response}")
        else:
            print("Operação cancelada.")

    def searchItem(self):
        id_livro = input("Diga o ID do livro: ")
        print("Procurando...")
        
        line = self.DB_obj.search(id_livro)
        
        if line["status"] != 200:
            print(f"Erro ou nenhum livro encontrado (Status: {line['status']})")
        else:
            display_data = {k: v for k, v in line.items() if k != "status"}
            tabela = tabulate([display_data], headers="keys", tablefmt="fancy_grid")
            print(tabela)
    def updateItem(self):
        id_livro = input("Diga o ID do livro: ")
        busca = self.DB_obj.search(id_livro)
        if busca["status"] == 200:
            print(f"Informações atuais: name: {busca["name"]}, author: {busca["author"]}, price: {busca["price"]}")
            newName = input("Insira o novo nome: ")
            newAuthor = input("Insira o novo autor: ")
            newPrice = input("Insira o novo preço: ")
            self.DB_obj.updateElementFromDB(id_livro, newName, newAuthor, newPrice)
            busca = self.DB_obj.search(id_livro)
            print(f"Informações novas: name: {busca["name"]}, author: {busca["author"]}, price: {busca["price"]}")
        else:
            print("Nenhum livro encontrado com esse ID")
    def insertItem(self):
        name = input("Diga o nome do livro: ")
        author = input("Diga o nome do autor do livro: ")
        price = input("Diga o do preço do livro: ")
        self.DB_obj.insertIntoDB(name, author, price)
        print(f"Livro '{name}' inserido com sucesso")
    def searchAuthor(self):
        author = input("Diga o nome do autor: ")
        print("Procurando...")
        data = self.DB_obj.searchByAuthor(author)
        if data["status"] == 200:
            print(tabulate(data["outArray"], headers="keys", tablefmt="fancy_grid"))
        else:
            print("Erro ao encontrar autor")
    def searchName(self):
        name = input("Diga o nome do livro: ")
        print("Procurando...")
        data = self.DB_obj.searchByName(name)
        if data["status"] == 200:
            print(tabulate(data["outArray"], headers="keys", tablefmt="fancy_grid"))
        else:
            print("Erro ao encontrar livro")
Database = DB()
Screen = Menu(Database)
Screen.start()