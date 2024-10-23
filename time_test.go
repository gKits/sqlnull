package sqlnull_test

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/gkits/sqlnull"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_Time_MarshalJSON(t *testing.T) {
	cases := []struct {
		name string
		in   sqlnull.NullTime
		want string
	}{
		{
			name: "marshal zero timestamp to null",
			in: sqlnull.NullTime{
				Time:  time.Time{},
				Valid: false,
			},
			want: "null",
		},
		{
			name: "marshal non zero timestamp to null",
			in: sqlnull.NullTime{
				Time:  time.Date(2024, 10, 23, 17, 50, 0, 0, time.UTC),
				Valid: false,
			},
			want: "null",
		},
		{
			name: "marshal zero timestamp",
			in: sqlnull.NullTime{
				Time:  time.Time{},
				Valid: true,
			},
			want: `"0001-01-01T00:00:00Z"`,
		},
		{
			name: "marshal non zero timestamp",
			in: sqlnull.NullTime{
				Time:  time.Date(2024, 10, 23, 17, 50, 0, 0, time.UTC),
				Valid: true,
			},
			want: `"2024-10-23T17:50:00Z"`,
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
