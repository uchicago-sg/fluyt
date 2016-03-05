// +build !appengine

package fluyt

import (
	"encoding/gob"
	"github.com/fatlotus/scroll"
	"golang.org/x/net/context"
	"net/http"
)

func init() {
	gob.Register(Listing{})
}

func makeContext(req *http.Request) context.Context {
	return context.Background()
}

func makeBackend() scroll.Log {
	return scroll.MemoryLog()
}
