package main

import (
	"net/http"
	"strconv"
)

func redirect(w http.ResponseWriter, r *http.Request) {
	id := getId(w, r)

	http.Redirect(w, r, "http://allrecipes.com/recipe/"+strconv.Itoa(id)+"/", 307)
}
