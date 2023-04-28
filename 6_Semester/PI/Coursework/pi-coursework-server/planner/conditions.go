package planner

import (
	"errors"
	"strconv"
)

type WhereCondition struct {
	Sign     string `regroup:"where_sign"`
	Column   string `regroup:"where_column"`
	Value    string `regroup:"where_value"`
	ValueStr string `regroup:"where_value_str"`
	ValueInt int32  `regroup:"where_value_int"`
}

func (where *WhereCondition) ExtractValue() string {
	if where.Value[0] == '\'' {
		return where.ValueStr
	} else {
		return strconv.Itoa(int(where.ValueInt))
	}
}

func (where *WhereCondition) GetIntValue() (int32, error) {
	if where.Value[0] == '\'' {
		return 0, errors.New("where condition is not digit")
	} else {
		return where.ValueInt, nil
	}
}

type WhereConditionCheck struct {
	WhereStr string `regroup:"has_where"`
	WhereCondition
	HasWhere bool
}

type LimitCondition struct {
	Limit    int32  `regroup:"limit"`
	LimitStr string `regroup:"has_limit"`
	HasLimit bool
}

func NewWhereCondition(column, sign, value string) *WhereCondition {
	return &WhereCondition{
		Column:   column,
		Sign:     sign,
		Value:    value,
		ValueStr: value,
	}
}

func NewWhereConditionCheck(column, sign, value string) *WhereConditionCheck {
	return &WhereConditionCheck{
		HasWhere:       true,
		WhereStr:       "where",
		WhereCondition: *NewWhereCondition(column, sign, value),
	}
}

func NewLimitCondition(limit int32) *LimitCondition {
	return &LimitCondition{
		Limit:    limit,
		LimitStr: "limit",
		HasLimit: true,
	}
}
