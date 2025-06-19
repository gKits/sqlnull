package sqlnull

import (
	"database/sql"
	"encoding/json"
	"reflect"
)

type NullInt64 sql.NullInt64

func (n NullInt64) MarshalJSON() ([]byte, error) {
	if !n.Valid {
		return json.Marshal(nil)
	}
	return json.Marshal(n.Int64)
}

func (n *NullInt64) UnmarshalJSON(data []byte) error {
	var target *int64
	if err := json.Unmarshal(data, &target); err != nil {
		return err
	}

	n.Valid = target != nil
	if n.Valid {
		n.Int64 = *target
	} else {
		n.Int64 = 0
	}
	return nil
}

func (n *NullInt64) Scan(src any) error {
	var sqln sql.NullInt64
	if err := sqln.Scan(src); err != nil {
		return err
	}
	n.Int64 = sqln.Int64
	n.Valid = reflect.TypeOf(src) != nil
	return nil
}

func (n *NullInt64) IsZero() bool {
	return !n.Valid
}
