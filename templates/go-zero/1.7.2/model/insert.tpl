func (m *default{{.upperStartCamelObject}}Model) Insert(ctx context.Context, data *{{.upperStartCamelObject}}) (sql.Result,error) {
	{{if .withCache}}{{.keys}}
    ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values ({{.expression}})", m.table, {{.lowerStartCamelObject}}RowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, {{.expressionValues}})
	}, {{.keyValues}}){{else}}query := fmt.Sprintf("insert into %s (%s) values ({{.expression}})", m.table, {{.lowerStartCamelObject}}RowsExpectAutoSet)
    ret,err:=m.conn.ExecCtx(ctx, query, {{.expressionValues}}){{end}}
	return ret,err
}


func (m *custom{{.upperStartCamelObject}}Model) InsertOne(ctx context.Context, data *{{.upperStartCamelObject}}) (int64,error) {
	var lastInsertId int64
	{{if .withCache}}{{.keys}}
    ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values ({{.expression}}) returning id", m.table, {{.lowerStartCamelObject}}RowsExpectAutoSet)
		err := conn.QueryRowCtx(ctx, lastInsertId, query, {{.expressionValues}})
		return lastInsertId, err
	}, {{.keyValues}}){{else}}query := fmt.Sprintf("insert into %s (%s) values ({{.expression}}) returning id", m.table, {{.lowerStartCamelObject}}RowsExpectAutoSet)
	err := m.conn.QueryRowCtx(ctx, &lastInsertId, query, {{.expressionValues}})
	switch {
	case err == nil:
		return lastInsertId, nil
	default:
		return 0, err
	}{{end}}
}

