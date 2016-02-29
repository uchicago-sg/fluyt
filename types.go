package fluyt

import (
	"strings"
	"time"
)

type Person struct {
	Email   string `json:"email"`
	Roaming string `json:"roaming"` // "" if signed in via Shibboleth
}

type Photo struct {
	Small string `json:"small"`
	Large string `json:"large"`
}

type Listing struct {
	Permalink  string    `json:"key"`
	Title      string    `json:"title"`
	Body       string    `json:"body"`
	Seller     *Person   `json:"seller,omitempty"`
	Price      float32   `json:"price"`
	Categories []string  `json:"categories"`
	Approved   bool      `json:"approved"`
	Sold       bool      `json:"sold"`
	LastUpdate time.Time `json:"lastUpdate"`
	Photos     []Photo   `json:"photos"`
}

func (l *Listing) Key() string {
	return l.Permalink
}

func (l *Listing) Match(query string) bool {
	return strings.Contains(strings.ToLower(l.Title), query)
}
