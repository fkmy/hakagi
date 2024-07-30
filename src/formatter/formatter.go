package formatter

import (
	"fmt"
	"strings"
	"strconv"

	"github.com/syucream/hakagi/src/constraint"
)

const (
	baseSql = "ALTER TABLE %s ADD CONSTRAINT FOREIGN KEY (%s) REFERENCES %s(%s);"
	customSql = "ALTER TABLE %s ADD CONSTRAINT %s FOREIGN KEY (%s) REFERENCES %s(%s);"
)

func FormatSql(constraints []constraint.Constraint) string {
	var queries []string

	for i := 1, c := range constraints {
	  if len(c.Table) >= 57 {
			fk_name := c.Table[57] + "_ibfk_" + strconv.Itoa(i)
			q := fmt.Sprintf(customSql, c.Table, fk_name, c.Column, c.ReferedTable, c.ReferedColumn)
		} else {
			q := fmt.Sprintf(baseSql, c.Table, c.Column, c.ReferedTable, c.ReferedColumn)
		}

		queries = append(queries, q)
	}

	return strings.Join(queries, "\n")
}
