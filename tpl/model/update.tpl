func (m *default{{.upperStartCamelObject}}Model) Update(ctx context.Context, {{if .containsIndexCache}}newData{{else}}data{{end}} *{{.upperStartCamelObject}}) error {
{{if .withCache}}{{if .containsIndexCache}}data, err:=m.FindOne(ctx, newData.{{.upperStartCamelPrimaryKey}})
if err!=nil{
return err
}

{{end}}	{{.keys}}
_, {{if .containsIndexCache}}err{{else}}err:{{end}}= m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
query := fmt.Sprintf("update %s set %s where {{.originalPrimaryKey}} = {{if .postgreSql}}$1{{else}}?{{end}}", m.table, {{.lowerStartCamelObject}}RowsWithPlaceHolder)
return conn.ExecCtx(ctx, query, {{.expressionValues}})
}, {{.keyValues}}){{else}}query := fmt.Sprintf("update %s set %s where {{.originalPrimaryKey}} = {{if .postgreSql}}$1{{else}}?{{end}}", m.table, {{.lowerStartCamelObject}}RowsWithPlaceHolder)
_,err:=m.conn.ExecCtx(ctx, query, {{.expressionValues}}){{end}}
return err
}

func (m *default{{.upperStartCamelObject}}Model) FindCount(ctx context.Context, builder squirrel.SelectBuilder, field string) (int64,error) {

    builder = builder.Columns("COUNT(*)")

    query, values, err := builder.Where("deleted_at is null").ToSql()
    if err != nil {
        return 0, err
    }

    var resp int64
    {{if .withCache}}err = m.QueryRowNoCacheCtx(ctx,&resp, query, values...){{else}}
    err = m.conn.QueryRowCtx(ctx,&resp, query, values...)
    {{end}}
    switch err {
        case nil:
            return resp, nil
        default:
            return 0, err
    }
}

func (m *default{{.upperStartCamelObject}}Model) FindAll(ctx context.Context,builder squirrel.SelectBuilder,orderBy string) ([]*{{.upperStartCamelObject}},error) {

    builder = builder.Columns({{.lowerStartCamelObject}}Rows)

    if orderBy == ""{
        builder = builder.OrderBy("id DESC")
    }else{
        builder = builder.OrderBy(orderBy)
    }

    query, values, err := builder.Where("deleted_at is null").ToSql()
    if err != nil {
        return nil, err
    }

    var resp []*{{.upperStartCamelObject}}
    {{if .withCache}}err = m.QueryRowsNoCacheCtx(ctx,&resp, query, values...){{else}}
    err = m.conn.QueryRowsCtx(ctx,&resp, query, values...)
    {{end}}
    switch err {
        case nil:
            return resp, nil
        default:
            return nil, err
    }
}

func (m *default{{.upperStartCamelObject}}Model) FindPageListByPage(ctx context.Context,builder squirrel.SelectBuilder,page ,pageSize int64,orderBy string) ([]*{{.upperStartCamelObject}},error) {

    builder = builder.Columns({{.lowerStartCamelObject}}Rows)

    if orderBy == ""{
        builder = builder.OrderBy("id DESC")
    }else{
        builder = builder.OrderBy(orderBy)
    }

    if page < 1{
        page = 1
    }
    offset := (page - 1) * pageSize

    query, values, err := builder.Where("deleted_at is null").Offset(uint64(offset)).Limit(uint64(pageSize)).ToSql()
    if err != nil {
        return nil, err
    }

    var resp []*{{.upperStartCamelObject}}
    {{if .withCache}}err = m.QueryRowsNoCacheCtx(ctx,&resp, query, values...){{else}}
    err = m.conn.QueryRowsCtx(ctx,&resp, query, values...)
    {{end}}
    switch err {
        case nil:
            return resp, nil
        default:
            return nil, err
    }
}

func(m *default{{.upperStartCamelObject}}Model)  SelectBuilder() squirrel.SelectBuilder {
    return squirrel.Select().From(m.table)
}