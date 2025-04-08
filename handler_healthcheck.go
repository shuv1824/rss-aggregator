package main

import (
	"net/http"
)

func handlerHealthcheck(w http.ResponseWriter, r *http.Request) {
	respondWithJSON(w, 200, struct{}{})
}
