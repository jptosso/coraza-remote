package corazacenter

import (
	"os"
	"testing"
	"time"

	"github.com/corazawaf/coraza/v3"
	"github.com/corazawaf/coraza/v3/seclang"
	"github.com/jptosso/coraza-center/test"
	"github.com/stretchr/testify/assert"
)

func TestSecRemote(t *testing.T) {
	tmp2, err := os.MkdirTemp(os.TempDir(), "coraza-center-tests")
	assert.NoError(t, err)
	defer os.RemoveAll(tmp2)
	test.TestServer(t)
	waf := coraza.NewWaf()
	parser, _ := seclang.NewParser(waf)
	if err := parser.FromString(`
		SecAction "id:15"
		SecRemoteCredentials admin admin
		SecRemoteCacheDir /tmp
		SecWebAppId test
		SecRemote http://127.0.0.1:2022
	`); err != nil {
		t.Fatal(err)
	}
	time.Sleep(2 * time.Second)
	// we expect 1 rule from here and another rule from the server
	assert.Equal(t, 2, waf.Rules.Count())
}
