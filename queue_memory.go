// +build !appengine

package fluyt

import (
	"time"
)

// Sync pulls the latest commits into the backing store.
func Sync(r *http.Request, m *Marketplace) {}

// Add persists the given listing to the backend.
func Add(r *http.Request, m *Marketplace, l *Listing) error {
	if l.Key == "" {
		l.Key = randStringBytes(10)
	}
	l.LastUpdate = time.Now()
	m.added(*l)
	return nil
}
