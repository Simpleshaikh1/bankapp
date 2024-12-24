package api

import (
	"encoding/json"
	"fintechApp/helpers"
	"fintechApp/users"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
)

type Login struct {
	Username string
	Password string
}

type ErrorResponse struct {
	Message string
}

func login(w http.ResponseWriter, r *http.Request) {
	//Read body
	body, err := ioutil.ReadAll(r.Body)
	helpers.HandleError(err, "Failed to read request body")

	//Handle Login

	var formattedBody Login
	err = json.Unmarshal(body, &formattedBody)
	helpers.HandleError(err, "Failed to unmarshal request body")

	login := users.Login(formattedBody.Username, formattedBody.Password)

	//Prepare Response
	if login["message"] == "all is fine" {
		resp := login
		json.NewEncoder(w).Encode(resp)
	} else {
		//handle error
		resp := ErrorResponse{Message: login["message"].(string)}
		json.NewEncoder(w).Encode(resp)
	}
}

func StartServer() {
	router := mux.NewRouter()
	router.HandleFunc("/login", login).Methods("POST")
	fmt.Println("Server started on port: 8888")
	log.Fatal(http.ListenAndServe(":8888", router))
}
