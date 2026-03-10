package field

import (
	"gorm.io/gorm/clause"
)

// Gt creates a greater than comparison expression (field > value).
func (f Field[T]) Gt(value T) clause.Expression {
	return clause.Gt{Column: f.column, Value: value}
}

// GtExpr creates a greater than comparison expression (field > expression).
func (f Field[T]) GtExpr(expr clause.Expression) clause.Expression {
	return clause.Gt{Column: f.column, Value: expr}
}

// Gte creates a greater than or equal comparison expression (field >= value).
func (f Field[T]) Gte(value T) clause.Expression {
	return clause.Gte{Column: f.column, Value: value}
}

// GteExpr creates a greater than or equal comparison expression (field >= expression).
func (f Field[T]) GteExpr(expr clause.Expression) clause.Expression {
	return clause.Gte{Column: f.column, Value: expr}
}

// Lt creates a less than comparison expression (field < value).
func (f Field[T]) Lt(value T) clause.Expression {
	return clause.Lt{Column: f.column, Value: value}
}

// LtExpr creates a less than comparison expression (field < expression).
func (f Field[T]) LtExpr(expr clause.Expression) clause.Expression {
	return clause.Lt{Column: f.column, Value: expr}
}

// Lte creates a less than or equal comparison expression (field <= value).
func (f Field[T]) Lte(value T) clause.Expression {
	return clause.Lte{Column: f.column, Value: value}
}

// LteExpr creates a less than or equal comparison expression (field <= expression).
func (f Field[T]) LteExpr(expr clause.Expression) clause.Expression {
	return clause.Lte{Column: f.column, Value: expr}
}

// Like creates a LIKE pattern matching expression (field LIKE pattern).
func (f Field[T]) Like(pattern string) clause.Expression {
	return clause.Like{Column: f.column, Value: pattern}
}

// NotLike creates a NOT LIKE pattern matching expression (field NOT LIKE pattern).
func (f Field[T]) NotLike(pattern string) clause.Expression {
	return clause.Expr{SQL: "? NOT LIKE ?", Vars: []any{f.column, pattern}}
}

// ILike creates a case-insensitive LIKE pattern matching expression (field ILIKE pattern).
func (f Field[T]) ILike(pattern string) clause.Expression {
	return clause.Expr{SQL: "? ILIKE ?", Vars: []any{f.column, pattern}}
}

// NotILike creates a case-insensitive NOT LIKE pattern matching expression (field NOT ILIKE pattern).
func (f Field[T]) NotILike(pattern string) clause.Expression {
	return clause.Expr{SQL: "? NOT ILIKE ?", Vars: []any{f.column, pattern}}
}

// Regexp creates a regular expression matching expression (field REGEXP pattern).
func (f Field[T]) Regexp(pattern string) clause.Expression {
	return clause.Expr{SQL: "? REGEXP ?", Vars: []any{f.column, pattern}}
}

// NotRegexp creates a regular expression not matching expression (field NOT REGEXP pattern).
func (f Field[T]) NotRegexp(pattern string) clause.Expression {
	return clause.Expr{SQL: "? NOT REGEXP ?", Vars: []any{f.column, pattern}}
}

// Between creates a range comparison expression (field BETWEEN v1 AND v2).
func (f Field[T]) Between(v1, v2 T) clause.Expression {
	return clause.And(
		clause.Gte{Column: f.column, Value: v1},
		clause.Lte{Column: f.column, Value: v2},
	)
}

// In creates an IN comparison expression (field IN (values...)).
func (f Field[T]) In(values ...T) clause.Expression {
	interfaceValues := make([]any, len(values))
	for i, v := range values {
		interfaceValues[i] = v
	}
	return clause.IN{Column: f.column, Values: interfaceValues}
}

// NotIn creates a NOT IN comparison expression (field NOT IN (values...)).
func (f Field[T]) NotIn(values ...T) clause.Expression {
	interfaceValues := make([]any, len(values))
	for i, v := range values {
		interfaceValues[i] = v
	}
	return clause.Not(clause.IN{Column: f.column, Values: interfaceValues})
}

// Basic SQL expression functions for arithmetic operations

// Incr creates an increment expression (field + value).
func (f Field[T]) Incr(value T) AssignerExpression {
	return colOpExpr{col: f.column, sql: "? + ?", vars: []any{f.column, value}}
}

// Decr creates a decrement expression (field - value).
func (f Field[T]) Decr(value T) AssignerExpression {
	return colOpExpr{col: f.column, sql: "? - ?", vars: []any{f.column, value}}
}

// Mul creates a multiplication expression (field * value).
func (f Field[T]) Mul(value T) AssignerExpression {
	return colOpExpr{col: f.column, sql: "? * ?", vars: []any{f.column, value}}
}

// Div creates a division expression (field / value).
func (f Field[T]) Div(value T) AssignerExpression {
	return colOpExpr{col: f.column, sql: "? / ?", vars: []any{f.column, value}}
}
