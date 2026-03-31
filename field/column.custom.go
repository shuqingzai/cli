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

// SimpleColumnString returns the column name for selecting
//
// Simple build column name for selecting cases
func (f Field[T]) SimpleColumnString() string {
	return simpleColumnString(f.column)
}

// SimpleColumnString returns the column name for selecting
//
// Simple build column name for selecting cases
func (s String) SimpleColumnString() string {
	return simpleColumnString(s.column)
}

// SimpleColumnString returns the column name for selecting
//
// Simple build column name for selecting cases
func (b Bool) SimpleColumnString() string {
	return simpleColumnString(b.column)
}

// SimpleColumnString returns the column name for selecting
//
// Simple build column name for selecting cases
func (b Bytes) SimpleColumnString() string {
	return simpleColumnString(b.column)
}

// SimpleColumnString returns the column name for selecting
//
// Simple build column name for selecting cases
func (t Time) SimpleColumnString() string {
	return simpleColumnString(t.column)
}

// SimpleColumnString returns the column name for selecting
//
// Simple build column name for selecting cases
func (n Number[T]) SimpleColumnString() string {
	return simpleColumnString(n.column)
}
