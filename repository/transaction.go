package repository

import (
	"context"
	"database/sql"
	"fmt"

	sq "github.com/Masterminds/squirrel"
)

type transactionRepo struct {
	db *sql.DB
}

func (t transactionRepo) SumPrice(ctx context.Context, userID int64, productCategoryIDs []int) (int, error) {
	query, args, err := sq.Select("SUM(p.price)").
		From("shopping_cart sc").
		Join("product_category pc ON pc.id = sc.product_category_id").
		Join("product p ON p.id = pc.product_id").
		Where(sq.Eq{"sc.product_category_id": productCategoryIDs}).
		Where(sq.Eq{"sc.user_id": userID}).
		GroupBy("sc.user_id").
		ToSql()
	if err != nil {
		return 0, err
	}

	rows := t.db.QueryRowContext(ctx, query, args...)
	if err != nil {
		return 0, err
	}

	var totalPrice int
	err = rows.Scan(&totalPrice)
	if err != nil {
		return 0, err
	}

	return totalPrice, nil

}

func (t transactionRepo) CreateTransaction(ctx context.Context, userID int64, productCategoryIDs []int, sumPrice int) error {

	tx, err := t.db.Begin()
	if err != nil {
		return err
	}

	query, args, err := sq.Insert("transaction").
		Columns("user_id",
			"sum_price").
		Values(userID,
			sumPrice).
		ToSql()
	if err != nil {
		return err
	}

	var errRollback error
	res, err := tx.ExecContext(ctx, query, args...)
	if err != nil {
		errRollback = tx.Rollback()
		if errRollback != nil {
			fmt.Println("error:", errRollback)
		}

		return err
	}

	transactionID, err := res.LastInsertId()
	if err != nil {
		errRollback = tx.Rollback()
		if errRollback != nil {
			fmt.Println("error:", errRollback)
		}

		return err
	}

	queryInsert := sq.Insert("transaction_detail").
		Columns("transaction_id",
			"product_category_id")

	for _, productCategoryID := range productCategoryIDs {
		queryInsert = queryInsert.Values(transactionID, productCategoryID)
	}

	query, args, err = queryInsert.ToSql()
	if err != nil {
		return err
	}

	_, err = tx.ExecContext(ctx, query, args...)
	if err != nil {
		errRollback = tx.Rollback()
		if errRollback != nil {
			fmt.Println("error:", errRollback)
		}

		return err
	}

	query, args, err = sq.Delete("shopping_cart").
		Where(sq.Eq{"product_category_id": productCategoryIDs}).
		Where(sq.Eq{"user_id": userID}).
		ToSql()
	if err != nil {
		return err
	}

	_, err = tx.ExecContext(ctx, query, args...)
	if err != nil {
		errRollback = tx.Rollback()
		if errRollback != nil {
			fmt.Println("error:", errRollback)
		}

		return err
	}

	err = tx.Commit()
	if err != nil {
		errRollback = tx.Rollback()
		if errRollback != nil {
			fmt.Println("error:", errRollback)
		}

		return err
	}

	return nil
}

func NewTransactionRepository(db *sql.DB) TransactionRepository {
	return transactionRepo{
		db: db,
	}
}
