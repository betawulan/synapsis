package repository

import (
	"context"
	"database/sql"

	sq "github.com/Masterminds/squirrel"

	"github.com/betawulan/synapsis/model"
)

type authRepo struct {
	db *sql.DB
}

func (a authRepo) Register(ctx context.Context, user model.User) error {

	query, args, err := sq.Insert("user").
		Columns("name",
			"role",
			"email",
			"password").
		Values(user.Name,
			user.Role,
			user.Email,
			user.Password).
		ToSql()
	if err != nil {
		return err
	}

	res, err := a.db.ExecContext(ctx, query, args...)
	if err != nil {
		return err
	}

	user.ID, err = res.LastInsertId()
	if err != nil {
		return err
	}

	return nil
}

func (a authRepo) Login(ctx context.Context, role string, email string, password string) (model.User, error) {
	query, args, err := sq.Select("id",
		"name",
		"role",
		"email").
		From("user").
		Where(sq.Eq{"role": role}).
		Where(sq.Eq{"email": email}).
		Where(sq.Eq{"password": password}).
		ToSql()
	if err != nil {
		return model.User{}, nil
	}

	row := a.db.QueryRowContext(ctx, query, args...)
	var user model.User
	err = row.Scan(&user.ID,
		&user.Name,
		&user.Role,
		&user.Email)
	if err != nil {
		if err == sql.ErrNoRows {
			return model.User{}, err
		}

		return model.User{}, nil
	}

	return user, nil
}

func NewAuthRepository(db *sql.DB) AuthRepository {
	return authRepo{db: db}
}
