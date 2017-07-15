package chdbr

import "strings"

type joinType uint8
type strictness uint8

const (
	inner joinType = iota
	left
)

const (
	any strictness = iota
	all
)

// https://clickhouse.yandex/reference_en.html#JOIN clause
func join(s strictness, t joinType, table interface{}, usings ...string) Builder {
	return BuildFunc(func(d Dialect, buf Buffer) error {
		buf.WriteString(" ")
		switch s {
		case any:
			buf.WriteString("ANY ")
		case all:
			buf.WriteString("ALL ")
		}
		switch t {
		case inner:
			buf.WriteString("INNER ")
		case left:
			buf.WriteString("LEFT ")
		}
		buf.WriteString("JOIN ")
		switch table := table.(type) {
		case string:
			buf.WriteString(d.QuoteIdent(table))
		default:
			// buf.WriteString("(")
			buf.WriteString(placeholder)
			buf.WriteValue(table)
			// buf.WriteString(")")
		}
		buf.WriteString(" USING ")
		buf.WriteString(strings.Join(usings, ", "))
		return nil
	})
}
