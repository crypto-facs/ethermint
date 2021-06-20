package rpc

import (
	"github.com/cosmos/ethermint/ethereum/rpc/types"
	log "github.com/xlab/suplog"
)

// PublicTxPoolAPI offers and API for the transaction pool. It only operates on data that is non-confidential.
type PublicTxPoolAPI struct {
	logger  log.Logger
	backend Backend
}

// NewPublicTxPoolAPI creates a new tx pool service that gives information about the transaction pool.
func NewPublicTxPoolAPI(backend Backend) *PublicTxPoolAPI {
	return &PublicTxPoolAPI{
		logger:  log.WithField("module", "txpool"),
		backend: backend,
	}
}

// Content returns the transactions contained within the transaction pool
// NOTE: For more info about the current status of this endpoint https://github.com/tharsis/ethermint/issues/124
func (api *PublicTxPoolAPI) Content() (map[string]map[string]map[string]*types.RPCTransaction, error) {
	api.logger.Debug("txpool_content")
	return api.backend.TxPoolContent()
}
