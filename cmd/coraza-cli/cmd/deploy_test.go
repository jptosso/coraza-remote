package cmd

import (
	"os"
	"path"
	"testing"

	"github.com/jptosso/coraza-center/client"
	"github.com/jptosso/coraza-center/test"
	"github.com/stretchr/testify/assert"
)

func TestDeploy(t *testing.T) {
	tmp2, err := os.MkdirTemp(os.TempDir(), "coraza-center-tests")
	assert.NoError(t, err)
	defer os.RemoveAll(tmp2)
	test.TestServer(t)
	remote = client.NewRemote("http://127.0.0.1:2022", "admin", "admin")
	assert.NoError(t, os.WriteFile(path.Join(tmp2, "test.data"), []byte("test"), 0644))
	assert.NoError(t, remote.Upload(tmp2, "test"))
}
