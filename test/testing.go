package test

import (
	"encoding/base64"
	"testing"
	"time"

	"github.com/jptosso/coraza-center/database"
	"github.com/jptosso/coraza-center/server"
	"github.com/stretchr/testify/assert"
)

func TestServer(t *testing.T) {
	var err error
	if err := database.Connect(":memory:"); err != nil {
		t.Fatal(err)
	}
	go func() {
		opts := server.ServerOptions{
			Bind: "127.0.0.1:2022",
		}
		server.Start(opts)
	}()
	time.Sleep(2 * time.Second)
	tx := database.DB.Create(&database.User{
		ID:       "1",
		UserName: "admin",
		Password: "admin",
		Admin:    true,
	})
	assert.NoError(t, tx.Error)
	assert.Equal(t, int(tx.RowsAffected), 1)
	// this is just a basic .tar.gz file with one directive and the word "OK"
	payload := "H4sIALQR+2IAA8vPZqA5MDAwMDMxUQDR5mamYBoIYDQIGCsYmhiZm5kZGpiZmisYGBoZmRkwKBjQ3mkMDKXFJYlFQKdkFZTkFxfn41RXnpGamoPHHFRPKVDdnTQC/t5cA+2EUTCAoCS1uEQvOT8vjYZ2EMz/xibw/G9qagjM/8amJuaj+Z8eIDg12TG5JDM/T0EpM8XK0MBAJ7c43Uo9P1tdabRgGAWjYBSMgmEMACS75uYADAAA"
	// base64decode payload
	base64decoded, err := base64.StdEncoding.DecodeString(payload)
	assert.NoError(t, err)
	tx = database.DB.Create(&database.Waf{
		ID:   "1",
		Tag:  "test",
		Data: base64decoded,
	})
	assert.NoError(t, tx.Error)
	assert.Equal(t, int(tx.RowsAffected), 1)
}
