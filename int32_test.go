package sqlnull_test

import (
	"encoding/json"
	"testing"

	"github.com/gkits/sqlnull"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_Int32_MarshalJSON(t *testing.T) {
	cases := []struct {
		name string
		in   sqlnull.NullInt32
		want string
	}{
		{
			name: "marshal 0 int32 to null",
			in: sqlnull.NullInt32{
				Int32: 0,
				Valid: false,
			},
			want: "null",
		},
		{
			name: "marshal non 0 int32 to null",
			in: sqlnull.NullInt32{
				Int32: 2147483647,
				Valid: false,
			},
			want: "null",
		},
		{
			name: "marshal 0 int32",
			in: sqlnull.NullInt32{
				Int32: 0,
				Valid: true,
			},
			want: "0",
		},
		{
			name: "marshal non 0 int32",
			in: sqlnull.NullInt32{
				Int32: 2147483647,
				Valid: true,
			},
			want: "2147483647",
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
