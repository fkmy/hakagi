package formatter

import (
	"fmt"
	"strings"

	"github.com/syucream/hakagi/src/constraint"
)

const (
	baseSql = `ALTER TABLE
	%s
ADD CONSTRAINT
  %s_ibfk_%s%s
FOREIGN KEY
  (%s)
REFERENCES
  %s(%s);
`
)

func FormatSql(constraints []constraint.Constraint) string {
	var queries []string

	for _, c := range constraints {
		q := fmt.Sprintf(baseSql, c.Table, c.Table, c.ReferedTable, c.ReferedColumn, c.Column, c.ReferedTable, c.ReferedColumn)
		queries = append(queries, q)
	}

	return strings.Join(queries, "\n")
}
