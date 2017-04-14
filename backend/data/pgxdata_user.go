package data

// This file is automatically generated by pgxdata.

import (
	"strings"

	"github.com/jackc/pgx"
	"github.com/jackc/pgx/pgtype"
)

type User struct {
	ID             pgtype.Int4
	Name           pgtype.Varchar
	PasswordDigest pgtype.Bytea
	PasswordSalt   pgtype.Bytea
	Email          pgtype.Varchar
}

const countUserSQL = `select count(*) from "users"`

func CountUser(db Queryer) (int64, error) {
	var n int64
	err := prepareQueryRow(db, "pgxdataCountUser", countUserSQL).Scan(&n)
	return n, err
}

const SelectAllUserSQL = `select
  "id",
  "name",
  "password_digest",
  "password_salt",
  "email"
from "users"`

func SelectAllUser(db Queryer) ([]User, error) {
	var rows []User

	dbRows, err := prepareQuery(db, "pgxdataSelectAllUser", SelectAllUserSQL)
	if err != nil {
		return nil, err
	}

	for dbRows.Next() {
		var row User
		dbRows.Scan(
			&row.ID,
			&row.Name,
			&row.PasswordDigest,
			&row.PasswordSalt,
			&row.Email,
		)
		rows = append(rows, row)
	}

	if dbRows.Err() != nil {
		return nil, dbRows.Err()
	}

	return rows, nil
}

const selectUserByPKSQL = `select
  "id",
  "name",
  "password_digest",
  "password_salt",
  "email"
from "users"
where "id"=$1`

func SelectUserByPK(
	db Queryer,
	id int32,
) (*User, error) {
	var row User
	err := prepareQueryRow(db, "pgxdataSelectUserByPK", selectUserByPKSQL, id).Scan(
		&row.ID,
		&row.Name,
		&row.PasswordDigest,
		&row.PasswordSalt,
		&row.Email,
	)
	if err == pgx.ErrNoRows {
		return nil, ErrNotFound
	} else if err != nil {
		return nil, err
	}

	return &row, nil
}

func InsertUser(db Queryer, row *User) error {
	args := pgx.QueryArgs(make([]interface{}, 0, 5))

	var columns, values []string

	if row.ID.Status != pgtype.Undefined {
		columns = append(columns, `id`)
		values = append(values, args.Append(&row.ID))
	}
	if row.Name.Status != pgtype.Undefined {
		columns = append(columns, `name`)
		values = append(values, args.Append(&row.Name))
	}
	if row.PasswordDigest.Status != pgtype.Undefined {
		columns = append(columns, `password_digest`)
		values = append(values, args.Append(&row.PasswordDigest))
	}
	if row.PasswordSalt.Status != pgtype.Undefined {
		columns = append(columns, `password_salt`)
		values = append(values, args.Append(&row.PasswordSalt))
	}
	if row.Email.Status != pgtype.Undefined {
		columns = append(columns, `email`)
		values = append(values, args.Append(&row.Email))
	}

	sql := `insert into "users"(` + strings.Join(columns, ", ") + `)
values(` + strings.Join(values, ",") + `)
returning "id"
  `

	psName := preparedName("pgxdataInsertUser", sql)

	return prepareQueryRow(db, psName, sql, args...).Scan(&row.ID)
}

func UpdateUser(db Queryer,
	id int32,
	row *User,
) error {
	sets := make([]string, 0, 5)
	args := pgx.QueryArgs(make([]interface{}, 0, 5))

	if row.ID.Status != pgtype.Undefined {
		sets = append(sets, `id`+"="+args.Append(&row.ID))
	}
	if row.Name.Status != pgtype.Undefined {
		sets = append(sets, `name`+"="+args.Append(&row.Name))
	}
	if row.PasswordDigest.Status != pgtype.Undefined {
		sets = append(sets, `password_digest`+"="+args.Append(&row.PasswordDigest))
	}
	if row.PasswordSalt.Status != pgtype.Undefined {
		sets = append(sets, `password_salt`+"="+args.Append(&row.PasswordSalt))
	}
	if row.Email.Status != pgtype.Undefined {
		sets = append(sets, `email`+"="+args.Append(&row.Email))
	}

	if len(sets) == 0 {
		return nil
	}

	sql := `update "users" set ` + strings.Join(sets, ", ") + ` where ` + `"id"=` + args.Append(id)

	psName := preparedName("pgxdataUpdateUser", sql)

	commandTag, err := prepareExec(db, psName, sql, args...)
	if err != nil {
		return err
	}
	if commandTag.RowsAffected() != 1 {
		return ErrNotFound
	}
	return nil
}

func DeleteUser(db Queryer,
	id int32,
) error {
	args := pgx.QueryArgs(make([]interface{}, 0, 1))

	sql := `delete from "users" where ` + `"id"=` + args.Append(id)

	commandTag, err := prepareExec(db, "pgxdataDeleteUser", sql, args...)
	if err != nil {
		return err
	}
	if commandTag.RowsAffected() != 1 {
		return ErrNotFound
	}
	return nil
}
