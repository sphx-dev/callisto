package database

import (
	"fmt"

	"github.com/forbole/callisto/v4/types"
)

func (db *Db) SaveCreateMarginAccountEvent(event types.MarginAccountEvent) error {
	query := `
		INSERT INTO margin_account_event (
			margin_acc_addr, owner, acc_number
		) VALUES (
			$1, $2, $3
		);
	`
	_, err := db.SQL.Exec(
		query,
		event.MarginAccAddr,
		event.Owner,
		event.AccNumber,
	)

	if err != nil {
		return fmt.Errorf("failed to save margin account event: %w", err)
	}
	return nil
}

func (db *Db) SaveWithdrawEvent(event types.WithdrawEvent) error {
	query := `
		INSERT INTO withdraw_events (
			margin_acc_addr, recipient, withdraw_coin
		) VALUES (
			$1, $2, $3
		)
		ON CONFLICT (margin_acc_addr) DO UPDATE SET
			recipient = EXCLUDED.recipient,
			withdraw_coin = EXCLUDED.withdraw_coin;
	`
	_, err := db.SQL.Exec(
		query,
		event.MarginAccAddr,
		event.Recipient,
		event.WithdrawCoin,
	)

	if err != nil {
		return fmt.Errorf("failed to save withdraw event: %w", err)
	}
	return nil
}
