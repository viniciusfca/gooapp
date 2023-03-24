package main

import "net/http"

func main() {
	http.HandleFunc("/goapp", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("<h1>Hello World</h1>"))
	})

	http.ListenAndServe(":8080", nil)
}
