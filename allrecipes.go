package main

import (
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

func getId(w http.ResponseWriter, r *http.Request) int {
	rand.Seed(time.Now().UTC().UnixNano())

	for i := 0; i < 100; i++ {
		id := rand.Intn(262701)

		resp, err := http.Get("http://allrecipes.com/recipe/" + strconv.Itoa(id) + "/")
		if err != nil {
			panic(err)
		}
		defer resp.Body.Close()

		title := getTitle(resp.Body)
		if title != "Johnsonville® Three Cheese Italian Style Chicken Sausage Skillet Pizza Recipe - Allrecipes.com" &&
			title != "Allrecipes - Server Error" &&
			title != "Allrecipes - File Not Found" {
			return id
		}
	}

	return 0
}
