package cmd

import (
	"os"
	"path"
	"testing"

	"github.com/jptosso/coraza-center/client"
	"github.com/jptosso/coraza-center/test"
	"github.com/stretchr/testify/assert"
)

func TestDownload(t *testing.T) {
	tmp2, err := os.MkdirTemp(os.TempDir(), "coraza-center-tests")
	assert.NoError(t, err)
	defer os.RemoveAll(tmp2)
	test.TestServer(t)
	remote = client.NewRemote("http://127.0.0.1:2022", "admin", "admin")
	assert.NoError(t, remote.Download("test", tmp2))
	bts, err := os.ReadFile(path.Join(tmp2, "ok"))
	assert.NoError(t, err)
	assert.Equal(t, "OK\n", string(bts))
}
