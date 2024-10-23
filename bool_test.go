package sqlnull_test

import (
	"encoding/json"
	"testing"

	"github.com/gkits/sqlnull"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_Bool_MarshalJSON(t *testing.T) {
	cases := []struct {
		name string
		in   sqlnull.NullBool
		want string
	}{
		{
			name: "marshal false bool to null",
			in: sqlnull.NullBool{
				Bool:  false,
				Valid: false,
			},
			want: "null",
		},
		{
			name: "marshal true bool to null",
			in: sqlnull.NullBool{
				Bool:  true,
				Valid: false,
			},
			want: "null",
		},
		{
			name: "marshal false bool",
			in: sqlnull.NullBool{
				Bool:  false,
				Valid: true,
			},
			want: "false",
		},
		{
			name: "marshal true bool",
			in: sqlnull.NullBool{
				Bool:  true,
				Valid: true,
			},
			want: "true",
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
