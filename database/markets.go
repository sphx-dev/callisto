package database

import (
	"fmt"

	"github.com/forbole/callisto/v4/types"
)

func (db *Db) SaveRegisterMarketEvent(event types.MarketEvent) error {
	query := `
		INSERT INTO market_event (
			market_id, market_ticker, market_base_asset, market_quote_asset, market_status
		) VALUES (
			$1, $2, $3, $4, $5
		);
	`

	_, err := db.SQL.Exec(
		query,
		event.MarketId,
		event.MarketTicker,
		event.MarketBaseAsset,
		event.MarketQuoteAsset,
		event.MarketStatus,
	)

	if err != nil {
		return fmt.Errorf("failed to save market event: %w", err)
	}
	return nil
}
