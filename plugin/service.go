package corazacenter

import (
	"fmt"
	"path"
	"time"

	"github.com/corazawaf/coraza/v3"
	"github.com/corazawaf/coraza/v3/seclang"
	"github.com/jptosso/coraza-center/client"
)

type remoteService struct {
	waf      *coraza.Waf
	oldRules []*coraza.Rule
	interval int
	remote   *client.Remote
	cacheDir string
	tag      string
}

func (rs *remoteService) Init() error {
	if rs.interval <= 0 {
		rs.interval = 500
		rs.waf.Logger.Error("Remote service interval is not set, using default value of 500 seconds")
	}
	// first we save the previously created rules
	rs.oldRules = rs.waf.Rules.GetRules()
	if err := rs.PullConfig(); err != nil {
		rs.waf.Logger.Error("Failed to load remote rules, loading from cache: %s", err)
		// we get config from cache
		if err := rs.loadFromDir(path.Join(rs.cacheDir, rs.waf.WebAppID)); err != nil {
			return err
		}
	}
	ticker := time.NewTicker(time.Duration(rs.interval) * time.Second)
	quit := make(chan struct{})
	go func() {
		for {
			select {
			case <-ticker.C:
				if err := rs.PullConfig(); err != nil {
					rs.waf.Logger.Error(err.Error())
				}
			case <-quit:
				ticker.Stop()
				return
			}
		}
	}()
	return nil
}

func (rs *remoteService) PullConfig() error {
	p := path.Join(rs.cacheDir, rs.tag)
	if err := rs.remote.Download(rs.tag, p); err != nil {
		return err
	}
	return rs.loadFromDir(p)
}

func (rs *remoteService) loadFromDir(dir string) error {
	waf := coraza.NewWaf()
	for _, pr := range rs.oldRules {
		if err := waf.Rules.Add(pr); err != nil {
			return fmt.Errorf("Failed to parse remote config: %s", err)
		}
	}
	parser, _ := seclang.NewParser(waf)
	if err := parser.FromFile(path.Join(dir, "*.conf")); err != nil {
		return fmt.Errorf("Failed to parse remote config: %s", err)
	}
	*rs.waf = *waf
	return nil
}
