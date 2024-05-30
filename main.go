package main

import (
	// "fmt"
	"fmt"
	// "net/http"
	"github.com/Nico2220/go-httprouter/router"
)

func main() {
	// r := router.New()

	// r.Handle(http.MethodGet, "/submit/:id", func(w http.ResponseWriter, r *http.Request, params router.Params) {
	// 	w.Write([]byte("Form submitted!"))
	// })

	s := router.CleanPath("hello")
	fmt.Println(s)

	// http.ListenAndServe(":8080", r)
}
