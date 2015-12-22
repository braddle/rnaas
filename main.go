package main

import (
	"fmt"
	"log"
	"net/http"
	"romannumeral"
	"strings"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		n := strings.Replace(r.URL.Path, "/", "", -1)

		rn, err := romannumeral.NewRomanNumeral(n)

		if err != nil {
			fmt.Fprintf(w, "Error: %v", err)
		} else {
			fmt.Fprintf(w, "%v = %v", n, rn.ToInt())
		}
	})

	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}
