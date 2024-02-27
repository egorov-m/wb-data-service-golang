package goqu

import (
	"github.com/doug-martin/goqu/v9"
	_ "github.com/doug-martin/goqu/v9/dialect/postgres"
)

var Dialect goqu.DialectWrapper

// Aliases
type (
	Ex     = goqu.Ex
	Record = goqu.Record
)

var (
	C = goqu.C
)

func init() {
	Dialect = goqu.Dialect("postgres")
}
