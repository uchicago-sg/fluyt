package fluyt

import (
	"github.com/husobee/vestigo"
	"net/http"
)

func Serial(m *Marketplace, f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		m.Lock()
		defer m.Unlock()

		m.Sync(makeContext(r))

		f(w, r)
	}
}

func MarketplaceHandler() http.Handler {
	m := NewMarketplace()

	router := vestigo.NewRouter()
	router.SetGlobalCors(&vestigo.CorsAccessControl{
		AllowOrigin: []string{"*"},
	})

	router.Get("/listings", Serial(m, m.SearchListings))
	router.Post("/listings", Serial(m, m.CreateListing))
	router.Get("/listings/:listing", Serial(m, m.GetListing))
	router.Post("/listings/:listing", Serial(m, m.PatchListing))
	router.Patch("/listings/:listing", Serial(m, m.PatchListing))
	router.Get("/inquiries", Serial(m, m.SearchInquiries))
	router.Post("/listings/:listing/inquiries", Serial(m, m.SendInquiry))
	router.Handle("/", http.FileServer(http.Dir("static")))

	return router
}
