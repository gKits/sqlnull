package sqlnull_test

import (
	"encoding/json"
	"testing"

	"github.com/gkits/sqlnull"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_Int16_MarshalJSON(t *testing.T) {
	cases := []struct {
		name string
		in   sqlnull.NullInt16
		want string
	}{
		{
			name: "marshal 0 int16 to null",
			in: sqlnull.NullInt16{
				Int16: 0,
				Valid: false,
			},
			want: "null",
		},
		{
			name: "marshal non 0 int16 to null",
			in: sqlnull.NullInt16{
				Int16: 32767,
				Valid: false,
			},
			want: "null",
		},
		{
			name: "marshal 0 int16",
			in: sqlnull.NullInt16{
				Int16: 0,
				Valid: true,
			},
			want: "0",
		},
		{
			name: "marshal non 0 int16",
			in: sqlnull.NullInt16{
				Int16: 32767,
				Valid: true,
			},
			want: "32767",
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
