// +build !appengine

package fluyt

import (
	"time"
)

func Sync(r *http.Request, m *Marketplace) {}

func Add(r *http.Request, m *Marketplace, l *Listing) error {
	if l.Key == "" {
		l.Key = RandStringBytes(10)
	}
	l.LastUpdate = time.Now()
	m.added(*l)
	return nil
}
