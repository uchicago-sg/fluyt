package fluyt

import (
	"encoding/json"
	"github.com/fatlotus/scroll"
	"golang.org/x/net/context"
	"net/http"
	"sort"
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
	log := makeBackend()
	return &Marketplace{
		Listings:        make(map[string]Listing),
		PendingListings: make(map[Person]map[string]Listing),
		Backend:         log,
		Cursor:          log.Cursor(),
	}
}

func (m *Marketplace) Sync(ctx context.Context) error {
	listing := Listing{}
	for {
		err := m.Cursor.Next(ctx, &listing)
		if err == scroll.Done {
			break
		} else if err != nil {
			panic(err)
		}

		m.Listings[listing.Permalink] = listing
	}
	return nil
}

type ord []Listing

func (s ord) Len() int           { return len(s) }
func (s ord) Less(i, j int) bool { return s[i].LastUpdate.Before(s[j].LastUpdate) }
func (s ord) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }

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

	sort.Sort(ord(results.Listings))

	json.NewEncoder(w).Encode(&results)
}
