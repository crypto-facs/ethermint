package rpc

import log "github.com/xlab/suplog"

type PublicTxPoolApi struct {
	logger log.Logger
}

func NewTxPoolApi() *PublicTxPoolApi {
	return &PublicTxPoolApi{
		logger: log.WithField("module", "txpool"),
	}
}

func (api *PublicTxPoolApi) Content() {
	api.logger.Debug("txpool_content")
	return
}
