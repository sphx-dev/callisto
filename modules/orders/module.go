package orders

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/forbole/callisto/v4/database"
	"github.com/forbole/juno/v6/logging"
	"github.com/forbole/juno/v6/modules"
)

var (
	_ modules.Module            = &Module{}
	_ modules.TransactionModule = &Module{}
)

// Module represents the x/feemodel module
type Module struct {
	cdc    codec.Codec
	db     *database.Db
	logger logging.Logger
}

// NewModule returns a new Module instance
func NewModule(
	cdc codec.Codec, db *database.Db,
	logger logging.Logger,
) *Module {
	return &Module{
		cdc:    cdc,
		db:     db,
		logger: logger,
	}
}

// Name implements modules.Module
func (m *Module) Name() string {
	return "orders"
}
