package dialect

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestClickhouse(t *testing.T) {
	for _, test := range []struct {
		in   string
		want string
	}{
		{
			in:   "table.col",
			want: "`table`.`col`",
		},
		{
			in:   "col",
			want: "`col`",
		},
	} {
		assert.Equal(t, test.want, Clickhouse.QuoteIdent(test.in))
	}
}
