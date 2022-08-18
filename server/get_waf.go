package server

import (
	"bytes"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/jptosso/coraza-center/database"
)

func getWafHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	wafid := vars["waf_tag"]
	var waf *database.Waf
	tx := database.DB.Model(&database.Waf{}).First(&waf, "tag = ?", wafid)
	if tx.Error != nil {
		httpError(w, tx.Error)
		return
	}
	if waf == nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	lastTimestamp := r.Header.Get("If-Modified-Since")
	if lastTimestamp != "" {
		t, err := time.Parse(time.ANSIC, lastTimestamp)
		if err != nil {
			httpError(w, err)
			return
		}
		if !waf.UpdatedAt.After(t) {
			w.WriteHeader(http.StatusNotModified)
			return
		}
	}
	// now we return a file
	// w.Header().Add("Content-Type", "application/x-gzip")
	// w.Header().Add("Content-Disposition", fmt.Sprintf("attachment; filename=coraza_%s_%d.tar.gz", waf.Tag, waf.UpdatedAt.Unix()))
	// w.Header().Add("Content-Length", fmt.Sprintf("%d", len(waf.Data)))
	// w.WriteHeader(http.StatusOK)
	http.ServeContent(w, r, "", time.Time{}, bytes.NewReader(waf.Data))
}
