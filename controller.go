package fluyt

import (
	"github.com/husobee/vestigo"
	"google.golang.org/appengine"
	"net/http"
)

func Serial(m *Marketplace, f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		m.Lock()
		defer m.Unlock()

		m.Backend.SetContext(appengine.NewContext(r))
		m.Sync()

		f(w, r)
	}
}

func init() {
	m := NewMarketplace()

	router := vestigo.NewRouter()
	router.Get("/listings", Serial(m, m.SearchListings))
	router.Post("/listings", Serial(m, m.CreateListing))
	router.Get("/listings/:listing", Serial(m, m.GetListing))
	router.Post("/listings/:listing", Serial(m, m.PatchListing))
	router.Patch("/listings/:listing", Serial(m, m.PatchListing))
	router.Get("/inquiries", Serial(m, m.SearchInquiries))
	router.Post("/listings/:listing/inquiries", Serial(m, m.SendInquiry))

	http.Handle("/", router)
}
