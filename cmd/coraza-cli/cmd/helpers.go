package cmd

import (
	"fmt"
	"os"
	"path"
	"strings"

	"github.com/jptosso/coraza-center/client"
)

var localConfig *config
var userDir string

type config struct {
	WafTag string
	Prefix string
}

func loadProjectDir(dir string) {
	projectDir := path.Join(dir, ".coraza")
	if _, err := os.Stat(projectDir); !os.IsNotExist(err) {
		bts, err := os.ReadFile(path.Join(projectDir, "TAG"))
		if err != nil {
			panic(err)
		}
		localConfig = &config{
			WafTag: strings.TrimSpace(string(bts)),
			Prefix: projectDir,
		}
	}
}

func validateCredentials() error {
	bts, err := os.ReadFile(path.Join(userDir, "credentials"))
	if err != nil {
		return err
	}
	data := strings.TrimSpace(string(bts))
	if !strings.ContainsRune(data, ':') {
		return fmt.Errorf("Invalid credentials file")
	}
	return nil
}

func login(server string, username string, password string) error {
	if err := os.WriteFile(path.Join(prefixPath, "credentials"), []byte(fmt.Sprintf("%s:%s", username, password)), 0600); err != nil {
		return err
	}
	if err := os.WriteFile(path.Join(prefixPath, "server"), []byte(server), 0655); err != nil {
		return err
	}
	return nil
}

func init() {
	var err error
	userDir, err = os.UserHomeDir()
	if err != nil {
		panic(err)
	}
	userDir = path.Join(userDir, ".coraza")
	bts, err := os.ReadFile(path.Join(userDir, "credentials"))
	if err != nil {
		return
	}
	data := strings.TrimSpace(string(bts))
	spl := strings.Split(data, ":")
	if len(spl) != 2 {
		panic("Invalid credentials file")
	}
	username, password := spl[0], spl[1]
	server, err := os.ReadFile(path.Join(userDir, "server"))
	if err != nil {
		panic("No server configured")
	}
	remote = client.NewRemote(strings.TrimSpace(string(server)), strings.TrimSpace(username), strings.TrimSpace(password))
}
