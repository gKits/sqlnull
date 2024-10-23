package sqlnull_test

import (
	"encoding/json"
	"testing"

	"github.com/gkits/sqlnull"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_Byte_MarshalJSON(t *testing.T) {
	cases := []struct {
		name string
		in   sqlnull.NullByte
		want string
	}{
		{
			name: "marshal zero byte to null",
			in: sqlnull.NullByte{
				Byte:  0x00,
				Valid: false,
			},
			want: "null",
		},
		{
			name: "marshal non zero byte to null",
			in: sqlnull.NullByte{
				Byte:  0x99,
				Valid: false,
			},
			want: "null",
		},
		{
			name: "marshal zero byte",
			in: sqlnull.NullByte{
				Byte:  0x00,
				Valid: true,
			},
			want: "0",
		},
		{
			name: "marshal non zero byte",
			in: sqlnull.NullByte{
				Byte:  0x69,
				Valid: true,
			},
			want: "105",
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
