// +build appengine

package fluyt

import (
	"google.golang.org/appengine"
	"google.golang.org/appengine/datastore"
	"net/http"
	"time"
)

func Sync(r *http.Request, m *Marketplace) {
	context := appengine.NewContext(r)
	listings := make([]Listing, 0)
	q := datastore.NewQuery("Listing").Filter("LastUpdate >", m.LastUpdate)
	_, err := q.GetAll(context, &listings)
	if err != nil {
		panic(err)
	}
	for _, listing := range listings {
		m.added(listing)
	}
}

func Add(r *http.Request, m *Marketplace, listing *Listing) error {
	listing.LastUpdate = time.Now()
	context := appengine.NewContext(r)
	if listing.Key == "" {
		listing.Key = randStringBytes(30)
	}
	key := datastore.NewKey(context, "Listing", listing.Key, 0, nil)
	key, err := datastore.Put(context, key, listing)
	if err != nil {
		return err
	}
	m.added(*listing)
	return nil
}
