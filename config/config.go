package config

import (
	"time"

	"github.com/gogf/gf/v2/container/garray"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
)

var (
	PORT           = 8080
	PlusModels     = garray.NewStrArrayFrom([]string{"claude", "claude-3-opus-20240229", "claude-3-haiku-20240307", "claude-3-5-sonnet-20240620", "claude-default", "claude-3-5-sonnet-20241022"})
	ForbiddenWords = []string{}    // 禁止词
	LIMIT          = 40            // 限制次数
	PER            = time.Hour * 3 // 限制时间
	OAIKEY         = ""            // OAIKEY
	OAIKEYLOG      = ""            // OAIKEYLOG 隐藏
	MODERATION     = ""
)

func init() {
	ctx := gctx.GetInitCtx()
	port := g.Cfg().MustGetWithEnv(ctx, "PORT").Int()
	if port > 0 {
		PORT = port
	}
	g.Log().Info(ctx, "PORT:", PORT)
	limit := g.Cfg().MustGetWithEnv(ctx, "LIMIT").Int()
	if limit > 0 {
		LIMIT = limit
	}
	g.Log().Info(ctx, "LIMIT:", LIMIT)
	per := g.Cfg().MustGetWithEnv(ctx, "PER").Duration()
	if per > 0 {
		PER = per
	}
	g.Log().Info(ctx, "PER:", PER)
	oaikey := g.Cfg().MustGetWithEnv(ctx, "OAIKEY").String()
	// oaikey 不为空
	if oaikey != "" {
		OAIKEY = oaikey
		// 日志隐藏 oaikey，有 * 代表有值
		OAIKEYLOG = "******"
	}
	g.Log().Info(ctx, "OAIKEY:", OAIKEYLOG)
	moderation := g.Cfg().MustGetWithEnv(ctx, "MODERATION").String()
	if moderation != "" {
		MODERATION = moderation
	}
	g.Log().Info(ctx, "MODERATION:", MODERATION)
}
