package sqlnull_test

import (
	"encoding/json"
	"testing"

	"github.com/gkits/sqlnull"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_String_MarshalJSON(t *testing.T) {
	cases := []struct {
		name string
		in   sqlnull.NullString
		want string
	}{
		{
			name: "marshal empty string to null",
			in: sqlnull.NullString{
				String: "",
				Valid:  false,
			},
			want: "null",
		},
		{
			name: "marshal non empty string to null",
			in: sqlnull.NullString{
				String: "null",
				Valid:  false,
			},
			want: "null",
		},
		{
			name: "marshal empty string",
			in: sqlnull.NullString{
				String: "",
				Valid:  true,
			},
			want: `""`,
		},
		{
			name: "marshal non empty string",
			in: sqlnull.NullString{
				String: "null",
				Valid:  true,
			},
			want: `"null"`,
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
