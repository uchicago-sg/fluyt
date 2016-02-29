package fluyt

import (
	"encoding/json"
	"github.com/fatlotus/scroll"
	"net/http"
	"sync"
)

type Marketplace struct {
	Listings        map[string]Listing
	Backend         scroll.Log
	Cursor          scroll.Cursor
	PendingListings map[Person]map[string]Listing
	sync.Mutex
}

func NewMarketplace() *Marketplace {
	log := scroll.DatastoreLog(nil, "Operation")
	return &Marketplace{
		Listings:        make(map[string]Listing),
		PendingListings: make(map[Person]map[string]Listing),
		Backend:         log,
		Cursor:          log.Cursor(),
	}
}

func (m *Marketplace) Sync() error {
	listing := Listing{}
	for {
		err := m.Cursor.Next(&listing)
		if err == scroll.Done {
			break
		} else if err != nil {
			panic(err)
		}

		m.Listings[listing.Permalink] = listing
	}
	return nil
}

type Results struct {
	Listings []Listing `json:"listings"`
}

func (m *Marketplace) SearchListings(w http.ResponseWriter, r *http.Request) {
	query := r.FormValue("q")
	results := &Results{make([]Listing, 0)}

	for _, listing := range m.Listings {
		if listing.Match(query) {
			results.Listings = append(results.Listings, listing)
		}
	}

	json.NewEncoder(w).Encode(&results)
}
