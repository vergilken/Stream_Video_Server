package main

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
	"io"
)

func CreateUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	io.WriteString(w, "Create  User Handle")
}

func Login(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	userName := p.ByName("user_name")
	io.WriteString(w, userName)
}