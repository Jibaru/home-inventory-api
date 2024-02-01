package repositories

const (
	AndLogicalOperator LogicalOperator = "AND"
	OrLogicalOperator  LogicalOperator = "OR"
)

const (
	LikeComparisonOperator  ComparisonOperator = "LIKE"
	EqualComparisonOperator ComparisonOperator = "="
)

type ComparisonOperator string
type LogicalOperator string

type PageFilter struct {
	Offset int
	Limit  int
}

type QueryFilter struct {
	ConditionGroups []ConditionGroup
}

type ConditionGroup struct {
	Operator   LogicalOperator
	Conditions []Condition
}

type Condition struct {
	Field    string
	Operator ComparisonOperator
	Value    interface{}
}
