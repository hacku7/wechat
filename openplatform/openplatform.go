package openplatform

import (
	"net/http"

	"github.com/hacku7/wechat/officialaccount/server"
	"github.com/hacku7/wechat/openplatform/account"
	"github.com/hacku7/wechat/openplatform/config"
	"github.com/hacku7/wechat/openplatform/context"
	"github.com/hacku7/wechat/openplatform/miniprogram"
	"github.com/hacku7/wechat/openplatform/officialaccount"
)

// OpenPlatform 微信开放平台相关api
type OpenPlatform struct {
	*context.Context
}

// NewOpenPlatform new openplatform
func NewOpenPlatform(cfg *config.Config) *OpenPlatform {
	if cfg.Cache == nil {
		panic("cache 未设置")
	}
	ctx := &context.Context{
		Config: cfg,
	}
	return &OpenPlatform{ctx}
}

// GetServer get server
func (openPlatform *OpenPlatform) GetServer(req *http.Request, writer http.ResponseWriter) *server.Server {
	off := officialaccount.NewOfficialAccount(openPlatform.Context, "")
	return off.GetServer(req, writer)
}

// GetOfficialAccount 公众号代处理
func (openPlatform *OpenPlatform) GetOfficialAccount(appID string) *officialaccount.OfficialAccount {
	return officialaccount.NewOfficialAccount(openPlatform.Context, appID)
}

// GetMiniProgram 小程序代理
func (openPlatform *OpenPlatform) GetMiniProgram(appID string) *miniprogram.MiniProgram {
	return miniprogram.NewMiniProgram(openPlatform.Context, appID)
}

// GetAccountManager 账号管理
// TODO
func (openPlatform *OpenPlatform) GetAccountManager() *account.Account {
	return account.NewAccount(openPlatform.Context)
}
