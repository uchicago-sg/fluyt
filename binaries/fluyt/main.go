// +build !appengine

package main

import (
	_ "github.com/uchicago-sg/fluyt"
	"net/http"
)

func main() {
	http.ListenAndServe("localhost:9000", nil)
}
