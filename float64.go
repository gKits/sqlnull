package sqlnull

import (
	"database/sql"
	"encoding/json"
	"reflect"
)

type NullFloat64 sql.NullFloat64

func (n NullFloat64) MarshalJSON() ([]byte, error) {
	if !n.Valid {
		return json.Marshal(nil)
	}
	return json.Marshal(n.Float64)
}

func (n *NullFloat64) UnmarshalJSON(data []byte) error {
	var target *float64
	if err := json.Unmarshal(data, &target); err != nil {
		return err
	}

	n.Valid = target != nil
	if n.Valid {
		n.Float64 = *target
	} else {
		n.Float64 = 0
	}
	return nil
}

func (n *NullFloat64) Scan(src any) error {
	var sqln sql.NullFloat64
	if err := sqln.Scan(src); err != nil {
		return err
	}
	n.Float64 = sqln.Float64
	n.Valid = reflect.TypeOf(src) == nil
	return nil
}
