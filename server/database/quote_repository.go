package database

import (
	"context"
	"fmt"
	"time"

	"github.com/Ryan-18-system/desafio-client-server-goexpert/server/model"
	"github.com/Ryan-18-system/desafio-client-server-goexpert/server/util"
)

func InsertQuote(quote model.QuoteResponse) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Millisecond)
	defer cancel()
	if DB == nil {
		return fmt.Errorf("banco de dados não inicializado")
	}
	tx, err := DB.BeginTx(ctx, nil)
	if err != nil {
		return util.HandlerError(err)
	}
	// própria doc do go indica - https://go.dev/doc/database/execute-transactions
	defer tx.Rollback()
	stmt, err := tx.PrepareContext(ctx, `INSERT INTO quote
			(code, codein, name, high, low, varBid, pctChange, bid, ask, "timestamp", create_date)
			VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?);
	`)
	if err != nil {
		return util.HandlerError(err)
	}
	defer stmt.Close()
	_, err = stmt.ExecContext(ctx, quote.USDBRL.Code, quote.USDBRL.Codein, quote.USDBRL.Name, quote.USDBRL.High, quote.USDBRL.Low,
		quote.USDBRL.VarBid, quote.USDBRL.PctChange, quote.USDBRL.Bid, quote.USDBRL.Ask, quote.USDBRL.Timestamp, quote.USDBRL.CreateDate)
	if err != nil {
		return util.HandlerError(err)
	}
	err = tx.Commit()
	if err != nil {
		return util.HandlerError(err)
	}
	return nil
}
