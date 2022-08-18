package corazacenter

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/corazawaf/coraza/v3/seclang"
	"github.com/jptosso/coraza-center/client"
)

const (
	remoteConfigUsername = "remote_username"
	remoteConfigPassword = "remote_password"
	remoteConfigInterval = "remote_interval"
	remoteConfigCacheDir = "remote_cache_dir"
)

func secRemoteCredentialsDirective(options *seclang.DirectiveOptions) error {
	spl := strings.SplitN(options.Arguments, " ", 2)
	if len(spl) != 2 {
		return fmt.Errorf("Use: SecRemoteCredentials <username> <password>")
	}
	username, password := spl[0], spl[1]
	options.Config.Set(remoteConfigUsername, username)
	options.Config.Set(remoteConfigPassword, password)
	return nil
}

func secRemoteRefreshIntervalDirective(options *seclang.DirectiveOptions) error {
	interval, err := strconv.Atoi(options.Arguments)
	if err != nil {
		return fmt.Errorf("Use: SecRemoteRefreshInterval <interval>")
	}
	if interval <= 0 {
		return fmt.Errorf("Interval must be greater than 0")
	}
	options.Config.Set(remoteConfigInterval, interval)
	return nil
}

func secRemoteCacheDirDirective(options *seclang.DirectiveOptions) error {
	// first we validate if directory exists
	cacheDir := options.Arguments
	if _, err := os.Stat(cacheDir); os.IsNotExist(err) {
		return fmt.Errorf("Cache directory %s does not exist", cacheDir)
	}
	options.Config.Set(remoteConfigCacheDir, options.Arguments)
	return nil
}

func secRemoteDirective(options *seclang.DirectiveOptions) error {
	username := options.Config.Get(remoteConfigUsername, "").(string)
	password := options.Config.Get(remoteConfigPassword, "").(string)
	interval := options.Config.Get(remoteConfigInterval, 600).(int)
	cacheDir := options.Config.Get(remoteConfigCacheDir, "").(string)
	if username == "" || password == "" {
		return fmt.Errorf("Remote configuration requires SecRemoteCredentials <username> <password>")
	}
	if cacheDir == "" {
		return fmt.Errorf("Remote configuration requires SecRemoteCacheDir <cache_dir>")
	}
	server := options.Arguments
	if server == "" {
		return fmt.Errorf("Remote configuration requires <server>")
	}
	appid := options.Waf.WebAppID
	if appid == "" {
		return fmt.Errorf("Remote configuration requires \"SecWebAppID <appid>\"")
	}
	rs := &remoteService{
		waf:      options.Waf,
		oldRules: options.Waf.Rules.GetRules(),
		interval: interval,
		remote:   client.NewRemote(server, username, password),
		cacheDir: cacheDir,
		tag:      appid,
	}
	return rs.Init()
}

var (
	_ seclang.Directive = secRemoteDirective
	_ seclang.Directive = secRemoteCredentialsDirective
	_ seclang.Directive = secRemoteRefreshIntervalDirective
	_ seclang.Directive = secRemoteCacheDirDirective
)

func init() {
	seclang.RegisterDirective("SecRemote", secRemoteDirective)
	seclang.RegisterDirective("SecRemoteCredentials", secRemoteCredentialsDirective)
	seclang.RegisterDirective("SecRemoteRefreshInterval", secRemoteRefreshIntervalDirective)
	seclang.RegisterDirective("SecRemoteCacheDir", secRemoteCacheDirDirective)
}
