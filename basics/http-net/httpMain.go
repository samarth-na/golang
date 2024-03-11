package main

import (
	"alderaan/basics/http-net/api"
	"net/http"
)

func main() {
	srv := api.Newserver()
	err, _ := http.ListenAndServe(":8080", srv)
	if err != nil {
		fmt.Println(err)
	}
}
