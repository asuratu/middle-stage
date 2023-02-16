func (m *default{{.upperStartCamelObject}}Model) Update(ctx context.Context, session sqlx.Session, {{if .containsIndexCache}}newData{{else}}data{{end}} *{{.upperStartCamelObject}}) error {
	{{if .withCache}}
	{{if .containsIndexCache}}
	data, err:=m.FindOne(ctx, newData.{{.upperStartCamelPrimaryKey}})
	if err != nil{
		return err
	}
    {{end}}
    {{.keys}}
    _, {{if .containsIndexCache}}err{{else}}err:{{end}}= m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where {{.originalPrimaryKey}} = {{if .postgreSql}}$1{{else}}?{{end}}", m.table, {{.lowerStartCamelObject}}RowsWithPlaceHolder)
		if session != nil{
            return session.ExecCtx(ctx,query, {{.expressionValues}})
        } else {
    		return conn.ExecCtx(ctx, query, {{.expressionValues}})
        }
	}, {{.keyValues}})
	{{else}}
	query := fmt.Sprintf("update %s set %s where {{.originalPrimaryKey}} = {{if .postgreSql}}$1{{else}}?{{end}}", m.table, {{.lowerStartCamelObject}}RowsWithPlaceHolder)
	var err error
	if session != nil{
        _, err = session.ExecCtx(ctx,query, {{.expressionValues}})
    } else {
        _, err = m.conn.ExecCtx(ctx, query, {{.expressionValues}})
    }
    {{end}}
	return err
}

func (m *default{{.upperStartCamelObject}}Model) UpdateWithVersion(ctx context.Context, session sqlx.Session, {{if .containsIndexCache}}newData{{else}}data{{end}} *{{.upperStartCamelObject}}) error {
	oldVersion := {{if .containsIndexCache}}newData{{else}}data{{end}}.Version
	{{if .containsIndexCache}}newData{{else}}data{{end}}.Version += 1

	var sqlResult sql.Result
	var err error

	{{if .withCache}}
	{{if .containsIndexCache}}
	data, err:=m.FindOne(ctx, newData.{{.upperStartCamelPrimaryKey}})
    if err!=nil{
        return err
    }
    {{end}}
    {{.keys}}
	sqlResult,err =  m.ExecCtx(ctx,func(ctx context.Context,conn sqlx.SqlConn) (result sql.Result, err error) {
	query := fmt.Sprintf("update %s set %s where {{.originalPrimaryKey}} = {{if .postgreSql}}$1{{else}}?{{end}} and version = ? ", m.table, {{.lowerStartCamelObject}}RowsWithPlaceHolder)
	if session != nil{
        return session.ExecCtx(ctx,query, {{.expressionValues}}, oldVersion)
    } else {
    	return conn.ExecCtx(ctx,query, {{.expressionValues}}, oldVersion)
    }
	},
	{{.keyValues}})
	{{else}}
	query := fmt.Sprintf("update %s set %s where {{.originalPrimaryKey}} = {{if .postgreSql}}$1{{else}}?{{end}} and version = ? ", m.table, {{.lowerStartCamelObject}}RowsWithPlaceHolder)
	if session != nil{
        sqlResult,err  =  session.ExecCtx(ctx,query, {{.expressionValues}}, oldVersion)
    }else{
        sqlResult,err  =  m.conn.ExecCtx(ctx,query, {{.expressionValues}}, oldVersion)
    }
	{{end}}
	if err != nil {
		return err
	}
	updateCount , err := sqlResult.RowsAffected()
	if err != nil{
		return err
	}
	if updateCount == 0 {
		return  xerr.NewErrCode(xerr.DbUpdateAffectedZeroError)
	}

	return nil
}






