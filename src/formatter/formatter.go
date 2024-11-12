package formatter

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/syucream/hakagi/src/constraint"
)

const (
	baseSql   = "ALTER TABLE %s ADD CONSTRAINT FOREIGN KEY (%s) REFERENCES %s(%s);"
	customSql = "ALTER TABLE %s ADD CONSTRAINT %s FOREIGN KEY (%s) REFERENCES %s(%s);"
)

func FormatSql(constraints []constraint.Constraint) string {
	var queries []string

	for i, c := range constraints {
		var q string

		// If the table name is too long, the foreign key name will also be too long,
		// and an error will occur if the foreign key name exceeds 64 characters.
		// I think 4 digits is enough for numbering.
		if len(c.Table) >= 56 {
			fk_name := c.Table[:56] + "_fk_" + strconv.Itoa(i)
			q = fmt.Sprintf(customSql, c.Table, fk_name, c.Column, c.ReferedTable, c.ReferedColumn)
		} else {
			q = fmt.Sprintf(baseSql, c.Table, c.Column, c.ReferedTable, c.ReferedColumn)
		}

		queries = append(queries, q)
	}

	return strings.Join(queries, "\n")
}
