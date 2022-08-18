package server

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	"github.com/gorilla/mux"
	"github.com/jptosso/coraza-center/database"
	"github.com/jptosso/coraza-center/utils"
)

func postWafHandler(w http.ResponseWriter, r *http.Request) {
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
	if err := r.ParseMultipartForm(10 << 20); err != nil {
		httpError(w, err)
		return
	}
	file, header, err := r.FormFile("file")
	filename := header.Filename
	if !strings.HasSuffix(filename, ".tar.gz") {
		httpError(w, fmt.Errorf("filename must be type .tar.gz"))
		return
	}
	if err != nil {
		httpError(w, err)
		return
	}
	defer file.Close()
	bts, err := io.ReadAll(file)
	if err != nil {
		httpError(w, err)
		return
	}
	if err := validateWafFile(bts); err != nil {
		httpError(w, err)
		return
	}
	waf.Data = bts
	tx = database.DB.Save(&waf)
	if tx.Error != nil {
		httpError(w, tx.Error)
		return
	}
	if tx.RowsAffected == 0 {
		httpError(w, fmt.Errorf("Failed to update waf"))
		return
	}
	w.WriteHeader(http.StatusOK)
	fmt.Printf("[%s] Successfully updated WAF %s\n", r.RemoteAddr, waf.ID)
}

// validateWafFile untars the .tar.gz file from []byte and validates it
// To test this we create a temporary file and uncompress all files there
func validateWafFile(data []byte) error {
	tmpdir, err := ioutil.TempDir(os.TempDir(), "coraza-center")
	if err != nil {
		return err
	}
	defer os.RemoveAll(tmpdir)
	return utils.UntarTo(data, tmpdir)
}
