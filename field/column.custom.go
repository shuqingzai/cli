package field

import (
	"strings"

	"gorm.io/gorm/clause"
)

// simpleColumnString returns the column name for selecting
//
// Simple build column name for selecting cases
func simpleColumnString(column clause.Column) string {
	var buf strings.Builder
	if column.Table != "" {
		if column.Table != clause.CurrentTable {
			buf.WriteString(column.Table + ".")
		}
	}

	if column.Name == clause.PrimaryKey {
		// only use id for primary key
		buf.WriteString("id")
	} else {
		buf.WriteString(column.Name)
	}

	if column.Alias != "" {
		buf.WriteString(" AS ")
		buf.WriteString(column.Alias)
	}

	return buf.String()
}

// SimpleColumn returns the column name for selecting
//
// Simple build column name for selecting cases
func (f Field[T]) SimpleColumn() string {
	return simpleColumnString(f.column)
}

// SimpleColumn returns the column name for selecting
//
// Simple build column name for selecting cases
func (s String) SimpleColumn() string {
	return simpleColumnString(s.column)
}

// SimpleColumn returns the column name for selecting
//
// Simple build column name for selecting cases
func (b Bool) SimpleColumn() string {
	return simpleColumnString(b.column)
}

// SimpleColumn returns the column name for selecting
//
// Simple build column name for selecting cases
func (b Bytes) SimpleColumn() string {
	return simpleColumnString(b.column)
}

// SimpleColumn returns the column name for selecting
//
// Simple build column name for selecting cases
func (t Time) SimpleColumn() string {
	return simpleColumnString(t.column)
}

// SimpleColumn returns the column name for selecting
//
// Simple build column name for selecting cases
func (n Number[T]) SimpleColumn() string {
	return simpleColumnString(n.column)
}
