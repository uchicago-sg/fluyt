// +build appengine

package fluyt

import (
	"github.com/fatlotus/scroll"
	"golang.org/x/net/context"
	"google.golang.org/appengine"
	"net/http"
)

func makeContext(r *http.Request) context.Context {
	return appengine.NewContext(r)
}

func makeBackend() scroll.Log {
	return scroll.DatastoreLog("Operation")
}

func init() {
	http.Handle("/", MarketplaceHandler())
}
