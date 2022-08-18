package cmd

import (
	"os"
	"path"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDirectoryInitialization(t *testing.T) {
	tmp, err := os.MkdirTemp(os.TempDir(), "coraza-center-tests")
	assert.NoError(t, err)
	defer os.RemoveAll(tmp)
	err = initDirectory("sample", tmp)
	assert.NoError(t, err)
	p := path.Join(tmp, ".coraza", "TAG")
	bts, err := os.ReadFile(p)
	assert.NoError(t, err)
	assert.Equal(t, "sample", string(bts))
	loadProjectDir(tmp)
	assert.Equal(t, "sample", localConfig.WafTag)
}
