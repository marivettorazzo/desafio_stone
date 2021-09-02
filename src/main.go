package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"exemple.com/DesafioStone/src/controllers"
	"exemple.com/DesafioStone/src/model"
	"github.com/gorilla/mux"
)

func testeRoute(w http.ResponseWriter, r *http.Request) { //returns the port the server is running on
	fmt.Fprintf(w, "Server running on port 3000")
}

func showAccountsBalance(w http.ResponseWriter, r *http.Request) { //function that returns the balance looking for the id via the url
	w.Header().Set("Content-Type", "aplication/json") // get the response variable (w) set the type of the http header value to app/json
	params := mux.Vars(r)                             // Shows account balance with id specified in route

	for _, item := range controllers.DataAcc {
		id := strconv.Itoa(item.ID)
		if id == params["account_id"] { //makes a range traversing the dataset passed in the slice comparing if item is equal id passed in route

			json.NewEncoder(w).Encode(item.Balance)

			return
		}
	}

}

func CreateAccount(w http.ResponseWriter, r *http.Request) { //route that creates an account
	w.Header().Set("Content-Type", "aplication/json")
	var cria model.Account
	json.NewDecoder(r.Body).Decode(&cria)

	x, err := controllers.InsertLieDatabase(cria)
	if err != nil {
		err := "400 Bad Request"
		json.NewEncoder(w).Encode(err)
	}
	json.NewEncoder(w).Encode(x)
}

func showAccounts(w http.ResponseWriter, r *http.Request) { //shows all registered accounts
	//turns information into json

	json.NewEncoder(w).Encode(controllers.DataAcc)

}

// routes Transfer
func TransferRoute(w http.ResponseWriter, r *http.Request) { //route that performs the transfer transaction
	w.Header().Set("Content-type", "aplication/json")
	var b model.Transfer

	json.NewDecoder(r.Body).Decode(&b)
	result := controllers.TransferringUnique(b)

	json.NewEncoder(w).Encode(result)

}
func GetTransfers(w http.ResponseWriter, r *http.Request) { //route showing all transactions performed

	json.NewEncoder(w).Encode(controllers.DataTran)
}

//route login
func Login(w http.ResponseWriter, r *http.Request) { //login validation route

	var l model.Login
	json.NewDecoder(r.Body).Decode(&l)
	result := controllers.ValidaLogin(l)
	json.NewEncoder(w).Encode(result)

}

func configuringServer() {
	router := mux.NewRouter()
	router.HandleFunc("/", testeRoute).Methods("GET")
	router.HandleFunc("/accounts", showAccounts).Methods("GET")
	router.HandleFunc("/accounts/{account_id}/balance", showAccountsBalance).Methods("GET")
	router.HandleFunc("/accounts", CreateAccount).Methods("Post")
	// Routers Transfer
	router.HandleFunc("/transfer", TransferRoute).Methods("Post")
	router.HandleFunc("/transfer", GetTransfers).Methods("GET")
	// Route Login
	router.HandleFunc("/login", Login).Methods("Post")

	log.Fatal(http.ListenAndServe(":3000", router)) //DefaultServerMux localhost : 3000 , receives the port and the route of ok as stated above
}

func main() {

	configuringServer()

}

// O desafio será avaliado através de quatro critérios.

// Entrega
// O resultado final está completo para ser executado?
// O resultado final atende ao que se propõe fazer?
// O resultado final atende totalmente aos requisitos propostos?
// O código possui algum controle de dependências?
// Boas Práticas
// O código está de acordo com o guia de estilo do Go?
// O código está bem estruturado?
// O código está fluente na linguagem?
// O código faz o uso correto de Design Patterns?
// Documentação
// O código foi entregue com um arquivo de README claro de como se guiar?
// O código possui comentários pertinentes?
// O código está em algum controle de versão?
// Os commits são pequenos e consistentes?
// As mensagens de commit são claras?
// Código Limpo
// O código possibilita expansão para novas funcionalidades?
// O código é Don't Repeat Yourself?
// O código é fácil de compreender?
// O código possui testes?
// Material de Estudo
// Go
// Aprenda Go com Testes
// Curso Aprenda Go (@ellenkorbes)
// Learn Go
// Gophercise
// Build Web Application with Golang
// Error Handling
// DDD Lite in Go
// Repository Pattern in Go
// Rest
// Desing RESTful API's
// HTTP Status Code
// Boas praticas
// Boas Práticas na Stone
// Uber-go Guide
// Outros
// SOLID
// SOLID in GO
// Grupo de Estudos de Go (pt-br)
// Web Development with Go (samples) - Jon Calhoun
// Go Bootcamp from Gopherguides.tv
// Just for Func
// Go WEB Examples
// Dave Cheney Blog
// Ardan Labs Blog
// Comunidade Go
// Slack Gopher
// Telegram
// Sugestões
// Gorilla Mux
// Negroni
// Chi
// sirupsen/log
