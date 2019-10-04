package main

const (
	HEAD = `package main

	import (
		"fmt"
		"net/http"
	)
	
	func main() {`

	// Replace: ROUTE, STATUS, HeaderKey, HeaderValue, BODY
	HANDLER = `http.HandleFunc("{{ROUTE}}", func(w http.ResponseWriter, r *http.Request) {
		{{HEADERS}}
		w.WriteHeader({{STATUS}})
		fmt.Fprintf(w, ` + "`" + `{{BODY}}` + "`" + `)
	})`

	HEADER = `w.Header().Set("{{KEY}}", "{{VALUE}}")`

	TAIL = `err := http.ListenAndServe(":8080", nil)
		if err != nil {
			panic(err)
		}
	}`
)
