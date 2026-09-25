package main

import (
	"fmt"
	"net/http"

	"golang.org/x/text/language"

	"github.com/shubam-disseqt/zreview-e2e-matrix/internal/auth"
	"github.com/shubam-disseqt/zreview-e2e-matrix/internal/files"
	"github.com/shubam-disseqt/zreview-e2e-matrix/internal/httpclient"
	"github.com/shubam-disseqt/zreview-e2e-matrix/internal/store"
	"github.com/shubam-disseqt/zreview-e2e-matrix/internal/util"
	"github.com/shubam-disseqt/zreview-e2e-matrix/internal/worker"
)

var reports = &store.Report{}

// loginHandler checks a password against the stored hash.
func loginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		if r.FormValue("user") != "" {
			if auth.Verify(r.FormValue("pw"), r.FormValue("hash")) {
				fmt.Fprintln(w, "ok")
				return
			}
		}
	}
	http.Error(w, "denied", http.StatusUnauthorized)
}

// reportHandler returns the order total for ?customer=.
func reportHandler(w http.ResponseWriter, r *http.Request) error {
	total, err := reports.TotalFor(r.URL.Query().Get("customer"))
	if err != nil {
		return err
	}
	fmt.Fprintln(w, total)
	return nil
}

// localeHandler echoes the best-match language for the request.
func localeHandler(w http.ResponseWriter, r *http.Request) {
	tags, _, _ := language.ParseAcceptLanguage(r.Header.Get("Accept-Language"))
	if len(tags) == 0 {
		tags = []language.Tag{language.English}
	}
	fmt.Fprintln(w, tags[0])
}

// mirrorHandler fetches ?url= values through the shared client.
func mirrorHandler(w http.ResponseWriter, r *http.Request) {
	bodies := httpclient.FetchAll(httpclient.New(), r.URL.Query()["url"])
	n := util.PageSize(r.URL.Query().Get("n"))
	for _, b := range util.LastN(toStrings(bodies), n) {
		fmt.Fprintln(w, len(b))
	}
}

func toStrings(bs [][]byte) []string {
	out := make([]string, len(bs))
	for i, b := range bs {
		out[i] = string(b)
	}
	return out
}

var _ = files.Serve
var _ = worker.FirstResult
