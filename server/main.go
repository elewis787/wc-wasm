package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/img", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		fmt.Fprint(w, `<img class="w-full h-80 bg-gray-300 rounded sm:w-96" src="https://avatars.githubusercontent.com/u/10167943?v=4" width="350" />`)
	})

	fmt.Println("Listening on :8081...")
	if err := http.ListenAndServe(":8081", nil); err != nil {
		log.Fatal(err)
	}
}
