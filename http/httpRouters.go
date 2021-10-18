package http

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"exemple.com/DesafioStone/desafio_stone/domain"
	"exemple.com/DesafioStone/desafio_stone/domain/usecases"
	"github.com/gorilla/mux"
)

func testeRoute(w http.ResponseWriter, r *http.Request) { //returns the port the server is running on
	fmt.Fprintf(w, "Server running on port 3000")
}

func showAccountsBalance(w http.ResponseWriter, r *http.Request) { //function that returns the balance looking for the id via the url
	w.Header().Set("Content-Type", "aplication/json") // get the response variable (w) set the type of the http header value to app/json
	params := mux.Vars(r)                             // Shows account balance with id specified in route

	for _, item := range usecases.DataAcc {
		id := strconv.Itoa(item.ID)
		if id == params["account_id"] { //makes a range traversing the dataset passed in the slice comparing if item is equal id passed in route

			json.NewEncoder(w).Encode(item.Balance)

			return
		}
	}

}

func CreateAccount(w http.ResponseWriter, r *http.Request) { //route that creates an account
	w.Header().Set("Content-Type", "aplication/json")
	var cria domain.Account
	json.NewDecoder(r.Body).Decode(&cria)

	x, err := usecases.InsertLieDatabase(cria)
	if err != nil {
		err := "400 Bad Request"
		json.NewEncoder(w).Encode(err)
	}
	json.NewEncoder(w).Encode(x)
}

func showAccounts(w http.ResponseWriter, r *http.Request) { //shows all registered accounts
	//turns information into json

	json.NewEncoder(w).Encode(usecases.DataAcc)

}

// routes Transfer
func TransferRoute(w http.ResponseWriter, r *http.Request) { //route that performs the transfer transaction
	w.Header().Set("Content-type", "aplication/json")
	var b domain.Transfer

	json.NewDecoder(r.Body).Decode(&b)
	result := usecases.TransferringUnique(b)

	json.NewEncoder(w).Encode(result)

}
func GetTransfers(w http.ResponseWriter, r *http.Request) { //route showing all transactions performed

	json.NewEncoder(w).Encode(usecases.DataTran)
}

//route login
func Login(w http.ResponseWriter, r *http.Request) { //login validation route

	var l domain.Login
	json.NewDecoder(r.Body).Decode(&l)
	result := usecases.ValidaLogin(l)
	json.NewEncoder(w).Encode(result)

}

func ConfiguringServer() {
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
