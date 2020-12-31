package main

import (
	"fmt"
	"net/http"

	"github.com/yzlq99/yin/yinS"
)

func main() {
	e := yinS.New()
	e.GET("/", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "%s", req.URL)
	})

	e.Run(":8080")
}
