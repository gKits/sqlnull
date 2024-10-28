package sqlnull_test

import (
	"database/sql"
	"encoding/json"
	"testing"
	"time"

	"github.com/gkits/sqlnull"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type target struct {
	Generic sqlnull.Null[string] `json:"generic"`
	String  sqlnull.NullString   `json:"string"`
	Bool    sqlnull.NullBool     `json:"bool"`
	Byte    sqlnull.NullByte     `json:"byte"`
	Int16   sqlnull.NullInt16    `json:"int16"`
	Int32   sqlnull.NullInt32    `json:"int32"`
	Int64   sqlnull.NullInt64    `json:"int64"`
	Float64 sqlnull.NullFloat64  `json:"float64"`
	Time    sqlnull.NullTime     `json:"time"`
}

func Test_MarshalJSON(t *testing.T) {
	cases := []struct {
		name string
		in   target
		want string
	}{
		{
			name: "successfully marshal with values",
			in: target{
				Generic: sqlnull.Null[string]{V: "generic", Valid: true},
				String:  sqlnull.NullString{String: "string", Valid: true},
				Bool:    sqlnull.NullBool{Bool: true, Valid: true},
				Byte:    sqlnull.NullByte{Byte: 255, Valid: true},
				Int16:   sqlnull.NullInt16{Int16: 16, Valid: true},
				Int32:   sqlnull.NullInt32{Int32: 32, Valid: true},
				Int64:   sqlnull.NullInt64{Int64: 64, Valid: true},
				Float64: sqlnull.NullFloat64{Float64: 64.6464, Valid: true},
				Time:    sqlnull.NullTime{Time: time.Date(2024, 10, 23, 17, 50, 0, 0, time.UTC), Valid: true},
			},
			want: `{ 
                "generic": "generic",
                "string": "string",
                "bool": true,
                "byte": 255,
                "int16": 16,
                "int32": 32,
                "int64": 64,
                "float64": 64.6464,
                "time": "2024-10-23T17:50:00Z"
            }`,
		},
		{
			name: "successfully marshal with null",
			in: target{
				Generic: sqlnull.Null[string]{V: "generic", Valid: false},
				String:  sqlnull.NullString{String: "string", Valid: false},
				Bool:    sqlnull.NullBool{Bool: true, Valid: false},
				Byte:    sqlnull.NullByte{Byte: 255, Valid: false},
				Int16:   sqlnull.NullInt16{Int16: 16, Valid: false},
				Int32:   sqlnull.NullInt32{Int32: 32, Valid: false},
				Int64:   sqlnull.NullInt64{Int64: 64, Valid: false},
				Float64: sqlnull.NullFloat64{Float64: 64.6464, Valid: false},
				Time:    sqlnull.NullTime{Time: time.Date(2024, 10, 23, 17, 50, 0, 0, time.UTC), Valid: false},
			},
			want: `{ 
                "generic": null,
                "string": null,
                "bool": null,
                "byte": null,
                "int16": null,
                "int32": null,
                "int64": null,
                "float64": null,
                "time": null
            }`,
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

func Test_UnmarshalJSON(t *testing.T) {
	cases := []struct {
		name string
		in   []byte
		want target
	}{
		{
			name: "successfully marshal with values",
			in: []byte(`{ 
                "generic": "generic", 
                "string": "string",
                "bool": true,
                "byte": 255,
                "int16": 16,
                "int32": 32,
                "int64": 64,
                "float64": 64.6464,
                "time": "2024-10-23T17:50:00Z"
            }`),
			want: target{
				Generic: sqlnull.Null[string]{V: "generic", Valid: true},
				String:  sqlnull.NullString{String: "string", Valid: true},
				Bool:    sqlnull.NullBool{Bool: true, Valid: true},
				Byte:    sqlnull.NullByte{Byte: 255, Valid: true},
				Int16:   sqlnull.NullInt16{Int16: 16, Valid: true},
				Int32:   sqlnull.NullInt32{Int32: 32, Valid: true},
				Int64:   sqlnull.NullInt64{Int64: 64, Valid: true},
				Float64: sqlnull.NullFloat64{Float64: 64.6464, Valid: true},
				Time:    sqlnull.NullTime{Time: time.Date(2024, 10, 23, 17, 50, 0, 0, time.UTC), Valid: true},
			},
		},
		{
			name: "successfully marshal with null",
			in: []byte(`{ 
                "generic": null,
                "string": null,
                "bool": null,
                "byte": null,
                "int16": null,
                "int32": null,
                "int64": null,
                "float64": null,
                "time": null
            }`),
			want: target{
				Generic: sqlnull.Null[string]{V: "", Valid: false},
				String:  sqlnull.NullString{String: "", Valid: false},
				Bool:    sqlnull.NullBool{Bool: false, Valid: false},
				Byte:    sqlnull.NullByte{Byte: 0, Valid: false},
				Int16:   sqlnull.NullInt16{Int16: 0, Valid: false},
				Int32:   sqlnull.NullInt32{Int32: 0, Valid: false},
				Int64:   sqlnull.NullInt64{Int64: 0, Valid: false},
				Float64: sqlnull.NullFloat64{Float64: 0, Valid: false},
				Time:    sqlnull.NullTime{Time: time.Time{}, Valid: false},
			},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			var got target
			err := json.Unmarshal(c.in, &got)
			require.NoError(t, err)
			assert.Equal(t, c.want, got)
		})
	}
}

func Test_Scan(t *testing.T) {
	db, err := sql.Open("sqlite3", "/tmp/test.sqlite")
	require.NoError(t, err)

	_, err = db.Exec(`
        DROP TABLE IF EXISTS test;
        CREATE TABLE test (
            id INTEGER PRIMARY KEY,
            generic TEXT,
            string TEXT,
            bool INTEGER,
            byte INTEGER,
            int64 INTEGER,
            int32 INTEGER,
            int16 INTEGER,
            float64 REAL,
            time TIMESTAMP
        );

        INSERT INTO test (id, generic, string, bool, byte, int64, int32, int16, float64, time)
        VALUES  (1, 'generic', 'string', TRUE, 255, 64, 32, 16, 64.6464, '2024-10-23T17:50:00Z'),
                (2, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);`)
	require.NoError(t, err)

	query, err := db.Prepare(`SELECT * FROM test WHERE id = $1;`)
	require.NoError(t, err)

	cases := []struct {
		name string
		in   int
		want target
	}{
		{
			name: "scan values",
			in:   1,
			want: target{
				Generic: sqlnull.Null[string]{V: "generic", Valid: true},
				String:  sqlnull.NullString{String: "string", Valid: true},
				Bool:    sqlnull.NullBool{Bool: true, Valid: true},
				Byte:    sqlnull.NullByte{Byte: 255, Valid: true},
				Int16:   sqlnull.NullInt16{Int16: 16, Valid: true},
				Int32:   sqlnull.NullInt32{Int32: 32, Valid: true},
				Int64:   sqlnull.NullInt64{Int64: 64, Valid: true},
				Float64: sqlnull.NullFloat64{Float64: 64.6464, Valid: true},
				Time:    sqlnull.NullTime{Time: time.Date(2024, 10, 23, 17, 50, 0, 0, time.UTC), Valid: true},
			},
		},
		{
			name: "scan null",
			in:   2,
			want: target{
				Generic: sqlnull.Null[string]{V: "", Valid: false},
				String:  sqlnull.NullString{String: "", Valid: false},
				Bool:    sqlnull.NullBool{Bool: false, Valid: false},
				Byte:    sqlnull.NullByte{Byte: 0, Valid: false},
				Int16:   sqlnull.NullInt16{Int16: 0, Valid: false},
				Int32:   sqlnull.NullInt32{Int32: 0, Valid: false},
				Int64:   sqlnull.NullInt64{Int64: 0, Valid: false},
				Float64: sqlnull.NullFloat64{Float64: 0, Valid: false},
				Time:    sqlnull.NullTime{Time: time.Time{}, Valid: false},
			},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			var got target
			var id int
			err := query.QueryRow(c.in).Scan(
				&id,
				&got.Generic,
				&got.String,
				&got.Bool,
				&got.Byte,
				&got.Int64,
				&got.Int32,
				&got.Int16,
				&got.Float64,
				&got.Time,
			)
			require.NoError(t, err)
			assert.Equal(t, c.want, got)
		})
	}
}
