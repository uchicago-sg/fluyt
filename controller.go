package fluyt

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

const RootPage = `
<!HTML head>
<html>
<head>
	<title>Marketplace</title>
	<link rel="stylesheet" href="https://maxcdn.bootstrapcdn.com/bootstrap/3.3.6/css/bootstrap.min.css"/>
	<link rel="stylesheet" href="/assets/styles.css"/>	
	<script type="text/javascript" src="/assets/main.js"></script>
</head>
<body class="container">
	Marketplace requires enabling JavaScript.
</body>
</html>`

type Results struct {
	Count    int       `json:"count"`
	Listings []Listing `json:"listings"`
}

func writeJSON(w http.ResponseWriter, x interface{}) {
	json.NewEncoder(w).Encode(x)
}

func init() {
	http.Handle("/", NewMarketplace())
}

func (m *Marketplace) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get("Accept") != "text/javascript" {
		w.Header().Set("Content-type", "text/html")
		fmt.Fprintf(w, RootPage)
		return
	}

	Sync(r, m)

	if len(r.URL.Path) > 1 {
		listing := m.Lookup(r.URL.Path[1:])
		if listing == nil {
			http.NotFound(w, r)
			return
		}

		switch r.Method {
		case "GET":
			writeJSON(w, listing)

		case "POST":
			m.update(listing, w, r)
		case "DELETE":
			m.delete(listing, w, r)
		default:
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		}
	} else {
		switch r.Method {
		case "GET":
			query := r.FormValue("q")
			skip, _ := strconv.Atoi(r.FormValue("skip"))
			limit, _ := strconv.Atoi(r.FormValue("limit"))
			if limit == 0 || limit > 1000 {
				limit = 1000
			}

			listings := m.Search(query, skip, limit)

			writeJSON(w, &Results{Listings: listings, Count: len(listings)})

		case "POST":
			for i := 0; i < 1000; i++ {
				listing := &Listing{
					Title: "Hello world",
					Body:  "Body Goes Here",
				}
				Add(r, m, listing)
				if i == 999 {
					http.Redirect(w, r, "/"+listing.Key, http.StatusTemporaryRedirect)
				}
			}

		default:
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		}
	}
}

func (m *Marketplace) update(l *Listing, w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello!")
}

func (m *Marketplace) delete(l *Listing, w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello!")
}
