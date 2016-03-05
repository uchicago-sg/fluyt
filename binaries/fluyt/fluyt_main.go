// +build !appengine

package main

import (
	"github.com/uchicago-sg/fluyt"
	"net/http"
)

func main() {
	http.ListenAndServe(":8080", fluyt.MarketplaceHandler())
}
