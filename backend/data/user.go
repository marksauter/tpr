// This file is automatically generated.
package data

import (
	"strings"

	"github.com/jackc/pgx"
)

type User struct {
	ID             Int32
	Name           String
	PasswordDigest Bytes
	PasswordSalt   Bytes
	Email          String
}

func CountUser(db Queryer) (int64, error) {
	var n int64
	sql := `select count(*) from "users"`
	err := db.QueryRow(sql).Scan(&n)
	return n, err
}

func SelectAllUser(db Queryer) ([]User, error) {
	sql := `select
  "id",
  "name",
  "password_digest",
  "password_salt",
  "email"
from "users"`

	var rows []User

	dbRows, err := db.Query(sql)
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

func SelectUserByPK(
	db Queryer,
	id int32,
) (*User, error) {
	sql := `select
  "id",
  "name",
  "password_digest",
  "password_salt",
  "email"
from "users"
where "id"=$1`

	var row User
	err := db.QueryRow(sql, id).Scan(
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

	row.ID.addInsert(`id`, &columns, &values, &args)
	row.Name.addInsert(`name`, &columns, &values, &args)
	row.PasswordDigest.addInsert(`password_digest`, &columns, &values, &args)
	row.PasswordSalt.addInsert(`password_salt`, &columns, &values, &args)
	row.Email.addInsert(`email`, &columns, &values, &args)

	sql := `insert into "users"(` + strings.Join(columns, ", ") + `)
values(` + strings.Join(values, ",") + `)
returning "id"
  `

	return db.QueryRow(sql, args...).Scan(&row.ID)
}

func UpdateUser(db Queryer,
	id int32,
	row *User,
) error {
	sets := make([]string, 0, 5)
	args := pgx.QueryArgs(make([]interface{}, 0, 5))

	row.ID.addUpdate(`id`, &sets, &args)
	row.Name.addUpdate(`name`, &sets, &args)
	row.PasswordDigest.addUpdate(`password_digest`, &sets, &args)
	row.PasswordSalt.addUpdate(`password_salt`, &sets, &args)
	row.Email.addUpdate(`email`, &sets, &args)

	if len(sets) == 0 {
		return nil
	}

	sql := `update "users" set ` + strings.Join(sets, ", ") + ` where ` + `"id"=` + args.Append(id)

	commandTag, err := db.Exec(sql, args...)
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

	commandTag, err := db.Exec(sql, args...)
	if err != nil {
		return err
	}
	if commandTag.RowsAffected() != 1 {
		return pgx.ErrNoRows
	}
	return nil
}
