package fluyt

import (
	"encoding/json"
	"net/http"
)

func (m *Marketplace) CreateListing(w http.ResponseWriter, r *http.Request) {
	listing := Listing{}
	err := json.NewDecoder(r.Body).Decode(&listing)
	if err != nil {
		http.Error(w, err.Error(), 200)
	}

	if err = m.Backend.Append(listing); err != nil {
		http.Error(w, err.Error(), 200)
	}
}

func (m *Marketplace) GetListing(w http.ResponseWriter, r *http.Request) {
	listing, ok := m.Listings[r.FormValue(":listing")]
	if !ok {
		http.NotFound(w, r)
		return
	}

	json.NewEncoder(w).Encode(&listing)
}

func (m *Marketplace) PatchListing(w http.ResponseWriter, r *http.Request) {
	listing, ok := m.Listings[r.FormValue(":listing")]
	if !ok {
		http.NotFound(w, r)
		return
	}

	before, after := listing, listing
	if err := json.NewDecoder(r.Body).Decode(&after); err != nil {
		http.Error(w, err.Error(), 500)
	}

	if before.Seller != after.Seller {
		http.Error(w, "not allowed to change seller", 401)
	}

	if err := m.Backend.Append(after); err != nil {
		http.Error(w, err.Error(), 500)
	}

	json.NewEncoder(w).Encode(&listing)
}

func (m *Marketplace) SearchInquiries(w http.ResponseWriter, r *http.Request) {}
func (m *Marketplace) SendInquiry(w http.ResponseWriter, r *http.Request)     {}
