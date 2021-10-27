package http

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"exemple.com/DesafioStone/desafio_stone/domain"
	"exemple.com/DesafioStone/desafio_stone/domain/usecases"
	"github.com/gorilla/mux"
)

func testeRoute(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "aplication/json") //returns the port the server is running on
	fmt.Fprintf(w, "Server running on port 3000")
}

func showAccountsBalance(w http.ResponseWriter, r *http.Request) { //function that returns the balance looking for the id via the url
	w.Header().Set("Content-Type", "aplication/json") // get the response variable (w) set the type of the http header value to app/json
	params := mux.Vars(r)
	// Shows account balance with id specified in route

	z, x, err := usecases.GetBalance(params["account_id"])

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(err.Error())
		return
	}
	y := int(x)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(z)
	json.NewEncoder(w).Encode(y)

}

func CreateAccount(w http.ResponseWriter, r *http.Request) { //route that creates an account
	w.Header().Set("Content-Type", "aplication/json")
	var cria domain.Account
	json.NewDecoder(r.Body).Decode(&cria)

	resp := usecases.InsertLieDatabase(cria)
	if resp != nil {

		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(resp.Error())
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("success")

}

func showAccounts(w http.ResponseWriter, r *http.Request) { //shows all registered accounts
	w.Header().Set("Content-Type", "aplication/json")

	x := usecases.Accounts()
	if x != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(x.Error())
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(usecases.DataAcc)
}

// routes Transfer
func TransferRoute(w http.ResponseWriter, r *http.Request) { //route that performs the transfer transaction
	w.Header().Set("Content-type", "aplication/json")
	var b domain.Transfer

	json.NewDecoder(r.Body).Decode(&b)
	result := usecases.TransferringUnique(b)

	if result != nil {

		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(result.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("Success")

}

func GetTransfers(w http.ResponseWriter, r *http.Request) { //route showing all transactions performed
	w.Header().Set("Content-type", "aplication/json")
	x := usecases.Transfers()
	if x != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(x.Error())
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(usecases.DataTran)
}

//route login
func Login(w http.ResponseWriter, r *http.Request) { //login validation route
	w.Header().Set("Content-Type", "aplication/json")
	var l domain.Login
	json.NewDecoder(r.Body).Decode(&l)
	str, result := usecases.ValidaLogin(l)
	if result != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(result.Error())
		return
	}
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(str)

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
	router.HandleFunc("/login", Login).Methods("Get")

	log.Fatal(http.ListenAndServe(":3000", router)) //DefaultServerMux localhost : 3000 , receives the port and the route of ok as stated above
}
