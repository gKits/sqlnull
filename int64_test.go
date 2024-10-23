package sqlnull_test

import (
	"encoding/json"
	"testing"

	"github.com/gkits/sqlnull"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_Int64_MarshalJSON(t *testing.T) {
	cases := []struct {
		name string
		in   sqlnull.NullInt64
		want string
	}{
		{
			name: "marshal 0 int64 to null",
			in: sqlnull.NullInt64{
				Int64: 0,
				Valid: false,
			},
			want: "null",
		},
		{
			name: "marshal non 0 int64 to null",
			in: sqlnull.NullInt64{
				Int64: 9223372036854775807,
				Valid: false,
			},
			want: "null",
		},
		{
			name: "marshal 0 int64",
			in: sqlnull.NullInt64{
				Int64: 0,
				Valid: true,
			},
			want: "0",
		},
		{
			name: "marshal non 0 int64",
			in: sqlnull.NullInt64{
				Int64: 9223372036854775807,
				Valid: true,
			},
			want: "9223372036854775807",
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got, err := json.Marshal(c.in)
			require.NoError(t, err)
			assert.JSONEq(t, c.want, string(got))
		})
	}
}
