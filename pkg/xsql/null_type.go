package xsql

import (
	"database/sql"
	"time"
)

func NullTimeString(v sql.NullTime) string {
	if v.Valid {
		return v.Time.Format(time.RFC3339)
	}
	return ""
}
