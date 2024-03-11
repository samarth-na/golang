package main

import "net/http"

func Mmain() {
	content := []byte("hello port 8080")
	http.HandleFunc("/hello-8080", func(w http.ResponseWriter, r *http.Request) {
		w.Write(content)
	})
	http.ListenAndServe(":8080", nil)
}
