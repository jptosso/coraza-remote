package server

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/jptosso/coraza-center/database"
)

type wafModel struct {
	ID        string `json:"id"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	Tag       string `json:"tag"`
	Size      int64  `json:"size"`
}

func getAllWafHandler(w http.ResponseWriter, r *http.Request) {
	var waf []database.Waf
	tx := database.DB.Table("wafs").Select("id, created_at, updated_at, tag").Scan(&waf)
	if tx.Error != nil {
		httpError(w, tx.Error)
		return
	}
	// now we return a file
	w.Header().Add("Content-Type", "application/json")
	wl := make([]wafModel, len(waf))
	for i, v := range waf {
		wl[i] = wafModel{
			ID:        v.ID,
			CreatedAt: v.CreatedAt.Format(time.RFC3339),
			UpdatedAt: v.UpdatedAt.Format(time.RFC3339),
			Tag:       v.Tag,
			Size:      0,
		}
	}

	jsdata, err := json.Marshal(wl)
	if err != nil {
		httpError(w, err)
		return
	}
	w.Header().Add("Content-Length", fmt.Sprintf("%d", len(jsdata)))
	w.WriteHeader(http.StatusOK)
	if _, err := w.Write(jsdata); err != nil {
		httpError(w, err)
	}
}
