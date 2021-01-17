package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/RodrigoAlanisWeb/Go-Mysql/db"
	"github.com/gorilla/mux"
)

func getAllBreakPoint(w http.ResponseWriter, r *http.Request) {
	users := db.GetAll()

	json.NewEncoder(w).Encode(users)
}

func createBreakPoint(w http.ResponseWriter, r *http.Request) {
	var user db.User
	by, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err.Error())
	}
	json.Unmarshal(by, &user)
	db.Create(user)
	fmt.Fprint(w, "Contact Updated Successfully")
	return
}

func getOneBreakPoint(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		panic(err.Error())
	}
	user := db.GetOne(id)
	json.NewEncoder(w).Encode(user)
}

func updateBreakPoint(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		panic(err.Error())
	}
	var user db.User
	by, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err.Error())
	}
	json.Unmarshal(by, &user)
	db.Update(id, user)
	fmt.Fprint(w, "Contact Updated Successfully")
	return
}

func deleteBreakPoint(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		panic(err.Error())
	}
	db.Delete(id)
	fmt.Fprint(w, "Contact Delete Successfully")
	return
}

func handdleRequest() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", getAllBreakPoint).Methods("GET")
	router.HandleFunc("/create", createBreakPoint).Methods("POST")
	router.HandleFunc("/{id}", getOneBreakPoint).Methods("GET")
	router.HandleFunc("/{id}", updateBreakPoint).Methods("PUT")
	router.HandleFunc("/{id}", deleteBreakPoint).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":3000", router))
}

func main() {
	handdleRequest()
}
