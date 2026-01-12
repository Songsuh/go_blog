package dao

import (
	"fmt"
	"go_blog/internal/svc"
	"gorm.io/gorm"
)

// QueryBuilder 通用查询构建器
type QueryBuilder struct {
	dbAlias string
	model   interface{}
}

// NewQueryBuilder 创建一个新的查询构建器
func NewQueryBuilder(dbAlias string, model interface{}) *QueryBuilder {
	return &QueryBuilder{
		dbAlias: dbAlias,
		model:   model,
	}
}

// QueryOptions 查询选项
type QueryOptions struct {
	Conditions Conditions //查询条件
	OrderBy    string     //排序字段
	Limit      int        //每页记录数
	Offset     int        //偏移量
	Select     []string   //指定字段
}

// QueryResult 查询结果
type QueryResult struct {
	Data  interface{}
	Total int64
	Error error
}

// getDB 获取数据库连接
func (qb *QueryBuilder) getDB(fields ...string) *gorm.DB {
	svcCtx := svc.GetSvc()
	if svcCtx == nil {
		return nil
	}
	return svcCtx.GetDb(qb.dbAlias)
}

// applyConditions 应用查询条件
func (qb *QueryBuilder) applyConditions(query *gorm.DB, conditions Conditions) *gorm.DB {
	if len(conditions) == 0 {
		return query
	}

	expressions := conditions.ToWhere()
	for _, expr := range expressions {
		query = query.Where(expr.Sql, expr.Values...)
	}

	return query
}

// FindOne 查询单条记录
func (qb *QueryBuilder) FindOne(options QueryOptions) *QueryResult {
	db := qb.getDB()
	if db == nil {
		return &QueryResult{Error: fmt.Errorf("数据库连接 %s 不存在", qb.dbAlias)}
	}

	query := db.Model(qb.model)
	query = qb.applyConditions(query, options.Conditions)

	if options.OrderBy != "" {
		query = query.Order(options.OrderBy)
	}
	fmt.Printf("查询的sql: %+v\n", query)
	result := &QueryResult{}
	err := query.Limit(1).Find(qb.model).Error
	if err != nil {
		fmt.Printf("查询错误: %+v\n", err)
	}
	fmt.Printf("查询到的数据: %+v\n", qb.model)
	result.Error = err
	result.Data = qb.model

	return result
}

// FindList 查询多条记录
func (qb *QueryBuilder) FindList(options QueryOptions) *QueryResult {
	db := qb.getDB()
	if db == nil {
		return &QueryResult{Error: fmt.Errorf("数据库连接 %s 不存在", qb.dbAlias)}
	}

	query := db.Model(qb.model)
	query = qb.applyConditions(query, options.Conditions)

	if options.OrderBy != "" {
		query = query.Order(options.OrderBy)
	}

	if options.Limit > 0 {
		query = query.Limit(options.Limit)
	}

	if options.Offset > 0 {
		query = query.Offset(options.Offset)
	}

	if len(options.Select) > 0 {
		query = query.Select(options.Select)
	}

	// 创建一个新的切片来存储结果
	result := &QueryResult{}
	err := query.Find(qb.model).Error
	result.Error = err
	result.Data = qb.model
	return result
}
