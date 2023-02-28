package repositories

import (
	"context"
	"english_exam_go/utils/app_logger"
	"english_exam_go/utils/transaction"
	"fmt"
	"gorm.io/gorm"
)

type TxCtxKey string

const txKey TxCtxKey = "TRANSACTION_CONTEXT_ATTRIBUTE_KEY"

type txImpl struct {
}

func (t *txImpl) Required(ctx context.Context,
	txFunc func(ctx context.Context) (interface{},
		error)) (data interface{},
	err error) {
	fmt.Printf("Begin transaction")
	tx := GetConn().Begin()
	if tx.Error != nil {
		return nil, &RdbRuntimeError{
			ErrMsg:        "[infrastructure.Required] failed to begin TransactionImpl",
			OriginalError: tx.Error,
		}
	}
	app_logger.Logger.Debug("TransactionImpl Begin")

	ctx = context.WithValue(ctx, txKey, tx)

	defer func() {
		if p := recover(); p != nil {
			app_logger.Logger.Debug("TransactionImpl Rollback")
			tx.Rollback()
			panic(p)
		} else if err != nil {
			app_logger.Logger.Debug("TransactionImpl Rollback")
			tx.Rollback()
		} else {
			app_logger.Logger.Debug("TransactionImpl Commit")
			if commitErr := tx.Commit().Error; commitErr != nil {
				err = &RdbRuntimeError{
					ErrMsg:        "[infrastructure.Required] failed to commit TransactionImpl",
					OriginalError: commitErr,
				}
			}
		}
	}()

	data, err = txFunc(ctx)
	return
}

var shareTx = &txImpl{}

func TransactionImpl() transaction.Transaction {
	return shareTx
}

func GetTx(ctx context.Context) (*gorm.DB, bool) {
	tx, ok := ctx.Value(txKey).(*gorm.DB)
	return tx, ok
}
