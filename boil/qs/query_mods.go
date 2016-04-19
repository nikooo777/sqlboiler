package qs

import "github.com/pobri19/sqlboiler/boil"

type QueryMod func(q *boil.Query)

func Limit(limit int) QueryMod {
	return func(q *boil.Query) {
		boil.SetLimit(q, limit)
	}
}

func InnerJoin(on string, args ...interface{}) QueryMod {
	return func(q *boil.Query) {
		boil.SetInnerJoin(q, on, args...)
	}
}

func OuterJoin(on string, args ...interface{}) QueryMod {
	return func(q *boil.Query) {
		boil.SetOuterJoin(q, on, args...)
	}
}

func LeftOuterJoin(on string, args ...interface{}) QueryMod {
	return func(q *boil.Query) {
		boil.SetLeftOuterJoin(q, on, args...)
	}
}

func RightOuterJoin(on string, args ...interface{}) QueryMod {
	return func(q *boil.Query) {
		boil.SetRightOuterJoin(q, on, args...)
	}
}

func Select(columns ...string) QueryMod {
	return func(q *boil.Query) {
		boil.SetSelect(q, columns...)
	}
}

func Where(clause string, args ...interface{}) QueryMod {
	return func(q *boil.Query) {
		boil.SetWhere(q, clause, args...)
	}
}

func GroupBy(clause string) QueryMod {
	return func(q *boil.Query) {
		boil.SetGroupBy(q, clause)
	}
}

func OrderBy(clause string) QueryMod {
	return func(q *boil.Query) {
		boil.SetOrderBy(q, clause)
	}
}

func Having(clause string) QueryMod {
	return func(q *boil.Query) {
		boil.SetHaving(q, clause)
	}
}

func From(table string) QueryMod {
	return func(q *boil.Query) {
		boil.SetFrom(q, table)
	}
}