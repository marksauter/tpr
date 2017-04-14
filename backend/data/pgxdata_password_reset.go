package data

// This file is automatically generated by pgxdata.

import (
	"strings"

	"github.com/jackc/pgx"
	"github.com/jackc/pgx/pgtype"
)

type PasswordReset struct {
	Token          pgtype.Varchar
	Email          pgtype.Varchar
	RequestIP      pgtype.Inet
	RequestTime    pgtype.Timestamptz
	UserID         pgtype.Int4
	CompletionIP   pgtype.Inet
	CompletionTime pgtype.Timestamptz
}

const countPasswordResetSQL = `select count(*) from "password_resets"`

func CountPasswordReset(db Queryer) (int64, error) {
	var n int64
	err := prepareQueryRow(db, "pgxdataCountPasswordReset", countPasswordResetSQL).Scan(&n)
	return n, err
}

const SelectAllPasswordResetSQL = `select
  "token",
  "email",
  "request_ip",
  "request_time",
  "user_id",
  "completion_ip",
  "completion_time"
from "password_resets"`

func SelectAllPasswordReset(db Queryer) ([]PasswordReset, error) {
	var rows []PasswordReset

	dbRows, err := prepareQuery(db, "pgxdataSelectAllPasswordReset", SelectAllPasswordResetSQL)
	if err != nil {
		return nil, err
	}

	for dbRows.Next() {
		var row PasswordReset
		dbRows.Scan(
			&row.Token,
			&row.Email,
			&row.RequestIP,
			&row.RequestTime,
			&row.UserID,
			&row.CompletionIP,
			&row.CompletionTime,
		)
		rows = append(rows, row)
	}

	if dbRows.Err() != nil {
		return nil, dbRows.Err()
	}

	return rows, nil
}

const selectPasswordResetByPKSQL = `select
  "token",
  "email",
  "request_ip",
  "request_time",
  "user_id",
  "completion_ip",
  "completion_time"
from "password_resets"
where "token"=$1`

func SelectPasswordResetByPK(
	db Queryer,
	token string,
) (*PasswordReset, error) {
	var row PasswordReset
	err := prepareQueryRow(db, "pgxdataSelectPasswordResetByPK", selectPasswordResetByPKSQL, token).Scan(
		&row.Token,
		&row.Email,
		&row.RequestIP,
		&row.RequestTime,
		&row.UserID,
		&row.CompletionIP,
		&row.CompletionTime,
	)
	if err == pgx.ErrNoRows {
		return nil, ErrNotFound
	} else if err != nil {
		return nil, err
	}

	return &row, nil
}

func InsertPasswordReset(db Queryer, row *PasswordReset) error {
	args := pgx.QueryArgs(make([]interface{}, 0, 7))

	var columns, values []string

	if row.Token.Status != pgtype.Undefined {
		columns = append(columns, `token`)
		values = append(values, args.Append(&row.Token))
	}
	if row.Email.Status != pgtype.Undefined {
		columns = append(columns, `email`)
		values = append(values, args.Append(&row.Email))
	}
	if row.RequestIP.Status != pgtype.Undefined {
		columns = append(columns, `request_ip`)
		values = append(values, args.Append(&row.RequestIP))
	}
	if row.RequestTime.Status != pgtype.Undefined {
		columns = append(columns, `request_time`)
		values = append(values, args.Append(&row.RequestTime))
	}
	if row.UserID.Status != pgtype.Undefined {
		columns = append(columns, `user_id`)
		values = append(values, args.Append(&row.UserID))
	}
	if row.CompletionIP.Status != pgtype.Undefined {
		columns = append(columns, `completion_ip`)
		values = append(values, args.Append(&row.CompletionIP))
	}
	if row.CompletionTime.Status != pgtype.Undefined {
		columns = append(columns, `completion_time`)
		values = append(values, args.Append(&row.CompletionTime))
	}

	sql := `insert into "password_resets"(` + strings.Join(columns, ", ") + `)
values(` + strings.Join(values, ",") + `)
returning "token"
  `

	psName := preparedName("pgxdataInsertPasswordReset", sql)

	return prepareQueryRow(db, psName, sql, args...).Scan(&row.Token)
}

func UpdatePasswordReset(db Queryer,
	token string,
	row *PasswordReset,
) error {
	sets := make([]string, 0, 7)
	args := pgx.QueryArgs(make([]interface{}, 0, 7))

	if row.Token.Status != pgtype.Undefined {
		sets = append(sets, `token`+"="+args.Append(&row.Token))
	}
	if row.Email.Status != pgtype.Undefined {
		sets = append(sets, `email`+"="+args.Append(&row.Email))
	}
	if row.RequestIP.Status != pgtype.Undefined {
		sets = append(sets, `request_ip`+"="+args.Append(&row.RequestIP))
	}
	if row.RequestTime.Status != pgtype.Undefined {
		sets = append(sets, `request_time`+"="+args.Append(&row.RequestTime))
	}
	if row.UserID.Status != pgtype.Undefined {
		sets = append(sets, `user_id`+"="+args.Append(&row.UserID))
	}
	if row.CompletionIP.Status != pgtype.Undefined {
		sets = append(sets, `completion_ip`+"="+args.Append(&row.CompletionIP))
	}
	if row.CompletionTime.Status != pgtype.Undefined {
		sets = append(sets, `completion_time`+"="+args.Append(&row.CompletionTime))
	}

	if len(sets) == 0 {
		return nil
	}

	sql := `update "password_resets" set ` + strings.Join(sets, ", ") + ` where ` + `"token"=` + args.Append(token)

	psName := preparedName("pgxdataUpdatePasswordReset", sql)

	commandTag, err := prepareExec(db, psName, sql, args...)
	if err != nil {
		return err
	}
	if commandTag.RowsAffected() != 1 {
		return ErrNotFound
	}
	return nil
}

func DeletePasswordReset(db Queryer,
	token string,
) error {
	args := pgx.QueryArgs(make([]interface{}, 0, 1))

	sql := `delete from "password_resets" where ` + `"token"=` + args.Append(token)

	commandTag, err := prepareExec(db, "pgxdataDeletePasswordReset", sql, args...)
	if err != nil {
		return err
	}
	if commandTag.RowsAffected() != 1 {
		return ErrNotFound
	}
	return nil
}
