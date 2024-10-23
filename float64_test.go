package sqlnull_test

import (
	"encoding/json"
	"testing"

	"github.com/gkits/sqlnull"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_Float64_MarshalJSON(t *testing.T) {
	cases := []struct {
		name string
		in   sqlnull.NullFloat64
		want string
	}{
		{
			name: "marshal 0 float to null",
			in: sqlnull.NullFloat64{
				Float64: 0,
				Valid:   false,
			},
			want: "null",
		},
		{
			name: "marshal non 0 float to null",
			in: sqlnull.NullFloat64{
				Float64: 1234.56789,
				Valid:   false,
			},
			want: "null",
		},
		{
			name: "marshal 0 float",
			in: sqlnull.NullFloat64{
				Float64: 0.000,
				Valid:   true,
			},
			want: "0",
		},
		{
			name: "marshal non 0 float",
			in: sqlnull.NullFloat64{
				Float64: 10101.0101,
				Valid:   true,
			},
			want: "10101.0101",
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
