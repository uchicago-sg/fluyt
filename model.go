package fluyt

import (
	"strings"
	"sync"
	"time"
)

// A Listing is an advertisement posted to Marketplace.
type Listing struct {
	Key        string    `json:"key"`
	Title      string    `json:"title"`
	Body       string    `json:"body"`
	LastUpdate time.Time `json:"lastUpdate"`
}

// A Marketplace stores the in-memory representation of all listings.
type Marketplace struct {
	Listings   map[string]Listing
	LastUpdate time.Time
	sync.Mutex
}

// NewMarketplace creates and allocates a new Marketplace representation.
func NewMarketplace() *Marketplace {
	m := &Marketplace{
		Listings: make(map[string]Listing),
	}
	return m
}

func (m *Marketplace) added(listing Listing) {
	m.Listings[listing.Key] = listing
	m.LastUpdate = listing.LastUpdate
}

// Retrieves the given listing from this instance.
func (m *Marketplace) Lookup(path string) *Listing {
	m.Lock()
	defer m.Unlock()

	result, ok := m.Listings[path]
	if !ok {
		return nil
	} else {
		return &result
	}
}

// Search looks for listings matching the given query.
func (m *Marketplace) Search(query string, skip, limit int) []Listing {
	m.Lock()
	defer m.Unlock()

	results := make([]Listing, 0)

	for _, listing := range m.Listings {
		if strings.Contains(listing.Title, query) ||
			strings.Contains(listing.Body, query) ||
			strings.Contains(listing.Key, query) {
			if skip > 0 {
				skip -= 1
			} else {
				results = append(results, listing)
				if len(results) >= limit {
					break
				}
			}
		}
	}

	return results
}
