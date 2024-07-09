package main

import (
	"net/http"
	"os"
)

func main() {

	// for i := 0; i < len(os.Args); i++ {
	// 	fmt.Println(os.Args[i])
	// }

	var handler http.Handler

	handler = http.FileServer(http.Dir(os.Args[2]))

	if len(os.Args[3]) > 3 {
		// Should be in the form of "/foo/"
		handler = http.StripPrefix(os.Args[3], handler)
	}

	// Should be full addr: "0.0.0.0:8080"
	if err := http.ListenAndServe(os.Args[1], handler); err != nil {
		panic(err)
	}
}
