package postgresql

import (
	"database/sql"
	"go/internal/config"

	"github.com/davecgh/go-spew/spew"
	"github.com/labstack/echo/v4"
)

type IPostgresqlRepository interface {
	FindById(id int64) PostgresqlEntity
	FindAll() []PostgresqlEntity
	Create(p PostgresqlEntity, ctx echo.Context) (sql.Result, *sql.Tx)
	Update(id int64, p PostgresqlEntity, ctx echo.Context) (sql.Result, *sql.Tx)
	Delete(id int64, ctx echo.Context) (sql.Result, *sql.Tx)
}

type PostgresqlRepository struct {
	shared *config.GlobalShared
}

func NewPostgresqlRepository(s *config.GlobalShared) IPostgresqlRepository {
	return PostgresqlRepository{
		shared: s,
	}
}

func (r PostgresqlRepository) FindAll() []PostgresqlEntity {

	var res []PostgresqlEntity

	spew.Dump(r.shared.Psqlconn)
	rows, err := r.shared.Psqlconn.Query("SELECT name, address, avatar FROM users")
	if err != nil {
		panic(err.Error())
	}

	defer rows.Close()
	for rows.Next() {
		var temp PostgresqlEntity
		err := rows.Scan(&temp.Name, &temp.Address, &temp.Avatar)
		if err != nil {
			panic(err.Error())
		}
		res = append(res, temp)
	}

	return res
}

func (r PostgresqlRepository) FindById(id int64) PostgresqlEntity {
	var res PostgresqlEntity

	row := r.shared.Psqlconn.QueryRow(`SELECT name, address, avatar FROM users WHERE id=$1`, id)

	err := row.Scan(&res.Name, &res.Address, &res.Avatar)
	if err != nil {
		panic(err.Error())
	}

	return res
}

func (r PostgresqlRepository) Create(p PostgresqlEntity, ctx echo.Context) (sql.Result, *sql.Tx) {
	tx, err := r.shared.Psqlconn.BeginTx(ctx.Request().Context(), nil)
	if err != nil {
		panic(err.Error())
	}

	rows, err := tx.ExecContext(ctx.Request().Context(), `INSERT INTO users(name, address, avatar) VALUES($1,$2,$3)`, p.Name, p.Address, p.Avatar)
	if err != nil {
		tx.Rollback()
		panic(err.Error())
	}

	return rows, tx

}

func (r PostgresqlRepository) Update(id int64, p PostgresqlEntity, ctx echo.Context) (sql.Result, *sql.Tx) {
	tx, err := r.shared.Psqlconn.BeginTx(ctx.Request().Context(), nil)
	if err != nil {
		panic(err.Error())
	}

	_ = r.FindById(id)

	rows, err := tx.ExecContext(ctx.Request().Context(), `UPDATE users SET name=$1, address=$2, avatar=$3 WHERE id=$4`, p.Name, p.Address, p.Avatar, id)
	if err != nil {
		panic(err.Error())
	}

	return rows, tx

}

func (r PostgresqlRepository) Delete(id int64, ctx echo.Context) (sql.Result, *sql.Tx) {
	tx, err := r.shared.Psqlconn.BeginTx(ctx.Request().Context(), nil)
	if err != nil {
		panic(err.Error())
	}

	_ = r.FindById(id)

	rows, err := tx.ExecContext(ctx.Request().Context(), `DELETE FROM users WHERE id=$1`, id)
	if err != nil {
		panic(err.Error())
	}

	return rows, tx
}
