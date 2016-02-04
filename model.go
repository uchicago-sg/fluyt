package fluyt

import (
	"strings"
	"time"
)

type Listing struct {
	Key        string    `json:"key"`
	Title      string    `json:"title"`
	Body       string    `json:"body"`
	LastUpdate time.Time `json:"lastUpdate"`
}

type Marketplace struct {
	Listings   map[string]Listing
	LastUpdate time.Time
}

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

func (m *Marketplace) Lookup(path string) *Listing {
	result, ok := m.Listings[path]
	if !ok {
		return nil
	} else {
		return &result
	}
}

func (m *Marketplace) Search(query string, skip, limit int) []Listing {
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
