package main

import (
	"fmt"
	"github/zenith/goframework/framework"
	"net/http"
)

func main() {
	server := &http.Server{
		Addr:              ":8080",
		Handler:           framework.NewCore(),
	}
	server.ListenAndServe()
	fmt.Println("go framework run...")
}