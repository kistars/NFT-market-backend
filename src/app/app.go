package app

import (
	"context"

	"github.com/kistars/NFT-market-backend/src/config"
	"github.com/kistars/NFT-market-backend/src/service/svc"
	"github.com/gin-gonic/gin"
	"github.com/kistars/NFT-market-base/logger/xzap"
	"go.uber.org/zap"
)

type Platform struct {
	config    *config.Config
	router    *gin.Engine
	serverCtx *svc.ServerCtx
}

func NewPlatform(config *config.Config, router *gin.Engine, serverCtx *svc.ServerCtx) (*Platform, error) {
	return &Platform{
		config:    config,
		router:    router,
		serverCtx: serverCtx,
	}, nil
}

func (p *Platform) Start() {
	xzap.WithContext(context.Background()).Info("EasySwap-End run", zap.String("port", p.config.Api.Port))
	if err := p.router.Run(p.config.Api.Port); err != nil {
		panic(err)
	}
}
