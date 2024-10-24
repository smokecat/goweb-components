package fw

import (
	"context"
	"database/sql"

	"github.com/zeromicro/go-zero/core/stores/sqlx"

	"github.com/smokecat/goweb-components/pkg/xcode"
	"github.com/smokecat/goweb-components/pkg/xerr"
	"github.com/smokecat/goweb-components/pkg/xutil"
)

const PostgresTxConnKey = "GoZeroPostgresTxConnKey"

var (
	defaultSqlConn sqlx.SqlConn
)

func InitSqlConn(cfg DatabaseConf) error {
	db, err := sql.Open("postgres", cfg.Dsn)
	if err != nil {
		return err
	}
	defaultSqlConn = sqlx.NewSqlConnFromDB(db)

	return PingDb(context.Background())
}

func PingDb(ctx context.Context) error {
	var r int
	if err := defaultSqlConn.QueryRowCtx(ctx, &r, "SELECT 1"); err != nil {
		return err
	}
	return nil
}

func CtxWithTx(ctx context.Context, sqlConn sqlx.SqlConn) context.Context {
	return context.WithValue(ctx, PostgresTxConnKey, sqlConn)
}

// TxFromCtx returns the transaction session from the context.
// If the session is not found, it returns the default connection.
// Return exists as false if the session is not found.
func TxFromCtx(ctx context.Context) (sqlConn sqlx.SqlConn, exists bool) {
	sqlConnVal, ok := xutil.CtxValue[sqlx.SqlConn](ctx, PostgresTxConnKey)

	if ok {
		return sqlConnVal, true
	}

	return defaultSqlConn, false
}

func RemoveTxFromCtx(ctx context.Context) context.Context {
	return context.WithValue(ctx, PostgresTxConnKey, nil)
}

func WithTx(ctx context.Context, fn func(context.Context, sqlx.SqlConn) error) error {
	sqlConn, ok := TxFromCtx(ctx)
	if ok {
		return fn(ctx, sqlConn)
	}

	return defaultSqlConn.TransactCtx(ctx, func(ctx context.Context, session sqlx.Session) error {
		sqlConn = sqlx.NewSqlConnFromSession(session)
		ctxWithTx := CtxWithTx(ctx, sqlConn)
		return fn(ctxWithTx, sqlConn)
	})
}

func WithNewTx(ctx context.Context, fn func(context.Context, sqlx.SqlConn) error) error {
	return defaultSqlConn.TransactCtx(ctx, func(ctx context.Context, session sqlx.Session) error {
		sqlConn := sqlx.NewSqlConnFromSession(session)
		ctxWithTx := CtxWithTx(ctx, sqlConn)
		return fn(ctxWithTx, sqlConn)
	})
}

func DbErr(err error, msg ...string) error {
	if len(msg) > 0 {
		return xerr.Wrap(err, xcode.CodePrivate, msg[0])
	}
	return xerr.Wrap(err, xcode.CodePrivate, "Internal db error")
}
