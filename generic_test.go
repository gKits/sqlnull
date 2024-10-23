package sqlnull_test

import (
	"database/sql"
	"encoding/json"
	"testing"

	"github.com/gkits/sqlnull"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type sqlNullable interface {
	json.Marshaler
	json.Unmarshaler
	sql.Scanner
}

func Test_Generic_MarshalJSON(t *testing.T) {
	cases := []struct {
		name string
		in   sqlNullable
		want string
	}{
		{
			name: "marshal generic string to null",
			in: &sqlnull.Null[string]{
				V:     "",
				Valid: false,
			},
			want: "null",
		},
		{
			name: "marshal generic int to null",
			in: &sqlnull.Null[int]{
				V:     0,
				Valid: false,
			},
			want: "null",
		},
		{
			name: "marshal generic string to null with set value",
			in: &sqlnull.Null[string]{
				V:     "this does not matter",
				Valid: false,
			},
			want: "null",
		},
		{
			name: "marshal generic string",
			in: &sqlnull.Null[string]{
				V:     "this is sooo important",
				Valid: true,
			},
			want: `"this is sooo important"`,
		},
		{
			name: "marshal generic int",
			in: &sqlnull.Null[int]{
				V:     69,
				Valid: true,
			},
			want: "69",
		},
		{
			name: "marshal generic float64",
			in: &sqlnull.Null[float64]{
				V:     420.001,
				Valid: true,
			},
			want: "420.001",
		},
		{
			name: "marshal generic struct",
			in: &sqlnull.Null[struct {
				Name string
				Age  int
			}]{
				V: struct {
					Name string
					Age  int
				}{"Joe", 25},
				Valid: true,
			},
			want: `{ "Name": "Joe", "Age": 25 }`,
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
