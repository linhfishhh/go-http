package main

import (
	"fmt"
	"io"
	"net/http"
)

func httpServer(rw http.ResponseWriter, req *http.Request) {
	rw.Header().Set(
		"content-type",
		"text-html",
	)
	fmt.Println()
	io.WriteString(rw,
		`<html>
			<title>golang server</title>
			<body>
				<h1>golang server test</h1>
			</body>
		</html>`,
	)
}

func main() {
	http.HandleFunc("/hello", httpServer)
	http.ListenAndServe(":9000", nil)
}
