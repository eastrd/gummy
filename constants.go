package main

const (
	HEAD = `package main

	import (
		"fmt"
		"net/http"
	)
	
	func main() {`

	// Replace: ROUTE, BODY
	HANDLER = `http.HandleFunc("{{ROUTE}}", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, ` + "`" + `{{BODY}}` + "`" + `)
	})`

	TAIL = `err := http.ListenAndServe(":8080", nil)
		if err != nil {
			panic(err)
		}
	}`
)
