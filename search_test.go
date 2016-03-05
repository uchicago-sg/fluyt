package fluyt_test

import (
	"github.com/uchicago-sg/fluyt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSearchListings(t *testing.T) {
	ts := httptest.NewServer(fluyt.MarketplaceHandler())
	defer ts.Close()

	res, err := http.Get(ts.URL + "/listings")
	if err != nil {
		t.Fatal(err)
	}

	result, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		t.Fatal(err)
	}
	if string(result) != "{\"listings\":[]}\n" {
		t.Fatal(string(result))
	}
}
