package context

import (
	"github.com/hacku7/wechat/credential"
	"github.com/hacku7/wechat/work/config"
)

// Context struct
type Context struct {
	*config.Config
	credential.AccessTokenHandle
}
