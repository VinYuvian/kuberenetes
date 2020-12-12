package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
)

//Users ...
var Users []User
var tkSecret = []byte("vinay")
var tkRefSecret = []byte("refreshSecret")

//User ...
type User struct {
	ID           string `json:"userid"`
	Firstname    string `json:"firstname"`
	Lastname     string `json:"lastname"`
	Gender       string `json:"gender"`
	Email        string `json:"email"`
	Password     string `json:"password"`
	RefreshToken string `json:"refreshtoken"`
}

//LoginHandler ...
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var resp User
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept")
	dec := json.NewDecoder(r.Body)
	err := dec.Decode(&resp)
	IsRegistered := IsRegistered(resp)
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}else{
		fmt.Fprintf(w, "The user is logged in")
	}


}

//IsRegistered ...
func IsRegistered(user User) bool {
	var state bool
	count := 0
	for _, value := range Users {
		if value.Email == user.Email {
			count = count + 1
		}
		if count == 1 {
			state = true
		} else {
			state = false
		}
	}
	return state
}


//SignUpHandler ...
func SignUpHandler(w http.ResponseWriter, r *http.Request) {
	var resp User
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Context-Type","application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept")
	Dec := json.NewDecoder(r.Body)
	err := Dec.Decode(&resp)
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}
	IsRegistered := IsRegistered(resp)
	if IsRegistered == true {
		fmt.Fprintf(w, "The email is already Registered")

	} else {
		Users = append(Users, resp)
		fmt.Fprintf(w, "The account is now Registered\n")
	}

}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/login", LoginHandler)
	r.HandleFunc("/signup", SignUpHandler)
	log.Fatal(http.ListenAndServe(":8081", r))
}
