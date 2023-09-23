package repository

import (
	"context"
	"database/sql"

	sq "github.com/Masterminds/squirrel"

	"github.com/betawulan/synapsis/model"
)

type shoppingCartRepo struct {
	db *sql.DB
}

func (s shoppingCartRepo) Create(ctx context.Context, shoppingCart model.ShoppingCart) (model.ShoppingCart, error) {

	query, args, err := sq.Insert("shopping_cart").
		Columns("user_id",
			"product_category_id").
		Values(shoppingCart.UserID,
			shoppingCart.ProductCategoryID).
		ToSql()
	if err != nil {
		return model.ShoppingCart{}, err
	}

	res, err := s.db.ExecContext(ctx, query, args...)
	if err != nil {
		return model.ShoppingCart{}, err
	}

	shoppingCart.ID, err = res.LastInsertId()
	if err != nil {
		return model.ShoppingCart{}, err
	}

	return shoppingCart, nil
}

func (s shoppingCartRepo) Delete(ctx context.Context, ID int64) error {
	query, args, err := sq.Delete("shopping_cart").
		Where(sq.Eq{"id": ID}).
		ToSql()
	if err != nil {
		return err
	}

	_, err = s.db.ExecContext(ctx, query, args...)
	if err != nil {
		return err
	}

	return nil
}

func (s shoppingCartRepo) Read(ctx context.Context, userID int64) ([]model.ShoppingCart, error) {
	query, args, err := sq.Select("sc.id",
		"sc.user_id",
		"sc.product_category_id",
		"p.name as product_name",
		"p.price",
		"c.name as category_name").
		From("shopping_cart sc").
		Join("product_category pc on sc.product_category_id  = pc.id").
		Join("product p on pc.product_id = p.id").
		Join("category c on c.id =pc.category_id").
		Where(sq.Eq{"sc.user_id": userID}).
		ToSql()
	if err != nil {
		return nil, err
	}

	rows, err := s.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	shoppingCarts := make([]model.ShoppingCart, 0)
	for rows.Next() {
		var shoppingCart model.ShoppingCart

		err = rows.Scan(&shoppingCart.ID,
			&shoppingCart.UserID,
			&shoppingCart.ProductCategoryID,
			&shoppingCart.Product.Name,
			&shoppingCart.Product.Price,
			&shoppingCart.Category.Name)
		if err != nil {
			return nil, err
		}

		shoppingCarts = append(shoppingCarts, shoppingCart)
	}

	return shoppingCarts, nil
}

func NewShoppingCartRepository(db *sql.DB) ShoppingCartRepository {
	return shoppingCartRepo{
		db: db,
	}
}
