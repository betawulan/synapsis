package repository

import (
	"context"
	"database/sql"

	sq "github.com/Masterminds/squirrel"

	"github.com/betawulan/synapsis/model"
)

type productRepo struct {
	db *sql.DB
}

func (p productRepo) Fetch(ctx context.Context, filter model.ProductCategoryFilter) ([]model.ProductCategory, error) {
	qSelect := sq.Select("product_category.id",
		"product_category.product_id",
		"product_category.category_id",
		"product.name",
		"product.price",
		"category.name").
		From("product").
		Join("product_category ON product.id=product_category.product_id").
		Join("category ON category.id=product_category.category_id")
	if filter.Category != "" {
		qSelect = qSelect.Where(sq.Eq{"category.name": filter.Category})
	}

	query, args, err := qSelect.ToSql()
	if err != nil {
		return nil, err
	}

	rows, err := p.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	productCategories := make([]model.ProductCategory, 0)
	for rows.Next() {
		var productCategory model.ProductCategory

		err = rows.Scan(
			&productCategory.ID,
			&productCategory.ProductID,
			&productCategory.CategoryID,
			&productCategory.Product.Name,
			&productCategory.Product.Price,
			&productCategory.Category.Name)
		if err != nil {
			return nil, err
		}

		productCategories = append(productCategories, productCategory)
	}

	return productCategories, nil
}

func NewProductRepository(db *sql.DB) ProductRepository {
	return productRepo{
		db: db,
	}
}
