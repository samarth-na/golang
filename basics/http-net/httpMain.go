package main

import (
	"alderaan/basics/http-net/api"
	"fmt"
	"net/http"
)

func main() {
	srv := api.Newserver()
	err := http.ListenAndServe(":8080", srv)
	if err != nil {
		fmt.Println(err)
	}
}
