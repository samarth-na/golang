package exercise

import (
	"fmt"
	"net/http"
	"time"
)

func greet(w http.ResponseWriter, r *http.Request) {
	{
		fmt.Fprintf(w, "Hello World! %s", time.Now())
	}
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
}

func server() {
	http.HandleFunc("/", greet)
	http.ListenAndServe(":8080", nil)
}

// func main() {
// 	http.HandleFunc("/", greet)
// 	http.ListenAndServe(":8080", nil)
// }
