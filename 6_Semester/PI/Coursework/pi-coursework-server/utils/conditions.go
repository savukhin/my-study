package utils

type WhereCondition struct {
	Sign   string `regroup:"where_sign"`
	Column string `regroup:"where_column"`
	Value  string `regroup:"where_value"`
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
		Column: column,
		Sign:   sign,
		Value:  value,
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
