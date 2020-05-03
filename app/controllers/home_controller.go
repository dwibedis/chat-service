package controllers

import (
	"fmt"
	"log"
	"net/http"
)

func Welcome(w http.ResponseWriter, r *http.Request)   {
	_, err := fmt.Fprintf(w, "Welcome to Chat project!")
	if err != nil {
		log.Println("error occured ", err)
	}
}
