package xsql

import (
	"database/sql"
	"time"
)

func NewNull[T any](v T) (n sql.Null[T]) {
	_ = n.Scan(v)
	return n
}

func NewNullString(v string) (n sql.NullString) {
	_ = n.Scan(v)
	return n
}

func NewNilNullString() (n sql.NullString) {
	return sql.NullString{}
}

func NewNullInt64(v int64) (n sql.NullInt64) {
	_ = n.Scan(v)
	return n
}

func NewNilNullInt64() (n sql.NullInt64) {
	return sql.NullInt64{}
}

func NewNullTime(v time.Time) (n sql.NullTime) {
	_ = n.Scan(v)
	return n
}

func NewNilNullTime() (n sql.NullTime) {
	return sql.NullTime{}
}
