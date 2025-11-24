package dao

import (
	"fmt"
	"strings"
)

// Condition 定义一个条件结构体，用于存储单个条件
type Condition struct {
	Field    string // 字段名
	Operator string // 操作符，如 =, >, <, LIKE 等
	Value    []any  //  条件对应的值 (例如: "John Doe", 25, []int{1, 2, 3})
}

// Conditions 定义一个条件切片类型，用于存储多个条件
type Conditions []Condition

// Express 定义一个表达式结构体，用于存储单个表达式
type Express struct {
	Sql    string
	Values []any
}

func (ex *Express) condErr() error { return nil }

// ToWhere 方法将条件切片转换为SQL表达式切片
func (conds *Conditions) ToWhere() []Express {
	expresses := make([]Express, 0) // 初始化切片
	for _, cond := range *conds {
		if cond.Field == "" || cond.Operator == "" {
			continue
		}
		if cond.Value == nil {
			cond.Value = []any{cond.Operator}
			cond.Operator = "="
		}
		expresses = append(expresses, *cond.factory(&Express{})) // 调用factory方法创建表达式
	}
	return expresses
}

// factory 方法根据操作符创建对应的sql表达式
func (cond *Condition) factory(expr *Express) *Express {
	op := strings.ToUpper(strings.TrimSpace(cond.Operator)) // 确保操作符是大写的
	switch op {
	case "FORMAT":
		expr = &Express{
			Sql:    cond.Field,
			Values: cond.Value,
		}
	case "LIKE":
		expr = &Express{
			Sql:    fmt.Sprintf("`%s` LIKE ?", cond.Field),
			Values: cond.Value,
		}
	case "IN", "NOT IN":

		expr = &Express{
			Sql:    fmt.Sprintf("`%s` %s (?)", cond.Field, op),
			Values: cond.Value,
		}

	case "=", ">", "<", ">=", "<=", "<>":

		expr = &Express{
			Sql:    fmt.Sprintf("`%s` %s ?", cond.Field, op),
			Values: cond.Value,
		}
	}
	return expr
}
